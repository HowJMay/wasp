// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package dividend

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "dividend"
	ScDescription = "Simple dividend smart contract"
	HScName       = wasmtypes.ScHname(0xcce2e239)
)

const (
	ParamAddress = "address"
	ParamFactor  = "factor"
	ParamOwner   = "owner"
)

const (
	ResultFactor = "factor"
	ResultOwner  = "owner"
)

const (
	StateMemberList  = "memberList"
	StateMembers     = "members"
	StateOwner       = "owner"
	StateTotalFactor = "totalFactor"
)

const (
	FuncDivide    = "divide"
	FuncInit      = "init"
	FuncMember    = "member"
	FuncSetOwner  = "setOwner"
	ViewGetFactor = "getFactor"
	ViewGetOwner  = "getOwner"
)

const (
	HFuncDivide    = wasmtypes.ScHname(0xc7878107)
	HFuncInit      = wasmtypes.ScHname(0x1f44d644)
	HFuncMember    = wasmtypes.ScHname(0xc07da2cb)
	HFuncSetOwner  = wasmtypes.ScHname(0x2a15fe7b)
	HViewGetFactor = wasmtypes.ScHname(0x0ee668fe)
	HViewGetOwner  = wasmtypes.ScHname(0x137107a6)
)