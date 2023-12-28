package p2p

import (
	"net"
	"time"
)

const (
	DefaultDialTimeout = 15 * time.Second
)

type PeerDialer struct {
	dialer    net.Dialer
	p2pServer *P2PServer
}

func NewPeerDialer() *PeerDialer {

	return &PeerDialer{}
}

func (pd *PeerDialer) StartDialing() error {
	return nil
}

func (pd *PeerDialer) dialing(destNode *Node) error {
	destAddress := &net.TCPAddr{
		IP:   destNode.IP,
		Port: int(destNode.TCP),
	}
	conn, err := pd.dialer.Dial(TCP, destAddress.String())
	if err != nil {
		return err
	}
	go pd.p2pServer.setupConnection(conn, destNode)
	return nil
}
