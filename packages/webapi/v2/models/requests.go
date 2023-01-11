package models

import (
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/kv/dict"
)

type OffLedgerRequest struct {
	ChainID string `json:"chainId" swagger:"desc(The chain id)"`
	Request string `json:"request" swagger:"desc(Offledger Request (Hex))"`
}

type ContractCallViewRequest struct {
	ChainID       string        `json:"chainId" swagger:"desc(The chain id)"`
	ContractName  string        `json:"contractName" swagger:"desc(The contract name)"`
	ContractHName isc.Hname     `json:"contractHName" swagger:"desc(The contract name as HName)"`
	FunctionName  string        `json:"functionName" swagger:"desc(The function name)"`
	FunctionHName isc.Hname     `json:"functionHName" swagger:"desc(The function name as HName)"`
	Arguments     dict.JSONDict `json:"arguments" swagger:"desc(Encoded arguments to be passed to the function)"`
}