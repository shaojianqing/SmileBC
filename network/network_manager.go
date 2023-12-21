package p2p

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/shaojianqing/smilebc/config"
)

const (
	UDP        = "udp"
	BufferSize = 1280

	Expiration = 20 * time.Second
)

type NetworkManager struct {
	NetworkConfig config.NetworkConfig
	RouterTable   *RouterTable

	UDPConn *net.UDPConn
	Dialer  *net.Dialer

	SelfNode Node

	MessageHandler *MessageHandler
}

func NewNetworkManager(config config.Config) (*NetworkManager, error) {
	privateKey := config.CommonConfig.PrivateKey
	nodeID := PublicKey2NodeID(&privateKey.PublicKey)
	networkConfig := config.NetworkConfig
	listenAddress, err := net.ResolveUDPAddr(UDP, networkConfig.ListenAddress)
	if err != nil {
		return nil, fmt.Errorf("resolve address error:%v", err)
	}

	selfNode := NewNode(nodeID, listenAddress.IP, networkConfig.TCPPort, networkConfig.TCPPort)
	routerTable := NewRouterTable(selfNode)
	networkManager := &NetworkManager{
		NetworkConfig: config.NetworkConfig,
		RouterTable:   routerTable,
	}

	messageHandler := NewMessageHandler(networkManager)
	networkManager.MessageHandler = messageHandler

	return networkManager, nil
}

func (nm *NetworkManager) Start() error {
	listenAddress, err := net.ResolveUDPAddr(UDP, nm.NetworkConfig.ListenAddress)
	if err != nil {
		return fmt.Errorf("resolve address error:%v", err)
	}

	udpConn, err := net.ListenUDP(UDP, listenAddress)
	if err != nil {
		return fmt.Errorf("start to listen udp error:%v", err)
	}
	nm.UDPConn = udpConn

	go nm.stateLoop()
	return nil
}

func (nm *NetworkManager) stateLoop() {
	defer nm.UDPConn.Close()

	buffer := make([]byte, BufferSize)
	for {
		count, source, err := nm.UDPConn.ReadFromUDP(buffer)
		if nm.isTemporaryError(err) {
			continue
		} else if err != nil {
			log.Printf("read from UDP error:%v", err)
			return
		}

		err = nm.handle(source, buffer[0:count])
		if err != nil {
			log.Printf("handle UDP data error:%v", err)
		}
	}
}

func (nm *NetworkManager) handle(sourceAddress *net.UDPAddr, data []byte) error {
	message, err := decodeMessage(data)
	if err != nil {
		return fmt.Errorf("decode node message error:%v", err)
	}

	err = nm.MessageHandler.handleMessage(sourceAddress, message)
	if err != nil {
		return fmt.Errorf("handle node message error:%v", err)
	}

	return nil
}

func (nm *NetworkManager) ping(destNodeID NodeID, destAddress *net.UDPAddr) error {
	return nil
}

func (nm *NetworkManager) send(destinationAddress *net.UDPAddr, message NodeMessage) error {
	return nil
}

func (nm *NetworkManager) isTemporaryError(err error) bool {
	temporaryErr, ok := err.(interface {
		Temporary() bool
	})
	return ok && temporaryErr.Temporary()
}
