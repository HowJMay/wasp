package accounts

import (
	"fmt"
	"math/big"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/collections"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/util"
	"golang.org/x/xerrors"
)

const (
	// map with all accounts listed
	varAccountsMap = "A"
	// map with on-chain totals listed
	varTotalAssetsMap = "T"
)

var (
	ErrNotEnoughFunds = xerrors.New("not enough funds")
	ErrBadAmount      = xerrors.New("bad native asset amount")
)

// getAccount each account is a map with the name of its controlling agentID.
// - nil key is balance of iotas uint64 8 bytes little-endian
// - iotago.NativeTokenID key is a big.Int balance of the native token
func getAccount(state kv.KVStore, agentID *iscp.AgentID) *collections.Map {
	return collections.NewMap(state, string(agentID.Bytes()))
}

func getAccountR(state kv.KVStoreReader, agentID *iscp.AgentID) *collections.ImmutableMap {
	return collections.NewMapReadOnly(state, string(agentID.Bytes()))
}

// getTotalAssetsAccount is an account with totals by token type
func getTotalAssetsAccount(state kv.KVStore) *collections.Map {
	return collections.NewMap(state, varTotalAssetsMap)
}

func getTotalAssetsAccountR(state kv.KVStoreReader) *collections.ImmutableMap {
	return collections.NewMapReadOnly(state, varTotalAssetsMap)
}

// getAccountsMap is a map which contains all non-empty accounts
func getAccountsMap(state kv.KVStore) *collections.Map {
	return collections.NewMap(state, varAccountsMap)
}

func getAccountsMapR(state kv.KVStoreReader) *collections.ImmutableMap {
	return collections.NewMapReadOnly(state, varAccountsMap)
}

// touchAccount ensures that only non-empty accounts are kept in the accounts map
func touchAccount(state kv.KVStore, account *collections.Map) {
	if account.Name() == varTotalAssetsMap {
		return
	}
	agentid := []byte(account.Name())
	accounts := getAccountsMap(state)
	if account.MustLen() == 0 {
		accounts.MustDelAt(agentid)
	} else {
		accounts.MustSetAt(agentid, []byte{0xFF})
	}
}

// tokenBalanceMutation structure for handling mutations of the on-chain accounts
type tokenBalanceMutation struct {
	balance *big.Int
	delta   *big.Int
}

// loadAccountMutations traverses the assets of interest in the account and collects values for further processing
func loadAccountMutations(account *collections.Map, assets *iscp.Assets) (uint64, uint64, map[iotago.NativeTokenID]tokenBalanceMutation) {
	if assets == nil {
		return 0, 0, nil
	}

	addIotas := assets.Iotas
	fromIotas := uint64(0)
	if v := account.MustGetAt(nil); v != nil {
		fromIotas = util.MustUint64From8Bytes(v)
	}

	tokenMutations := make(map[iotago.NativeTokenID]tokenBalanceMutation)
	zero := big.NewInt(0)
	for _, nt := range assets.Tokens {
		if nt.Amount.Cmp(zero) < 0 {
			panic(ErrBadAmount)
		}
		bal := big.NewInt(0)
		if v := account.MustGetAt(nt.ID[:]); v != nil {
			bal.SetBytes(v)
		}
		tokenMutations[nt.ID] = tokenBalanceMutation{
			balance: bal,
			delta:   nt.Amount,
		}
	}
	return fromIotas, addIotas, tokenMutations
}

// CreditToAccount brings new funds to the on chain ledger
func CreditToAccount(state kv.KVStore, agentID *iscp.AgentID, assets *iscp.Assets) {
	if assets == nil || (assets.Iotas == 0 && len(assets.Tokens) == 0) {
		return
	}
	account := getAccount(state, agentID)
	defer touchAccount(state, account)

	checkLedger(state, "CreditToAccount IN")
	defer checkLedger(state, "CreditToAccount OUT")

	creditToAccount(account, assets)
	creditToAccount(getTotalAssetsAccount(state), assets)
}

// creditToAccount adds assets to the internal account map
func creditToAccount(account *collections.Map, assets *iscp.Assets) {
	iotasBalance, iotasAdd, tokenMutations := loadAccountMutations(account, assets)
	if iotasAdd > 0 {
		account.MustSetAt(nil, util.Uint64To8Bytes(iotasBalance+iotasAdd))
	}
	for assetID, m := range tokenMutations {
		if util.IsZeroBigInt(m.delta) {
			continue
		}
		m.balance.Add(m.balance, m.delta)
		account.MustSetAt(assetID[:], m.balance.Bytes())
	}
}

// DebitFromAccount takes out assets balance the on chain ledger. If not enough it panics
func DebitFromAccount(state kv.KVStore, agentID *iscp.AgentID, assets *iscp.Assets) {
	if assets == nil || (assets.Iotas == 0 && len(assets.Tokens) == 0) {
		return
	}
	account := getAccount(state, agentID)
	defer touchAccount(state, account)

	checkLedger(state, "DebitFromAccount IN")
	defer checkLedger(state, "DebitFromAccount OUT")

	if !debitFromAccount(account, assets) {
		panic(ErrNotEnoughFunds)
	}
	if !debitFromAccount(getTotalAssetsAccount(state), assets) {
		panic("debitFromAccount: inconsistent ledger state")
	}
}

// debitFromAccount debits assets from the internal accounts map
func debitFromAccount(account *collections.Map, assets *iscp.Assets) bool {
	iotasBalance, iotasSub, tokenMutations := loadAccountMutations(account, assets)
	// check if enough
	if iotasBalance < iotasSub {
		return false
	}
	for _, m := range tokenMutations {
		if m.balance.Cmp(m.delta) < 0 {
			return false
		}
	}
	if iotasSub > 0 {
		if iotasBalance == iotasSub {
			account.MustDelAt(nil)
		} else {
			account.MustSetAt(nil, util.Uint64To8Bytes(iotasBalance-iotasSub))
		}
	}
	for id, m := range tokenMutations {
		m.balance.Add(m.balance, m.delta)
		if util.IsZeroBigInt(m.balance) {
			account.MustDelAt(id[:])
		} else {
			account.MustSetAt(id[:], m.balance.Bytes())
		}
	}
	return true
}

// MoveBetweenAccounts moves assets between on-chain accounts. Returns if it was a success (= enough funds in the source)
func MoveBetweenAccounts(state kv.KVStore, fromAgentID, toAgentID *iscp.AgentID, transfer *iscp.Assets) bool {
	checkLedger(state, "MoveBetweenAccounts.IN")
	defer checkLedger(state, "MoveBetweenAccounts.OUT")

	if fromAgentID.Equals(toAgentID) {
		// no need to move
		return true
	}
	// total assets doesn't change
	if !debitFromAccount(getAccount(state, fromAgentID), transfer) {
		return false
	}
	creditToAccount(getAccount(state, toAgentID), transfer)
	return true
}

func MustMoveBetweenAccounts(state kv.KVStore, fromAgentID, toAgentID *iscp.AgentID, assets *iscp.Assets) {
	if !MoveBetweenAccounts(state, fromAgentID, toAgentID, assets) {
		panic(ErrNotEnoughFunds)
	}
}

func getAccountAssetBalance(account *collections.ImmutableMap, assetID []byte) *big.Int {
	v := account.MustGetAt(assetID)
	if v == nil {
		return big.NewInt(0)
	}
	return new(big.Int).SetBytes(account.MustGetAt(assetID))
}

// GetIotaBalance return iota balance. 0 means it does not exist
func GetIotaBalance(state kv.KVStoreReader, agentID *iscp.AgentID) uint64 {
	return getIotaBalance(getAccountR(state, agentID))
}

func GetIotaBalanceTotal(state kv.KVStoreReader) uint64 {
	return getIotaBalance(getTotalAssetsAccountR(state))
}

func getIotaBalance(account *collections.ImmutableMap) uint64 {
	if v := account.MustGetAt(nil); v != nil {
		return util.MustUint64From8Bytes(v)
	}
	return 0
}

// GetNativeTokenBalance returns balance or nil if it does not exist
func GetNativeTokenBalance(state kv.KVStoreReader, agentID *iscp.AgentID, tokenID iotago.NativeTokenID) *big.Int {
	return getNativeTokenBalance(getAccountR(state, agentID), tokenID)
}

func GetNativeTokenBalanceTotal(state kv.KVStoreReader, tokenID iotago.NativeTokenID) *big.Int {
	return getNativeTokenBalance(getTotalAssetsAccountR(state), tokenID)
}

func getNativeTokenBalance(account *collections.ImmutableMap, tokenID iotago.NativeTokenID) *big.Int {
	ret := big.NewInt(0)
	if v := account.MustGetAt(tokenID[:]); v != nil {
		return ret.SetBytes(v)
	}
	return ret
}

// GetAssets returns all assets owned by agentID. Returns nil if account does not exist
func GetAssets(state kv.KVStoreReader, agentID *iscp.AgentID) *iscp.Assets {
	acc := getAccountR(state, agentID)
	ret := iscp.NewEmptyAssets()
	acc.MustIterate(func(k []byte, v []byte) bool {
		if len(k) == 0 {
			// iota
			ret.Iotas = util.MustUint64From8Bytes(v)
			return true
		}
		token := iotago.NativeToken{
			ID:     iscp.MustNativeTokenIDFromBytes(k),
			Amount: new(big.Int).SetBytes(v),
		}
		ret.Tokens = append(ret.Tokens, &token)
		return true
	})
	return ret
}

func getAccountsIntern(state kv.KVStoreReader) dict.Dict {
	ret := dict.New()
	getAccountsMapR(state).MustIterate(func(agentID []byte, val []byte) bool {
		ret.Set(kv.Key(agentID), []byte{})
		return true
	})
	return ret
}

func getAccountAssets(account *collections.ImmutableMap) *iscp.Assets {
	ret := iscp.NewEmptyAssets()
	account.MustIterate(func(idBytes []byte, val []byte) bool {
		if len(idBytes) == 0 {
			ret.Iotas = new(big.Int).SetBytes(val).Uint64()
			return true
		}
		token := iotago.NativeToken{
			ID:     iscp.MustNativeTokenIDFromBytes(idBytes),
			Amount: new(big.Int).SetBytes(val),
		}
		ret.Tokens = append(ret.Tokens, &token)
		return true
	})
	return ret
}

// GetAccountAssets returns all assets belonging to the agentID on the state
func GetAccountAssets(state kv.KVStoreReader, agentID *iscp.AgentID) (*iscp.Assets, bool) {
	account := getAccountR(state, agentID)
	if account.MustLen() == 0 {
		return nil, false
	}
	return getAccountAssets(account), true
}

func GetTotalAssets(state kv.KVStoreReader) *iscp.Assets {
	return getAccountAssets(getTotalAssetsAccountR(state))
}

// calcTotalAssets traverses the ledger and sums up all assets
func calcTotalAssets(state kv.KVStoreReader) *iscp.Assets {
	ret := iscp.NewEmptyAssets()

	getAccountsMapR(state).MustIterateKeys(func(key []byte) bool {
		agentID, err := iscp.AgentIDFromBytes(key)
		if err != nil {
			panic(xerrors.Errorf("calcTotalAssets: %w", err))
		}
		accBalances := getAccountAssets(getAccountR(state, agentID))
		ret.Add(accBalances)
		return true
	})
	return ret
}

func checkLedger(state kv.KVStore, checkpoint string) {
	a := GetTotalAssets(state)
	c := calcTotalAssets(state)
	if !a.Equals(c) {
		panic(fmt.Sprintf("inconsistent on-chain account ledger @ checkpoint '%s'\n total assets: %s\ncalc total: %s\n",
			checkpoint, a.String(), c.String()))
	}
}

func getAccountBalanceDict(account *collections.ImmutableMap) dict.Dict {
	return getAccountAssets(account).ToDict()
}

// DecodeBalances TODO move to iscp package
func DecodeBalances(balances dict.Dict) (*iscp.Assets, error) {
	return iscp.NewAssetsFromDict(balances)
}

// MaxAssumedNonce is maintained for each L1 address with the purpose of replay protection of off-ledger requests

const prefixMaxAssumedNonceKey = "nonce_"

func GetMaxAssumedNonce(state kv.KVStoreReader, address iotago.Address) uint64 {
	nonce, err := codec.DecodeUint64(
		state.MustGet(prefixMaxAssumedNonceKey+kv.Key(iscp.BytesFromAddress(address))),
		0,
	)
	if err != nil {
		panic(xerrors.Errorf("GetMaxAssumedNonce: %w", err))
	}
	return nonce
}

func SaveMaxAssumedNonce(state kv.KVStore, address iotago.Address, nonce uint64) {
	next := GetMaxAssumedNonce(state, address) + 1
	if nonce > next {
		next = nonce
	}
	state.Set(prefixMaxAssumedNonceKey+kv.Key(iscp.BytesFromAddress(address)), codec.EncodeUint64(next))
}

//
//func SetAssetsUtxoIndices(state kv.KVStore, stateIndex uint32, tokenUtxoIndices []iotago.NativeTokenID) {
//	mapping := getUtxoMapping(state)
//	for index, assetID := range tokenUtxoIndices {
//		entry := codec.EncodeUint16(uint16(index))
//		entry = append(entry, codec.EncodeUint32(stateIndex)...)
//		mapping.MustSetAt(assetID[:], entry)
//	}
//}
//
//func GetUtxoForAsset(state kv.KVStore, id iotago.NativeTokenID) (stateIndex uint32, outputIndex uint16, err error) {
//	mapping := getUtxoMapping(state)
//	entry := mapping.MustGetAt(id[:])
//	outputIndex, err = codec.DecodeUint16(entry[:2])
//	if err != nil {
//		return 0, 0, err
//	}
//	stateIndex, err = codec.DecodeUint32(entry[2:])
//	if err != nil {
//		return 0, 0, err
//	}
//	return stateIndex, outputIndex, nil
//}
//
//func removeUTXOMapping(state kv.KVStore, id iotago.NativeTokenID) {
//	mapping := getUtxoMapping(state)
//	mapping.DelAt(id[:]) //nolint:errcheck // No need to check this error
//}
