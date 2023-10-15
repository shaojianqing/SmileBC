package stat

import (
	"github.com/shaojianqing/smilebc/core/chain"
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/core/statement"
)

type Processor interface {
	Process(block *model.Block) *stat.StateResult
}

type StateProcessor struct {
	blockchain *chain.Blockchain
}

func NewStateProcessor(blockchain *chain.Blockchain) *StateProcessor {
	return &StateProcessor{
		blockchain: blockchain,
	}
}

func (p *StateProcessor) Process(block *model.Block) *stat.StateResult {
	return nil
}
