package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type validator interface {
	valid(string) bool
}

var required = map[string]validator{
	"byr": &numrange{
		min: 1920,
		max: 2002,
	},
	"iyr": &numrange{
		min: 2010,
		max: 2020,
	},
	"eyr": &numrange{
		min: 2020,
		max: 2030,
	},
	"hgt": &height{
		cm: numrange{
			min: 150,
			max: 193,
		},
		in: numrange{
			min: 59,
			max: 76,
		},
	},
	"hcl": &hair{},
	"ecl": &eye{
		color: map[string]interface{}{
			"amb": nil,
			"blu": nil,
			"brn": nil,
			"gry": nil,
			"grn": nil,
			"hzl": nil,
			"oth": nil,
		},
	},
	"pid": &id{},
}

func main() {
	passports := entries()
	valid := 0
	for _, p := range passports {
		if p.valid(required) {
			valid++
		}
	}
	fmt.Printf("Valid passports %d\n", valid)
}

type passport struct {
	fields map[string]string
}

func (p passport) valid(required map[string]validator) bool {
	for field, f := range required {
		if value, has := p.fields[field]; has == false || f.valid(value) == false {
			return false
		}
	}
	return true
}

type numrange struct {
	min int
	max int
}

func (r numrange) valid(str string) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= r.min && num <= r.max
}

type height struct {
	cm numrange
	in numrange
}

func (h height) valid(str string) bool {
	switch {
	case strings.HasSuffix(str, "in"):
		return h.in.valid(strings.ReplaceAll(str, "in", ""))
	case strings.HasSuffix(str, "cm"):
		return h.cm.valid(strings.ReplaceAll(str, "cm", ""))
	default:
		return false
	}
}

type hair struct{}

func (h hair) valid(str string) bool {
	if len(str) != 7 {
		return false
	}
	if str[0] != '#' {
		return false
	}
	for i := 1; i < len(str); i++ {
		switch {
		case str[i] >= '0' && str[i] <= '9':
		case str[i] >= 'a' && str[i] <= 'f':
		default:
			return false
		}
	}
	return true
}

type eye struct {
	color map[string]interface{}
}

func (e eye) valid(str string) bool {
	_, has := e.color[str]
	return has
}

type id struct{}

func (i id) valid(str string) bool {
	if len(str) != 9 {
		return false
	}
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func entries() []passport {
	in, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	passports := []passport{}
	lines := []string{}
	for scan.Scan() {
		line := scan.Text()
		if len(line) == 0 {
			pp := passport{
				fields: map[string]string{},
			}
			for _, l := range lines {
				entries := strings.Split(l, " ")
				for _, entry := range entries {
					kv := strings.Split(entry, ":")
					pp.fields[kv[0]] = kv[1]
				}
			}
			passports = append(passports, pp)
			lines = []string{}
			continue
		}
		lines = append(lines, line)
	}
	return passports
}
