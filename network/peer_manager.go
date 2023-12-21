package p2p

import (
	"net"
	"sync"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/model"
)

type Peer struct {
	Conn net.Conn
	Key  string
}

func (p *Peer) GetKey() string {
	return p.Key
}

type PeerManager struct {
	mutex   sync.RWMutex
	peerMap map[string]*Peer
}

func NewPeerManager(config config.SyncConfig) *PeerManager {
	return &PeerManager{
		peerMap: make(map[string]*Peer),
	}
}

func (pm *PeerManager) BroadcastTransaction(tx *model.Transaction) error {
	return nil
}

func (pm *PeerManager) AddPeer(peer *Peer) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if _, ok := pm.peerMap[peer.Key]; !ok {
		pm.peerMap[peer.Key] = peer
	}
}
