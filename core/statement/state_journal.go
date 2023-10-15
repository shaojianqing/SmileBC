package stat

import (
	"math/big"

	"github.com/shaojianqing/smilebc/common"
)

type StateJournal struct {
	changeItems   []JournalItem
	dirtyCountMap map[common.Address]uint
}

type JournalItem interface {
	revert(db *StateDB)

	dirty() common.Address
}

func NewStateJournal() *StateJournal {
	return &StateJournal{
		changeItems:   make([]JournalItem, 0),
		dirtyCountMap: make(map[common.Address]uint),
	}
}

func (j *StateJournal) append(item JournalItem) {
	j.changeItems = append(j.changeItems, item)
	j.dirtyCountMap[item.dirty()]++
}

type CodeChangeItem struct {
	account  common.Address
	codeHash common.Hash
	code     common.Code
}

func (ch *CodeChangeItem) revert(db *StateDB) {
	db.GetAccount(ch.account).Account.CodeHash = ch.codeHash
	db.GetAccount(ch.account).Code = ch.code
}

func (ch *CodeChangeItem) dirty() common.Address {
	return ch.account
}

type NonceChangeItem struct {
	account  common.Address
	previous uint64
}

func (ch *NonceChangeItem) revert(db *StateDB) {
	db.GetAccount(ch.account).Account.Nonce = ch.previous
}

func (ch *NonceChangeItem) dirty() common.Address {
	return ch.account
}

type BalanceChangeItem struct {
	account  common.Address
	previous big.Int
}

func (ch *BalanceChangeItem) revert(db *StateDB) {
	db.GetAccount(ch.account).Account.Balance = ch.previous
}

func (ch *BalanceChangeItem) dirty() common.Address {
	return ch.account
}

type CreateAccountChangeItem struct {
	account common.Address
}

func (ch *CreateAccountChangeItem) revert(db *StateDB) {
	delete(db.stateObjectMap, ch.account)
	delete(db.stateObjectDirty, ch.account)
}

func (ch *CreateAccountChangeItem) dirty() common.Address {
	return ch.account
}

type ResetAccountChangeItem struct {
	account         common.Address
	prevStateObject *StateObject
}

func (ch *ResetAccountChangeItem) revert(db *StateDB) {
	db.stateObjectMap[ch.account] = ch.prevStateObject
	delete(db.stateObjectDirty, ch.account)
}

func (ch *ResetAccountChangeItem) dirty() common.Address {
	return common.Address{}
}

type SuicideChangeItem struct {
	account     common.Address
	prevBalance big.Int
	prevSuicide bool
}

func (ch *SuicideChangeItem) revert(db *StateDB) {
	db.GetAccount(ch.account).Account.Balance = ch.prevBalance
	db.GetAccount(ch.account).suicided = ch.prevSuicide
	delete(db.stateObjectDirty, ch.account)
}

func (ch *SuicideChangeItem) dirty() common.Address {
	return ch.account
}
