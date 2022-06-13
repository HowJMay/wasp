// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package evmimpl

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/iotaledger/wasp/packages/evm/evmtypes"
	"github.com/iotaledger/wasp/packages/evm/evmutil"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/core/evm"
	"github.com/iotaledger/wasp/packages/vm/core/evm/emulator"
	"github.com/iotaledger/wasp/packages/vm/gas"
)

var Processor = evm.Contract.Processor(initialize,
	evm.FuncSetGasRatio.WithHandler(setGasRatio),
	evm.FuncGetGasRatio.WithHandler(getGasRatio),
	evm.FuncSendTransaction.WithHandler(applyTransaction),
	evm.FuncGetBalance.WithHandler(getBalance),
	evm.FuncCallContract.WithHandler(callContract),
	evm.FuncEstimateGas.WithHandler(estimateGas),
	evm.FuncGetNonce.WithHandler(getNonce),
	evm.FuncGetReceipt.WithHandler(getReceipt),
	evm.FuncGetCode.WithHandler(getCode),
	evm.FuncGetBlockNumber.WithHandler(getBlockNumber),
	evm.FuncGetBlockByNumber.WithHandler(getBlockByNumber),
	evm.FuncGetBlockByHash.WithHandler(getBlockByHash),
	evm.FuncGetTransactionByHash.WithHandler(getTransactionByHash),
	evm.FuncGetTransactionByBlockHashAndIndex.WithHandler(getTransactionByBlockHashAndIndex),
	evm.FuncGetTransactionByBlockNumberAndIndex.WithHandler(getTransactionByBlockNumberAndIndex),
	evm.FuncGetTransactionCountByBlockHash.WithHandler(getTransactionCountByBlockHash),
	evm.FuncGetTransactionCountByBlockNumber.WithHandler(getTransactionCountByBlockNumber),
	evm.FuncGetStorage.WithHandler(getStorage),
	evm.FuncGetLogs.WithHandler(getLogs),
	evm.FuncGetChainID.WithHandler(getChainID),
)

func initialize(ctx iscp.Sandbox) dict.Dict {
	genesisAlloc := core.GenesisAlloc{}
	var err error
	if ctx.Params().MustHas(evm.FieldGenesisAlloc) {
		genesisAlloc, err = evmtypes.DecodeGenesisAlloc(ctx.Params().MustGet(evm.FieldGenesisAlloc))
		ctx.RequireNoError(err)
	}

	gasLimit, err := codec.DecodeUint64(ctx.Params().MustGet(evm.FieldBlockGasLimit), evm.BlockGasLimitDefault)
	ctx.RequireNoError(err)

	blockKeepAmount, err := codec.DecodeInt32(ctx.Params().MustGet(evm.FieldBlockKeepAmount), evm.BlockKeepAmountDefault)
	ctx.RequireNoError(err)

	// add the standard ISC contract at arbitrary address 0x1074
	deployISCContractOnGenesis(genesisAlloc)

	chainID := evmtypes.MustDecodeChainID(ctx.Params().MustGet(evm.FieldChainID), evm.DefaultChainID)
	emulator.Init(
		evmStateSubrealm(ctx.State()),
		chainID,
		blockKeepAmount,
		gasLimit,
		timestamp(ctx),
		genesisAlloc,
	)

	gasRatio := codec.MustDecodeRatio32(ctx.Params().MustGet(evm.FieldGasRatio), evmtypes.DefaultGasRatio)
	ctx.State().Set(keyGasRatio, gasRatio.Bytes())
	// storing hname as a terminal value of the contract's state nil key.
	// This way we will be able to retrieve commitment to the contract's state
	ctx.State().Set("", ctx.Contract().Bytes())

	return nil
}

func applyTransaction(ctx iscp.Sandbox) dict.Dict {
	// we only want to charge gas for the actual execution of the ethereum tx
	ctx.Privileged().GasBurnEnable(false)
	defer ctx.Privileged().GasBurnEnable(true)

	tx := &types.Transaction{}
	err := tx.UnmarshalBinary(ctx.Params().MustGet(evm.FieldTransaction))
	ctx.RequireNoError(err)

	ctx.RequireCaller(iscp.NewEthereumAddressAgentID(evmutil.MustGetSender(tx)))

	// next block will be minted when the ISC block is closed
	bctx := getBlockContext(ctx)

	ctx.Requiref(tx.ChainId().Uint64() == uint64(bctx.emu.BlockchainDB().GetChainID()), "chainId mismatch")

	receipt, result, err := bctx.emu.SendTransaction(tx, ctx.Privileged().GasBurnEnable)

	// burn EVM gas as ISC gas
	if result != nil {
		gasRatio := codec.MustDecodeRatio32(ctx.State().MustGet(keyGasRatio), evmtypes.DefaultGasRatio)
		ctx.Privileged().GasBurnEnable(true)
		ctx.Gas().Burn(gas.BurnCodeEVM1P, evmtypes.EVMGasToISC(result.UsedGas, &gasRatio))
		ctx.Privileged().GasBurnEnable(false)
	}

	ctx.RequireNoError(err)

	// If EVM execution was reverted we must revert the ISC request as well.
	// Failed txs will be stored when closing the block context.
	bctx.txs = append(bctx.txs, tx)
	bctx.receipts = append(bctx.receipts, receipt)
	ctx.Requiref(receipt.Status == types.ReceiptStatusSuccessful, GetRevertErrorMessage(result, ctx.Contract()))

	return nil
}

func getBalance(ctx iscp.SandboxView) dict.Dict {
	addr := common.BytesToAddress(ctx.Params().MustGet(evm.FieldAddress))
	emu := createEmulatorR(ctx)
	_ = paramBlockNumberOrHashAsNumber(ctx, emu, false)
	return result(emu.StateDB().GetBalance(addr).Bytes())
}

func getBlockNumber(ctx iscp.SandboxView) dict.Dict {
	emu := createEmulatorR(ctx)
	return result(new(big.Int).SetUint64(emu.BlockchainDB().GetNumber()).Bytes())
}

func getBlockByNumber(ctx iscp.SandboxView) dict.Dict {
	return blockResult(blockByNumber(ctx))
}

func getBlockByHash(ctx iscp.SandboxView) dict.Dict {
	return blockResult(blockByHash(ctx))
}

func getTransactionByHash(ctx iscp.SandboxView) dict.Dict {
	return txResult(transactionByHash(ctx))
}

func getTransactionByBlockHashAndIndex(ctx iscp.SandboxView) dict.Dict {
	return txResult(transactionByBlockHashAndIndex(ctx))
}

func getTransactionByBlockNumberAndIndex(ctx iscp.SandboxView) dict.Dict {
	return txResult(transactionByBlockNumberAndIndex(ctx))
}

func getTransactionCountByBlockHash(ctx iscp.SandboxView) dict.Dict {
	return txCountResult(blockByHash(ctx))
}

func getTransactionCountByBlockNumber(ctx iscp.SandboxView) dict.Dict {
	return txCountResult(blockByNumber(ctx))
}

func getReceipt(ctx iscp.SandboxView) dict.Dict {
	txHash := common.BytesToHash(ctx.Params().MustGet(evm.FieldTransactionHash))
	emu := createEmulatorR(ctx)
	r := emu.BlockchainDB().GetReceiptByTxHash(txHash)
	if r == nil {
		return nil
	}
	return result(evmtypes.EncodeReceiptFull(r))
}

func getNonce(ctx iscp.SandboxView) dict.Dict {
	emu := createEmulatorR(ctx)
	addr := common.BytesToAddress(ctx.Params().MustGet(evm.FieldAddress))
	_ = paramBlockNumberOrHashAsNumber(ctx, emu, false)
	return result(codec.EncodeUint64(emu.StateDB().GetNonce(addr)))
}

func getCode(ctx iscp.SandboxView) dict.Dict {
	emu := createEmulatorR(ctx)
	addr := common.BytesToAddress(ctx.Params().MustGet(evm.FieldAddress))
	_ = paramBlockNumberOrHashAsNumber(ctx, emu, false)
	return result(emu.StateDB().GetCode(addr))
}

func getStorage(ctx iscp.SandboxView) dict.Dict {
	emu := createEmulatorR(ctx)
	addr := common.BytesToAddress(ctx.Params().MustGet(evm.FieldAddress))
	key := common.BytesToHash(ctx.Params().MustGet(evm.FieldKey))
	_ = paramBlockNumberOrHashAsNumber(ctx, emu, false)
	data := emu.StateDB().GetState(addr, key)
	return result(data[:])
}

func getLogs(ctx iscp.SandboxView) dict.Dict {
	q, err := evmtypes.DecodeFilterQuery(ctx.Params().MustGet(evm.FieldFilterQuery))
	ctx.RequireNoError(err)
	emu := createEmulatorR(ctx)
	logs := emu.FilterLogs(q)
	return result(evmtypes.EncodeLogs(logs))
}

func getChainID(ctx iscp.SandboxView) dict.Dict {
	emu := createEmulatorR(ctx)
	return result(evmtypes.EncodeChainID(emu.BlockchainDB().GetChainID()))
}

func callContract(ctx iscp.SandboxView) dict.Dict {
	callMsg, err := evmtypes.DecodeCallMsg(ctx.Params().MustGet(evm.FieldCallMsg))
	ctx.RequireNoError(err)
	emu := createEmulatorR(ctx)
	_ = paramBlockNumberOrHashAsNumber(ctx, emu, false)
	res, err := emu.CallContract(callMsg, nil)
	ctx.RequireNoError(err)
	ctx.Requiref(res.Err == nil, GetRevertErrorMessage(res, ctx.Contract()))
	return result(res.Return())
}

// TODO: For some reason, when EstimateGasMode == true the gas burned is less. How to automatically calculate this?
var additionalGasBurned = gas.BurnCodeReadFromState1P.Cost(2)

func estimateGas(ctx iscp.Sandbox) dict.Dict {
	// we only want to charge gas for the actual execution of the ethereum tx
	ctx.Privileged().GasBurnEnable(false)
	defer ctx.Privileged().GasBurnEnable(true)

	callMsg, err := evmtypes.DecodeCallMsg(ctx.Params().MustGet(evm.FieldCallMsg))
	ctx.RequireNoError(err)
	ctx.RequireCaller(iscp.NewEthereumAddressAgentID(callMsg.From))

	emu := createEmulator(ctx)
	res, err := emu.CallContract(callMsg, ctx.Privileged().GasBurnEnable)
	ctx.RequireNoError(err)
	ctx.Requiref(res.Err == nil, GetRevertErrorMessage(res, ctx.Contract()))

	// TODO: this assumes that the initial budget was gas.MaxGasPerCall
	// see evmOffLedgerEstimateGasRequest::GasBudget()
	// and VMContext::calculateAffordableGasBudget() when EstimateGasMode == true
	iscGasBurned := gas.MaxGasPerCall - ctx.Gas().Budget()
	gasRatio := codec.MustDecodeRatio32(ctx.State().MustGet(keyGasRatio), evmtypes.DefaultGasRatio)
	evmGasBurnedInISCCalls := evmtypes.ISCGasBurnedToEVM(iscGasBurned, &gasRatio) + additionalGasBurned
	return result(codec.EncodeUint64(res.UsedGas + evmGasBurnedInISCCalls))
}