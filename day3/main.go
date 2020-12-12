package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := entries()
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result := 1
	for _, s := range slopes {
		trees := slope(s[0], s[1], lines)
		fmt.Printf("Number of trees %d\n", trees)
		result *= trees
	}
	fmt.Printf("Result %d\n", result)
}

func slope(right, down int, lines []string) int {
	trees := 0
	pos := 0
	idx := 0

	for idx < len(lines) {
		line := lines[idx]
		pos = pos % len(line)
		if line[pos] == '#' {
			trees++
		}
		pos += right
		idx += down
	}
	return trees
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
