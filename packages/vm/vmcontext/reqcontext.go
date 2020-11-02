package vmcontext

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/sctransaction/txbuilder"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/processors"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
)

// context for one request
type VMContext struct {
	// same for the block
	chainID            coretypes.ChainID
	processors         *processors.ProcessorCache
	ownerAddress       address.Address
	rewardAddress      address.Address
	minimumReward      int64
	nodeRewardsEnabled bool
	txBuilder          *txbuilder.Builder // mutated
	saveTxBuilder      *txbuilder.Builder // for rollback
	virtualState       state.VirtualState // mutated
	log                *logger.Logger
	// request context
	entropy     hashing.HashValue // mutates with each request
	reqRef      sctransaction.RequestRef
	timestamp   int64
	stateUpdate state.StateUpdate // mutated
	callStack   []*callContext
}

type callContext struct {
	contractIndex uint16
	params        codec.ImmutableCodec
	budget        map[balance.Color]int64
}

func (vmctx *VMContext) PushCallContext(contractIndex uint16, params codec.ImmutableCodec, budget map[balance.Color]int64) error {
	vmctx.callStack = append(vmctx.callStack, &callContext{
		contractIndex: contractIndex,
		params:        params,
		budget:        budget,
	})
	// TODO check budget
	return nil
}

func (vmctx *VMContext) PopCallContext() {
	vmctx.callStack = vmctx.callStack[:len(vmctx.callStack)-1]
}

// handleNodeRewards return true if to continue with request processing
// rewards are "rewards for node", so smart contract sending request to itself might need
// to pay rewards too
func (vmctx *VMContext) handleNodeRewards() bool {
	if !vmctx.nodeRewardsEnabled {
		// nothing to do
		return true
	}
	var err error

	reqTxId := vmctx.reqRef.Tx.ID()
	// determining how many iotas have been left in the request transaction
	availableIotas := vmctx.txBuilder.GetInputBalanceFromTransaction(balance.ColorIOTA, reqTxId)

	var proceed bool
	// taking into account 1 request token which will be recolored back to iota
	// and will be send to the node reward address (if enabled)
	var sendToRewardAddress int64
	if availableIotas+1 >= vmctx.minimumReward {
		sendToRewardAddress = vmctx.minimumReward - 1
		proceed = true
	} else {
		sendToRewardAddress = availableIotas
		// if reward is not enough, the state update will be empty, i.e. NOP (the fee will be taken)
		proceed = false
	}
	err = vmctx.txBuilder.MoveToAddressFromTransaction(vmctx.rewardAddress, balance.ColorIOTA, sendToRewardAddress, reqTxId)

	if err != nil {
		vmctx.log.Error("can't move reward tokens: %v", err)
		proceed = false
	}
	return proceed
}

func (vmctx *VMContext) getEntryPoint() (vmtypes.EntryPoint, bool) {
	return nil, false
}

func (vmctx *VMContext) FinalizeTransaction(blockIndex uint32, stateHash *hashing.HashValue, timestamp int64) (*sctransaction.Transaction, error) {
	// add state block
	err := vmctx.txBuilder.SetStateParams(blockIndex, stateHash, timestamp)
	if err != nil {
		return nil, err
	}
	// create result transaction
	tx, err := vmctx.txBuilder.Build(false)
	if err != nil {
		return nil, err
	}
	// check semantic just in case
	if _, err := tx.Properties(); err != nil {
		return nil, err
	}
	return tx, nil
}
