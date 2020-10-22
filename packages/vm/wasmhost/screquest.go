package wasmhost

type ScRequest struct {
	MapObject
}

func (o *ScRequest) GetInt(keyId int32) int64 {
	switch keyId {
	case KeyTimestamp:
		return o.vm.ctx.GetTimestamp()
	}
	return o.MapObject.GetInt(keyId)
}

func (o *ScRequest) GetObjectId(keyId int32, typeId int32) int32 {
	return o.GetMapObjectId(keyId, typeId, map[int32]MapObjDesc{
		KeyColors:  {OBJTYPE_INT_ARRAY, func() WaspObject { return &ScColors{requestOnly: true} }},
		KeyBalance: {OBJTYPE_MAP, func() WaspObject { return &ScBalance{requestOnly: true} }},
		KeyParams:  {OBJTYPE_MAP, func() WaspObject { return &ScRequestParams{} }},
	})
}

func (o *ScRequest) GetString(keyId int32) string {
	switch keyId {
	case KeyAddress:
		return o.vm.ctx.AccessRequest().Sender().String()
	case KeyHash: // used to determine minted color
		id := o.vm.ctx.AccessRequest().ID()
		return id.TransactionId().String()
	case KeyId:
		id := o.vm.ctx.AccessRequest().ID()
		return id.String()
	}
	return o.MapObject.GetString(keyId)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestParams struct {
	MapObject
}

func (o *ScRequestParams) GetBytes(keyId int32) []byte {
	key := o.vm.GetKey(keyId)
	value, _ := o.vm.ctx.AccessRequest().Args().Get(key)
	return value
}

func (o *ScRequestParams) GetInt(keyId int32) int64 {
	key := o.vm.GetKey(keyId)
	value, _, _ := o.vm.ctx.AccessRequest().Args().GetInt64(key)
	return value
}

func (o *ScRequestParams) GetString(keyId int32) string {
	key := o.vm.GetKey(keyId)
	value, _, _ := o.vm.ctx.AccessRequest().Args().GetString(key)
	return value
}
