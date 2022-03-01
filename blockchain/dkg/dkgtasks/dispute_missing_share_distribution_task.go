package dkgtasks

import (
	"context"
	"github.com/MadBase/MadNet/blockchain/dkg"
	"github.com/MadBase/MadNet/blockchain/interfaces"
	"github.com/MadBase/MadNet/blockchain/objects"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math/big"
)

// DisputeMissingShareDistributionTask stores the data required to dispute shares
type DisputeMissingShareDistributionTask struct {
	*DkgTask
}

// asserting that DisputeMissingShareDistributionTask struct implements interface interfaces.Task
var _ interfaces.Task = &DisputeMissingShareDistributionTask{}

// asserting that DisputeMissingShareDistributionTask struct implements DkgTaskIfase
var _ DkgTaskIfase = &DisputeMissingShareDistributionTask{}

// NewDisputeMissingShareDistributionTask creates a new task
func NewDisputeMissingShareDistributionTask(state *objects.DkgState, start uint64, end uint64) *DisputeMissingShareDistributionTask {
	return &DisputeMissingShareDistributionTask{
		DkgTask: NewDkgTask(state, start, end),
	}
}

// Initialize begins the setup phase for DisputeMissingShareDistributionTask.
func (t *DisputeMissingShareDistributionTask) Initialize(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum, state interface{}) error {

	logger.Info("DisputeMissingShareDistributionTask Initializing...")

	return nil
}

// DoWork is the first attempt at disputing distributed shares
func (t *DisputeMissingShareDistributionTask) DoWork(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

// DoRetry is subsequent attempts at disputing distributed shares
func (t *DisputeMissingShareDistributionTask) DoRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

func (t *DisputeMissingShareDistributionTask) doTask(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("DisputeMissingShareDistributionTask doTask()")

	accusableParticipants, err := t.getAccusableParticipants(ctx, eth, logger)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask doTask() error getting accusableParticipants: %v", err)
	}

	// accuse missing validators
	if len(accusableParticipants) > 0 {
		logger.Warnf("Accusing missing distributed shares: %v", accusableParticipants)

		txnOpts, err := eth.GetTransactionOpts(ctx, t.State.Account)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask doTask() error getting txnOpts: %v", err)
		}

		// If the TxReplOpts exists, meaning the Tx replacement timeout was reached,
		// we increase the Gas to have priority for the next blocks
		if t.TxReplOpts != nil && t.TxReplOpts.Nonce != nil {
			logger.Info("txnOpts Replaced")
			txnOpts.Nonce = t.TxReplOpts.Nonce
			txnOpts.GasFeeCap = t.TxReplOpts.GasFeeCap
			txnOpts.GasTipCap = t.TxReplOpts.GasTipCap
		}

		txn, err := eth.Contracts().Ethdkg().AccuseParticipantDidNotDistributeShares(txnOpts, accusableParticipants)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask doTask() error accusing missing key shares: %v", err)
		}
		t.TxReplOpts.TxHash = txn.Hash()
		t.TxReplOpts.GasFeeCap = txn.GasFeeCap()
		t.TxReplOpts.GasTipCap = txn.GasTipCap()
		t.TxReplOpts.Nonce = big.NewInt(int64(txn.Nonce()))

		logger.WithFields(logrus.Fields{
			"GasFeeCap": t.TxReplOpts.GasFeeCap,
			"GasTipCap": t.TxReplOpts.GasTipCap,
			"Nonce":     t.TxReplOpts.Nonce,
			"Hash":      t.TxReplOpts.TxHash.Hex(),
		}).Info("missing share dispute fees")

		// Waiting for receipt
		receipt, err := eth.Queue().QueueAndWait(ctx, txn)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask doTask() error waiting for receipt failed: %v", err)
		}
		if receipt == nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask doTask() error missing share dispute receipt")
		}

		// Check receipt to confirm we were successful
		if receipt.Status != uint64(1) {
			return dkg.LogReturnErrorf(logger, "missing share distribution dispute status (%v) indicates failure: %v", receipt.Status, receipt.Logs)
		}
	} else {
		logger.Info("No accusations for missing distributed shares")
	}

	t.Success = true
	return nil
}

// ShouldRetry checks if it makes sense to try again
// if the DKG process is in the right phase and blocks
// range and there still someone to accuse, the retry
// is executed
func (t *DisputeMissingShareDistributionTask) ShouldRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) bool {

	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("DisputeMissingShareDistributionTask ShouldRetry()")

	generalRetry := GeneralTaskShouldRetry(ctx, logger, eth, t.Start, t.End)
	if !generalRetry {
		return false
	}

	if t.State.Phase != objects.ShareDistribution {
		return false
	}

	accusableParticipants, err := t.getAccusableParticipants(ctx, eth, logger)
	if err != nil {
		logger.Errorf("DisputeMissingShareDistributionTask ShouldRetry() error getting accusable participants: %v", err)
		return true
	}

	if len(accusableParticipants) > 0 {
		return true
	}

	return false
}

// DoDone creates a log entry saying task is complete
func (t *DisputeMissingShareDistributionTask) DoDone(logger *logrus.Entry) {
	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("Success", t.Success).Info("DisputeMissingShareDistributionTask done")
}

func (t *DisputeMissingShareDistributionTask) GetDkgTask() *DkgTask {
	return t.DkgTask
}

func (t *DisputeMissingShareDistributionTask) SetDkgTask(dkgTask *DkgTask) {
	t.DkgTask = dkgTask
}

func (t *DisputeMissingShareDistributionTask) getAccusableParticipants(ctx context.Context, eth interfaces.Ethereum, logger *logrus.Entry) ([]common.Address, error) {
	var accusableParticipants []common.Address
	callOpts := eth.GetCallOpts(ctx, t.State.Account)

	validators, err := dkg.GetValidatorAddressesFromPool(callOpts, eth, logger)
	if err != nil {
		return nil, dkg.LogReturnErrorf(logger, "DisputeMissingShareDistributionTask getAccusableParticipants() error getting validators: %v", err)
	}

	validatorsMap := make(map[common.Address]bool)
	for _, validator := range validators {
		validatorsMap[validator] = true
	}

	// find participants who did not submit their shares
	var emptySharesHash [32]byte
	for _, p := range t.State.Participants {
		_, isValidator := validatorsMap[p.Address]
		if isValidator && (p.Nonce != t.State.Nonce ||
			p.Phase != uint8(objects.ShareDistribution) ||
			p.DistributedSharesHash == emptySharesHash) {
			// did not distribute shares
			accusableParticipants = append(accusableParticipants, p.Address)
		}
	}

	return accusableParticipants, nil
}
