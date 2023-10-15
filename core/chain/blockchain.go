package chain

import (
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/core/processor"
	"github.com/shaojianqing/smilebc/storage"
)

type Blockchain struct {
	chainDB   storage.Database
	processor stat.Processor
}

func NewBlockchain(chainDB storage.Database) *Blockchain {

	blockchain := &Blockchain{
		chainDB: chainDB,
	}

	processor := stat.NewStateProcessor(blockchain)
	blockchain.processor = processor

	return blockchain
}

func (bc *Blockchain) InsertTransactions(trxes []model.Transaction) error {
	return nil
}

func (bc *Blockchain) SendTransaction(trx model.Transaction) error {
	return nil
}
