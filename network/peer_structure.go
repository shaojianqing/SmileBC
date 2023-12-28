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

	BlockType        = "block"
	TransactionsType = "transactions"
)

type Peer struct {
	id    string
	conn  MessageConn
	mutex sync.RWMutex

	knownBlockSet       *set.Set
	knownTransactionSet *set.Set
}

func NewPeer(nodeID NodeID, connection MessageConn) *Peer {
	peerId := string(nodeID[:])
	return &Peer{
		id:   peerId,
		conn: connection,

		knownBlockSet:       set.New(),
		knownTransactionSet: set.New(),
	}
}

func (p *Peer) StartRunning() {

}

func (p *Peer) GetId() string {
	return p.id
}

func (p *Peer) SendBlock(block *model.Block) error {
	if err := Send(p.conn, BlockType, block); err != nil {
		return fmt.Errorf("send block err:%v", err)
	}
	p.MarkBlock(block.GetHash())
	return nil
}

func (p *Peer) SendTransactions(transactions []*model.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	if err := Send(p.conn, TransactionsType, transactions); err != nil {
		return fmt.Errorf("send transactions err:%v", err)
	}

	for _, transaction := range transactions {
		p.MarkTransaction(transaction.GetHash())
	}
	return nil
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
