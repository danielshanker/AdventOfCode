package main

import (
	"flag"
	"fmt"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 19, part1, part2, 6, 16)

}

func part1(lines []string) int {
	answer := 0

	towels := strings.Split(lines[0], ", ")
	desired := []string{}

	for i, line := range lines {
		if i < 2 {
			continue
		}
		desired = append(desired, line)
	}

	for _, t := range desired {
		if isPossible(towels, t, map[string]bool{}) {
			answer++
		}
	}

	return answer
}

func isPossible(towels []string, desired string, badCache map[string]bool) bool {
	if len(desired) == 0 {
		return true
	}

	if badCache[desired] {
		return false
	}

	for _, t := range towels {
		l := len(t)
		if len(desired) < l {
			continue
		}
		if desired[:l] == t {
			if isPossible(towels, desired[l:], badCache) {
				return true
			} else {
				badCache[desired[l:]] = true
			}
		}
	}

	return false
}

func part2(lines []string) int {
	answer := 0
	towels := strings.Split(lines[0], ", ")
	desired := []string{}

	for i, line := range lines {
		if i < 2 {
			continue
		}
		desired = append(desired, line)
	}

	for _, t := range desired {
		count := isPossiblePart2(towels, t, map[string]bool{}, map[string]int{})
		answer += count
	}

	return answer
}

func isPossiblePart2(towels []string, desired string, badCache map[string]bool, goodCache map[string]int) int {
	if len(desired) == 0 {
		return 1
	}
	if desired == "rrbgbr" {
		fmt.Println()
	}

	if badCache[desired] {
		return 0
	}

	if val, ok := goodCache[desired]; ok {
		return val
	}

	count := 0
	for _, t := range towels {
		l := len(t)
		if len(desired) < l {
			continue
		}
		if desired[:l] == t {
			count += isPossiblePart2(towels, desired[l:], badCache, goodCache)
			if count == 0 {
				badCache[desired[l:]] = true
			}
		}
	}

	goodCache[desired] = count

	return count
}
