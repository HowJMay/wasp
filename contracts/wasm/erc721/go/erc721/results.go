// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package erc721

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableBalanceOfResults struct {
	proxy wasmtypes.Proxy
}

// amount of tokens owned by owner
func (s ImmutableBalanceOfResults) Amount() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ResultAmount))
}

type MutableBalanceOfResults struct {
	proxy wasmtypes.Proxy
}

// amount of tokens owned by owner
func (s MutableBalanceOfResults) Amount() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ResultAmount))
}

type ImmutableGetApprovedResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetApprovedResults) Approved() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultApproved))
}

type MutableGetApprovedResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetApprovedResults) Approved() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultApproved))
}

type ImmutableIsApprovedForAllResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableIsApprovedForAllResults) Approval() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ResultApproval))
}

type MutableIsApprovedForAllResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableIsApprovedForAllResults) Approval() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ResultApproval))
}

type ImmutableNameResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableNameResults) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultName))
}

type MutableNameResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableNameResults) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultName))
}

type ImmutableOwnerOfResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableOwnerOfResults) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultOwner))
}

type MutableOwnerOfResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableOwnerOfResults) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultOwner))
}

type ImmutableSymbolResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSymbolResults) Symbol() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultSymbol))
}

type MutableSymbolResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableSymbolResults) Symbol() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultSymbol))
}

type ImmutableTokenURIResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTokenURIResults) TokenURI() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ResultTokenURI))
}

type MutableTokenURIResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableTokenURIResults) TokenURI() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ResultTokenURI))
}