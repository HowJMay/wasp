// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableAddressMapOfAddressArrayAppendParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableAddressMapOfAddressArrayAppendParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableAddressMapOfAddressArrayClearParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class MutableAddressMapOfAddressArrayClearParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class ImmutableAddressMapOfAddressArraySetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableAddressMapOfAddressArraySetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableAddressMapOfAddressMapClearParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class MutableAddressMapOfAddressMapClearParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class ImmutableAddressMapOfAddressMapSetParams extends wasmtypes.ScProxy {
	keyAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableAddressMapOfAddressMapSetParams extends wasmtypes.ScProxy {
	keyAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}

	valueAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableArrayOfAddressArrayAppendParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	valueAddr(): sc.ArrayOfImmutableAddress {
		return new sc.ArrayOfImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableArrayOfAddressArrayAppendParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	valueAddr(): sc.ArrayOfMutableAddress {
		return new sc.ArrayOfMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableArrayOfAddressArraySetParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex1));
	}

	valueAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableArrayOfAddressArraySetParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex1));
	}

	valueAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableArrayOfAddressMapSetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	keyAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	valueAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class MutableArrayOfAddressMapSetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	keyAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	valueAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamValueAddr));
	}
}

export class ImmutableArrayOfStringArrayAppendParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	value(): sc.ArrayOfImmutableString {
		return new sc.ArrayOfImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableArrayOfStringArrayAppendParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	value(): sc.ArrayOfMutableString {
		return new sc.ArrayOfMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class ImmutableArrayOfStringArraySetParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex1));
	}

	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableArrayOfStringArraySetParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex1));
	}

	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class ImmutableArrayOfStringMapSetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	key(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamKey));
	}

	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableArrayOfStringMapSetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	key(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamKey));
	}

	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MapStringToImmutableBytes extends wasmtypes.ScProxy {

	getBytes(key: string): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class ImmutableParamTypesParams extends wasmtypes.ScProxy {
	address(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamAddress));
	}

	agentID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
	}

	bool(): wasmtypes.ScImmutableBool {
		return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamBool));
	}

	bytes(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamBytes));
	}

	chainID(): wasmtypes.ScImmutableChainID {
		return new wasmtypes.ScImmutableChainID(this.proxy.root(sc.ParamChainID));
	}

	hash(): wasmtypes.ScImmutableHash {
		return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamHash));
	}

	hname(): wasmtypes.ScImmutableHname {
		return new wasmtypes.ScImmutableHname(this.proxy.root(sc.ParamHname));
	}

	int16(): wasmtypes.ScImmutableInt16 {
		return new wasmtypes.ScImmutableInt16(this.proxy.root(sc.ParamInt16));
	}

	int32(): wasmtypes.ScImmutableInt32 {
		return new wasmtypes.ScImmutableInt32(this.proxy.root(sc.ParamInt32));
	}

	int64(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamInt64));
	}

	int8(): wasmtypes.ScImmutableInt8 {
		return new wasmtypes.ScImmutableInt8(this.proxy.root(sc.ParamInt8));
	}

	nftID(): wasmtypes.ScImmutableNftID {
		return new wasmtypes.ScImmutableNftID(this.proxy.root(sc.ParamNftID));
	}

	// special hook to be able to pass key/values as raw bytes
	param(): sc.MapStringToImmutableBytes {
		return new sc.MapStringToImmutableBytes(this.proxy);
	}

	requestID(): wasmtypes.ScImmutableRequestID {
		return new wasmtypes.ScImmutableRequestID(this.proxy.root(sc.ParamRequestID));
	}

	string(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamString));
	}

	tokenID(): wasmtypes.ScImmutableTokenID {
		return new wasmtypes.ScImmutableTokenID(this.proxy.root(sc.ParamTokenID));
	}

	uint16(): wasmtypes.ScImmutableUint16 {
		return new wasmtypes.ScImmutableUint16(this.proxy.root(sc.ParamUint16));
	}

	uint32(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamUint32));
	}

	uint64(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ParamUint64));
	}

	uint8(): wasmtypes.ScImmutableUint8 {
		return new wasmtypes.ScImmutableUint8(this.proxy.root(sc.ParamUint8));
	}
}

export class MapStringToMutableBytes extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getBytes(key: string): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MutableParamTypesParams extends wasmtypes.ScProxy {
	address(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamAddress));
	}

	agentID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
	}

	bool(): wasmtypes.ScMutableBool {
		return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamBool));
	}

	bytes(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamBytes));
	}

	chainID(): wasmtypes.ScMutableChainID {
		return new wasmtypes.ScMutableChainID(this.proxy.root(sc.ParamChainID));
	}

	hash(): wasmtypes.ScMutableHash {
		return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamHash));
	}

	hname(): wasmtypes.ScMutableHname {
		return new wasmtypes.ScMutableHname(this.proxy.root(sc.ParamHname));
	}

	int16(): wasmtypes.ScMutableInt16 {
		return new wasmtypes.ScMutableInt16(this.proxy.root(sc.ParamInt16));
	}

	int32(): wasmtypes.ScMutableInt32 {
		return new wasmtypes.ScMutableInt32(this.proxy.root(sc.ParamInt32));
	}

	int64(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamInt64));
	}

	int8(): wasmtypes.ScMutableInt8 {
		return new wasmtypes.ScMutableInt8(this.proxy.root(sc.ParamInt8));
	}

	nftID(): wasmtypes.ScMutableNftID {
		return new wasmtypes.ScMutableNftID(this.proxy.root(sc.ParamNftID));
	}

	// special hook to be able to pass key/values as raw bytes
	param(): sc.MapStringToMutableBytes {
		return new sc.MapStringToMutableBytes(this.proxy);
	}

	requestID(): wasmtypes.ScMutableRequestID {
		return new wasmtypes.ScMutableRequestID(this.proxy.root(sc.ParamRequestID));
	}

	string(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamString));
	}

	tokenID(): wasmtypes.ScMutableTokenID {
		return new wasmtypes.ScMutableTokenID(this.proxy.root(sc.ParamTokenID));
	}

	uint16(): wasmtypes.ScMutableUint16 {
		return new wasmtypes.ScMutableUint16(this.proxy.root(sc.ParamUint16));
	}

	uint32(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamUint32));
	}

	uint64(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ParamUint64));
	}

	uint8(): wasmtypes.ScMutableUint8 {
		return new wasmtypes.ScMutableUint8(this.proxy.root(sc.ParamUint8));
	}
}

export class ImmutableStringMapOfStringArrayAppendParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableStringMapOfStringArrayAppendParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class ImmutableStringMapOfStringArrayClearParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableStringMapOfStringArrayClearParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}

export class ImmutableStringMapOfStringArraySetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableStringMapOfStringArraySetParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class ImmutableStringMapOfStringMapClearParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableStringMapOfStringMapClearParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}

export class ImmutableStringMapOfStringMapSetParams extends wasmtypes.ScProxy {
	key(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamKey));
	}

	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamValue));
	}
}

export class MutableStringMapOfStringMapSetParams extends wasmtypes.ScProxy {
	key(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamKey));
	}

	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}

	value(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamValue));
	}
}

export class ImmutableTriggerEventParams extends wasmtypes.ScProxy {
	address(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamAddress));
	}

	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableTriggerEventParams extends wasmtypes.ScProxy {
	address(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamAddress));
	}

	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}

export class ImmutableAddressMapOfAddressArrayLengthParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class MutableAddressMapOfAddressArrayLengthParams extends wasmtypes.ScProxy {
	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class ImmutableAddressMapOfAddressArrayValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class MutableAddressMapOfAddressArrayValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class ImmutableAddressMapOfAddressMapValueParams extends wasmtypes.ScProxy {
	keyAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	nameAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class MutableAddressMapOfAddressMapValueParams extends wasmtypes.ScProxy {
	keyAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}

	nameAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamNameAddr));
	}
}

export class ImmutableArrayOfAddressArrayValueParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex1));
	}
}

export class MutableArrayOfAddressArrayValueParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex1));
	}
}

export class ImmutableArrayOfAddressMapValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	keyAddr(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}
}

export class MutableArrayOfAddressMapValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	keyAddr(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamKeyAddr));
	}
}

export class ImmutableArrayOfStringArrayValueParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex1));
	}
}

export class MutableArrayOfStringArrayValueParams extends wasmtypes.ScProxy {
	index0(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex0));
	}

	index1(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex1));
	}
}

export class ImmutableArrayOfStringMapValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	key(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamKey));
	}
}

export class MutableArrayOfStringMapValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	key(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamKey));
	}
}

export class ImmutableBigIntAddParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class MutableBigIntAddParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class ImmutableBigIntDivParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class MutableBigIntDivParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class ImmutableBigIntModParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class MutableBigIntModParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class ImmutableBigIntMulParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class MutableBigIntMulParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class ImmutableBigIntShlParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	shift(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamShift));
	}
}

export class MutableBigIntShlParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	shift(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamShift));
	}
}

export class ImmutableBigIntShrParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	shift(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamShift));
	}
}

export class MutableBigIntShrParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	shift(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamShift));
	}
}

export class ImmutableBigIntSubParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScImmutableBigInt {
		return new wasmtypes.ScImmutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class MutableBigIntSubParams extends wasmtypes.ScProxy {
	lhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamLhs));
	}

	rhs(): wasmtypes.ScMutableBigInt {
		return new wasmtypes.ScMutableBigInt(this.proxy.root(sc.ParamRhs));
	}
}

export class ImmutableBlockRecordParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}

	recordIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamRecordIndex));
	}
}

export class MutableBlockRecordParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}

	recordIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamRecordIndex));
	}
}

export class ImmutableBlockRecordsParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class MutableBlockRecordsParams extends wasmtypes.ScProxy {
	blockIndex(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamBlockIndex));
	}
}

export class ImmutableCheckAddressParams extends wasmtypes.ScProxy {
	addressBytes(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAddressBytes));
	}

	addressString(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamAddressString));
	}

	scAddress(): wasmtypes.ScImmutableAddress {
		return new wasmtypes.ScImmutableAddress(this.proxy.root(sc.ParamScAddress));
	}
}

export class MutableCheckAddressParams extends wasmtypes.ScProxy {
	addressBytes(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAddressBytes));
	}

	addressString(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamAddressString));
	}

	scAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.root(sc.ParamScAddress));
	}
}

export class ImmutableCheckAgentIDParams extends wasmtypes.ScProxy {
	agentBytes(): wasmtypes.ScImmutableBytes {
		return new wasmtypes.ScImmutableBytes(this.proxy.root(sc.ParamAgentBytes));
	}

	agentString(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamAgentString));
	}

	scAgentID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamScAgentID));
	}
}

export class MutableCheckAgentIDParams extends wasmtypes.ScProxy {
	agentBytes(): wasmtypes.ScMutableBytes {
		return new wasmtypes.ScMutableBytes(this.proxy.root(sc.ParamAgentBytes));
	}

	agentString(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamAgentString));
	}

	scAgentID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamScAgentID));
	}
}

export class ImmutableStringMapOfStringArrayLengthParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableStringMapOfStringArrayLengthParams extends wasmtypes.ScProxy {
	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}

export class ImmutableStringMapOfStringArrayValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamIndex));
	}

	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableStringMapOfStringArrayValueParams extends wasmtypes.ScProxy {
	index(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamIndex));
	}

	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}

export class ImmutableStringMapOfStringMapValueParams extends wasmtypes.ScProxy {
	key(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamKey));
	}

	name(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
	}
}

export class MutableStringMapOfStringMapValueParams extends wasmtypes.ScProxy {
	key(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamKey));
	}

	name(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
	}
}
