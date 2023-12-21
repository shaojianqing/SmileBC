package sync

import (
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	"github.com/shaojianqing/smilebc/network"
)

type ProtocolManager struct {
	syncConfig     config.SyncConfig
	blockchain     *chain.Blockchain
	peerManager    *p2p.PeerManager
	networkManager *p2p.NetworkManager
}

func NewProtocolManager(config config.Config, blockchain *chain.Blockchain) (*ProtocolManager, error) {
	peerManager := p2p.NewPeerManager(config.SyncConfig)
	networkManager, err := p2p.NewNetworkManager(config)
	if err != nil {
		return nil, err
	}

	return &ProtocolManager{
		syncConfig:     config.SyncConfig,
		blockchain:     blockchain,
		peerManager:    peerManager,
		networkManager: networkManager,
	}, nil
}

func (pm *ProtocolManager) Start() error {
	err := pm.networkManager.Start()
	if err != nil {
		return err
	}
	return nil
}
