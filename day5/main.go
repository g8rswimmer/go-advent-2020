package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	assignments := entries()

	result := 0
	seats := []int{}
	for _, assignment := range assignments {
		id := planeRow(assignment[:7])*8 + planeSeat(assignment[7:])
		if id > result {
			result = id
		}
		seats = append(seats, id)
	}

	fmt.Printf("Highest seat id: %d\n", result)
	sort.Ints(seats)
	for i := 1; i < len(seats); i++ {
		if diff := seats[i] - seats[i-1]; diff != 1 {
			fmt.Println(seats[i], seats[i-1])
			break
		}
	}
}

func planeRow(assignment string) int {
	rows := 128
	row := 0
	for _, r := range assignment {
		rows >>= 1
		if r == 'B' {
			row += rows
		}
	}
	return row
}

func planeSeat(assignment string) int {
	seats := 8
	seat := 0
	for _, s := range assignment {
		seats >>= 1
		if s == 'R' {
			seat += seats
		}
	}
	return seat
}

func entries() []string {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	assignments := []string{}
	for scan.Scan() {
		assignments = append(assignments, scan.Text())
	}
	return assignments
}
