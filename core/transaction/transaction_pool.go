package trx

import (
	"fmt"
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

func (pool *TransactionPool) AddTransactions(trxList []*model.Transaction) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

}

func (pool *TransactionPool) GetTransaction() []*model.Transaction {
	return nil
}

func (pool *TransactionPool) addTransaction(trx *model.Transaction) error {
	if pool.pendingTransactionMap[trx.GetHash()] != nil {
		return fmt.Errorf("transaction has existed already, trx hash:%s", trx.GetHash())
	}

	if err := pool.validateTransaction(trx); err != nil {
		return err
	}
	return nil
}

func (pool *TransactionPool) validateTransaction(trx *model.Transaction) error {
	return nil
}
