// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const (
	ScName        = "testcore"
	ScDescription = "Core test for ISCP wasmlib Rust/Wasm library"
	HScName       = wasmlib.ScHname(0x370d33ad)
)

const (
	ParamAddress         = wasmlib.Key("address")
	ParamAgentID         = wasmlib.Key("agentID")
	ParamCaller          = wasmlib.Key("caller")
	ParamChainID         = wasmlib.Key("chainID")
	ParamChainOwnerID    = wasmlib.Key("chainOwnerID")
	ParamContractCreator = wasmlib.Key("contractCreator")
	ParamContractID      = wasmlib.Key("contractID")
	ParamCounter         = wasmlib.Key("counter")
	ParamHash            = wasmlib.Key("Hash")
	ParamHname           = wasmlib.Key("Hname")
	ParamHnameContract   = wasmlib.Key("hnameContract")
	ParamHnameEP         = wasmlib.Key("hnameEP")
	ParamHnameZero       = wasmlib.Key("Hname-0")
	ParamInt64           = wasmlib.Key("int64")
	ParamInt64Zero       = wasmlib.Key("int64-0")
	ParamIntValue        = wasmlib.Key("intParamValue")
	ParamName            = wasmlib.Key("intParamName")
	ParamString          = wasmlib.Key("string")
	ParamStringZero      = wasmlib.Key("string-0")
)

const (
	ResultChainOwnerID = wasmlib.Key("chainOwnerID")
	ResultCounter      = wasmlib.Key("counter")
	ResultIntValue     = wasmlib.Key("intParamValue")
	ResultMintedColor  = wasmlib.Key("mintedColor")
	ResultMintedSupply = wasmlib.Key("mintedSupply")
	ResultSandboxCall  = wasmlib.Key("sandboxCall")
)

const (
	StateCounter      = wasmlib.Key("counter")
	StateHnameEP      = wasmlib.Key("hnameEP")
	StateInts         = wasmlib.Key("ints")
	StateMintedColor  = wasmlib.Key("mintedColor")
	StateMintedSupply = wasmlib.Key("mintedSupply")
)

const (
	FuncCallOnChain                 = "callOnChain"
	FuncCheckContextFromFullEP      = "checkContextFromFullEP"
	FuncDoNothing                   = "doNothing"
	FuncGetMintedSupply             = "getMintedSupply"
	FuncIncCounter                  = "incCounter"
	FuncInit                        = "init"
	FuncPassTypesFull               = "passTypesFull"
	FuncRunRecursion                = "runRecursion"
	FuncSendToAddress               = "sendToAddress"
	FuncSetInt                      = "setInt"
	FuncTestCallPanicFullEP         = "testCallPanicFullEP"
	FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull"
	FuncTestChainOwnerIDFull        = "testChainOwnerIDFull"
	FuncTestEventLogDeploy          = "testEventLogDeploy"
	FuncTestEventLogEventData       = "testEventLogEventData"
	FuncTestEventLogGenericData     = "testEventLogGenericData"
	FuncTestPanicFullEP             = "testPanicFullEP"
	FuncWithdrawToChain             = "withdrawToChain"
	ViewCheckContextFromViewEP      = "checkContextFromViewEP"
	ViewFibonacci                   = "fibonacci"
	ViewGetCounter                  = "getCounter"
	ViewGetInt                      = "getInt"
	ViewJustView                    = "justView"
	ViewPassTypesView               = "passTypesView"
	ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView"
	ViewTestChainOwnerIDView        = "testChainOwnerIDView"
	ViewTestPanicViewEP             = "testPanicViewEP"
	ViewTestSandboxCall             = "testSandboxCall"
)

const (
	HFuncCallOnChain                 = wasmlib.ScHname(0x95a3d123)
	HFuncCheckContextFromFullEP      = wasmlib.ScHname(0xa56c24ba)
	HFuncDoNothing                   = wasmlib.ScHname(0xdda4a6de)
	HFuncGetMintedSupply             = wasmlib.ScHname(0x0c2d113c)
	HFuncIncCounter                  = wasmlib.ScHname(0x7b287419)
	HFuncInit                        = wasmlib.ScHname(0x1f44d644)
	HFuncPassTypesFull               = wasmlib.ScHname(0x733ea0ea)
	HFuncRunRecursion                = wasmlib.ScHname(0x833425fd)
	HFuncSendToAddress               = wasmlib.ScHname(0x63ce4634)
	HFuncSetInt                      = wasmlib.ScHname(0x62056f74)
	HFuncTestCallPanicFullEP         = wasmlib.ScHname(0x4c878834)
	HFuncTestCallPanicViewEPFromFull = wasmlib.ScHname(0xfd7e8c1d)
	HFuncTestChainOwnerIDFull        = wasmlib.ScHname(0x2aff1167)
	HFuncTestEventLogDeploy          = wasmlib.ScHname(0x96ff760a)
	HFuncTestEventLogEventData       = wasmlib.ScHname(0x0efcf939)
	HFuncTestEventLogGenericData     = wasmlib.ScHname(0x6a16629d)
	HFuncTestPanicFullEP             = wasmlib.ScHname(0x24fdef07)
	HFuncWithdrawToChain             = wasmlib.ScHname(0x437bc026)
	HViewCheckContextFromViewEP      = wasmlib.ScHname(0x88ff0167)
	HViewFibonacci                   = wasmlib.ScHname(0x7940873c)
	HViewGetCounter                  = wasmlib.ScHname(0xb423e607)
	HViewGetInt                      = wasmlib.ScHname(0x1887e5ef)
	HViewJustView                    = wasmlib.ScHname(0x33b8972e)
	HViewPassTypesView               = wasmlib.ScHname(0x1a5b87ea)
	HViewTestCallPanicViewEPFromView = wasmlib.ScHname(0x91b10c99)
	HViewTestChainOwnerIDView        = wasmlib.ScHname(0x26586c33)
	HViewTestPanicViewEP             = wasmlib.ScHname(0x22bc4d72)
	HViewTestSandboxCall             = wasmlib.ScHname(0x42d72b63)
)