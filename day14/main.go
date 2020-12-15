package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	and          int
	or           int
	floating     int
	instructions []instruction
}

type instruction struct {
	location int
	value    int
}

func main() {
	programs := entries()
	addr := map[int]int{}
	for _, prog := range programs {
		for _, inst := range prog.instructions {
			val := inst.value
			val &= prog.and
			val |= prog.or
			addr[inst.location] = val
		}
	}

	sum := 0
	for _, val := range addr {
		sum += val
	}
	fmt.Printf("Instruction sum: %d\n", sum)

	addrv2 := map[int]int{}
	for _, prog := range programs {
		for _, inst := range prog.instructions {
			inst.location |= prog.or
			locations := addresses([]int{}, inst.location, prog.floating, 0x10000000000)
			for _, loc := range locations {
				addrv2[loc] = inst.value
			}
		}
	}

	sumv2 := 0
	for _, val := range addrv2 {
		sumv2 += val
	}
	fmt.Printf("Instruction sum v2: %d\n", sumv2)

}

func addresses(addrs []int, base int, floating int, mask int) []int {
	if mask == 0 {
		return addrs
	}

	addrs = append(addrs, base)
	addrs = addresses(addrs, base, floating, mask>>1)

	bit := floating & mask
	if bit > 0 {
		base ^= bit
		addrs = append(addrs, base)
		addrs = addresses(addrs, base, floating, mask>>1)
	}
	return addrs
}
func entries() []program {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	programs := []program{}
	prog := program{}
	for scan.Scan() {
		line := scan.Text()
		switch {
		case strings.HasPrefix(line, "mask"):
			if len(prog.instructions) > 0 {
				programs = append(programs, prog)
				prog = program{}
			}
			mask := strings.Split(line, " = ")[1]
			for _, r := range mask {
				prog.and <<= 1
				prog.or <<= 1
				prog.floating <<= 1
				switch {
				case r == '0':
					prog.and |= 1
				case r == '1':
					prog.or |= 1
				default:
					prog.floating |= 1
				}
			}
			prog.and = prog.and ^ 0xFFFFFFFFFF
		default:
			parts := strings.Split(line, " = ")
			location := strings.ReplaceAll(parts[0], "mem[", "")
			location = strings.ReplaceAll(location, "]", "")
			loc, _ := strconv.Atoi(location)
			val, _ := strconv.Atoi(parts[1])
			ins := instruction{
				location: loc,
				value:    val,
			}
			prog.instructions = append(prog.instructions, ins)
		}
	}
	if len(prog.instructions) > 0 {
		programs = append(programs, prog)
	}

	return programs
}
