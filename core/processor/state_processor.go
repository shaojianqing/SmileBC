package stat

import (
	"github.com/shaojianqing/smilebc/core/chain"
	core "github.com/shaojianqing/smilebc/core/general"
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/core/statement"
)

type StateProcessor struct {
	blockchain *chain.Blockchain
}

func NewStateProcessor(blockchain *chain.Blockchain) core.Processor {
	return &StateProcessor{
		blockchain: blockchain,
	}
}

func (p *StateProcessor) Process(block *model.Block) *stat.StateResult {

	_, err := stat.NewStateDB(block.BlockHeader.RootHash, p.blockchain.ChainDB)
	if err != nil {

	}

	return nil
}
