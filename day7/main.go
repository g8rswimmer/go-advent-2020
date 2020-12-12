package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bags := entries()

	innerGold := outerBags(bags, map[string]interface{}{"shiny gold": nil}, map[string]interface{}{})

	fmt.Printf("number of bags that fit shiny gold %d\n", len(innerGold))

	innerBags := innerBags(bags, "shiny gold", 0, 1)
	fmt.Printf("number of inner bags %d\n", innerBags)
}

func outerBags(bags map[string]map[string]int, inner map[string]interface{}, curr map[string]interface{}) map[string]interface{} {
	if len(inner) == 0 {
		return curr
	}
	next := map[string]interface{}{}

	for k, v := range bags {
		for in := range inner {
			if _, has := v[in]; has {
				if _, hasC := curr[k]; hasC == false {
					next[k] = nil
					curr[k] = nil
				}
			}
		}
	}
	return outerBags(bags, next, curr)
}

func innerBags(bags map[string]map[string]int, bag string, curr, level int) int {
	inner := bags[bag]
	if len(inner) == 0 {
		return 1
	}
	incurr := curr
	for in, cnt := range inner {
		add := level * cnt
		incurr += add
		curr += innerBags(bags, in, incurr, add)
	}
	return curr
}
func entries() map[string]map[string]int {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	bags := map[string]map[string]int{}
	for scan.Scan() {
		line := scan.Text()
		line = strings.ReplaceAll(line, ".", "")

		values := strings.Split(line, " contain ")

		key := strings.ReplaceAll(values[0], "bags", "")
		key = strings.Trim(key, " ")
		bags[key] = map[string]int{}

		if strings.Contains(values[1], "no other bags") {
			continue
		}

		inner := strings.Split(values[1], ", ")
		for _, in := range inner {
			cnts := strings.Split(in, " ")
			cnt, _ := strconv.Atoi(cnts[0])
			in = strings.ReplaceAll(in, "bags", "")
			in = strings.ReplaceAll(in, "bag", "")
			in = strings.Trim(in, cnts[0])
			in = strings.Trim(in, " ")
			bags[key][in] = cnt
		}
	}

	return bags
}
