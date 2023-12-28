package sync

import (
	"fmt"
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/network"
)

type ProtocolManager struct {
	syncConfig     config.SyncConfig
	blockchain     *chain.Blockchain
	peerManager    *p2p.PeerManager
	p2pServer      *p2p.P2PServer
	networkManager *p2p.NetworkManager
}

func NewProtocolManager(config config.Config, blockchain *chain.Blockchain) (*ProtocolManager, error) {
	peerManager := p2p.NewPeerManager(config.SyncConfig)
	networkManager, err := p2p.NewNetworkManager(config)
	if err != nil {
		return nil, err
	}

	p2pServer, err := p2p.NewP2PServer(config, peerManager)
	if err != nil {
		return nil, err
	}

	return &ProtocolManager{
		syncConfig:     config.SyncConfig,
		blockchain:     blockchain,
		p2pServer:      p2pServer,
		peerManager:    peerManager,
		networkManager: networkManager,
	}, nil
}

func (pm *ProtocolManager) Start() error {

	err := pm.networkManager.Start()
	if err != nil {
		return fmt.Errorf("network manager start error:%v", err)
	}

	err = pm.p2pServer.StartServer()
	if err != nil {
		return fmt.Errorf("p2p server start error:%v", err)
	}

	return nil
}

func (pm *ProtocolManager) BroadcastBlock(block *model.Block) error {
	return nil
}

func (pm *ProtocolManager) BroadcastTransaction(transaction *model.Transaction) error {
	return nil
}
