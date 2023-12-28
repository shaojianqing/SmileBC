package p2p

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/crypto"
	"net"
	"time"
)

const (
	Ping      = "ping"
	Pong      = "pong"
	FindNode  = "findNode"
	Neighbors = "neighbors"
)

type Endpoint struct {
	IP     net.IP
	UDP    uint16
	TCP    uint16
	NodeID NodeID
}

func NewEndpoint(address *net.UDPAddr, tcpPort uint16) Endpoint {
	ip := address.IP.To4()
	if ip == nil {
		ip = address.IP.To16()
	}
	return Endpoint{
		IP:  ip,
		TCP: tcpPort,
		UDP: uint16(address.Port),
	}
}

func NewEndpointWithNodeID(ip net.IP, udpPort, tcpPort uint16, nodeID NodeID) Endpoint {
	return Endpoint{
		IP:     ip,
		TCP:    tcpPort,
		UDP:    udpPort,
		NodeID: nodeID,
	}
}

type NodeMessage interface {
	InitMessage() error

	GetNodeID() NodeID
	GetMessageType() string
	GetSignature() string
	GetExpiration() int64
}

type GenericMessage struct {
	NodeID      NodeID      `json:"nodeID"`
	MessageType string      `json:"messageType"`
	Signature   string      `json:"signature"`
	Expiration  int64       `json:"expiration"`
	MessageHash common.Hash `json:"messageHash"`
}

func (gm *GenericMessage) GetMessageType() string {
	return gm.MessageType
}

func (gm *GenericMessage) GetMessageHash() common.Hash {
	return gm.MessageHash
}

func (gm *GenericMessage) GetSignature() string {
	return gm.Signature
}

func (gm *GenericMessage) GetExpiration() int64 {
	return gm.Expiration
}

func (gm *GenericMessage) GetNodeID() NodeID {
	return gm.NodeID
}

type PingMessage struct {
	GenericMessage

	Source      Endpoint
	Destination Endpoint
}

func NewPingMessage(selfNodeID NodeID, sourceAddress, destinationAddress *net.UDPAddr) *PingMessage {
	expiration := time.Now().Add(Expiration).UnixMilli()
	pingMessage := &PingMessage{
		GenericMessage: GenericMessage{
			NodeID:      selfNodeID,
			Expiration:  expiration,
			MessageType: Ping,
		},
		Source:      NewEndpoint(sourceAddress, uint16(sourceAddress.Port)),
		Destination: NewEndpoint(destinationAddress, uint16(destinationAddress.Port)),
	}
	pingMessage.InitMessage()
	return pingMessage
}

func (msg *PingMessage) InitMessage() error {
	msg.MessageHash = crypto.Keccak256Hash(msg.NodeID[:])
	return nil
}

type PongMessage struct {
	GenericMessage

	Destination Endpoint
	ReplyData   common.Hash
}

func (msg *PongMessage) InitMessage() error {
	return nil
}

type FindNodeMessage struct {
	GenericMessage

	Target NodeID `json:"targetNodeID"`
}

func NewFindNodeMessage(selfNodeID NodeID, targetNodeID NodeID) *FindNodeMessage {
	expiration := time.Now().Add(Expiration).UnixMilli()
	findNodeMessage := &FindNodeMessage{
		GenericMessage: GenericMessage{
			NodeID:      selfNodeID,
			Expiration:  expiration,
			MessageType: FindNode,
		},
		Target: targetNodeID,
	}
	findNodeMessage.InitMessage()
	return findNodeMessage
}

func (msg *FindNodeMessage) InitMessage() error {
	msg.MessageHash = crypto.Keccak256Hash(msg.NodeID[:])
	return nil
}

type NeighborsMessage struct {
	GenericMessage

	NodeList []Endpoint
}

func NewNeighborsMessage(selfNodeID NodeID, nodeList []Endpoint) *NeighborsMessage {
	expiration := time.Now().Add(Expiration).UnixMilli()
	neighborsMessage := &NeighborsMessage{
		GenericMessage: GenericMessage{
			NodeID:      selfNodeID,
			Expiration:  expiration,
			MessageType: Neighbors,
		},
		NodeList: nodeList,
	}
	neighborsMessage.InitMessage()
	return neighborsMessage
}

func (msg *NeighborsMessage) InitMessage() error {
	return nil
}

func decodeMessage(data []byte) (NodeMessage, error) {
	genericMessage := &GenericMessage{}
	if err := json.Unmarshal(data, genericMessage); err != nil {
		return nil, err
	}

	var err error
	var nodeMessage NodeMessage
	if genericMessage.GetMessageType() == Ping {
		nodeMessage := &PingMessage{}
		err = json.Unmarshal(data, nodeMessage)
	} else if genericMessage.GetMessageType() == Pong {
		nodeMessage := &PongMessage{}
		err = json.Unmarshal(data, nodeMessage)
	} else if genericMessage.GetMessageType() == FindNode {
		nodeMessage := &FindNodeMessage{}
		err = json.Unmarshal(data, nodeMessage)
	} else if genericMessage.GetMessageType() == Neighbors {
		nodeMessage := &NeighborsMessage{}
		err = json.Unmarshal(data, nodeMessage)
	} else {
		return nil, errors.New("message type is not valid")
	}

	if err != nil {
		return nil, fmt.Errorf("decode message error:%v", err)
	}
	return nodeMessage, nil
}

func encodeMessage(message NodeMessage) ([]byte, error) {
	if err := message.InitMessage(); err != nil {
		return nil, err
	}
	return json.Marshal(message)
}
