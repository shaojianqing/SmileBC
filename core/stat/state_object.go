package stat

import (
	"math/big"

	"github.com/shaojianqing/smilebc/common"
)

type StateObject struct {
}

func (so *StateObject) GetAddress() common.Address {
	return common.Address{}
}

func (so *StateObject) AddBalance(amount big.Int) {

}

func (so *StateObject) SubBalance(amount big.Int) {

}

func (so *StateObject) GetBalance() big.Int {
	return *big.NewInt(0)
}

func (so *StateObject) SetBalance(amount big.Int) {

}

func (so *StateObject) GetNonce() uint64 {
	return 0
}

func (so *StateObject) SetNonce(nonce uint64) {

}

func (so *StateObject) GetCode() common.Code {
	return common.Code{}
}

func (so *StateObject) SetCode(code common.Code) {

}
