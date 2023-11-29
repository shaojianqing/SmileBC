package core

import (
	"github.com/shaojianqing/smilebc/core/model"
	stat "github.com/shaojianqing/smilebc/core/statement"
)

type Processor interface {
	Process(block *model.Block) *stat.StateResult
}
