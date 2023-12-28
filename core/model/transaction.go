package model

import "github.com/shaojianqing/smilebc/common"

type Transaction struct {
}

func (tx *Transaction) GetHash() common.Hash {
	return common.Hash{}
}
