package trie

import (
	"encoding/json"
	"fmt"

	"github.com/shaojianqing/smilebc/common"
	"github.com/shaojianqing/smilebc/crypto/sha3"
)

const (
	BranchNodeSize = 17
)

const (
	NodeTypeHash      = "Hash"
	NodeTypeValue     = "Value"
	NodeTypeBranch    = "Branch"
	NodeTypeExtension = "Extension"
)

type Node interface {
	Type() string
	Hash() common.Hash
}

type HashNode struct {
	Key common.Key `json:"Key"`
}

type ValueNode struct {
	Value common.Data `json:"Value"`
}

type BranchNode struct {
	Children [BranchNodeSize]Node `json:"Children"`
}

type ExtensionNode struct {
	Key   common.Key `json:"Key"`
	Value Node       `json:"Value"`
}

func Parse(data common.Data) (Node, error) {
	internalNode := struct {
		Type string `json:"Type"`
	}{}

	if internalNode.Type == NodeTypeHash {
		hashNode := &HashNode{}
		if err := json.Unmarshal(data, hashNode); err != nil {
			return nil, err
		}
		return hashNode, nil
	} else if internalNode.Type == NodeTypeValue {
		valueNode := &ValueNode{}
		if err := json.Unmarshal(data, valueNode); err != nil {
			return nil, err
		}
		return valueNode, nil
	} else if internalNode.Type == NodeTypeBranch {
		branchNode := &BranchNode{}
		if err := json.Unmarshal(data, branchNode); err != nil {
			return nil, err
		}
		return branchNode, nil
	} else if internalNode.Type == NodeTypeExtension {
		extensionNode := &ExtensionNode{}
		if err := json.Unmarshal(data, extensionNode); err != nil {
			return nil, err
		}
		return extensionNode, nil
	} else {
		return nil, fmt.Errorf("system type does not match,type:%s", internalNode.Type)
	}
}

func (n *HashNode) Type() string {
	return NodeTypeHash
}

func (n *HashNode) Hash() common.Hash {
	return common.BytesToHash(n.Key)
}

func (n *ValueNode) Type() string {
	return NodeTypeValue
}

func (n *ValueNode) Hash() common.Hash {
	hash := sha3.NewKeccak256()
	hash.Write(n.Value)
	hashValue := hash.Sum(nil)
	return common.BytesToHash(hashValue)
}

func (n *BranchNode) Type() string {
	return NodeTypeBranch
}

func (n *BranchNode) Hash() common.Hash {
	data, _ := json.Marshal(n.Children)
	hash := sha3.NewKeccak256()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	return common.BytesToHash(hashValue)
}

func (n *ExtensionNode) Type() string {
	return NodeTypeExtension
}

func (n *ExtensionNode) Hash() common.Hash {
	data, _ := json.Marshal(n.Value)
	hash := sha3.NewKeccak256()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	return common.BytesToHash(hashValue)
}
