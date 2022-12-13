package main

import (
	"encoding/json"
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day13/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day13/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0

	pairs := [][][]any{}
	pair := [][]any{}
	for _, line := range lines {
		if line == "" {
			pairs = append(pairs, pair)
			pair = [][]any{}
			continue
		}
		var result []any
		json.Unmarshal([]byte(line), &result)
		pair = append(pair, result)

	}

	for i, p := range pairs {
		if a := compare(p[0], p[1]); a == 1 {
			answer += i + 1
		}
	}

	fmt.Printf("Answer 1 : %d\n", answer)
}

func compare(left []any, right []any) int {

	min := 0
	if len(left) < len(right) {
		min = len(left)
	} else {
		min = len(right)
	}

	for i := 0; i < min; i++ {
		l := left[i]
		r := right[i]
		var lT string
		var rT string
		switch l.(type) {
		case float64:
			lT = "f"
		default:
			lT = "a"
		}

		switch r.(type) {
		case float64:
			rT = "f"
		default:
			rT = "a"
		}

		if lT == "f" && rT == "f" {
			if l.(float64) > r.(float64) {
				return 0
			} else if l.(float64) < r.(float64) {
				return 1
			}
		} else if lT == "f" && rT == "a" {
			newL := []any{l}
			a := compare(newL, r.([]any))
			if a != 2 {
				return a
			}
		} else if lT == "a" && rT == "f" {
			newR := []any{r}
			a := compare(l.([]any), newR)
			if a != 2 {
				return a
			}
		} else if lT == "a" && rT == "a" {
			a := compare(l.([]any), r.([]any))
			if a != 2 {
				return a
			}
		}
	}

	if len(left) < len(right) {
		return 1
	} else if len(left) > len(right) {
		return 0
	} else {
		return 2
	}
}

func part2(lines []string) {
	answer := 1

	list := [][]any{}
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var result []any
		json.Unmarshal([]byte(line), &result)
		list = append(list, result)

	}

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if a := compare(list[i], list[j]); a == 1 {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}
	for i := 0; i < len(list); i++ {
		val, _ := json.Marshal(list[i])
		if string(val) == "[[6]]" || string(val) == "[[2]]" {
			answer *= (i + 1)
		}
	}
	fmt.Printf("Answer 2 : %d\n", answer)
}
