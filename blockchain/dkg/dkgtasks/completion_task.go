package dkgtasks

import (
	"context"
	"fmt"
	"github.com/MadBase/MadNet/blockchain/dkg"
	"github.com/MadBase/MadNet/blockchain/interfaces"
	"github.com/MadBase/MadNet/blockchain/objects"
	"github.com/sirupsen/logrus"
	"math/big"
)

// CompletionTask contains required state for safely performing a registration
type CompletionTask struct {
	*DkgTask
}

// asserting that CompletionTask struct implements interface interfaces.Task
var _ interfaces.Task = &CompletionTask{}

// asserting that CompletionTask struct implements DkgTaskIfase
var _ DkgTaskIfase = &CompletionTask{}

// NewCompletionTask creates a background task that attempts to call Complete on ethdkg
func NewCompletionTask(state *objects.DkgState, start uint64, end uint64) *CompletionTask {
	return &CompletionTask{
		DkgTask: &DkgTask{
			State:             state,
			Start:             start,
			End:               end,
			Success:           false,
			TxReplacementOpts: &TxReplacementOpts{},
		},
	}
}

// Initialize prepares for work to be done in the Completion phase
func (t *CompletionTask) Initialize(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum, state interface{}) error {
	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("StateLocation", fmt.Sprintf("%p", t.State)).Info("CompletionTask Initialize()...")

	if t.State.Phase != objects.DisputeGPKJSubmission {
		return fmt.Errorf("%w because it's not in DisputeGPKJSubmission phase", objects.ErrCanNotContinue)
	}

	return nil
}

// DoWork is the first attempt
func (t *CompletionTask) DoWork(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

// DoRetry is all subsequent attempts
func (t *CompletionTask) DoRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {
	return t.doTask(ctx, logger, eth)
}

func (t *CompletionTask) doTask(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) error {

	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("CompletionTask doTask()")

	if t.isTaskCompleted(ctx, eth) {
		return nil
	}

	// Setup
	c := eth.Contracts()
	txnOpts, err := eth.GetTransactionOpts(ctx, t.State.Account)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "getting txn opts failed: %v", err)
	}

	logger.WithFields(logrus.Fields{
		"GasFeeCap": txnOpts.GasFeeCap,
		"GasTipCap": txnOpts.GasTipCap,
		"Nonce":     txnOpts.Nonce,
	}).Info("complete fees")

	// Register
	txn, err := c.Ethdkg().Complete(txnOpts)
	if err != nil {
		return dkg.LogReturnErrorf(logger, "completion failed: %v", err)
	}

	t.TxReplacementOpts.TxHash = txn.Hash()
	t.TxReplacementOpts.GasFeeCap = txn.GasFeeCap()
	t.TxReplacementOpts.GasTipCap = txn.GasTipCap()
	t.TxReplacementOpts.Nonce = big.NewInt(int64(txn.Nonce()))

	logger.WithFields(logrus.Fields{
		"GasFeeCap": t.TxReplacementOpts.GasFeeCap,
		"GasTipCap": t.TxReplacementOpts.GasTipCap,
		"Nonce":     t.TxReplacementOpts.Nonce,
		"Hash":      t.TxReplacementOpts.TxHash.Hex(),
	}).Info("complete fees2")

	logger.Info("CompletionTask sent completed call")

	// Queue transaction
	eth.Queue().QueueTransaction(ctx, txn)

	logger.Info("CompletionTask complete!")
	t.Success = true

	return nil
}

// ShouldRetry checks if it makes sense to try again
// Predicates:
// -- we haven't passed the last block
// -- the registration open hasn't moved, i.e. ETHDKG has not restarted
func (t *CompletionTask) ShouldRetry(ctx context.Context, logger *logrus.Entry, eth interfaces.Ethereum) bool {

	t.State.Lock()
	defer t.State.Unlock()

	logger.Info("CompletionTask ShouldRetry()")

	generalRetry := GeneralTaskShouldRetry(ctx, logger, eth, t.Start, t.End)
	if !generalRetry {
		return false
	}

	if t.isTaskCompleted(ctx, eth) {
		logger.WithFields(logrus.Fields{
			"t.State.Phase":      t.State.Phase,
			"t.State.PhaseStart": t.State.PhaseStart,
		}).Info("CompletionTask ShouldRetry - will not retry")
		return false
	}

	return true
}

// DoDone creates a log entry saying task is complete
func (t *CompletionTask) DoDone(logger *logrus.Entry) {
	t.State.Lock()
	defer t.State.Unlock()

	logger.WithField("Success", t.Success).Infof("CompletionTask done")
}

func (t *CompletionTask) GetDkgTask() *DkgTask {
	return t.DkgTask
}

func (t *CompletionTask) SetDkgTask(dkgTask *DkgTask) {
	t.DkgTask = dkgTask
}

func (t *CompletionTask) isTaskCompleted(ctx context.Context, eth interfaces.Ethereum) bool {
	c := eth.Contracts()
	phase, err := c.Ethdkg().GetETHDKGPhase(eth.GetCallOpts(ctx, t.State.Account))
	if err != nil {
		return false
	}

	return phase == uint8(objects.Completion)
}
