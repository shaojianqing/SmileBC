package p2p

import "github.com/syndtr/goleveldb/leveldb"

type NodeDB struct {
	db *leveldb.DB
}

func NewNodeDB(dbPath string, nodeID NodeID) *NodeDB {
	return &NodeDB{}
}
