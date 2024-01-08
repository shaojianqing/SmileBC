package proto

import (
	"fmt"

	p2p "github.com/shaojianqing/smilebc/network"
)

type handler func(message *p2p.Message) error

type ProtocolHandler struct {
	manager    *ProtocolManager
	handlerMap map[string]handler
}

func NewProtocolHandler() (*ProtocolHandler, error) {
	protocolHandler := &ProtocolHandler{
		handlerMap: make(map[string]handler),
	}

	// Construct the handler function map for message type
	protocolHandler.handlerMap[p2p.BlockBodiesReq] = protocolHandler.handleBlockBodiesReq
	protocolHandler.handlerMap[p2p.BlockBodiesResp] = protocolHandler.handleBlockBodiesResp
	protocolHandler.handlerMap[p2p.BlockHeadersReq] = protocolHandler.handleBlockHeadersReq
	protocolHandler.handlerMap[p2p.BlockHeadersResp] = protocolHandler.handleBlockHeadersResp
	protocolHandler.handlerMap[p2p.NodeDataReq] = protocolHandler.handleNodeDataReq
	protocolHandler.handlerMap[p2p.NodeDataResp] = protocolHandler.handleNodeDataResp
	protocolHandler.handlerMap[p2p.ReceiptsReq] = protocolHandler.handleReceiptsReq
	protocolHandler.handlerMap[p2p.ReceiptsResp] = protocolHandler.handleReceiptsResp

	return protocolHandler, nil
}

func (ph *ProtocolHandler) Handle(peer *p2p.Peer) error {
	message, err := peer.ReadMessage()
	if err != nil {
		return err
	}

	// Get the message handler function from the
	// handler map inside protocol handler instance
	messageHandler := ph.handlerMap[message.MessageType]
	if messageHandler != nil {
		return fmt.Errorf("handler does not exist for message type[MessageType:%v]", message.MessageType)
	}

	// handle the peer message according to logic
	// implemented by the corresponding handler function
	if err = messageHandler(message); err != nil {
		return err
	}
	return nil
}

func (ph *ProtocolHandler) handleBlockHeadersReq(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleBlockHeadersResp(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleBlockBodiesReq(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleBlockBodiesResp(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleNodeDataReq(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleNodeDataResp(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleReceiptsReq(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) handleReceiptsResp(message *p2p.Message) error {
	return nil
}

func (ph *ProtocolHandler) SetProtocolManager(manager *ProtocolManager) {
	ph.manager = manager
}
