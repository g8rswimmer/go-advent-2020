package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	eqs := entries()
	basicResult := 0
	for _, eq := range eqs {
		basicResult += basic(eq)
	}
	fmt.Printf("Sum of basic homework: %d\n", basicResult)

	advancedResult := 0
	for _, eq := range eqs {
		advancedResult += advanced(eq)
	}
	fmt.Printf("Sum of advanced homework: %d\n", advancedResult)
}

func basic(eq string) int {
	nums := []int{}
	operations := []rune{}
	num := 0

	params := []rune(eq)
	idx := 0
	for idx < len(params) {
		r := params[idx]
		switch {
		case r >= '0' && r <= '9':
			num *= 10
			num += int(r - '0')
		case r == '+' || r == '*':
			operations = append(operations, r)
			nums = append(nums, num)
			num = 0
		case r == '(':
			inner := []rune{}
			p := 0
			for {
				ir := params[idx]
				switch {
				case ir == '(':
					p++
				case ir == ')':
					p--
				default:
				}
				inner = append(inner, ir)
				if p == 0 {
					break
				}
				idx++
			}
			num = basic(string(inner[1 : len(inner)-1]))
		default:
		}
		idx++

	}
	nums = append(nums, num)

	for len(operations) > 0 {
		switch operations[0] {
		case '+':
			nums[1] += nums[0]
		case '*':
			nums[1] *= nums[0]
		}
		operations = operations[1:]
		nums = nums[1:]
	}
	return nums[0]
}

func advanced(eq string) int {
	nums := []int{}
	operations := []rune{}
	num := 0

	params := []rune(eq)
	idx := 0
	for idx < len(params) {
		r := params[idx]
		switch {
		case r >= '0' && r <= '9':
			num *= 10
			num += int(r - '0')
		case r == '+' || r == '*':
			operations = append(operations, r)
			nums = append(nums, num)
			num = 0
		case r == '(':
			inner := []rune{}
			p := 0
			for {
				ir := params[idx]
				switch {
				case ir == '(':
					p++
				case ir == ')':
					p--
				default:
				}
				inner = append(inner, ir)
				if p == 0 {
					break
				}
				idx++
			}
			num = advanced(string(inner[1 : len(inner)-1]))
		default:
		}
		idx++

	}
	nums = append(nums, num)

	snums := []int{}
	sop := []rune{}

	for len(operations) > 0 {
		switch operations[0] {
		case '+':
			nums[1] += nums[0]
		case '*':
			snums = append(snums, nums[0])
			sop = append(sop, '*')
		}
		operations = operations[1:]
		nums = nums[1:]
	}
	snums = append(snums, nums[0])

	for len(sop) > 0 {
		switch sop[0] {
		case '*':
			snums[1] *= snums[0]
		}
		sop = sop[1:]
		snums = snums[1:]
	}
	return snums[0]
}

func entries() []string {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	lines := []string{}
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines
}
