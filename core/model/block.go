package model

import "github.com/shaojianqing/smilebc/common"

type BlockHeader struct {
	BlockNumber uint64

	ParentHash common.Hash
	RootHash   common.Hash
	CoinBase   common.Address

	GasLimit uint64
	GasUsed  uint64

	Extra common.Content

	CreateTime uint64
}

type Block struct {
	BlockHeader  *BlockHeader
	Transactions []*Transaction
}

func (b *Block) GetHash() common.Hash {
	return b.BlockHeader.RootHash
}
