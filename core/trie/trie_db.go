package trie

import (
	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/storage"
)

type Database struct {
}

func NewTrieDB(database storage.Database) (*Database, error) {
	return nil, nil
}

func (db *Database) TryGet(key storage.Key) (storage.Value, error) {
	return nil, nil
}

func (db *Database) TryUpdate(key storage.Key, value storage.Value) error {
	return nil
}

func (db *Database) TryDelete(key storage.Key) error {
	return nil
}

func (db *Database) CommitTo(database storage.Database) (common.Hash, error) {
	return common.Hash{}, nil
}
