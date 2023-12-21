package stat

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/core/trie"
	"github.com/shaojianqing/smilebc/crypto"
	"github.com/shaojianqing/smilebc/storage"
)

type StateDB struct {
	root     common.Hash
	trie     *trie.TrieDB
	journal  *StateJournal
	database storage.Database

	stateObjectMap   map[common.Address]*StateObject
	stateObjectDirty map[common.Address]*StateObject

	lockControl sync.Mutex

	databaseError error
}

func NewStateDB(root common.Hash, database storage.Database) (*StateDB, error) {

	journal := NewStateJournal()
	trie, err := trie.NewTrieDB(root, database)
	if err != nil {
		return nil, fmt.Errorf("fail to create stateDB instance, err:%w", err)
	}

	stateDB := StateDB{
		root:     root,
		trie:     trie,
		journal:  journal,
		database: database,

		stateObjectMap:   make(map[common.Address]*StateObject),
		stateObjectDirty: make(map[common.Address]*StateObject),
	}
	return &stateDB, nil
}

func (db *StateDB) GetAccount(address common.Address) *StateObject {
	db.lockControl.Lock()
	object := db.stateObjectMap[address]
	db.lockControl.Unlock()
	if object != nil {
		if object.deleted {
			return nil
		}
		return object
	}

	// Get stateAccount from trieDB and it still comes from storage database finally.
	value, err := db.trie.TryGet(address[:])
	if err != nil {
		db.databaseError = err
		return nil
	}

	stateAccount := StateAccount{}
	if err := json.Unmarshal(value, &stateAccount); err != nil {
		log.Printf("fail to unmarshal the state account, address:%s, err:%v", address, err)
		return nil
	}

	// Create state object from stateAccount and stateDB instance
	stateObject := NewStateObject(db, address, stateAccount)

	db.lockControl.Lock()
	db.stateObjectMap[address] = stateObject
	db.lockControl.Unlock()

	return stateObject
}

func (db *StateDB) CreateAccount(address common.Address) *StateObject {
	stateObject := db.GetAccount(address)
	if stateObject == nil || stateObject.deleted {
		stateObject = db.createStateObject(address)
	}
	return stateObject
}

func (db *StateDB) GetOrCreateAccount(address common.Address) *StateObject {
	stateObject := db.GetAccount(address)
	if stateObject == nil || stateObject.deleted {
		stateObject = db.createStateObject(address)
	}
	return stateObject
}

func (db *StateDB) createStateObject(address common.Address) *StateObject {
	prevStateObject := db.GetAccount(address)

	newStateAccount := StateAccount{
		Nonce:   0,
		Balance: *big.NewInt(0),
	}

	newStateObject := NewStateObject(db, address, newStateAccount)
	if prevStateObject == nil {
		db.journal.append(&CreateAccountChangeItem{
			account: address,
		})
	} else {
		db.journal.append(&ResetAccountChangeItem{
			account:         address,
			prevStateObject: prevStateObject,
		})
	}
	db.lockControl.Lock()
	db.stateObjectMap[address] = newStateObject
	db.stateObjectDirty[address] = newStateObject
	db.lockControl.Unlock()

	return newStateObject
}

func (db *StateDB) AddBalance(address common.Address, amount big.Int) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.AddBalance(amount)
	}
}

func (db *StateDB) SubBalance(address common.Address, amount big.Int) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.SubBalance(amount)
	}
}

func (db *StateDB) GetBalance(address common.Address) big.Int {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		return stateObject.GetBalance()
	}
	return *big.NewInt(0)
}

func (db *StateDB) SetBalance(address common.Address, amount big.Int) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.SetBalance(amount)
	}
}

func (db *StateDB) GetNonce(address common.Address) uint64 {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		return stateObject.GetNonce()
	}
	return 0
}

func (db *StateDB) SetNonce(address common.Address, nonce uint64) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.SetNonce(nonce)
	}
}

func (db *StateDB) SetCode(address common.Address, code common.Code) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		hash := crypto.Keccak256Hash(code)
		stateObject.SetCode(hash, code)
	}
}

func (db *StateDB) GetCode(address common.Address) common.Code {
	if stateObject := db.GetAccount(address); stateObject != nil {
		return stateObject.GetCode()
	}
	return nil
}

func (db *StateDB) SetState(address common.Address, key common.Hash, value common.Hash) {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.SetState(key, value)
	}
}

func (db *StateDB) GetContractCode(codeHash common.Hash) (common.Code, error) {
	value, err := db.database.Get(codeHash[:])
	return common.Code(value), err
}

func (db *StateDB) Suicide(address common.Address) bool {
	if stateObject := db.GetOrCreateAccount(address); stateObject != nil {
		stateObject.Suicide()
		return true
	}
	return false
}

func (db *StateDB) Snapshot() uint64 {
	return 0
}

func (db *StateDB) RevertToSnapshot(versionId uint64) {

}

func (db *StateDB) Finalise() bool {
	return true
}

func (db *StateDB) Commit() (common.Hash, error) {

	return common.Hash{}, nil
}
