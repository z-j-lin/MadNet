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

// DisputeMissingRegistrationTask contains required state for accusing missing registrations
type DisputeMissingRegistrationTask struct {
	*DkgTask
}

// asserting that DisputeMissingRegistrationTask struct implements interface interfaces.Task
var _ interfaces.Task = &DisputeMissingRegistrationTask{}

// asserting that DisputeMissingRegistrationTask struct implements DkgTaskIfase
var _ DkgTaskIfase = &DisputeMissingRegistrationTask{}

// NewDisputeMissingRegistrationTask creates a background task to accuse missing registrations during ETHDKG
func NewDisputeMissingRegistrationTask(state *objects.DkgState, start uint64, end uint64) *DisputeMissingRegistrationTask {
	return &DisputeMissingRegistrationTask{
		DkgTask: NewDkgTask(state, start, end),
	}
}

// Initialize begins the setup phase for Dispute Registration.
func (t *DisputeMissingRegistrationTask) Initialize(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum, state interface{}) error {

	logger.Info("DisputeMissingRegistrationTask Initializing...")

	return nil
}

// DoWork is the first attempt at Disputing Missing Registrations with ethdkg
func (t *DisputeMissingRegistrationTask) DoWork(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

// DoRetry is all subsequent attempts at registering with ethdkg
func (t *DisputeMissingRegistrationTask) DoRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

func (t *DisputeMissingRegistrationTask) doTask(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("DisputeMissingRegistrationTask doTask()")

	accusableParticipants, err := t.getAccusableParticipants(ctx, eth, logger)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask doTask() error getting accusable participants: %v", err)
	}

	// accuse missing validators
	if len(accusableParticipants) > 0 {
		logger.Warnf("Accusing missing registrations: %v", accusableParticipants)

		txnOpts, err := eth.GetTransactionOpts(ctx, t.State.Account)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask doTask() error getting txnOpts: %v", err)
		}

		// If the TxReplOpts exists, meaning the Tx replacement timeout was reached,
		// we increase the Gas to have priority for the next blocks
		if t.TxReplOpts != nil && t.TxReplOpts.Nonce != nil {
			logger.Info("txnOpts Replaced")
			txnOpts.Nonce = t.TxReplOpts.Nonce
			txnOpts.GasFeeCap = t.TxReplOpts.GasFeeCap
			txnOpts.GasTipCap = t.TxReplOpts.GasTipCap
		}

		txn, err := eth.Contracts().Ethdkg().AccuseParticipantNotRegistered(txnOpts, accusableParticipants)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask doTask() error accusing missing key shares: %v", err)
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
		}).Info("missing registration dispute fees")

		// Waiting for receipt
		receipt, err := eth.Queue().QueueAndWait(ctx, txn)
		if err != nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask doTask() error waiting for receipt failed: %v", err)
		}
		if receipt == nil {
			return dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask doTask() error missing share dispute receipt")
		}

		// Check receipt to confirm we were successful
		if receipt.Status != uint64(1) {
			return dkg.LogReturnErrorf(logger, "missing registration dispute status (%v) indicates failure: %v", receipt.Status, receipt.Logs)
		}
	} else {
		logger.Info("No accusations for missing registrations")
	}

	t.Success = true
	return nil
}

// ShouldRetry checks if it makes sense to try again
// Predicates:
// -- we haven't passed the last block
// -- the registration open hasn't moved, i.e. ETHDKG has not restarted
// -- We have unregistered participants
func (t *DisputeMissingRegistrationTask) ShouldRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) bool {
	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("DisputeMissingRegistrationTask ShouldRetry()")

	generalRetry := GeneralTaskShouldRetry(ctx, logger, eth, t.Start, t.End)
	if !generalRetry {
		return false
	}

	if t.State.Phase != objects.RegistrationOpen {
		return false
	}

	accusableParticipants, err := t.getAccusableParticipants(ctx, eth, logger)
	if err != nil {
		logger.Errorf("DisputeMissingRegistrationTask ShouldRetry() error getting accusable participants: %v", err)
		return true
	}

	if len(accusableParticipants) > 0 {
		return true
	}

	return false
}

// DoDone just creates a log entry saying task is complete
func (t *DisputeMissingRegistrationTask) DoDone(logger *logrus.Entry) {
	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("Success", t.Success).Infof("DisputeMissingRegistrationTask done")
}

func (t *DisputeMissingRegistrationTask) GetDkgTask() *DkgTask {
	return t.DkgTask
}

func (t *DisputeMissingRegistrationTask) SetDkgTask(dkgTask *DkgTask) {
	t.DkgTask = dkgTask
}

func (t *DisputeMissingRegistrationTask) getAccusableParticipants(ctx context.Context, eth interfaces.Ethereum, logger *logrus.Entry) ([]common.Address, error) {
	var accusableParticipants []common.Address
	callOpts := eth.GetCallOpts(ctx, t.State.Account)

	validators, err := dkg.GetValidatorAddressesFromPool(callOpts, eth, logger)
	if err != nil {
		return nil, dkg.LogReturnErrorf(logger, "DisputeMissingRegistrationTask getAccusableParticipants() error getting validators: %v", err)
	}

	validatorsMap := make(map[common.Address]bool)
	for _, validator := range validators {
		validatorsMap[validator] = true
	}

	// find participants who did not register
	for _, addr := range t.State.ValidatorAddresses {

		participant, ok := t.State.Participants[addr]
		_, isValidator := validatorsMap[addr]

		if isValidator && (!ok ||
			participant.Nonce != t.State.Nonce ||
			participant.Phase != uint8(objects.RegistrationOpen) ||
			(participant.PublicKey[0].Cmp(big.NewInt(0)) == 0 &&
				participant.PublicKey[1].Cmp(big.NewInt(0)) == 0)) {

			// did not register
			accusableParticipants = append(accusableParticipants, addr)
		}
	}

	return accusableParticipants, nil
}
