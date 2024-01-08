package proto

import (
	"fmt"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	"github.com/shaojianqing/smilebc/core/model"
	"github.com/shaojianqing/smilebc/network"
)

type ProtocolManager struct {
	config          config.Config
	blockchain      *chain.Blockchain
	peerManager     *p2p.PeerManager
	p2pServer       *p2p.P2PServer
	networkManager  *p2p.NetworkManager
	protocolHandler *ProtocolHandler
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

	protocolHandler, err := NewProtocolHandler()
	if err != nil {
		return nil, err
	}

	protocolManager := &ProtocolManager{
		config:          config,
		blockchain:      blockchain,
		p2pServer:       p2pServer,
		peerManager:     peerManager,
		networkManager:  networkManager,
		protocolHandler: protocolHandler,
	}
	peerManager.SetProtocolHandler(protocolManager)
	protocolHandler.SetProtocolManager(protocolManager)

	return protocolManager, nil
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

func (pm *ProtocolManager) HandlePeerMessage(peer *p2p.Peer) error {
	commonConfig := pm.config.CommonConfig
	genesisBlock := pm.blockchain.GetGenesisBlock()
	if err := peer.PerformHandshake(commonConfig.Network,
		commonConfig.Version, genesisBlock.BlockHeader.RootHash); err != nil {
		return fmt.Errorf("peer performs handshake error:%v", err)
	}

	for {
		// We handle all the peer message in a loop in separate goroutine, and
		// it returns error directly in case of error raising
		if err := pm.protocolHandler.Handle(peer); err != nil {
			return fmt.Errorf("protocol handle peer message error:%v", err)
		}
	}
}

func (pm *ProtocolManager) BroadcastBlock(block *model.Block) error {
	return nil
}

func (pm *ProtocolManager) BroadcastTransaction(transaction *model.Transaction) error {
	return nil
}
