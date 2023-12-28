package trx

import (
	"sync"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/core/model"
)

type TransactionPool struct {
	mutex sync.RWMutex

	pendingTransactionMap map[common.Hash]*model.Transaction
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{}
}
