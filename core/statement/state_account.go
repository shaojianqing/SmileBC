package stat

import (
	"math/big"

	"github.com/shaojianqing/smilebc/common"
)

type StateAccount struct {
	Nonce    uint64
	Balance  big.Int
	RootHash common.Hash
	CodeHash common.Hash
}
