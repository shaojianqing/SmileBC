package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/core/chain"
	"github.com/shaojianqing/smilebc/storage"
)

const (
	CreateAccount = "createAccount"
	GetAccount    = "getAccount"
	DeleteAccount = "deleteAccount"

	GetCurrentBlock  = "getCurrentBlock"
	GetBlockByNumber = "getBlockByNum"
	GetBlockHeader   = "getBlockHeader"

	SendTransaction  = "sendTransaction"
	SignTransaction  = "signTransaction"
	GetTransaction   = "getTransaction"
	EstimateTransGas = "estimateTransGas"

	GetWalletInfo     = "getWalletInfo"
	GetAccountBalance = "getAccountBalance"

	GetNetworkInfo = "getNetworkInfo"
	ListPeerSet    = "listPeerSet"
	ConnectPeer    = "connectPeer"
	DisconnectPeer = "disConnectPeer"
)

type HttpServer struct {
	serverPort string
	httpRouter *gin.Engine
	blockchain *chain.Blockchain
	chainDB    storage.Database
}

func NewHttpServer(config config.HttpConfig, chainDB storage.Database, blockchain *chain.Blockchain) (*HttpServer, error) {
	server := &HttpServer{
		serverPort: config.ServerPort,
		httpRouter: gin.Default(),
		blockchain: blockchain,
		chainDB:    chainDB,
	}

	server.initiate()
	return server, nil
}

func (hs *HttpServer) initiate() {
	//Initiate account related interface
	hs.httpRouter.POST(GetAccount, hs.getAccount)
	hs.httpRouter.POST(CreateAccount, hs.createAccount)
	hs.httpRouter.POST(DeleteAccount, hs.deleteAccount)

	//Initiate block related interface
	hs.httpRouter.POST(GetCurrentBlock, hs.getCurrentBlock)
	hs.httpRouter.POST(GetBlockByNumber, hs.getBlockByNumber)
	hs.httpRouter.POST(GetBlockHeader, hs.getBlockHeader)

	//Initiate transaction related interface
	hs.httpRouter.POST(SendTransaction, hs.sendTransaction)
	hs.httpRouter.POST(SignTransaction, hs.signTransaction)
	hs.httpRouter.POST(GetTransaction, hs.getTransaction)
	hs.httpRouter.POST(GetTransaction, hs.getTransaction)
	hs.httpRouter.POST(EstimateTransGas, hs.estimateTransactionGas)

	//Initiate wallet and balance related interface
	hs.httpRouter.POST(GetWalletInfo, hs.getWalletInfo)
	hs.httpRouter.POST(GetAccountBalance, hs.getAccountBalance)

	//Initiate network related interface
	hs.httpRouter.POST(GetNetworkInfo, hs.getNetworkInfo)
	hs.httpRouter.POST(ListPeerSet, hs.listPeers)
	hs.httpRouter.POST(ConnectPeer, hs.connectPeer)
	hs.httpRouter.POST(DisconnectPeer, hs.disconnectPeer)
}

func (hs *HttpServer) StartService() error {
	if err := hs.httpRouter.Run(hs.serverPort); err != nil {
		return fmt.Errorf("fail to start the http server:%w", err)
	}
	return nil
}
