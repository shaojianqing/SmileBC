package vmachine

import "github.com/shaojianqing/smilebc/core/model"

type Param []byte

type Result []byte

type VMachine struct {
}

func NewVMachine() *VMachine {
	return nil
}

func (vm *VMachine) Run(contract model.Contract, input Param, readOnly bool) (Result, error) {
	return Result{}, nil
}
