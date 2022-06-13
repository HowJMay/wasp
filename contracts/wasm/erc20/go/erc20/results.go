// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package erc20

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableAllowanceResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAllowanceResults) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultAmount))
}

type MutableAllowanceResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableAllowanceResults) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultAmount))
}

type ImmutableBalanceOfResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBalanceOfResults) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultAmount))
}

type MutableBalanceOfResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableBalanceOfResults) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultAmount))
}

type ImmutableTotalSupplyResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTotalSupplyResults) Supply() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultSupply))
}

type MutableTotalSupplyResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableTotalSupplyResults) Supply() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultSupply))
}
