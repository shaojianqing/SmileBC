package p2p

import (
	"sync"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/model"
)

type ProtocolMessageHandler interface {
	HandlePeerMessage(peer *Peer) error
}

type PeerManager struct {
	mutex   sync.RWMutex
	peerMap map[NodeID]*Peer
	handler ProtocolMessageHandler
}

func NewPeerManager(config config.SyncConfig) *PeerManager {
	return &PeerManager{
		peerMap: make(map[NodeID]*Peer),
	}
}

func (pm *PeerManager) BroadcastTransaction(tx *model.Transaction) error {
	return nil
}

func (pm *PeerManager) AddPeer(peer *Peer) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if success := pm.addPeer(peer); success {
		// Once the peer is added successfully, start to
		// handle messages from the peer immediately in a loop
		if err := pm.handler.HandlePeerMessage(peer); err != nil {
			// If error raises when handling peer message, directly
			// remove the corresponding peer from peer manager
			pm.RemovePeer(peer.GetID())
		}
	}
}

func (pm *PeerManager) RemovePeer(peerId NodeID) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if _, ok := pm.peerMap[peerId]; ok {
		delete(pm.peerMap, peerId)
	}
}

func (pm *PeerManager) addPeer(peer *Peer) bool {
	if _, ok := pm.peerMap[peer.id]; !ok {
		pm.peerMap[peer.id] = peer
		return true
	}
	return false
}

func (pm *PeerManager) SetProtocolHandler(messageHandler ProtocolMessageHandler) {
	pm.handler = messageHandler
}
