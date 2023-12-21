package trie

import (
	"fmt"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/storage"
)

type TrieDB struct {
	rootNode Node
	rootHash common.Hash
	database storage.Database
}

func NewTrieDB(root common.Hash, database storage.Database) (*TrieDB, error) {
	tr := &TrieDB{
		rootHash: root,
		database: database,
	}
	rootNode, err := tr.resolveHash(root[:])
	if err != nil {
		return nil, err
	}
	tr.rootNode = rootNode
	return tr, nil
}

func (tr *TrieDB) TryGet(key common.Key) (common.Data, error) {
	return nil, nil
}

func (tr *TrieDB) TryUpdate(key storage.Key, value storage.Value) error {
	return nil
}

func (tr *TrieDB) TryDelete(key storage.Key) error {
	return nil
}

func (tr *TrieDB) CommitTo(database storage.Database) (common.Hash, error) {
	return common.Hash{}, nil
}

func (tr *TrieDB) resolveHash(key common.Key) (Node, error) {
	data, err := tr.database.Get(storage.Key(key))
	if err != nil || data == nil {
		return nil, fmt.Errorf("missing system with key:%s", key)
	}
	node, err := Parse(common.Data(data))
	if err != nil {
		return nil, fmt.Errorf("fail to parse data:%s", data)
	}

	return node, nil
}
