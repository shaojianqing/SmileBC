package system

import (
	"fmt"
	"log"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	stat "github.com/shaojianqing/smilebc/core/processor"
	sync "github.com/shaojianqing/smilebc/protocol"
	"github.com/shaojianqing/smilebc/server"
	"github.com/shaojianqing/smilebc/storage"
)

type Smile struct {
	protocolManager *sync.ProtocolManager
	httpServer      *server.HttpServer
	blockchain      *chain.Blockchain
}

func NewSmile(config *config.Config) *Smile {
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

	manager, err := sync.NewProtocolManager(config.SyncConfig, blockchain)
	if err != nil {
		log.Fatalf("fail to initiate sync manager,error:%v", err)
	}

	node := &Smile{
		blockchain:      blockchain,
		httpServer:      server,
		protocolManager: manager,
	}
	return node
}

func (sm *Smile) StartService() error {
	if err := sm.httpServer.StartService(); err != nil {
		return fmt.Errorf("fail to start http server, err:%w", err)
	}

	if err := sm.protocolManager.Start(); err != nil {
		return fmt.Errorf("fail to start sync process, err:%w", err)
	}

	return nil
}
