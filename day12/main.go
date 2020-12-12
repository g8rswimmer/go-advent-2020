package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ship struct {
	facing int
	pos    []int
}

type instruction struct {
	cmd   string
	value int
}

type waypoint struct {
	facing int
	dis    int
}

type wayward struct {
	pnt []waypoint
}

var rotation = []int{1, 1, -1, -1}

func (w *wayward) rotateRight(degrees int) {
	rotations := degrees / 90

	for i := range w.pnt {
		w.pnt[i].facing += rotations
		w.pnt[i].facing = w.pnt[i].facing % len(rotation)
	}
}
func (w *wayward) rotateLeft(degrees int) {
	rotations := degrees / 90
	for i := range w.pnt {
		w.pnt[i].facing -= rotations
		if w.pnt[i].facing < 0 {
			w.pnt[i].facing += len(rotation)
		}
	}
}

func (w *wayward) move(units []int) {
	for i := range w.pnt {
		if w.pnt[i].facing == 0 || w.pnt[i].facing == 2 {
			w.pnt[i].dis += (units[0] * rotation[w.pnt[i].facing])
			continue
		}
		w.pnt[i].dis += (units[1] * rotation[w.pnt[i].facing])
	}
}

func (w wayward) point() []int {
	pnt := make([]int, 2)
	for i := range w.pnt {
		if w.pnt[i].facing == 0 || w.pnt[i].facing == 2 {
			pnt[0] = w.pnt[i].dis * rotation[w.pnt[i].facing]
			continue
		}
		pnt[1] = w.pnt[i].dis * rotation[w.pnt[i].facing]
	}
	return pnt
}

var direction = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {
	instructions := entries()

	md := naviagte(instructions, ship{facing: 1, pos: make([]int, 2)})
	fmt.Printf("Manhatten Distance: %d\n", md)

	ww := &wayward{
		pnt: []waypoint{
			{
				facing: 1,
				dis:    10,
			},
			{
				facing: 0,
				dis:    1,
			},
		},
	}
	wwmd := navigateWayPoints(instructions, ship{pos: make([]int, 2)}, ww)
	fmt.Printf("Way Manhatten Distance: %d\n", wwmd)

}

func navigateWayPoints(instructions []instruction, s ship, ww *wayward) int {
	for _, ins := range instructions {
		switch ins.cmd {
		case "F":
			pnt := ww.point()
			s.pos[0] += (ins.value * pnt[0])
			s.pos[1] += (ins.value * pnt[1])
		case "N":
			ww.move([]int{ins.value, 0})
		case "S":
			ww.move([]int{-ins.value, 0})
		case "E":
			ww.move([]int{0, ins.value})
		case "W":
			ww.move([]int{0, -ins.value})
		case "R":
			ww.rotateRight(ins.value)
		case "L":
			ww.rotateLeft(ins.value)
		}
	}
	return abs(s.pos[0]) + abs(s.pos[1])
}

func naviagte(instructions []instruction, s ship) int {
	for _, ins := range instructions {
		move := make([]int, 2)
		facing := 0
		switch ins.cmd {
		case "F":
			move[0] = direction[s.facing][0] * ins.value
			move[1] = direction[s.facing][1] * ins.value
		case "N":
			move[0] = ins.value
		case "S":
			move[0] = ins.value * -1
		case "E":
			move[1] = ins.value
		case "W":
			move[1] = ins.value * -1
		case "R":
			facing = ins.value / 90
		case "L":
			facing = (ins.value / 90) * -1
		}
		facing += s.facing
		switch {
		case facing >= 0:
			s.facing = facing % len(direction)
		default:
			s.facing = len(direction) + facing
		}
		s.pos[0] += move[0]
		s.pos[1] += move[1]
	}
	return abs(s.pos[0]) + abs(s.pos[1])
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func entries() []instruction {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	instructions := []instruction{}
	for scan.Scan() {
		line := scan.Text()
		cmd := line[:1]
		value, _ := strconv.Atoi(line[1:])
		ins := instruction{
			cmd:   cmd,
			value: value,
		}
		instructions = append(instructions, ins)
	}

	return instructions
}
