package main

import "fmt"

var seed = []int{9, 19, 1, 6, 0, 5, 4}

func main() {
	seen := map[int]int{}
	for i := 0; i < len(seed)-1; i++ {
		seen[seed[i]] = i + 1
	}
	last := seed[len(seed)-1]
	turn := len(seed)
	fmt.Printf("2020th number spoken %d\n", game(turn, last, 2020, seen))
	fmt.Printf("30000000th number spoken %d\n", game(turn, last, 30000000, seen))
}

func game(turn, last, length int, seen map[int]int) int {
	s := map[int]int{}
	for k, v := range seen {
		s[k] = v
	}
	for turn < length {
		last, seen = next(turn, last, s)
		turn++
	}
	return last
}
func next(turn, last int, seen map[int]int) (int, map[int]int) {
	lt, has := seen[last]
	seen[last] = turn
	if has == false {
		return 0, seen
	}
	return turn - lt, seen
}
