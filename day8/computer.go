package main

import "math"

type opcode struct {
	instruction string
	arg         int
}

type registers struct {
	accumlator     int
	programCounter int
}

func nop(arg int, reg *registers) {
	reg.programCounter++
}

func acc(arg int, reg *registers) {
	reg.accumlator += arg
	reg.programCounter++
}

func jmp(arg int, reg *registers) {
	reg.programCounter += arg
}

var instructions = map[string]func(int, *registers){
	"nop": nop,
	"acc": acc,
	"jmp": jmp,
}

func execute(opcodes []opcode) (int, bool) {
	seen := map[int]interface{}{}
	regs := &registers{}

	for regs.programCounter < len(opcodes) {
		if _, has := seen[regs.programCounter]; has {
			return regs.accumlator, false
		}
		seen[regs.programCounter] = regs.programCounter

		oc := opcodes[regs.programCounter]
		instruction, has := instructions[oc.instruction]
		if has == false {
			return math.MinInt64, false
		}
		instruction(oc.arg, regs)

	}
	return regs.accumlator, true
}
