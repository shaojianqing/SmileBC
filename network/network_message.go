package p2p

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shaojianqing/smilebc/common"
	"net"
)

const (
	Ping      = "ping"
	Pong      = "pong"
	FindNode  = "findNode"
	Neighbors = "neighbors"
)

type Endpoint struct {
	IP  net.IP
	UDP uint16
	TCP uint16
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

func (msg *PingMessage) InitMessage() error {
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
}

func (msg *FindNodeMessage) InitMessage() error {
	return nil
}

type NeighborsMessage struct {
	GenericMessage
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
