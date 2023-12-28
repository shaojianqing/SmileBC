package p2p

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/crypto"
	"github.com/shaojianqing/smilebc/crypto/ecies"
	"net"
	"sync"
)

const (
	EncryptEchoMessage = "Received AES Encrypt Key"
)

type P2PServer struct {
	running      bool
	listenerAddr string

	mutex sync.Mutex

	listener net.Listener

	selfNodeID NodeID
	privateKey *ecdsa.PrivateKey

	peerDialer  *PeerDialer
	peerManager *PeerManager
}

func NewP2PServer(config config.Config, peerManager *PeerManager) (*P2PServer, error) {
	privateKey, err := crypto.HexToECDSA(config.CommonConfig.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("create P2PServer instance error:%v", err)
	}

	listenAddress := config.P2PConfig.ListenAddress
	selfNodeID := PublicKey2NodeID(&privateKey.PublicKey)
	peerDialer := NewPeerDialer()

	server := &P2PServer{
		running:      false,
		selfNodeID:   selfNodeID,
		privateKey:   privateKey,
		peerDialer:   peerDialer,
		peerManager:  peerManager,
		listenerAddr: listenAddress,
	}
	peerDialer.p2pServer = server

	return server, nil
}

func (ps *P2PServer) StartServer() error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	// start to listen the p2p requests from
	// other peer nodes in the network
	if err := ps.startListening(); err != nil {
		return err
	}

	// start to dial and connect to other peer
	// nodes in the network periodically
	if err := ps.peerDialer.StartDialing(); err != nil {
		return err
	}

	ps.running = true
	return nil
}

func (ps *P2PServer) startListening() error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	listener, err := net.Listen(TCP, ps.listenerAddr)
	if err != nil {
		return err
	}
	ps.listener = listener
	go ps.listenLoop()
	return nil
}

func (ps *P2PServer) listenLoop() {
	for {
		conn, err := ps.listener.Accept()
		if err != nil {
			return
		}
		go ps.setupConnection(conn, nil)
	}
}

func (ps *P2PServer) setupConnection(conn net.Conn, destNode *Node) {
	secret, err := ps.doEncryptionHandshake(conn, destNode)
	if err != nil {
		conn.Close()
		return
	}

	messageConn := NewMConnection(conn, secret)
	peer := NewPeer(destNode.ID, messageConn)
	ps.peerManager.AddPeer(peer)
	peer.StartRunning()
}

func (ps *P2PServer) doEncryptionHandshake(conn net.Conn, destNode *Node) (*Secret, error) {
	var err error
	var secret *Secret
	if destNode == nil {
		secret, err = ps.receiveEncryption(conn)
	} else {
		secret, err = ps.initiateEncryption(conn, destNode)
	}
	return secret, err
}

func (ps *P2PServer) receiveEncryption(conn net.Conn) (*Secret, error) {
	buffer := make([]byte, AuthRequestLength)
	count, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	buffer = buffer[:count]
	privateKey := ecies.ImportECDSA(ps.privateKey)
	plain, err := privateKey.Decrypt(rand.Reader, buffer, nil, nil)
	if err != nil {
		return nil, err
	}

	authRequest := &AuthRequest{}
	err = json.Unmarshal(plain, authRequest)
	if err != nil {
		return nil, err
	}

	remotePublicKey, err := NodeID2PublicKey(authRequest.remoteNodeID)
	if err != nil {
		return nil, err
	}

	secret, err := NewSecretWithAESKey(authRequest.remoteNodeID, ps.privateKey, remotePublicKey, authRequest.aesEncryptKey)
	if err != nil {
		return nil, err
	}

	authResponse := NewAuthResponse(ps.selfNodeID, authRequest.remoteNodeID, EncryptEchoMessage)
	buffer, err = json.Marshal(authResponse)
	if err != nil {
		return nil, err
	}

	cipher, err := ecies.Encrypt(rand.Reader, secret.remotePublicKey, buffer, nil, nil)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write(cipher)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (ps *P2PServer) initiateEncryption(conn net.Conn, destNode *Node) (*Secret, error) {
	remotePublicKey, err := NodeID2PublicKey(destNode.ID)
	if err != nil {
		return nil, err
	}

	secret, err := NewSecret(destNode.ID, ps.privateKey, remotePublicKey)
	if err != nil {
		return nil, err
	}

	authRequest := NewAuthRequest(ps.selfNodeID, secret.shareEncryptKey)
	buffer, err := json.Marshal(authRequest)
	if err != nil {
		return nil, err
	}

	cipher, err := ecies.Encrypt(rand.Reader, secret.remotePublicKey, buffer, nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(cipher)
	if err != nil {
		return nil, err
	}

	buffer = make([]byte, AuthResponseLength)
	count, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	buffer = buffer[:count]

	authResponse := &AuthResponse{}
	err = json.Unmarshal(buffer, authResponse)
	if err != nil {
		return nil, err
	}
	if authResponse.remoteNodeID != ps.selfNodeID {
		return nil, fmt.Errorf("the peer nodeID is not correct, remoteNodeID:%s", authResponse.remoteNodeID)
	}
	if authResponse.echoMessage != EncryptEchoMessage {
		return nil, fmt.Errorf("the encrypt echo message is not correct, echoMessage:%s", authResponse.echoMessage)
	}

	return secret, nil
}
