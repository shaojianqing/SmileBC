package vmachine

var STOP = OpCode{code: 0, name: "Stop"}
var ADD = OpCode{code: 1, name: "Add"}
var MUL = OpCode{code: 2, name: "Mul"}
var SUB = OpCode{code: 3, name: "Sub"}
var DIV = OpCode{code: 4, name: "Div"}
var SDIV = OpCode{code: 5, name: "SDiv"}
var MOD = OpCode{code: 6, name: "Mod"}
var SMOD = OpCode{code: 7, name: "SMod"}
var ADDMOD = OpCode{code: 8, name: "AddMod"}
var MULMOD = OpCode{code: 9, name: "MulMod"}
var EXP = OpCode{code: 10, name: "Exp"}
var SIGNEXTEND = OpCode{code: 11, name: "SignExtend"}

var LT = OpCode{code: 16, name: "Lt"}
var GT = OpCode{code: 17, name: "Gt"}
