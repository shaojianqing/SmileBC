package node

import (
	"github.com/shaojianqing/smilebc/protocol"
	"github.com/shaojianqing/smilebc/server"
)

type SmileNode struct {
	SyncManager protocol.SyncManager
	HttpServer  server.HttpServer
}

func NewSmileNode() *SmileNode {
	return nil
}

func (sn *SmileNode) StartService() error {
	return nil
}
