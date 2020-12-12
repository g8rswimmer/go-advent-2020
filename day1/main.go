package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	nums := entries()

	fmt.Printf("Two Sum: Answer: %d\n", twoSum(2020, nums))

	for i := 0; i < len(nums); i++ {
		sum := 2020 - nums[i]
		if result := twoSum(sum, nums[i+1:]); result != math.MinInt32 {
			fmt.Printf("Three Sum Answer: %d\n", nums[i]*result)
		}
	}
}

func twoSum(sum int, nums []int) int {
	seen := map[int]interface{}{}
	for _, num := range nums {
		need := sum - num
		if _, has := seen[need]; has {
			return num * need
		}
		seen[num] = nil
	}
	return math.MinInt32
}

func entries() []int {
	in, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	nums := []int{}
	for scan.Scan() {
		num, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
