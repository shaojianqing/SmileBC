package stat

import (
	"fmt"
	"github.com/shaojianqing/smilebc/crypto"
	"math/big"

	"github.com/shaojianqing/smilebc/common"
)

var EmptyCodeHash = crypto.Keccak256Hash(nil)

type StateObject struct {
	Address  common.Address
	AddrHash common.Hash
	Code     common.Code
	Account  StateAccount

	DatabaseError error

	StateDB *StateDB

	dirtyCode bool
	suicided  bool
	deleted   bool
}

func NewStateObject(stateDB *StateDB, address common.Address, account StateAccount) *StateObject {
	stateObject := &StateObject{
		StateDB:  stateDB,
		Address:  address,
		Account:  account,
		AddrHash: crypto.Keccak256Hash(address[:]),
	}
	return stateObject
}

func (so *StateObject) GetAddress() common.Address {
	return so.Address
}

func (so *StateObject) AddBalance(amount big.Int) {
	prevBalance := so.GetBalance()
	newBalance := prevBalance.Add(&prevBalance, &amount)
	so.SetBalance(*newBalance)
}

func (so *StateObject) SubBalance(amount big.Int) {
	prevBalance := so.GetBalance()
	newBalance := prevBalance.Sub(&prevBalance, &amount)
	so.SetBalance(*newBalance)
}

func (so *StateObject) GetBalance() big.Int {
	return so.Account.Balance
}

func (so *StateObject) SetBalance(balance big.Int) {
	prevBalance := so.GetBalance()
	so.StateDB.journal.append(&BalanceChangeItem{
		account:  so.Address,
		previous: prevBalance,
	})
	so.Account.Balance = balance
}

func (so *StateObject) GetNonce() uint64 {
	return so.Account.Nonce
}

func (so *StateObject) SetNonce(nonce uint64) {
	prevNonce := so.GetNonce()
	so.StateDB.journal.append(&NonceChangeItem{
		account:  so.Address,
		previous: prevNonce,
	})
	so.Account.Nonce = nonce
}

func (so *StateObject) GetCode() common.Code {
	if so.Code != nil {
		return so.Code
	}

	// If the codeHash value is equal to empty codeHash, it means this state
	// object does not contain any code, return nil directly.
	if common.HashEqual(so.GetCodeHash(), EmptyCodeHash) {
		return nil
	}

	// Try to get contract code from levelDB with the codeHash as the key and parameter.
	// If getting the code content, set the code field in state object as the cache for temporary.
	code, err := so.StateDB.GetContractCode(so.GetCodeHash())
	if err != nil {
		so.DatabaseError = fmt.Errorf("can not get contract code, codeHash:%s,err:%w", so.GetCodeHash(), err)
	}
	so.Code = code
	return code
}

func (so *StateObject) GetCodeHash() common.Hash {
	return so.Account.CodeHash
}

func (so *StateObject) SetCode(codeHash common.Hash, code common.Code) {
	prevCodeData := so.GetCode()
	prevCodeHash := so.GetCodeHash()

	so.StateDB.journal.append(&CodeChangeItem{
		account:  so.Address,
		codeHash: prevCodeHash,
		code:     prevCodeData,
	})

	so.Code = code
	so.Account.CodeHash = codeHash
}

func (so *StateObject) GetState(key common.Hash) common.Hash {
	return common.Hash{}
}

func (so *StateObject) SetState(key, value common.Hash) {

}

func (so *StateObject) Suicide() {

	so.StateDB.journal.append(&SuicideChangeItem{
		account:     so.Address,
		prevSuicide: so.suicided,
		prevBalance: so.Account.Balance,
	})

	so.Account.Balance = *big.NewInt(0)
	so.suicided = true
}
