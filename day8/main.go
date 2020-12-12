package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	opcodes := entries()
	accumlator, _ := execute(opcodes)
	fmt.Printf("Accumulator %d\n", accumlator)

	for i, op := range opcodes {
		if accumlator, finish := execute(opcodes); finish {
			fmt.Printf("Finished Accumlator %d", accumlator)
			break
		}

		instruction := op.instruction
		switch op.instruction {
		case "nop":
			opcodes[i].instruction = "jmp"
		case "jmp":
			opcodes[i].instruction = "nop"
		default:
			continue
		}
		if accumlator, finish := execute(opcodes); finish {
			fmt.Printf("Finished Accumlator %d\n", accumlator)
			break
		}
		opcodes[i].instruction = instruction
	}
	fmt.Println("Done")
}

func entries() []opcode {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	opcodes := []opcode{}
	for scan.Scan() {
		line := scan.Text()
		sep := strings.Split(line, " ")
		arg, err := strconv.Atoi(sep[1])
		if err != nil {
			panic(err)
		}
		op := opcode{
			instruction: sep[0],
			arg:         arg,
		}
		opcodes = append(opcodes, op)
	}

	return opcodes
}
