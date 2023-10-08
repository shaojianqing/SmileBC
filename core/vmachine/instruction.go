package vmachine

const ByteSize = 256

type OpCode struct {
	code byte
	name string
}

type Executor func(vm *VMachine, inst *Instruction) error

type DynamicGas func(vm *VMachine, inst *Instruction) uint64

type Instruction struct {
	opCode     OpCode
	staticGas  uint64
	dynamicGas DynamicGas
	executor   Executor
}

type InstructionTable [ByteSize]Instruction

func NewInstruction(code OpCode, executor Executor, dynamicGas DynamicGas) Instruction {
	return Instruction{
		opCode:     code,
		executor:   executor,
		dynamicGas: dynamicGas,
	}
}

func InitiateInstTable() InstructionTable {
	instTable := InstructionTable{}
	instTable[STOP.code] = NewInstruction(STOP, opStopExecutor, DummyDynamicGas)
	instTable[ADD.code] = NewInstruction(ADD, opAddExecutor, DummyDynamicGas)
	instTable[MUL.code] = NewInstruction(MUL, opMulExecutor, DummyDynamicGas)
	instTable[SUB.code] = NewInstruction(SUB, opSubExecutor, DummyDynamicGas)
	instTable[DIV.code] = NewInstruction(DIV, opDivExecutor, DummyDynamicGas)
	instTable[SDIV.code] = NewInstruction(SDIV, opSDivExecutor, DummyDynamicGas)
	instTable[MOD.code] = NewInstruction(MOD, opModExecutor, DummyDynamicGas)
	instTable[SMOD.code] = NewInstruction(SMOD, opSModExecutor, DummyDynamicGas)
	instTable[ADDMOD.code] = NewInstruction(ADDMOD, opAddModExecutor, DummyDynamicGas)
	instTable[MULMOD.code] = NewInstruction(MULMOD, opMulModExecutor, DummyDynamicGas)
	instTable[EXP.code] = NewInstruction(EXP, opExpExecutor, DummyDynamicGas)
	instTable[SIGNEXTEND.code] = NewInstruction(SIGNEXTEND, opSignExtendExecutor, DummyDynamicGas)

	instTable[LT.code] = NewInstruction(LT, opLtExecutor, DummyDynamicGas)
	instTable[GT.code] = NewInstruction(GT, opGtExecutor, DummyDynamicGas)

	return instTable
}
