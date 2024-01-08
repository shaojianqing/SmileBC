package p2p

import "github.com/shaojianqing/smilebc/common"

const (
	NewBlockResp     = "newBlockResp"
	TransactionsResp = "transactionsResp"
	PeerStateReq     = "peerStateReq"
	PeerStateResp    = "peerStateResp"
	BlockHeadersReq  = "blockHeadersReq"
	BlockHeadersResp = "blockHeadersResp"
	BlockBodiesReq   = "blockBodiesReq"
	BlockBodiesResp  = "blockBodiesResp"
	NodeDataReq      = "nodeDataReq"
	NodeDataResp     = "nodeDataResp"
	ReceiptsReq      = "receiptReq"
	ReceiptsResp     = "receiptsResp"
)

type Message struct {
	MessageType string      `json:"type"`
	MessageBody interface{} `json:"body"`
}

type internalMessage struct {
	MessageType string `json:"type"`
}

type PeerState struct {
	Network string `json:"network"`
	Version string `json:"version"`

	Genesis common.Hash `json:"genesis"`
}

func (m *Message) UnmarshalJSON(buffer []byte) error {
	return nil
}
