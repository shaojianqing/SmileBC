package protocol

import "github.com/shaojianqing/smilebc/config"

type SyncManager struct {
}

func NewSyncManager(config config.SyncConfig) (*SyncManager, error) {
	return nil, nil
}

func (s *SyncManager) StartSync() error {
	return nil
}
