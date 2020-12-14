package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type schedule struct {
	timestamp int
	busIDs    []int
	routes    [][]int
}

func main() {
	busSchedule := entries()

	diff := math.MaxInt32
	diffID := -1
	for _, id := range busSchedule.busIDs {
		trips := busSchedule.timestamp / id
		if busSchedule.timestamp%id > 0 {
			trips++
		}

		depart := trips * id
		d := depart - busSchedule.timestamp
		if diff > d {
			diff = d
			diffID = id
		}
	}
	fmt.Printf("Earlist Depart (%d) * Bus ID (%d): %d\n", diff, diffID, diff*diffID)

	sort.Slice(busSchedule.routes, func(i, j int) bool {
		return busSchedule.routes[i][0] > busSchedule.routes[j][0]
	})
	fmt.Printf("%+v\n", busSchedule)

	lcm := -1
	ts := -1
	idx := 0
	for {
		route := busSchedule.routes[idx]
		if lcm == -1 {
			lcm = route[0]
			ts = route[0] - route[1]
			idx++
			continue
		}

		if t := ts + route[1]; t%route[0] == 0 {
			idx++
			if idx == len(busSchedule.routes) {
				break
			}
			lcm *= route[0]
			continue
		}
		ts += lcm
	}
	fmt.Printf("Earlist Seq timestamp %d\n", ts)
}

func entries() schedule {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	sched := schedule{}
	if scan.Scan() {
		txt := scan.Text()
		ts, _ := strconv.Atoi(txt)
		sched.timestamp = ts
	}
	if scan.Scan() {
		txt := scan.Text()
		ids := strings.Split(txt, ",")
		sched.routes = [][]int{}
		for i, id := range ids {
			if id == "x" {
				continue
			}
			num, _ := strconv.Atoi(id)
			sched.busIDs = append(sched.busIDs, num)
			sched.routes = append(sched.routes, []int{num, i})
		}
	}
	return sched
}
