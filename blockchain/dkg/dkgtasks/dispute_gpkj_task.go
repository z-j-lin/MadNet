package dkgtasks

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MadBase/MadNet/blockchain/dkg"
	"github.com/MadBase/MadNet/blockchain/dkg/math"
	"github.com/MadBase/MadNet/blockchain/interfaces"
	"github.com/MadBase/MadNet/blockchain/objects"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/crypto/bn256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

// DisputeGPKjTask contains required state for performing a group accusation
type DisputeGPKjTask struct {
	Start   uint64
	End     uint64
	State   *objects.DkgState
	Success bool
}

// asserting that DisputeGPKjTask struct implements interface interfaces.Task
var _ interfaces.Task = &DisputeGPKjTask{}

// NewDisputeGPKjTask creates a background task that attempts perform a group accusation if necessary
func NewDisputeGPKjTask(state *objects.DkgState, start uint64, end uint64) *DisputeGPKjTask {
	return &DisputeGPKjTask{
		Start:   start,
		End:     end,
		State:   state,
		Success: false,
	}
}

// Initialize prepares for work to be done in the GPKjDispute phase.
// Here, we determine if anyone submitted an invalid gpkj.
func (t *DisputeGPKjTask) Initialize(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum, state interface{}) error {
	dkgState, validState := state.(*objects.DkgState)
	if !validState {
		panic(fmt.Errorf("%w invalid state type", objects.ErrCanNotContinue))
	}

	t.State = dkgState

	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("StateLocation", fmt.Sprintf("%p", t.State)).Info("GPKJDisputeTask Initialize()...")

	if t.State.Phase != objects.DisputeGPKJSubmission {
		return fmt.Errorf("%w because it's not in DisputeGPKJSubmission phase", objects.ErrCanNotContinue)
	}

	var (
		groupPublicKeys  [][4]*big.Int
		groupCommitments [][][2]*big.Int
	)

	// callOpts := eth.GetCallOpts(ctx, t.State.Account)
	var participantList = t.State.GetSortedParticipants()

	for _, participant := range participantList {
		// Retrieve values all group keys and signatures from contract
		// groupPublicKey, err := dkg.RetrieveGroupPublicKey(callOpts, eth, participant.Address)
		// if err != nil {
		// 	return dkg.LogReturnErrorf(logger, "Failed to retrieve group public key for %v", participant.Address.Hex())
		// }

		// Save the values
		// t.State.GroupPublicKeys[participant.Address] = groupPublicKey

		// Build array
		groupPublicKeys = append(groupPublicKeys, participant.GPKj)
		groupCommitments = append(groupCommitments, participant.Commitments)
	}

	//
	honest, dishonest, missing, err := math.CategorizeGroupSigners(groupPublicKeys, participantList, groupCommitments)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "Failed to determine honest vs dishonest validators: %v", err)
	}

	inverse, err := math.InverseArrayForUserCount(int(t.State.NumberOfValidators))
	if err != nil {
		return dkg.LogReturnErrorf(logger, "Failed to calculate inversion: %v", err)
	}

	logger.Debugf("   Honest indices: %v", honest.ExtractIndices())
	logger.Debugf("Dishonest indices: %v", dishonest.ExtractIndices())
	logger.Debugf("  Missing indices: %v", missing.ExtractIndices())

	t.State.DishonestValidators = dishonest
	t.State.HonestValidators = honest
	t.State.Inverse = inverse

	return nil
}

// DoWork is the first attempt at submitting an invalid gpkj accusation
func (t *DisputeGPKjTask) DoWork(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

// DoRetry is all subsequent attempts at submitting an invalid gpkj accusation
func (t *DisputeGPKjTask) DoRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

func (t *DisputeGPKjTask) doTask(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("GPKJDisputeTask doTask()")

	// Setup
	txnOpts, err := eth.GetTransactionOpts(ctx, t.State.Account)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "getting txn opts failed: %v", err)
	}

	// Perform group accusation
	logger.Infof("   Honest indices: %v", t.State.HonestValidators.ExtractIndices())
	logger.Infof("Dishonest indices: %v", t.State.DishonestValidators.ExtractIndices())

	var groupEncryptedSharesHash [][32]byte
	var groupCommitments [][][2]*big.Int
	var validatorAddresses []common.Address
	var participantList = t.State.GetSortedParticipants()

	for _, participant := range participantList {
		// Get group encrypted shares
		es := participant.EncryptedShares
		encryptedSharesBin, err := bn256.MarshalBigIntSlice(es)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "group accusation failed: %v", err)
		}
		hashSlice := crypto.Hasher(encryptedSharesBin)
		var hashSlice32 [32]byte
		copy(hashSlice32[:], hashSlice)
		groupEncryptedSharesHash = append(groupEncryptedSharesHash, hashSlice32)
		// Get group commitments
		com := participant.Commitments
		groupCommitments = append(groupCommitments, com)
		validatorAddresses = append(validatorAddresses, participant.Address)
	}

	// Loop through dishonest participants and perform accusation
	for _, dishonestParticipant := range t.State.DishonestValidators {
		txn, err := eth.Contracts().Ethdkg().AccuseParticipantSubmittedBadGPKJ(txnOpts, validatorAddresses, groupEncryptedSharesHash, groupCommitments, dishonestParticipant.Address)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "group accusation failed: %v", err)
		}

		// Waiting for receipt
		receipt, err := eth.Queue().QueueAndWait(ctx, txn)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "waiting for receipt failed: %v", err)
		}
		if receipt == nil {
			return dkg.LogReturnErrorf(logger, "missing receipt")
		}

		// Check receipt to confirm we were successful
		if receipt.Status != uint64(1) {
			return dkg.LogReturnErrorf(logger, "bad gpkj error (%v) indicates failure: %v", receipt.Status, receipt.Logs)
		}
	}

	t.Success = true

	return nil
}

// ShouldRetry checks if it makes sense to try again
// Predicates:
// -- we haven't passed the last block
// -- the registration open hasn't moved, i.e. ETHDKG has not restarted
func (t *DisputeGPKjTask) ShouldRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) bool {

	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("GPKJDisputeTask ShouldRetry()")

	currentBlock, err := eth.GetCurrentHeight(ctx)
	if err != nil {
		return true
	}
	//logger = logger.WithField("CurrentHeight", currentBlock)

	if t.State.Phase == objects.DisputeGPKJSubmission &&
		t.Start <= currentBlock &&
		currentBlock < t.End {
		return true
	}

	// This wraps the retry logic for every phase, _except_ registration
	// return GeneralTaskShouldRetry(ctx, t.State.Account, logger, eth,
	// 	t.State.TransportPublicKey, t.Start, t.End)
	return false
}

// DoDone creates a log entry saying task is complete
func (t *DisputeGPKjTask) DoDone(logger *logrus.Entry) {
	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("Success", t.Success).Infof("GPKJDisputeTask done")
}
