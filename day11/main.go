package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	orig := entries()
	layout := orig
	for {
		next := arrange(layout)
		if compareLayout(layout, next) {
			break
		}
		layout = next
	}
	fmt.Printf("Occupied seats: %d\n", occupiedSeats(layout))

	viewLayout := orig
	for {
		next := arrangeView(viewLayout)
		if compareLayout(viewLayout, next) {
			break
		}
		viewLayout = next
	}
	fmt.Printf("Occupied view seats: %d\n", occupiedSeats(viewLayout))
}

func occupiedSeats(layout [][]rune) int {
	seats := 0
	for _, row := range layout {
		for _, seat := range row {
			if seat == '#' {
				seats++
			}
		}
	}
	return seats
}
func arrange(layout [][]rune) [][]rune {
	next := make([][]rune, len(layout))
	for r, row := range layout {
		next[r] = make([]rune, len(row))
		for c, seat := range row {
			seats := adjacentSeats(layout, r, c)
			occ := occupied(seats)
			switch {
			case seat == 'L' && occ == 0:
				next[r][c] = '#'
			case seat == '#' && occ >= 4:
				next[r][c] = 'L'
			default:
				next[r][c] = seat
			}
		}
	}
	return next
}

func arrangeView(layout [][]rune) [][]rune {
	next := make([][]rune, len(layout))
	for r, row := range layout {
		next[r] = make([]rune, len(row))
		for c, seat := range row {
			seats := adjacentSeatsView(layout, r, c)
			occ := occupied(seats)
			switch {
			case seat == 'L' && occ == 0:
				next[r][c] = '#'
			case seat == '#' && occ >= 5:
				next[r][c] = 'L'
			default:
				next[r][c] = seat
			}
		}
	}
	return next
}

func occupied(seats []rune) int {
	o := 0
	for _, seat := range seats {
		if seat == '#' {
			o++
		}
	}
	return o
}

func adjacentSeatsView(layout [][]rune, r, c int) []rune {
	views := [][]int{{-1, 0}, {-1, -1}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	seats := []rune{}
	for _, view := range views {
		nr := r
		nc := c
		for {
			nr = nr + view[0]
			if nr < 0 || nr >= len(layout) {
				break
			}
			nc = nc + view[1]
			if nc < 0 || nc >= len(layout[r]) {
				break
			}

			if layout[nr][nc] == '.' {
				continue
			}
			seats = append(seats, layout[nr][nc])
			break
		}
	}
	return seats
}

func adjacentSeats(layout [][]rune, r, c int) []rune {
	seats := []rune{}
	up := r - 1
	down := r + 1
	left := c - 1
	right := c + 1

	if up >= 0 {
		if left >= 0 {
			seats = append(seats, layout[up][left])
		}
		seats = append(seats, layout[up][c])
		if right < len(layout[r]) {
			seats = append(seats, layout[up][right])
		}
	}
	if left >= 0 {
		seats = append(seats, layout[r][left])
	}
	if right < len(layout[r]) {
		seats = append(seats, layout[r][right])
	}
	if down < len(layout) {
		if left >= 0 {
			seats = append(seats, layout[down][left])
		}
		seats = append(seats, layout[down][c])
		if right < len(layout[r]) {
			seats = append(seats, layout[down][right])
		}
	}

	return seats
}
func compareLayout(layout1, layout2 [][]rune) bool {
	if len(layout1) != len(layout2) {
		return false
	}
	for i := range layout1 {
		if len(layout1[i]) != len(layout2[i]) {
			return false
		}
		for j := range layout1[i] {
			if layout1[i][j] != layout2[i][j] {
				return false
			}
		}
	}
	return true
}

func entries() [][]rune {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	layout := [][]rune{}
	for scan.Scan() {
		line := scan.Text()
		row := make([]rune, len(line))
		for i, r := range line {
			row[i] = r
		}
		layout = append(layout, row)
	}
	return layout
}
