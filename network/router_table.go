package p2p

import (
	"sync"
	"time"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/crypto"
)

const (
	BucketCount = 256
	BucketSize  = 16
)

type Bucket struct {
	activeEntries []*Node
	backupEntries []*Node
}

func (bucket *Bucket) bump(node *Node) bool {
	for i := range bucket.activeEntries {
		// If the bucket contains the node, directly
		// move the node to the first position
		if bucket.activeEntries[i].ID == node.ID {
			copy(bucket.activeEntries[1:], bucket.activeEntries[:i])
			bucket.activeEntries[0] = node
			return true
		}
	}
	return false
}

type ClosestNodeSet struct {
	TargetHash   common.Hash
	ClosestNodes [BucketSize]*Node
}

func (set *ClosestNodeSet) AcceptNode(node *Node) {

	// Iterate the node list and set the value when meeting free space
	for i := range set.ClosestNodes {
		if set.ClosestNodes[i] == nil {
			set.ClosestNodes[i] = node
			return
		} else if set.ClosestNodes[i].ID == node.ID {
			return
		}
	}

	// Try to find the farthest node within the closestNodes list
	index := 0
	farthest := set.ClosestNodes[index]
	for i := range set.ClosestNodes {
		closestNode := set.ClosestNodes[i]
		if CompareDistance(set.TargetHash, closestNode.Hash, farthest.Hash) > 0 {
			index = i
			farthest = closestNode
		}
	}

	// If the farthest node is still farther than the input
	// node, replace the farthest node with the input node
	if CompareDistance(set.TargetHash, farthest.Hash, node.Hash) > 0 {
		set.ClosestNodes[index] = node
	}
}

type RouterTable struct {
	bucketList [BucketCount]*Bucket
	mutex      sync.Mutex
	count      uint32
	self       *Node
}

func NewRouterTable(self *Node) *RouterTable {
	table := &RouterTable{
		self:  self,
		count: 0,
	}

	for i := range table.bucketList {
		table.bucketList[i] = &Bucket{}
	}

	return table
}

func (table *RouterTable) StartRefresh() {

}

func (table *RouterTable) GetClosestNodes(nodeID NodeID) *ClosestNodeSet {
	closestNodeSet := &ClosestNodeSet{
		TargetHash: crypto.Keccak256Hash(nodeID[:]),
	}

	// Iterate all the entries in the buckets, select the closest nodes and inject them into
	// closestNodeSet, so that client can accept them
	for _, bucket := range table.bucketList {
		for _, entry := range bucket.activeEntries {
			closestNodeSet.AcceptNode(entry)
		}
	}

	return closestNodeSet
}

func (table *RouterTable) AddNode(node *Node) {
	table.mutex.Lock()
	defer table.mutex.Unlock()

	distance := LogDistance(node.Hash, table.self.Hash)
	bucket := table.bucketList[distance]
	if bucket.bump(node) {
		return
	}

	// If free space is still available, add the node into the activeEntries
	// and delete the same node from the backupEntries
	if len(bucket.activeEntries) < BucketSize {
		bucket.activeEntries = pushNode(bucket.activeEntries, node)
		bucket.backupEntries = deleteNode(bucket.backupEntries, node)
		node.AddedTime = time.Now()
		table.count++
		return
	}

	// If no more free space for adding new node, add the node into the backupEntries instead
	if !nodeInBucketBackup(node, bucket) {
		bucket.backupEntries = pushNode(bucket.backupEntries, node)
	}
}

func (table *RouterTable) InjectNodes(nodes []*Node) {
	table.mutex.Lock()
	defer table.mutex.Unlock()

	for _, node := range nodes {
		// Calculate the logarithmic distance of self node and remote note, and then
		// get the corresponding bucket and check whether it contains the remote node
		distance := LogDistance(node.Hash, table.self.Hash)
		bucket := table.bucketList[distance]
		if nodeInBucketActive(node, bucket) {
			continue
		}

		// If the bucket does not contain the remote node, we directly add the node into
		// the corresponding bucket, so that the kademlia table could be constructed
		if len(bucket.activeEntries) < BucketSize {
			bucket.activeEntries = append(bucket.activeEntries, node)
			table.count++
		}
	}
}

// Check whether the node has been included within the activeEntries in bucket, and the logic is simple here
func nodeInBucketActive(node *Node, bucket *Bucket) bool {
	for _, entry := range bucket.activeEntries {
		if node.ID == entry.ID {
			return true
		}
	}
	return false
}

// Check whether the node has been included within the backupEntries in bucket, and the logic is simple here
func nodeInBucketBackup(node *Node, bucket *Bucket) bool {
	for _, entry := range bucket.backupEntries {
		if node.ID == entry.ID {
			return true
		}
	}
	return false
}

// push the node and put it on the first position in the list
func pushNode(list []*Node, n *Node) []*Node {
	result := make([]*Node, len(list)+1)
	copy(result[1:], list)
	result[0] = n
	return result
}

// delete the node from list.
func deleteNode(list []*Node, n *Node) []*Node {
	for i := range list {
		if list[i].ID == n.ID {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}
