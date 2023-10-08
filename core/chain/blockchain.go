package chain

import (
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/core/stat"
	"github.com/shaojianqing/smilebc/storage"
)

type Blockchain struct {
	chainDB storage.Database
	stateDB *stat.StateDB
}

func NewBlockchain(chainDB storage.Database, stateDB *stat.StateDB) *Blockchain {
	return &Blockchain{
		chainDB: chainDB,
		stateDB: stateDB,
	}
}

func (bc *Blockchain) InsertTransactions(trxes []model.Transaction) error {
	return nil
}

func (bc *Blockchain) SendTransaction(trx model.Transaction) error {
	return nil
}
