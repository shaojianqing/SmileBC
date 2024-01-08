package chain

import (
	core "github.com/shaojianqing/smilebc/core/general"
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/storage"
)

type Blockchain struct {
	ChainDB   storage.Database
	Processor core.Processor
}

func NewBlockchain(chainDB storage.Database) *Blockchain {
	blockchain := &Blockchain{
		ChainDB: chainDB,
	}

	return blockchain
}

func (bc *Blockchain) InsertTransactions(transactions []model.Transaction) error {
	return nil
}

func (bc *Blockchain) SendTransaction(transaction model.Transaction) error {
	return nil
}

func (bc *Blockchain) GetGenesisBlock() *model.Block {
	return nil
}
