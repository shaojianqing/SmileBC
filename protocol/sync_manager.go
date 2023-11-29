package protocol

import (
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
)

type SyncManager struct {
}

func NewSyncManager(config config.SyncConfig, blockChain *chain.Blockchain) (*SyncManager, error) {
	return nil, nil
}

func (s *SyncManager) StartSync() error {
	return nil
}
