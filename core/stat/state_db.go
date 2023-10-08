package stat

import (
	"github.com/shaojianqing/smilebc/core/trie"
	"github.com/shaojianqing/smilebc/storage"
	"math/big"

	"github.com/shaojianqing/smilebc/common"
)

type StateDB struct {
	chainDB storage.Database
	trieDB  *trie.Database
	journal *Journal

	stateObjectMap   map[common.Address]*StateObject
	stateObjectDirty map[common.Address]*StateObject
}

func NewStateDB(chainDB storage.Database, trieDB *trie.Database) (*StateDB, error) {

	journal := NewJournal()

	stateDB := StateDB{
		chainDB: chainDB,
		trieDB:  trieDB,
		journal: journal,

		stateObjectMap:   make(map[common.Address]*StateObject),
		stateObjectDirty: make(map[common.Address]*StateObject),
	}
	return &stateDB, nil
}

func (db *StateDB) GetAccount(address common.Address) *StateObject {
	return nil
}

func (db *StateDB) CreateAccount(address common.Address) *StateObject {
	return nil
}

func (db *StateDB) AddBalance(address common.Address, amount big.Int) {

}

func (db *StateDB) SubBalance(address common.Address, amount big.Int) {

}

func (db *StateDB) GetBalance(address common.Address) big.Int {
	return *big.NewInt(0)
}

func (db *StateDB) SetBalance(address common.Address, amount big.Int) {

}

func (db *StateDB) GetNonce(address common.Address) uint64 {
	return 0
}

func (db *StateDB) SetNonce(address common.Address, nonce uint64) {

}

func (db *StateDB) SetCode(address common.Address, code common.Code) {

}

func (db *StateDB) GetCode(address common.Address) common.Code {
	return common.Code{}
}

func (db *StateDB) SetState(address common.Address, key common.Hash, value common.Hash) {

}

func (db *StateDB) Suicide(address common.Address) bool {
	return true
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
