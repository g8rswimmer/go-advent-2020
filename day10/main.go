package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	jolts := entries()
	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	distrubution := map[int]int{}
	for i := 0; i < len(jolts)-1; i++ {
		diff := jolts[i+1] - jolts[i]
		distrubution[diff]++
	}
	result := distrubution[1] * distrubution[3]
	fmt.Printf("1 diff: %d 3 diff: %d answer %d\n", distrubution[1], distrubution[3], result)

	mem := make([]int, len(jolts))
	for i := range mem {
		mem[i] = -1
	}
	ways := paths(jolts, 0, mem)
	fmt.Printf("number of ways: %d\n", ways)
}

func paths(jolts []int, idx int, mem []int) int {
	if idx == len(jolts)-1 {
		return 1
	}
	if mem[idx] != -1 {
		return mem[idx]
	}
	result := 0
	nidx := idx + 1
	for nidx < len(jolts) {
		diff := jolts[nidx] - jolts[idx]
		if diff > 3 {
			break
		}
		result += paths(jolts, nidx, mem)
		nidx++
	}
	mem[idx] = result
	return result
}
func entries() []int {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	jolts := []int{}
	for scan.Scan() {
		line := scan.Text()
		jolt, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		jolts = append(jolts, jolt)
	}
	return jolts
}
