package objs

import (
	"testing"

	"github.com/MadBase/MadNet/crypto"
)

func TestPreCommit(t *testing.T) {
	bnVal := &crypto.BNGroupValidator{}
	secpVal := &crypto.Secp256k1Validator{}
	bclaimsList, txHashListList, err := generateChain(2)
	if err != nil {
		t.Fatal(err)
	}
	bclaims := bclaimsList[0]
	bhsh, err := bclaims.BlockHash()
	if err != nil {
		t.Fatal(err)
	}
	gk := crypto.BNGroupSigner{}
	err = gk.SetPrivk(crypto.Hasher([]byte("secret")))
	if err != nil {
		t.Fatal(err)
	}
	sig, err := gk.Sign(bhsh)
	if err != nil {
		t.Fatal(err)
	}
	secpSigner := &crypto.Secp256k1Signer{}
	err = secpSigner.SetPrivk(crypto.Hasher([]byte("secret")))
	if err != nil {
		t.Fatal(err)
	}
	secpSigner2 := &crypto.Secp256k1Signer{}
	err = secpSigner2.SetPrivk(crypto.Hasher([]byte("secret2")))
	if err != nil {
		t.Fatal(err)
	}
	bh := &BlockHeader{
		BClaims:  bclaims,
		SigGroup: sig,
		TxHshLst: txHashListList[0],
	}
	err = bh.ValidateSignatures(bnVal)
	if err != nil {
		t.Fatal(err)
	}
	rcert, err := bh.GetRCert()
	if err != nil {
		t.Fatal(err)
	}
	err = rcert.ValidateSignature(bnVal)
	if err != nil {
		t.Fatal(err)
	}
	pclms := &PClaims{
		BClaims: bclaimsList[1],
		RCert:   rcert,
	}
	pclmsBytes, err := pclms.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	sig, err = secpSigner.Sign(pclmsBytes)
	if err != nil {
		t.Fatal(err)
	}
	prop := &Proposal{
		PClaims:   pclms,
		Signature: sig,
		TxHshLst:  txHashListList[1],
	}
	pv1, err := prop.PreVote(secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	pv2, err := prop.PreVote(secpSigner2)
	if err != nil {
		t.Fatal(err)
	}
	pvl := PreVoteList{pv1, pv2}
	pc, err := pvl.MakePreCommit(secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = pc.ValidateSignatures(secpVal, bnVal)
	if err != nil {
		t.Fatal(err)
	}
	pcbytes, err := pc.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	pc2 := &PreCommit{}
	err = pc2.UnmarshalBinary(pcbytes)
	if err != nil {
		t.Fatal(err)
	}
	bclaimsEqual(t, pc.Proposal.PClaims.BClaims, pc2.Proposal.PClaims.BClaims)
	err = pc2.ValidateSignatures(secpVal, bnVal)
	if err != nil {
		t.Fatal(err)
	}

	preVotes, err := pc2.MakeImplPreVotes()
	if err != nil {
		t.Fatal(err)
	}

	if len(preVotes) != 2 {
		t.Fatal("invalid preVotes length")
	}
}
