package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rule, compositeRules := entries()
	for len(compositeRules) > 0 {
		rule, compositeRules = filter(rule, compositeRules)
	}

	zeros := map[string]interface{}{}
	for _, msg := range rule[0] {
		zeros[msg] = nil
	}

	msgs := messages()
	valid := 0
	for _, msg := range msgs {
		if _, has := zeros[msg]; has {
			valid++
		}
	}
	fmt.Printf("Valid message for rule 0: %d\n", valid)

	updateRule, updateCompositRules := updateEntries()
	for {
		currComposite := len(updateCompositRules)
		updateRule, updateCompositRules = filter(updateRule, updateCompositRules)
		if len(updateCompositRules) == currComposite {
			break
		}
	}
	updateValid := 0
	for _, msg := range msgs {
		pre := prefixes(msg, rule[42])
		suf := suffixes(msg, rule[31])

		l := (pre * len(rule[42][0])) + (suf * len(rule[31][0]))
		if l != len(msg) {
			continue
		}
		if suf == 0 {
			continue
		}
		if pre <= suf {
			continue
		}
		updateValid++
	}
	fmt.Printf("Update valid message for rule 0: %d\n", updateValid)

}

func prefixes(msg string, prefix []string) int {
	found := false
	for _, p := range prefix {
		if strings.HasPrefix(msg, p) {
			found = true
			break
		}
	}
	if found == false {
		return 0
	}
	return prefixes(msg[len(prefix[0]):], prefix) + 1
}

func suffixes(msg string, suffix []string) int {
	found := false
	for _, p := range suffix {
		if strings.HasSuffix(msg, p) {
			found = true
			break
		}
	}
	if found == false {
		return 0
	}
	end := len(msg) - len(suffix[0])
	return suffixes(msg[:end], suffix) + 1
}

func filter(rule map[int][]string, compositeRules map[int][][]int) (map[int][]string, map[int][][]int) {
	remove := []int{}
	for id, crs := range compositeRules {
		can := true
		for _, cr := range crs {
			if contains(rule, cr) == false {
				can = false
				break
			}
		}

		if can {
			rules := []string{}
			for _, cr := range crs {
				rules = append(rules, generate(rule, []string{""}, 0, cr)...)
			}
			rule[id] = rules
			remove = append(remove, id)
		}
	}
	for _, id := range remove {
		delete(compositeRules, id)
	}
	return rule, compositeRules
}
func contains(rule map[int][]string, rules []int) bool {
	for _, r := range rules {
		if _, has := rule[r]; has == false {
			return false
		}
	}
	return true
}

func generate(rule map[int][]string, curr []string, idx int, rules []int) []string {
	switch {
	case idx == len(rules):
		return curr
	default:
	}

	results := []string{}
	additions := rule[rules[idx]]

	for _, c := range curr {
		for _, add := range additions {
			results = append(results, c+add)
		}
	}
	return generate(rule, results, idx+1, rules)
}

func messages() []string {
	in, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	msgs := []string{}

	for scan.Scan() {
		msgs = append(msgs, scan.Text())
	}
	return msgs
}

func entries() (map[int][]string, map[int][][]int) {
	in, err := os.Open("rules.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	rules := map[int][]string{}
	composites := map[int][][]int{}

	for scan.Scan() {
		entry := scan.Text()
		if isSingle(entry) {
			id, lines := single(entry)
			rules[id] = lines
			continue
		}
		id, compositeRules := composite(entry)
		composites[id] = compositeRules
	}
	return rules, composites
}

func updateEntries() (map[int][]string, map[int][][]int) {
	in, err := os.Open("rules_update.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	rules := map[int][]string{}
	composites := map[int][][]int{}

	for scan.Scan() {
		entry := scan.Text()
		if isSingle(entry) {
			id, lines := single(entry)
			rules[id] = lines
			continue
		}
		id, compositeRules := composite(entry)
		composites[id] = compositeRules
	}
	return rules, composites
}

func isSingle(entry string) bool {
	return strings.Contains(entry, `"`)
}

func single(entry string) (int, []string) {
	parts := strings.Split(entry, ":")
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	valid := strings.Trim(parts[1], " ")
	valid = strings.ReplaceAll(valid, `"`, "")
	return id, []string{valid}
}

func composite(entry string) (int, [][]int) {
	parts := strings.Split(entry, ":")
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	rs := strings.Split(parts[1], "|")
	crs := make([][]int, len(rs))
	for i, r := range rs {
		crs[i] = rules(r)
	}
	return id, crs
}

func rules(r string) []int {
	r = strings.Trim(r, " ")
	numStrs := strings.Split(r, " ")
	nums := make([]int, len(numStrs))
	for i, ns := range numStrs {
		n, err := strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}
