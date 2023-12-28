package p2p

import (
	"errors"
	"fmt"
	"net"
	"time"
)

const (
	MaxNeighborsCount = 8
)

type handleNodeMessage func(sourceAddress *net.UDPAddr, message NodeMessage) error

type MessageHandler struct {
	networkManager *NetworkManager
	handlerMap     map[string]handleNodeMessage
}

func NewMessageHandler(manager *NetworkManager) *MessageHandler {
	handlerMap := make(map[string]handleNodeMessage, 4)
	messageHandler := &MessageHandler{
		networkManager: manager,
		handlerMap:     handlerMap,
	}

	handlerMap[Ping] = messageHandler.handlePingMessage
	handlerMap[Pong] = messageHandler.handlePongMessage
	handlerMap[FindNode] = messageHandler.handleFindNodeMessage
	handlerMap[Neighbors] = messageHandler.handleNeighborsMessage

	return messageHandler
}

func (h *MessageHandler) handleMessage(sourceAddress *net.UDPAddr, message NodeMessage) error {
	expiration := time.UnixMilli(message.GetExpiration())
	if expiration.Before(time.Now()) {
		return errors.New("message has already expired")
	}

	if handler, ok := h.handlerMap[message.GetMessageType()]; ok {
		return handler(sourceAddress, message)
	}

	return errors.New("message type is not valid")
}

func (h *MessageHandler) handlePingMessage(sourceAddress *net.UDPAddr, message NodeMessage) error {
	pingMessage, ok := message.(*PingMessage)
	if !ok {
		return errors.New("message is not the expected PingMessage")
	}
	selfNodeID := h.networkManager.SelfNode.ID
	expiration := time.Now().Add(Expiration).UnixMilli()
	pongMessage := &PongMessage{
		GenericMessage: GenericMessage{
			NodeID:      selfNodeID,
			Expiration:  expiration,
			MessageType: Pong,
		},
		Destination: NewEndpoint(sourceAddress, pingMessage.Source.TCP),
		ReplyData:   pingMessage.GetMessageHash(),
	}
	if err := h.networkManager.send(sourceAddress, pongMessage); err != nil {
		return fmt.Errorf("send pong message error:%v", err)
	}

	return nil
}

func (h *MessageHandler) handlePongMessage(sourceAddress *net.UDPAddr, message NodeMessage) error {
	return nil
}

func (h *MessageHandler) handleFindNodeMessage(sourceAddress *net.UDPAddr, message NodeMessage) error {
	findNodeMessage, ok := message.(*FindNodeMessage)
	if !ok {
		return errors.New("message is not the expected FindNodeMessage")
	}
	manager := h.networkManager
	routeTable := manager.RouterTable
	if !routeTable.ExistNodeInTable(findNodeMessage.NodeID) {
		return fmt.Errorf("unknown node from FindNodeMessage, nodeID:%s", findNodeMessage.NodeID)
	}

	closestNodeSet := routeTable.GetClosestNodes(findNodeMessage.Target)
	nodeList := make([]Endpoint, 0)
	for _, node := range closestNodeSet.ClosestNodes {
		endpoint := NewEndpointWithNodeID(node.IP, node.UDP, node.TCP, node.ID)
		nodeList = append(nodeList, endpoint)
		if len(nodeList) > MaxNeighborsCount {
			neighborsMessage := NewNeighborsMessage(manager.SelfNode.ID, nodeList)
			if err := manager.send(sourceAddress, neighborsMessage); err != nil {
				return fmt.Errorf("send neighbors message error:%v", err)
			}
			nodeList = nodeList[:0]
		}
	}

	return nil
}

func (h *MessageHandler) handleNeighborsMessage(sourceAddress *net.UDPAddr, message NodeMessage) error {
	neighborsMessage, ok := message.(*NeighborsMessage)
	if !ok {
		return errors.New("message is not the expected NeighborsMessage")
	}

	manager := h.networkManager
	routeTable := manager.RouterTable
	if !routeTable.ExistNodeInTable(neighborsMessage.NodeID) {
		return fmt.Errorf("unknown node from NeighborsMessage, nodeID:%s", neighborsMessage.NodeID)
	}

	for _, endpoint := range neighborsMessage.NodeList {
		if IsReserveIP(endpoint.IP) {
			continue
		}
		node := NewNode(endpoint.NodeID, endpoint.IP, endpoint.TCP, endpoint.UDP)
		routeTable.AddNode(node)
	}

	return nil
}
