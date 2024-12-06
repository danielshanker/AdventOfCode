package main

import (
	"flag"
	"sort"
	"strings"
	. "utils"

	"golang.org/x/exp/slices"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()
	Start(test, 5, part1, part2)
}

type rule struct {
	first  int
	second int
}

func part1(lines []string) int {
	answer := 0

	correct, _, _ := getCorrectAndIncorrect(lines)

	for _, val := range correct {
		answer += val[(len(val)-1)/2]
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	_, incorrect, rules := getCorrectAndIncorrect(lines)
	for _, val := range incorrect {
		sort.Slice(val, func(i, j int) bool {
			ret := false
			for _, rule := range rules {
				if val[i] == rule.first && val[j] == rule.second {
					ret = true
					break
				}
			}
			return ret
		})
	}

	for _, val := range incorrect {
		answer += val[(len(val)-1)/2]
	}

	return answer
}

func getCorrectAndIncorrect(lines []string) ([][]int, [][]int, []rule) {
	correct := [][]int{}
	incorrect := [][]int{}

	rules := []rule{}
	updates := [][]int{}

	afterSpace := false
	for _, line := range lines {
		if line == "" {
			afterSpace = true
			continue
		}
		if !afterSpace {
			vals := strings.Split(line, "|")
			r := rule{
				first:  S2i(vals[0]),
				second: S2i(vals[1]),
			}
			rules = append(rules, r)
		} else {

			vals := strings.Split(line, ",")
			update := []int{}
			for _, val := range vals {
				update = append(update, S2i(val))
			}
			updates = append(updates, update)
		}
	}

	for _, update := range updates {
		good := true
		for i, inst := range update {
			for _, rule := range rules {
				if inst == rule.second {
					before := update[:i]
					after := update[i:]
					if !slices.Contains(before, rule.first) && slices.Contains(after, rule.first) {
						good = false
						break
					}
				}
			}
		}
		if !good {
			incorrect = append(incorrect, update)
		} else {
			correct = append(correct, update)
		}
	}

	return correct, incorrect, rules
}
