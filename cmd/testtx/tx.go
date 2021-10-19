package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/localrpc"
	"github.com/MadBase/MadNet/utils"
)

var numEpochs uint32 = 1

type worker struct {
	f      *funder
	client *localrpc.Client
	signer aobjs.Signer
	acct   []byte
	idx    int
}

type funder struct {
	ctx         context.Context
	cf          func()
	wg          sync.WaitGroup
	signer      *crypto.Secp256k1Signer
	client      *localrpc.Client
	acct        []byte
	numChildren int
	children    []*worker
	nodeList    []string
	baseIdx     int
}

func (f *funder) setupDataStoreTransaction3(ctx context.Context, signer aobjs.Signer, ownerAcct []byte, msg string, ind string) (*aobjs.Tx, error) {
	index := crypto.Hasher([]byte(ind))
	deposit, err := aobjs.BaseDepositEquation(uint32(len(msg)), numEpochs)
	if err != nil {
		return nil, err
	}

	curveSpec := f.getCurveSpec(signer)

	var ds *aobjs.TXOut
	for {
		resp, err := f.client.PaginateDataStoreUTXOByOwner(ctx, curveSpec, ownerAcct, 1, utils.CopySlice(index))
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		utxoIDs := [][]byte{}
		for i := 0; i < len(resp); i++ {
			utxoIDs = append(utxoIDs, resp[i].UTXOID)
		}
		if len(utxoIDs) != 1 {
			time.Sleep(1 * time.Second)
			continue
		}

		resp2, err := f.client.GetUTXO(ctx, utxoIDs)
		if err != nil {
			continue
		}

		if len(resp2) != 1 {
			time.Sleep(1 * time.Second)
			continue
		} else {
			ds = resp2[0]
			break
		}
	}
	bn, err := f.client.GetBlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	v, err := ds.RemainingValue(bn + constants.EpochLength)
	if err != nil {
		return nil, err
	}
	depositClone := deposit.Clone()
	valueNeeded, err := depositClone.Sub(depositClone, v)
	if err != nil {
		panic(err)
	}

	consumedUtxos, consumedValue, err := f.blockingGetFunding(ctx, f.client, f.getCurveSpec(f.signer), f.acct, valueNeeded)
	if err != nil {
		panic(err)
	}
	if consumedValue.Lt(valueNeeded) {
		fmt.Printf("ACCOUNT DOES NOT HAVE ENOUGH FUNDING: REQUIRES:%v    HAS:%v\n", valueNeeded, consumedValue)
	}

	tx := &aobjs.Tx{
		Vin:  aobjs.Vin{},
		Vout: aobjs.Vout{},
	}
	chainID := uint32(42)
	for _, utxo := range consumedUtxos {
		consumedVS, err := utxo.ValueStore()
		if err != nil {
			return nil, err
		}
		fmt.Println(consumedVS.Value())
		chainID, err = consumedVS.ChainID()
		if err != nil {
			return nil, err
		}
		txIn, err := utxo.MakeTxIn()
		if err != nil {
			return nil, err
		}
		tx.Vin = append(tx.Vin, txIn)
	}

	dsTxIn, err := ds.MakeTxIn()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	tx.Vin = append(tx.Vin, dsTxIn)
	consumedUtxos = append(consumedUtxos, ds)
	fmt.Printf("THE LEN OF VIN %v \n", len(tx.Vin))

	consumedValue.Add(consumedValue, v)
	valueOut := uint256.Zero()
	{
		en, err := f.client.GetEpochNumber(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Printf("REQUIRED DEPOSIT %v \n", deposit)
		valueOut.Add(valueOut, deposit)
		newOwner := &aobjs.DataStoreOwner{}
		newOwner.New(ownerAcct, f.getCurveSpec(signer))
		newDataStore := &aobjs.DataStore{
			DSLinker: &aobjs.DSLinker{
				DSPreImage: &aobjs.DSPreImage{
					ChainID:  chainID,
					Index:    index,
					IssuedAt: en + 1,
					Deposit:  deposit,
					RawData:  []byte(msg),
					TXOutIdx: 0,
					Owner:    newOwner,
					Fee:      new(uint256.Uint256).SetZero(),
				},
				TxHash: make([]byte, constants.HashLen),
			},
		}
		eoe, err := newDataStore.EpochOfExpiration()
		if err != nil {
			return nil, err
		}
		fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
		fmt.Printf("DS:  index:%x    deposit:%v    EpochOfExpire:%v    msg:%s\n", index, deposit, eoe, msg)
		newUTXO := &aobjs.TXOut{}
		newUTXO.NewDataStore(newDataStore)
		tx.Vout = append(tx.Vout, newUTXO)
	}
	fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
	if consumedValue.Gt(valueOut) {
		diff, err := new(uint256.Uint256).Sub(consumedValue.Clone(), valueOut.Clone())
		if err != nil {
			panic(err)
		}
		newOwner := &aobjs.ValueStoreOwner{}
		newOwner.New(ownerAcct, f.getCurveSpec(signer))
		newValueStore := &aobjs.ValueStore{
			VSPreImage: &aobjs.VSPreImage{
				ChainID:  chainID,
				Value:    diff,
				Owner:    newOwner,
				TXOutIdx: 0,
				Fee:      new(uint256.Uint256).SetZero(),
			},
			TxHash: make([]byte, constants.HashLen),
		}
		newUTXO := &aobjs.TXOut{}
		newUTXO.NewValueStore(newValueStore)
		tx.Vout = append(tx.Vout, newUTXO)
		valueOut.Add(valueOut, diff)
	}
	fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
	tx.Vout.SetTxOutIdx()
	fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
	tx.SetTxHash()
	for _, newUtxo := range tx.Vout {
		switch {
		case newUtxo.HasDataStore():
			ds, err := newUtxo.DataStore()
			if err != nil {
				return nil, err
			}
			if err := ds.PreSign(signer); err != nil {
				return nil, err
			}
		default:
			continue
		}
	}
	for idx, consumedUtxo := range consumedUtxos {
		switch {
		case consumedUtxo.HasValueStore():
			consumedVS, err := consumedUtxo.ValueStore()
			if err != nil {
				return nil, err
			}
			txIn := tx.Vin[idx]
			consumedVS.Sign(txIn, signer)
		case consumedUtxo.HasDataStore():
			consumedDS, err := consumedUtxo.DataStore()
			if err != nil {
				return nil, err
			}
			txIn := tx.Vin[idx]
			consumedDS.Sign(txIn, signer)
		}
	}
	fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
	bn, err = f.client.GetBlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Consumed Next:%v    ValueOut:%v\n", consumedValue, valueOut)
	txb, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	fmt.Printf("TX SIZE: %v\n", len(txb))
	return tx, nil
}

func main() {
	dPtr := flag.Bool("d", false, "DataStore mode.")
	sPtr := flag.Bool("s", false, "Spam mode.")
	nPtr := flag.Int("n", 100, "Number workers.")
	bPtr := flag.Int("b", 100, "Base privk offset for workers. This should not overlap such that another test group is in same range.")
	mPtr := flag.String("m", "", "Data to write to data store.")
	iPtr := flag.String("i", "", "Index of data to write to data store.")
	flag.Parse()
	datastoreMode := *dPtr
	spamMode := *sPtr
	privk := "6aea45ee1273170fb525da34015e4f20ba39fe792f486ba74020bcacc9badfc1"
	nodeList := []string{"127.0.0.1:8887", "127.0.0.1:8888"}
	if spamMode {
		base := *bPtr
		num := 63
		for {
			f := &funder{}
			if err := f.doSpam(privk, num, nodeList, base); err != nil {
				panic(err)
			}
			base = base + num + 1
		}
	}
	f := &funder{}
	ctx := context.Background()
	if datastoreMode {
		fmt.Println("Running in DataStore Mode")
		if err := f.setupDataStoreMode(privk, nodeList); err != nil {
			panic(err)
		}
		tx, err := f.setupDataStoreTransaction(ctx, f.signer, f.acct, *mPtr, *iPtr)
		if err != nil {
			panic(err)
		}
		f.blockingSendTx(ctx, f.client, tx)
		time.Sleep(5 * time.Second)
		tx, err = f.setupDataStoreTransaction2(ctx, f.signer, f.acct, strings.Join([]string{*mPtr, "two"}, "-"), *iPtr)
		if err != nil {
			panic(err)
		}
		f.blockingSendTx(ctx, f.client, tx)
		time.Sleep(5 * time.Second)
		tx, err = f.setupDataStoreTransaction3(ctx, f.signer, f.acct, strings.Join([]string{*mPtr, "three"}, "-"), *iPtr)
		if err != nil {
			panic(err)
		}
		f.blockingSendTx(ctx, f.client, tx)
	} else {
		numChildren := *nPtr
		baseIdx := *bPtr
		if err := f.init(privk, numChildren, nodeList, baseIdx); err != nil {
			panic(err)
		}
	}
}
