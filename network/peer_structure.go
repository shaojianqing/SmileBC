package p2p

import (
	"fmt"
	"sync"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/core/model"
	"gopkg.in/fatih/set.v0"
)

const (
	MaxKnownBlocks       = 1024
	MaxKnownTransactions = 32768
)

type Peer struct {
	id    NodeID
	conn  MessageConn
	mutex sync.RWMutex

	knownBlockSet       *set.Set
	knownTransactionSet *set.Set
}

func NewPeer(nodeID NodeID, connection MessageConn) *Peer {
	return &Peer{
		id:   nodeID,
		conn: connection,

		knownBlockSet:       set.New(),
		knownTransactionSet: set.New(),
	}
}

func (p *Peer) GetID() NodeID {
	return p.id
}

func (p *Peer) StartRunning() {

}

func (p *Peer) SendBlock(block *model.Block) error {
	message := &Message{MessageType: NewBlockResp, MessageBody: block}
	if err := p.WriteMessage(message); err != nil {
		return fmt.Errorf("send block error:%v", err)
	}

	p.MarkBlock(block.GetHash())
	return nil
}

func (p *Peer) SendTransactions(transactions []*model.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	message := &Message{MessageType: TransactionsResp, MessageBody: transactions}
	if err := p.WriteMessage(message); err != nil {
		return fmt.Errorf("send transactions error:%v", err)
	}

	for _, transaction := range transactions {
		p.MarkTransaction(transaction.GetHash())
	}
	return nil
}

func (p *Peer) ReadMessage() (*Message, error) {
	return p.conn.Read()
}

func (p *Peer) WriteMessage(message *Message) error {
	return p.conn.Write(message)
}

func (p *Peer) MarkBlock(hash common.Hash) {
	if p.knownBlockSet.Size() > MaxKnownBlocks {
		p.knownBlockSet.Pop()
	}
	p.knownBlockSet.Add(hash)
}

func (p *Peer) MarkTransaction(hash common.Hash) {
	if p.knownTransactionSet.Size() > MaxKnownTransactions {
		p.knownTransactionSet.Pop()
	}
	p.knownTransactionSet.Add(hash)
}
