package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	codes := entries()

	fw := firstWrong(codes, 25)
	fmt.Printf("First wrong code %d\n", fw)

	ew := weakness(codes, fw)
	fmt.Printf("Weakness %d\n", ew)

}

func firstWrong(codes []int, preamble int) int {
	positions := map[int][]int{}
	for i, code := range codes {
		positions[code] = append(positions[code], i)
	}

	for i := preamble; i < len(codes); i++ {
		found := false
		start := i - preamble
		end := i + preamble
		for j := start; j < end; j++ {
			need := codes[i] - codes[j]
			if idxs, has := positions[need]; has {
				for _, idx := range idxs {
					if idx >= start && idx < end {
						found = true
						break
					}
				}
			}
		}
		if found == false {
			return codes[i]
		}
	}
	return math.MinInt32
}

func weakness(codes []int, wrong int) int {
	for i := 0; i < len(codes); i++ {
		sum := 0
		nums := []int{}
		for j := i; j < len(codes); j++ {
			sum += codes[j]
			if sum >= wrong {
				break
			}
			nums = append(nums, codes[j])
		}
		if sum == wrong {
			sort.Ints(nums)
			return nums[0] + nums[len(nums)-1]
		}
	}
	return math.MinInt32
}
func entries() []int {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	codes := []int{}
	for scan.Scan() {
		line := scan.Text()
		code, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		codes = append(codes, code)
	}
	return codes
}
