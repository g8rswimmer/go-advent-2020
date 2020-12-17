package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numRange struct {
	min int
	max int
}

func (n numRange) in(num int) bool {
	return num >= n.min && num <= n.max
}

type rule struct {
	field  string
	ranges []numRange
}

func (r rule) in(num int) bool {
	for _, nr := range r.ranges {
		if nr.in(num) {
			return true
		}
	}
	return false
}

type ticket struct {
	values []int
}

type input struct {
	rules         []rule
	myTicket      ticket
	nearbyTickets []ticket
}

func main() {
	ticketInputs := entries()

	nums := map[int]interface{}{}
	for _, r := range ticketInputs.rules {
		for _, nr := range r.ranges {
			for val := nr.min; val <= nr.max; val++ {
				nums[val] = nil
			}
		}
	}

	errorRate := 0
	validTickets := []ticket{}
	for _, nearby := range ticketInputs.nearbyTickets {
		valid := true
		for _, val := range nearby.values {
			if _, has := nums[val]; has == false {
				errorRate += val
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, nearby)
		}
	}
	ticketInputs.nearbyTickets = validTickets

	fmt.Printf("Ticket Scanning error rate: %d\n", errorRate)

	departures := 1
	fields := fields(ticketInputs.rules, ticketInputs.nearbyTickets)
	for i, field := range fields {
		if strings.HasPrefix(field, "departure") {
			fmt.Println(i+1, field, ticketInputs.myTicket.values[i])
			departures *= ticketInputs.myTicket.values[i]
		}
	}
	fmt.Printf("Departure Ticket Value: %d\n", departures)
}

func fields(rules []rule, tickets []ticket) []string {
	possible := make([]map[string]interface{}, len(tickets[0].values))
	for i := range possible {
		possible[i] = mapFields(rules, tickets, i)
	}

	fields := make([]string, len(possible))
	found := 0
	for {
		for i := range possible {
			if len(possible[i]) == 1 {
				for f := range possible[i] {
					fields[i] = f
				}
				found++
			}
		}

		if found == len(fields) {
			break
		}

		for i := range possible {
			for _, f := range fields {
				if len(f) == 0 {
					continue
				}
				if _, has := possible[i][f]; has {
					delete(possible[i], f)
				}
			}
		}

	}

	return fields
}

func mapFields(rules []rule, tickets []ticket, idx int) map[string]interface{} {
	fieldRules := map[string]interface{}{}
	for _, r := range rules {
		found := true
		for _, t := range tickets {
			if r.in(t.values[idx]) == false {
				found = false
			}
		}
		if found {
			fieldRules[r.field] = nil
		}
	}
	return fieldRules
}

func entries() input {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	ip := input{}
	myTicket := false
	nearTickets := false
	for scan.Scan() {
		line := scan.Text()

		switch {
		case len(line) == 0:
		case myTicket:
			mvals := strings.Split(line, ",")
			for _, val := range mvals {
				num, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				ip.myTicket.values = append(ip.myTicket.values, num)
			}
			myTicket = false
		case nearTickets:
			nvals := strings.Split(line, ",")
			vals := []int{}
			for _, val := range nvals {
				num, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				vals = append(vals, num)
			}
			ip.nearbyTickets = append(ip.nearbyTickets, ticket{values: vals})
		case line == "your ticket:":
			myTicket = true
		case line == "nearby tickets:":
			nearTickets = true
		default:
			parts := strings.Split(line, ":")
			r := rule{
				field: parts[0],
			}
			parts[1] = strings.Trim(parts[1], " ")
			ranges := strings.Split(parts[1], " ")
			for i := 0; i < len(ranges); i++ {
				if strings.Contains(ranges[i], "-") == false {
					continue
				}
				strs := strings.Split(ranges[i], "-")
				min, err := strconv.Atoi(strs[0])
				if err != nil {
					panic(err)
				}
				max, err := strconv.Atoi(strs[1])
				if err != nil {
					panic(err)
				}
				r.ranges = append(r.ranges, numRange{min: min, max: max})
			}
			ip.rules = append(ip.rules, r)
		}
	}
	return ip
}
