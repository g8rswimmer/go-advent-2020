package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	validators := entries()
	result := 0
	for _, v := range validators {
		if v.valid() {
			result++
		}
	}
	fmt.Printf("Number of valid passwords %d\n", result)
}

type validator struct {
	min      int
	max      int
	char     rune
	password string
}

func (v validator) valid() bool {
	times := 0
	for i, r := range v.password {
		switch {
		case i+1 == v.min && r == v.char:
			times++
		case i+1 == v.max && r == v.char:
			times++
		default:
		}
	}
	return times == 1
}

func entries() []validator {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	validators := []validator{}
	for scan.Scan() {
		line := scan.Text()
		sections := strings.Split(line, " ")
		rs := strings.Split(sections[0], "-")
		v := validator{
			min: func() int {
				n, _ := strconv.Atoi(rs[0])
				return n
			}(),
			max: func() int {
				n, _ := strconv.Atoi(rs[1])
				return n
			}(),
			char:     rune(sections[1][0]),
			password: sections[2],
		}
		validators = append(validators, v)
	}
	return validators
}
