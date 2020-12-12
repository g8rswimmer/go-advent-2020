package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	groups := entries()
	allYes := 0
	for _, g := range groups {
		allYes += g.answers()
	}
	fmt.Printf("Number of yeses %d\n", allYes)

	commonYes := 0
	for _, g := range groups {
		commonYes += g.common()
	}
	fmt.Printf("Common yeses %d\n", commonYes)

}

type group struct {
	forms []string
}

func (g group) answers() int {
	yes := map[rune]struct{}{}
	for _, form := range g.forms {
		for _, a := range form {
			yes[a] = struct{}{}
		}
	}
	return len(yes)
}

func (g group) common() int {
	yes := map[rune]int{}
	for _, form := range g.forms {
		for _, a := range form {
			yes[a]++
		}
	}
	c := 0
	people := len(g.forms)
	for _, cnt := range yes {
		if cnt == people {
			c++
		}
	}
	return c
}

func entries() []group {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	groups := []group{}
	g := group{}
	for scan.Scan() {
		line := scan.Text()
		if len(line) == 0 {
			groups = append(groups, g)
			g = group{}
			continue
		}
		g.forms = append(g.forms, line)
	}
	if len(g.forms) > 0 {
		groups = append(groups, g)
	}
	return groups
}
