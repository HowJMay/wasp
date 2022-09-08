// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package testcore

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "testcore"
	ScDescription = "Wasm equivalent of built-in TestCore contract"
	HScName       = wasmtypes.ScHname(0x370d33ad)
)

const (
	ParamAddress              = "address"
	ParamAgentID              = "agentID"
	ParamBaseTokensWithdrawal = "baseTokensWithdrawal"
	ParamCaller               = "caller"
	ParamChainID              = "chainID"
	ParamChainOwnerID         = "chainOwnerID"
	ParamContractID           = "contractID"
	ParamCounter              = "counter"
	ParamFail                 = "initFailParam"
	ParamGasBudget            = "gasBudget"
	ParamHash                 = "Hash"
	ParamHname                = "Hname"
	ParamHnameContract        = "hnameContract"
	ParamHnameEP              = "hnameEP"
	ParamHnameZero            = "Hname-0"
	ParamInt64                = "int64"
	ParamInt64Zero            = "int64-0"
	ParamIntValue             = "intParamValue"
	ParamN                    = "n"
	ParamName                 = "intParamName"
	ParamProgHash             = "progHash"
	ParamString               = "string"
	ParamStringZero           = "string-0"
	ParamVarName              = "varName"
)

const (
	ResultChainOwnerID = "chainOwnerID"
	ResultCounter      = "counter"
	ResultN            = "n"
	ResultSandboxCall  = "sandboxCall"
	ResultValues       = "this"
	ResultVars         = "this"
)

const (
	StateCounter = "counter"
	StateInts    = "ints"
	StateStrings = "strings"
)

const (
	FuncCallOnChain                 = "callOnChain"
	FuncCheckContextFromFullEP      = "checkContextFromFullEP"
	FuncClaimAllowance              = "claimAllowance"
	FuncDoNothing                   = "doNothing"
	FuncEstimateMinStorageDeposit   = "estimateMinStorageDeposit"
	FuncIncCounter                  = "incCounter"
	FuncInfiniteLoop                = "infiniteLoop"
	FuncInit                        = "init"
	FuncPassTypesFull               = "passTypesFull"
	FuncPingAllowanceBack           = "pingAllowanceBack"
	FuncRunRecursion                = "runRecursion"
	FuncSendLargeRequest            = "sendLargeRequest"
	FuncSendNFTsBack                = "sendNFTsBack"
	FuncSendToAddress               = "sendToAddress"
	FuncSetInt                      = "setInt"
	FuncSpawn                       = "spawn"
	FuncSplitFunds                  = "splitFunds"
	FuncSplitFundsNativeTokens      = "splitFundsNativeTokens"
	FuncTestBlockContext1           = "testBlockContext1"
	FuncTestBlockContext2           = "testBlockContext2"
	FuncTestCallPanicFullEP         = "testCallPanicFullEP"
	FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull"
	FuncTestChainOwnerIDFull        = "testChainOwnerIDFull"
	FuncTestEventLogDeploy          = "testEventLogDeploy"
	FuncTestEventLogEventData       = "testEventLogEventData"
	FuncTestEventLogGenericData     = "testEventLogGenericData"
	FuncTestPanicFullEP             = "testPanicFullEP"
	FuncWithdrawFromChain           = "withdrawFromChain"
	ViewCheckContextFromViewEP      = "checkContextFromViewEP"
	ViewFibonacci                   = "fibonacci"
	ViewFibonacciIndirect           = "fibonacciIndirect"
	ViewGetCounter                  = "getCounter"
	ViewGetInt                      = "getInt"
	ViewGetStringValue              = "getStringValue"
	ViewInfiniteLoopView            = "infiniteLoopView"
	ViewJustView                    = "justView"
	ViewPassTypesView               = "passTypesView"
	ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView"
	ViewTestChainOwnerIDView        = "testChainOwnerIDView"
	ViewTestPanicViewEP             = "testPanicViewEP"
	ViewTestSandboxCall             = "testSandboxCall"
)

const (
	HFuncCallOnChain                 = wasmtypes.ScHname(0x95a3d123)
	HFuncCheckContextFromFullEP      = wasmtypes.ScHname(0xa56c24ba)
	HFuncClaimAllowance              = wasmtypes.ScHname(0x40bec0e6)
	HFuncDoNothing                   = wasmtypes.ScHname(0xdda4a6de)
	HFuncEstimateMinStorageDeposit   = wasmtypes.ScHname(0xf5d411cb)
	HFuncIncCounter                  = wasmtypes.ScHname(0x7b287419)
	HFuncInfiniteLoop                = wasmtypes.ScHname(0xf571430a)
	HFuncInit                        = wasmtypes.ScHname(0x1f44d644)
	HFuncPassTypesFull               = wasmtypes.ScHname(0x733ea0ea)
	HFuncPingAllowanceBack           = wasmtypes.ScHname(0x66f43c0b)
	HFuncRunRecursion                = wasmtypes.ScHname(0x833425fd)
	HFuncSendLargeRequest            = wasmtypes.ScHname(0xfdaaca3c)
	HFuncSendNFTsBack                = wasmtypes.ScHname(0x8f6ef428)
	HFuncSendToAddress               = wasmtypes.ScHname(0x63ce4634)
	HFuncSetInt                      = wasmtypes.ScHname(0x62056f74)
	HFuncSpawn                       = wasmtypes.ScHname(0xec929d12)
	HFuncSplitFunds                  = wasmtypes.ScHname(0xc7ea86c9)
	HFuncSplitFundsNativeTokens      = wasmtypes.ScHname(0x16532a28)
	HFuncTestBlockContext1           = wasmtypes.ScHname(0x796d4136)
	HFuncTestBlockContext2           = wasmtypes.ScHname(0x758b0452)
	HFuncTestCallPanicFullEP         = wasmtypes.ScHname(0x4c878834)
	HFuncTestCallPanicViewEPFromFull = wasmtypes.ScHname(0xfd7e8c1d)
	HFuncTestChainOwnerIDFull        = wasmtypes.ScHname(0x2aff1167)
	HFuncTestEventLogDeploy          = wasmtypes.ScHname(0x96ff760a)
	HFuncTestEventLogEventData       = wasmtypes.ScHname(0x0efcf939)
	HFuncTestEventLogGenericData     = wasmtypes.ScHname(0x6a16629d)
	HFuncTestPanicFullEP             = wasmtypes.ScHname(0x24fdef07)
	HFuncWithdrawFromChain           = wasmtypes.ScHname(0x405c0b0a)
	HViewCheckContextFromViewEP      = wasmtypes.ScHname(0x88ff0167)
	HViewFibonacci                   = wasmtypes.ScHname(0x7940873c)
	HViewFibonacciIndirect           = wasmtypes.ScHname(0x6dd98513)
	HViewGetCounter                  = wasmtypes.ScHname(0xb423e607)
	HViewGetInt                      = wasmtypes.ScHname(0x1887e5ef)
	HViewGetStringValue              = wasmtypes.ScHname(0xcf0a4d32)
	HViewInfiniteLoopView            = wasmtypes.ScHname(0x1a383295)
	HViewJustView                    = wasmtypes.ScHname(0x33b8972e)
	HViewPassTypesView               = wasmtypes.ScHname(0x1a5b87ea)
	HViewTestCallPanicViewEPFromView = wasmtypes.ScHname(0x91b10c99)
	HViewTestChainOwnerIDView        = wasmtypes.ScHname(0x26586c33)
	HViewTestPanicViewEP             = wasmtypes.ScHname(0x22bc4d72)
	HViewTestSandboxCall             = wasmtypes.ScHname(0x42d72b63)
)