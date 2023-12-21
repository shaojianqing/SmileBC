package p2p

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"net"
	"time"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/crypto"
)

const (
	NodeIDLength = 64
)

type NodeID [NodeIDLength]byte

type Node struct {
	ID  NodeID `json:"id"`
	IP  net.IP `json:"ip"`
	TCP uint16 `json:"tcp"`
	UDP uint16 `json:"udp"`

	Contest   bool        `json:"contest"`
	Hash      common.Hash `json:"hash"`
	AddedTime time.Time   `json:"addedTime"`
}

func NewNode(id NodeID, ip net.IP, tcp, udp uint16) *Node {
	hash := crypto.Keccak256Hash(id[:])
	return &Node{
		ID:   id,
		IP:   ip,
		TCP:  tcp,
		UDP:  udp,
		Hash: hash,
	}
}

func PublicKey2NodeID(pub *ecdsa.PublicKey) NodeID {
	var id NodeID
	content := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	if len(content)-1 != len(id) {
		panic(fmt.Errorf("need %d bit publicKey, got %d bits", (len(id)+1)*8, len(content)))
	}
	copy(id[:], content[1:])
	return id
}

var leadingZeroCount = [256]int{
	8, 7, 6, 6, 5, 5, 5, 5,
	4, 4, 4, 4, 4, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

func LogDistance(a, b common.Hash) int {
	leadingZero := 0
	for i := range a {
		x := a[i] ^ b[i]
		if x == 0 {
			leadingZero += 8
		} else {
			leadingZero += leadingZeroCount[x]
			break
		}
	}
	return len(a)*8 - leadingZero
}

func CompareDistance(target, a, b common.Hash) int {
	for i := range target {
		da := a[i] ^ target[i]
		db := b[i] ^ target[i]
		if da > db {
			return 1
		} else if da < db {
			return -1
		}
	}
	return 0
}
