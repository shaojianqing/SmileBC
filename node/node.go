package node

import (
	"fmt"
	"log"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	stat "github.com/shaojianqing/smilebc/core/processor"
	"github.com/shaojianqing/smilebc/protocol"
	"github.com/shaojianqing/smilebc/server"
	"github.com/shaojianqing/smilebc/storage"
)

type SmileNode struct {
	syncManager *protocol.SyncManager
	httpServer  *server.HttpServer
	blockchain  *chain.Blockchain
}

func NewSmileNode(config *config.Config) *SmileNode {
	chainDB, err := storage.NewDatabase(config.DBConfig)
	if err != nil {
		log.Fatalf("fail to initiate chain database storage,error:%v", err)
	}

	blockchain := chain.NewBlockchain(chainDB)

	processor := stat.NewStateProcessor(blockchain)
	blockchain.Processor = processor

	server, err := server.NewHttpServer(config.HttpConfig, chainDB, blockchain)
	if err != nil {
		log.Fatalf("fail to initiate http server,error:%v", err)
	}

	manager, err := protocol.NewSyncManager(config.SyncConfig, blockchain)
	if err != nil {
		log.Fatalf("fail to initiate sync manager,error:%v", err)
	}

	node := &SmileNode{
		httpServer:  server,
		syncManager: manager,
		blockchain:  blockchain,
	}
	return node
}

func (sn *SmileNode) StartService() error {
	if err := sn.httpServer.StartService(); err != nil {
		return fmt.Errorf("fail to start http server, err:%w", err)
	}

	if err := sn.syncManager.StartSync(); err != nil {
		return fmt.Errorf("fail to start sync process, err:%w", err)
	}

	return nil
}
