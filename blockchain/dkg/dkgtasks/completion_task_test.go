package dkgtasks_test

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MadBase/MadNet/blockchain/dkg/dkgevents"
	"github.com/MadBase/MadNet/blockchain/dkg/dkgtasks"
	"github.com/MadBase/MadNet/blockchain/dkg/dtest"
	"github.com/MadBase/MadNet/blockchain/objects"
	"github.com/MadBase/MadNet/consensus/objs"
	"github.com/MadBase/MadNet/crypto"
	bn256A "github.com/MadBase/MadNet/crypto/bn256"
	bn256 "github.com/MadBase/MadNet/crypto/bn256/cloudflare"
	"github.com/MadBase/MadNet/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// We begin by submitting invalid information.
// This test is meant to raise an error resulting from an invalid argument
// for the Ethereum interface.
func TestCompletionBad1(t *testing.T) {
	n := 4
	_, ecdsaPrivateKeys := dtest.InitializeNewDetDkgStateInfo(n)
	logger := logging.GetLogger("ethereum")
	logger.SetLevel(logrus.DebugLevel)
	eth := connectSimulatorEndpoint(t, ecdsaPrivateKeys, 333*time.Millisecond)
	defer eth.Close()

	acct := eth.GetKnownAccounts()[0]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a task to share distribution and make sure it succeeds
	state := objects.NewDkgState(acct)
	task := dkgtasks.NewCompletionTask(state)
	log := logger.WithField("TaskID", "foo")

	defer func() {
		// If we didn't get here by recovering from a panic() we failed
		if reason := recover(); reason == nil {
			t.Log("No panic in sight")
			t.Fatal("Should have panicked")
		} else {
			t.Logf("Good panic because: %v", reason)
		}
	}()
	task.Initialize(ctx, log, eth, nil)
}

// We test to ensure that everything behaves correctly.
func TestCompletionBad2(t *testing.T) {
	n := 4
	_, ecdsaPrivateKeys := dtest.InitializeNewDetDkgStateInfo(n)
	logger := logging.GetLogger("ethereum")
	logger.SetLevel(logrus.DebugLevel)
	eth := connectSimulatorEndpoint(t, ecdsaPrivateKeys, 333*time.Millisecond)
	defer eth.Close()

	acct := eth.GetKnownAccounts()[0]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Do bad Completion task
	state := objects.NewDkgState(acct)
	log := logger.WithField("TaskID", "foo")
	task := dkgtasks.NewCompletionTask(state)
	err := task.Initialize(ctx, log, eth, state)
	if err == nil {
		t.Fatal("Should have raised error")
	}
}

func makeSigners(t *testing.T) ([]byte, []*crypto.BNGroupSigner, [][]byte, []*crypto.Secp256k1Signer, [][]byte) {
	s := new(crypto.BNGroupSigner)
	msg := []byte("A message to sign")

	secret1 := big.NewInt(100)
	secret2 := big.NewInt(101)
	secret3 := big.NewInt(102)
	secret4 := big.NewInt(103)

	msk := big.NewInt(0)
	msk.Add(msk, secret1)
	msk.Add(msk, secret2)
	msk.Add(msk, secret3)
	msk.Add(msk, secret4)
	msk.Mod(msk, bn256.Order)
	mpk := new(bn256.G2).ScalarBaseMult(msk)

	big1 := big.NewInt(1)
	big2 := big.NewInt(2)

	privCoefs1 := []*big.Int{secret1, big1, big2}
	privCoefs2 := []*big.Int{secret2, big1, big2}
	privCoefs3 := []*big.Int{secret3, big1, big2}
	privCoefs4 := []*big.Int{secret4, big1, big2}

	share1to1 := bn256.PrivatePolyEval(privCoefs1, 1)
	share1to2 := bn256.PrivatePolyEval(privCoefs1, 2)
	share1to3 := bn256.PrivatePolyEval(privCoefs1, 3)
	share1to4 := bn256.PrivatePolyEval(privCoefs1, 4)
	share2to1 := bn256.PrivatePolyEval(privCoefs2, 1)
	share2to2 := bn256.PrivatePolyEval(privCoefs2, 2)
	share2to3 := bn256.PrivatePolyEval(privCoefs2, 3)
	share2to4 := bn256.PrivatePolyEval(privCoefs2, 4)
	share3to1 := bn256.PrivatePolyEval(privCoefs3, 1)
	share3to2 := bn256.PrivatePolyEval(privCoefs3, 2)
	share3to3 := bn256.PrivatePolyEval(privCoefs3, 3)
	share3to4 := bn256.PrivatePolyEval(privCoefs3, 4)
	share4to1 := bn256.PrivatePolyEval(privCoefs4, 1)
	share4to2 := bn256.PrivatePolyEval(privCoefs4, 2)
	share4to3 := bn256.PrivatePolyEval(privCoefs4, 3)
	share4to4 := bn256.PrivatePolyEval(privCoefs4, 4)

	groupShares := make([][]byte, 4)
	for k := 0; k < len(groupShares); k++ {
		groupShares[k] = make([]byte, len(mpk.Marshal()))
	}

	listOfSS1 := []*big.Int{share1to1, share2to1, share3to1, share4to1}
	gsk1 := bn256.GenerateGroupSecretKeyPortion(listOfSS1)
	gpk1 := new(bn256.G2).ScalarBaseMult(gsk1)
	groupShares[0] = gpk1.Marshal()
	s1 := new(crypto.BNGroupSigner)
	err := s1.SetPrivk(gsk1.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	sig1, err := s1.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	listOfSS2 := []*big.Int{share1to2, share2to2, share3to2, share4to2}
	gsk2 := bn256.GenerateGroupSecretKeyPortion(listOfSS2)
	gpk2 := new(bn256.G2).ScalarBaseMult(gsk2)
	groupShares[1] = gpk2.Marshal()
	s2 := new(crypto.BNGroupSigner)
	err = s2.SetPrivk(gsk2.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	sig2, err := s2.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	listOfSS3 := []*big.Int{share1to3, share2to3, share3to3, share4to3}
	gsk3 := bn256.GenerateGroupSecretKeyPortion(listOfSS3)
	gpk3 := new(bn256.G2).ScalarBaseMult(gsk3)
	groupShares[2] = gpk3.Marshal()
	s3 := new(crypto.BNGroupSigner)
	err = s3.SetPrivk(gsk3.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	sig3, err := s3.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	listOfSS4 := []*big.Int{share1to4, share2to4, share3to4, share4to4}
	gsk4 := bn256.GenerateGroupSecretKeyPortion(listOfSS4)
	gpk4 := new(bn256.G2).ScalarBaseMult(gsk4)
	groupShares[3] = gpk4.Marshal()
	s4 := new(crypto.BNGroupSigner)
	err = s4.SetPrivk(gsk4.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	sig4, err := s4.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}

	sigs := make([][]byte, 4)
	for k := 0; k < len(sigs); k++ {
		sigs[k] = make([]byte, 192)
	}
	sigs[0] = sig1
	sigs[1] = sig2
	sigs[2] = sig3
	sigs[3] = sig4

	err = s.SetGroupPubk(mpk.Marshal())
	if err != nil {
		t.Fatal(err)
	}

	// Make bad sigs array
	sigsBad := make([][]byte, 2)
	for k := 0; k < len(sigsBad); k++ {
		sigsBad[k] = make([]byte, 192)
	}
	sigsBad[0] = sig1
	sigsBad[1] = sig2
	_, err = s.Aggregate(sigsBad, groupShares)
	if err == nil {
		t.Fatal("Should have raised an error for too few signatures!")
	}

	// Finally submit signature
	grpsig, err := s.Aggregate(sigs, groupShares)
	if err != nil {
		t.Fatal(err)
	}

	bnVal := &crypto.BNGroupValidator{}
	groupk, err := bnVal.PubkeyFromSig(grpsig)
	if err != nil {
		t.Fatal(err)
	}

	bnSigners := []*crypto.BNGroupSigner{}
	bnSigners = append(bnSigners, s1)
	bnSigners = append(bnSigners, s2)
	bnSigners = append(bnSigners, s3)
	bnSigners = append(bnSigners, s4)

	secpSigners := []*crypto.Secp256k1Signer{}
	secpPubks := [][]byte{}
	for _, share := range groupShares {
		signer, pubk := makeSecpSigner(share)
		secpPubks = append(secpPubks, pubk)
		secpSigners = append(secpSigners, signer)
	}

	for _, signer := range bnSigners {
		err := signer.SetGroupPubk(groupk)
		if err != nil {
			t.Fatal(err)
		}
	}
	mpkBin := mpk.Marshal()
	fmt.Printf("MPK: %x\n", mpkBin)
	fmt.Printf("GroupK: %x\n", groupk)
	if bytes.Equal(mpkBin, groupk) {
		t.Fatal("Error")
	}

	return groupk, bnSigners, groupShares, secpSigners, secpPubks
}

func makeSecpSigner(seed []byte) (*crypto.Secp256k1Signer, []byte) {
	secpSigner := &crypto.Secp256k1Signer{}
	err := secpSigner.SetPrivk(crypto.Hasher(seed))
	if err != nil {
		panic(err)
	}
	secpKey, _ := secpSigner.Pubkey()
	return secpSigner, secpKey
}

// We complete everything correctly, but we do not complete in time
func TestCompletionBad3(t *testing.T) {
	n := 4
	dkgStates, ecdsaPrivateKeys := dtest.InitializeNewDetDkgStateInfo(n)

	eth := connectSimulatorEndpoint(t, ecdsaPrivateKeys, 333*time.Millisecond)
	assert.NotNil(t, eth)
	defer eth.Close()

	ctx := context.Background()

	accounts := eth.GetKnownAccounts()

	owner := accounts[0]
	err := eth.UnlockAccount(owner)
	assert.Nil(t, err)

	// Start EthDKG
	ownerOpts, err := eth.GetTransactionOpts(ctx, owner)
	assert.Nil(t, err)

	// Check Stake
	callOpts := eth.GetCallOpts(ctx, owner)
	for _, acct := range accounts {
		stakeBalance, err := eth.Contracts().Staking().BalanceStakeFor(callOpts, acct.Address)
		assert.Nil(t, err)

		t.Logf("stakeBalance:%v", stakeBalance.String())
	}

	// Shorten ethdkg phase for testing purposes
	txn, err := eth.Contracts().Ethdkg().UpdatePhaseLength(ownerOpts, big.NewInt(100))
	assert.Nil(t, err)
	rcpt, err := eth.Queue().QueueAndWait(ctx, txn)
	assert.Nil(t, err)

	txn, err = eth.Contracts().Ethdkg().InitializeState(ownerOpts)
	assert.Nil(t, err)

	eth.Commit()
	rcpt, err = eth.Queue().QueueAndWait(ctx, txn)
	assert.Nil(t, err)

	for _, log := range rcpt.Logs {
		if log.Topics[0].String() == "0x9c6f8368fe7e77e8cb9438744581403bcb3f53298e517f04c1b8475487402e97" {
			event, err := eth.Contracts().Ethdkg().ParseRegistrationOpen(*log)
			assert.Nil(t, err)

			for _, dkgState := range dkgStates {
				dkgevents.PopulateSchedule(dkgState, event)
			}
		}
	}

	// Do Register task
	registerTasks := make([]*dkgtasks.RegisterTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		registerTasks[idx] = dkgtasks.NewRegisterTask(state)
		err = registerTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = registerTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, registerTasks[idx].Success)

		// Set Registration success to true
		dkgStates[idx].Registration = true
	}

	// Advance to share distribution phase
	advanceTo(t, eth, dkgStates[0].ShareDistributionStart)

	// Check public keys are present and valid
	for idx, acct := range accounts {
		callOpts := eth.GetCallOpts(context.Background(), acct)
		k0, err := eth.Contracts().Ethdkg().PublicKeys(callOpts, acct.Address, common.Big0)
		assert.Nil(t, err)
		k1, err := eth.Contracts().Ethdkg().PublicKeys(callOpts, acct.Address, common.Big1)
		assert.Nil(t, err)

		// check points
		publicKey := dkgStates[idx].TransportPublicKey
		if (publicKey[0].Cmp(k0) != 0) || (publicKey[1].Cmp(k1) != 0) {
			t.Fatal("Invalid public key")
		}
	}

	// Do Share Distribution task
	shareDistributionTasks := make([]*dkgtasks.ShareDistributionTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		shareDistributionTasks[idx] = dkgtasks.NewShareDistributionTask(state)
		err = shareDistributionTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = shareDistributionTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, shareDistributionTasks[idx].Success)

		// Set ShareDistribution success to true
		dkgStates[idx].ShareDistribution = true
	}
	// Ensure all participants have valid share information
	dtest.PopulateEncryptedSharesAndCommitments(dkgStates)

	// Advance to share dispute phase
	advanceTo(t, eth, dkgStates[0].DisputeStart)

	// Do Dispute task
	disputeTasks := make([]*dkgtasks.DisputeTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		disputeTasks[idx] = dkgtasks.NewDisputeTask(state)
		err = disputeTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = disputeTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, disputeTasks[idx].Success)

		// Set Dispute success to true
		dkgStates[idx].Dispute = true
	}

	// Confirm number of BadShares is zero
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		if len(state.BadShares) != 0 {
			t.Fatalf("Idx %v has incorrect number of BadShares", idx)
		}
	}

	// Advance to key share distribution phase
	advanceTo(t, eth, dkgStates[0].KeyShareSubmissionStart)
	dtest.GenerateKeyShares(dkgStates)

	// Do Key Share Submission task
	keyShareSubmissionTasks := make([]*dkgtasks.KeyshareSubmissionTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		keyShareSubmissionTasks[idx] = dkgtasks.NewKeyshareSubmissionTask(state)
		err = keyShareSubmissionTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = keyShareSubmissionTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, keyShareSubmissionTasks[idx].Success)

		// Set KeyShareSubmission success to true
		dkgStates[idx].KeyShareSubmission = true
	}

	// Double check to Make sure all transactions were good
	rcpts, err := eth.Queue().WaitGroupTransactions(ctx, 1)
	assert.Nil(t, err)

	for _, rcpt := range rcpts {
		assert.NotNil(t, rcpt)
		assert.Equal(t, uint64(1), rcpt.Status)
	}

	eth.Commit()

	// Check key shares are present and valid
	for idx, acct := range accounts {
		callOpts := eth.GetCallOpts(context.Background(), acct)
		k0, err := eth.Contracts().Ethdkg().KeyShares(callOpts, acct.Address, common.Big0)
		assert.Nil(t, err)
		k1, err := eth.Contracts().Ethdkg().KeyShares(callOpts, acct.Address, common.Big1)
		assert.Nil(t, err)

		// check points
		keyShareG1 := dkgStates[idx].KeyShareG1s[acct.Address]
		if (keyShareG1[0].Cmp(k0) != 0) || (keyShareG1[1].Cmp(k1) != 0) {
			t.Fatal("Invalid key share")
		}
	}

	// Advance to mpk submission phase
	advanceTo(t, eth, dkgStates[0].MPKSubmissionStart)
	dtest.GenerateMasterPublicKey(dkgStates)

	// Do MPK Submission task
	mpkSubmissionTasks := make([]*dkgtasks.MPKSubmissionTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		mpkSubmissionTasks[idx] = dkgtasks.NewMPKSubmissionTask(state)
		err = mpkSubmissionTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = mpkSubmissionTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, mpkSubmissionTasks[idx].Success)

		// Set MPK submission success to true
		dkgStates[idx].MPKSubmission = true
	}

	// Double check to Make sure all transactions were good
	rcpts, err = eth.Queue().WaitGroupTransactions(ctx, 1)
	assert.Nil(t, err)

	for _, rcpt := range rcpts {
		assert.NotNil(t, rcpt)
		assert.Equal(t, uint64(1), rcpt.Status)
	}

	eth.Commit()

	// Validate MPK
	for idx, acct := range accounts {
		callOpts := eth.GetCallOpts(context.Background(), acct)
		k0, err := eth.Contracts().Ethdkg().MasterPublicKey(callOpts, common.Big0)
		assert.Nil(t, err)
		k1, err := eth.Contracts().Ethdkg().MasterPublicKey(callOpts, common.Big1)
		assert.Nil(t, err)
		k2, err := eth.Contracts().Ethdkg().MasterPublicKey(callOpts, common.Big2)
		assert.Nil(t, err)
		k3, err := eth.Contracts().Ethdkg().MasterPublicKey(callOpts, common.Big3)
		assert.Nil(t, err)

		// check mpk
		mpk := dkgStates[idx].MasterPublicKey
		if (mpk[0].Cmp(k0) != 0) || (mpk[1].Cmp(k1) != 0) || (mpk[2].Cmp(k2) != 0) || (mpk[3].Cmp(k3) != 0) {
			t.Fatal("Invalid master public key")
		}
	}

	// Advance to gpkj submission phase
	advanceTo(t, eth, dkgStates[0].GPKJSubmissionStart)
	dtest.GenerateGPKJ(dkgStates)

	// Do MPK Submission task
	gpkjSubmitTasks := make([]*dkgtasks.GPKSubmissionTask, n)
	myStr := `
		import { ValidatorRawData } from "../../ethdkg/setup";
		export const validatorsSnapshots: ValidatorRawData[] = [
	`
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		adminHandler := new(adminHandlerMock)
		gpkjSubmitTasks[idx] = dkgtasks.NewGPKSubmissionTask(state, adminHandler)
		err = gpkjSubmitTasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = gpkjSubmitTasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, gpkjSubmitTasks[idx].Success)

		// Set GpkjSubmission success to true
		dkgStates[idx].GPKJSubmission = true
		encryptedShares := `[`
		for _, share := range state.EncryptedShares[accounts[idx].Address] {
			encryptedShares += fmt.Sprintf(`"0x%x",`, share)
		}
		encryptedShares += "]"
		commitments := `[`
		for _, commit := range state.Commitments[accounts[idx].Address] {
			commitments += fmt.Sprintf(`["0x%x","0x%x"],`, commit[0], commit[1])
		}
		commitments += "]"
		myStr += fmt.Sprintf(`
			{
			privateKey: "0x%x",
			address: "0x%x",
			madNetPublicKey:["0x%x", "0x%x"],
			encryptedShares:%s,
			commitments:%s,
			keyShareG1: ["0x%x", "0x%x"],
			keyShareG1CorrectnessProof: ["0x%x", "0x%x"],
			keyShareG2: ["0x%x", "0x%x", "0x%x", "0x%x"],
			mpk: ["0x%x", "0x%x", "0x%x", "0x%x"],
			gpkj: ["0x%x", "0x%x", "0x%x", "0x%x"],
			},
		`,
			ecdsaPrivateKeys[idx].D,
			state.Account.Address,
			state.TransportPublicKey[0], state.TransportPublicKey[1],
			encryptedShares,
			commitments,
			state.KeyShareG1s[accounts[idx].Address][0], state.KeyShareG1s[accounts[idx].Address][1],
			state.KeyShareG1CorrectnessProofs[accounts[idx].Address][0], state.KeyShareG1CorrectnessProofs[accounts[idx].Address][1],
			state.KeyShareG2s[accounts[idx].Address][0], state.KeyShareG2s[accounts[idx].Address][1], state.KeyShareG2s[accounts[idx].Address][2], state.KeyShareG2s[accounts[idx].Address][3],
			state.MasterPublicKey[0], state.MasterPublicKey[1], state.MasterPublicKey[2], state.MasterPublicKey[3],
			state.GroupPublicKey[0], state.GroupPublicKey[1], state.GroupPublicKey[2], state.GroupPublicKey[3],
		)
	}
	myStr += "];\n"
	fmt.Print(myStr)

	// Double check to Make sure all transactions were good
	rcpts, err = eth.Queue().WaitGroupTransactions(ctx, 1)
	assert.Nil(t, err)

	for _, rcpt := range rcpts {
		assert.NotNil(t, rcpt)
		assert.Equal(t, uint64(1), rcpt.Status)
	}

	eth.Commit()

	// Check gpkjs and signatures are present and valid
	for idx, acct := range accounts {
		callOpts := eth.GetCallOpts(context.Background(), acct)
		s0, err := eth.Contracts().Ethdkg().InitialSignatures(callOpts, acct.Address, common.Big0)
		assert.Nil(t, err)
		s1, err := eth.Contracts().Ethdkg().InitialSignatures(callOpts, acct.Address, common.Big1)
		assert.Nil(t, err)

		// check signature
		signature := dkgStates[idx].GroupSignature
		if (signature[0].Cmp(s0) != 0) || (signature[1].Cmp(s1) != 0) {
			t.Fatal("Invalid signature")
		}

		k0, err := eth.Contracts().Ethdkg().GpkjSubmissions(callOpts, acct.Address, common.Big0)
		assert.Nil(t, err)
		k1, err := eth.Contracts().Ethdkg().GpkjSubmissions(callOpts, acct.Address, common.Big1)
		assert.Nil(t, err)
		k2, err := eth.Contracts().Ethdkg().GpkjSubmissions(callOpts, acct.Address, common.Big2)
		assert.Nil(t, err)
		k3, err := eth.Contracts().Ethdkg().GpkjSubmissions(callOpts, acct.Address, common.Big3)
		assert.Nil(t, err)

		// check gpkj
		gpkj := dkgStates[idx].GroupPublicKey
		if (gpkj[0].Cmp(k0) != 0) || (gpkj[1].Cmp(k1) != 0) || (gpkj[2].Cmp(k2) != 0) || (gpkj[3].Cmp(k3) != 0) {
			t.Fatal("Invalid gpkj")
		}
	}

	// Advance to gpkj dispute phase
	advanceTo(t, eth, dkgStates[0].GPKJGroupAccusationStart)

	// Do GPKjDispute task
	tasks := make([]*dkgtasks.GPKJDisputeTask, n)
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		logger := logging.GetLogger("test").WithField("Validator", accounts[idx].Address.String())

		tasks[idx] = dkgtasks.NewGPKJDisputeTask(state)
		logger.Errorf("Idx: %v\n", idx)
		err = tasks[idx].Initialize(ctx, logger, eth, state)
		assert.Nil(t, err)
		err = tasks[idx].DoWork(ctx, logger, eth)
		assert.Nil(t, err)

		eth.Commit()
		assert.True(t, tasks[idx].Success)

		// Set GPKjDispute success to true
		dkgStates[idx].GPKJGroupAccusation = true
	}

	// Double check to Make sure all transactions were good
	rcpts, err = eth.Queue().WaitGroupTransactions(ctx, 1)
	assert.Nil(t, err)

	for _, rcpt := range rcpts {
		assert.NotNil(t, rcpt)
		assert.Equal(t, uint64(1), rcpt.Status)
	}

	// Advance to Completion phase
	advanceTo(t, eth, dkgStates[0].CompleteStart)

	// // Advance to end of Completion phase
	// advanceTo(t, eth, dkgStates[0].CompleteEnd)
	// eth.Commit()

	// Do bad Completion task; this should fail because we are past
	state := dkgStates[0]
	logger := logging.GetLogger("test").WithField("Validator", accounts[0].Address.String())
	task := dkgtasks.NewCompletionTask(state)
	err = task.Initialize(ctx, logger, eth, state)
	if err != nil {
		t.Fatal(err)
	}
	err = task.DoWork(ctx, logger, eth)

	bnSigners := []*crypto.BNGroupSigner{}
	for idx := 0; idx < n; idx++ {
		state := dkgStates[idx]
		signer := &crypto.BNGroupSigner{}
		signer.SetPrivk(state.GroupPrivateKey.Bytes())
		bnSigners = append(bnSigners, signer)
		groupKey, err := bn256A.MarshalBigIntSlice(state.MasterPublicKey[:])
		if err != nil {
			t.Fatal(err)
		}
		err = signer.SetGroupPubk(groupKey)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Valid at 1024
	grpSig1024, bClaimsBin1024, err := GenerateSnapshotData(1, 1024, bnSigners, n, state.MasterPublicKey[:], false)
	if err != nil {
		t.Fatal(err)
	}

	//Valid block at 2048
	grpSig2048, bClaimsBin2048, err := GenerateSnapshotData(1, 2048, bnSigners, n, state.MasterPublicKey[:], false)
	if err != nil {
		t.Fatal(err)
	}

	// Incorrect height at 500
	grpSig500, bClaimsBin500, err := GenerateSnapshotData(1, 500, bnSigners, n, state.MasterPublicKey[:], false)
	if err != nil {
		t.Fatal(err)
	}

	// Incorrect chanid 2 at 1024
	grpSigChain2, bClaimsBinChain2, err := GenerateSnapshotData(2, 1024, bnSigners, n, state.MasterPublicKey[:], false)
	if err != nil {
		t.Fatal(err)
	}

	// Incorrect signature in a valid block at 1024
	grpSigIncorrect, bClaimsBinIncorrect, err := GenerateSnapshotData(2, 1024, bnSigners, n, state.MasterPublicKey[:], true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf(`
		export const validSnapshot1024: Snapshot = {
			GroupSignature: "0x%x",
			BClaims: "0x%x"
		};

		export const validSnapshot2048: Snapshot = {
			GroupSignature: "0x%x",
			BClaims: "0x%x"
		};

		export const invalidSnapshot500: Snapshot = {
			GroupSignature: "0x%x",
			BClaims: "0x%x"
		};

		export const invalidSnapshotChainID2: Snapshot = {
			GroupSignature: "0x%x",
			BClaims: "0x%x"
		};

		export const invalidSnapshotIncorrectSig: Snapshot = {
			GroupSignature: "0x%x",
			BClaims: "0x%x"
		};`,
		grpSig1024, bClaimsBin1024,
		grpSig2048, bClaimsBin2048,
		grpSig500, bClaimsBin500,
		grpSigChain2, bClaimsBinChain2,
		grpSigIncorrect, bClaimsBinIncorrect,
	)

}

func GenerateSnapshotData(chainID uint32, height uint32, bnSigners []*crypto.BNGroupSigner, n int, mpkI []*big.Int, fakeSig bool) ([]byte, []byte, error) {
	bclaims := &objs.BClaims{
		ChainID:    chainID,
		Height:     height,
		TxCount:    0,
		PrevBlock:  crypto.Hasher([]byte("")),
		TxRoot:     crypto.Hasher([]byte("")),
		StateRoot:  crypto.Hasher([]byte("")),
		HeaderRoot: crypto.Hasher([]byte("")),
	}

	blockHash, err := bclaims.BlockHash()
	if err != nil {
		return nil, nil, err
	}

	bClaimsBin, err := bclaims.MarshalBinary()
	if err != nil {
		return nil, nil, err
	}

	grpsig := []byte{}
	if fakeSig {
		grpsig, err = bnSigners[0].Sign(blockHash)
		if err != nil {
			return nil, nil, err
		}
	} else {
		grpsig, err = GenerateBlockSignature(bnSigners, n, blockHash, mpkI)
		if err != nil {
			return nil, nil, err
		}
	}

	bnVal := &crypto.BNGroupValidator{}
	_, err = bnVal.Validate(blockHash, grpsig)
	if err != nil {
		return nil, nil, err
	}

	return grpsig, bClaimsBin, nil
}

func GenerateBlockSignature(bnSigners []*crypto.BNGroupSigner, n int, blockHash []byte, mpkI []*big.Int) ([]byte, error) {
	sigs := [][]byte{}
	groupShares := [][]byte{}
	for idx := 0; idx < n; idx++ {
		sig, err := bnSigners[idx].Sign(blockHash)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Sig: %x\n", sig)
		sigs = append(sigs, sig)
		pkShare, err := bnSigners[idx].PubkeyShare()
		if err != nil {
			return nil, err
		}
		groupShares = append(groupShares, pkShare)
		fmt.Printf("Pkshare: %x\n", pkShare)
	}
	s := new(crypto.BNGroupSigner)
	mpk, err := bn256A.MarshalBigIntSlice(mpkI)
	err = s.SetGroupPubk(mpk)
	if err != nil {
		return nil, err
	}

	// Finally submit signature
	grpsig, err := s.Aggregate(sigs, groupShares)
	if err != nil {
		return nil, err
	}
	return grpsig, nil

}
