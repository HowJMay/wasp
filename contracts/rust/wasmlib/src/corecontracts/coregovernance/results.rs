// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coregovernance::*;
use crate::host::*;

pub struct ArrayOfImmutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableBytes {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetAllowedStateControllerAddressesResults {
    pub(crate) id: i32,
}

impl ImmutableGetAllowedStateControllerAddressesResults {
    pub fn allowed_state_controller_addresses(&self) -> ArrayOfImmutableBytes {
        let arr_id = get_object_id(self.id, RESULT_ALLOWED_STATE_CONTROLLER_ADDRESSES.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBytes { obj_id: arr_id }
    }
}

pub struct ArrayOfMutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableBytes {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetAllowedStateControllerAddressesResults {
    pub(crate) id: i32,
}

impl MutableGetAllowedStateControllerAddressesResults {
    pub fn allowed_state_controller_addresses(&self) -> ArrayOfMutableBytes {
        let arr_id = get_object_id(self.id, RESULT_ALLOWED_STATE_CONTROLLER_ADDRESSES.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBytes { obj_id: arr_id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetChainInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetChainInfoResults {
    pub fn chain_id(&self) -> ScImmutableChainID {
        ScImmutableChainID::new(self.id, RESULT_CHAIN_ID.get_key_id())
    }

    pub fn chain_owner_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, RESULT_CHAIN_OWNER_ID.get_key_id())
    }

    pub fn default_owner_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, RESULT_DEFAULT_OWNER_FEE.get_key_id())
    }

    pub fn default_validator_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, RESULT_DEFAULT_VALIDATOR_FEE.get_key_id())
    }

    pub fn description(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, RESULT_DESCRIPTION.get_key_id())
    }

    pub fn fee_color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, RESULT_FEE_COLOR.get_key_id())
    }

    pub fn max_blob_size(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, RESULT_MAX_BLOB_SIZE.get_key_id())
    }

    pub fn max_event_size(&self) -> ScImmutableInt16 {
        ScImmutableInt16::new(self.id, RESULT_MAX_EVENT_SIZE.get_key_id())
    }

    pub fn max_events_per_req(&self) -> ScImmutableInt16 {
        ScImmutableInt16::new(self.id, RESULT_MAX_EVENTS_PER_REQ.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetChainInfoResults {
    pub(crate) id: i32,
}

impl MutableGetChainInfoResults {
    pub fn chain_id(&self) -> ScMutableChainID {
        ScMutableChainID::new(self.id, RESULT_CHAIN_ID.get_key_id())
    }

    pub fn chain_owner_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, RESULT_CHAIN_OWNER_ID.get_key_id())
    }

    pub fn default_owner_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, RESULT_DEFAULT_OWNER_FEE.get_key_id())
    }

    pub fn default_validator_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, RESULT_DEFAULT_VALIDATOR_FEE.get_key_id())
    }

    pub fn description(&self) -> ScMutableString {
        ScMutableString::new(self.id, RESULT_DESCRIPTION.get_key_id())
    }

    pub fn fee_color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, RESULT_FEE_COLOR.get_key_id())
    }

    pub fn max_blob_size(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, RESULT_MAX_BLOB_SIZE.get_key_id())
    }

    pub fn max_event_size(&self) -> ScMutableInt16 {
        ScMutableInt16::new(self.id, RESULT_MAX_EVENT_SIZE.get_key_id())
    }

    pub fn max_events_per_req(&self) -> ScMutableInt16 {
        ScMutableInt16::new(self.id, RESULT_MAX_EVENTS_PER_REQ.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetFeeInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetFeeInfoResults {
    pub fn fee_color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, RESULT_FEE_COLOR.get_key_id())
    }

    pub fn owner_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, RESULT_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, RESULT_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetFeeInfoResults {
    pub(crate) id: i32,
}

impl MutableGetFeeInfoResults {
    pub fn fee_color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, RESULT_FEE_COLOR.get_key_id())
    }

    pub fn owner_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, RESULT_OWNER_FEE.get_key_id())
    }

    pub fn validator_fee(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, RESULT_VALIDATOR_FEE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetMaxBlobSizeResults {
    pub(crate) id: i32,
}

impl ImmutableGetMaxBlobSizeResults {
    pub fn max_blob_size(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, RESULT_MAX_BLOB_SIZE.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetMaxBlobSizeResults {
    pub(crate) id: i32,
}

impl MutableGetMaxBlobSizeResults {
    pub fn max_blob_size(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, RESULT_MAX_BLOB_SIZE.get_key_id())
    }
}