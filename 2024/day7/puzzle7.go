package main

import (
	"flag"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	debug := false
	test := flag.Bool("t", debug, "use sample")
	flag.Parse()

	Start(test, 7, part1, part2, 3749, 11387)

}

func part1(lines []string) int {
	answer := 0

	for _, line := range lines {
		segs := strings.Split(line, ":")
		target := S2i(segs[0])
		vals := strings.Fields(segs[1])
		combos := generateCombo("*+", len(vals)-1, "", []string{})
		for _, combo := range combos {
			total := S2i(vals[0])
			for i, r := range combo {
				command := string(r)
				if command == "+" {
					total += S2i(vals[i+1])
				} else {
					total *= S2i(vals[i+1])
				}
			}
			if total == target {
				answer += total
				break
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	for _, line := range lines {
		segs := strings.Split(line, ":")
		target := S2i(segs[0])
		vals := strings.Fields(segs[1])
		combos := generateCombo("*+|", len(vals)-1, "", []string{})
		for _, combo := range combos {
			total := S2i(vals[0])
			for i, r := range combo {
				command := string(r)
				if command == "+" {
					total += S2i(vals[i+1])
				} else if command == "*" {
					total *= S2i(vals[i+1])
				} else {
					sVal := strconv.Itoa(total) + vals[i+1]
					total = S2i(sVal)
				}
			}
			if total == target {
				answer += total
				break
			}
		}
	}

	return answer
}

func generateCombo(chars string, length int, current string, total []string) []string {
	if len(current) == length {
		total = append(total, current)
		return total
	}

	for _, char := range chars {
		total = generateCombo(chars, length, current+string(char), total)
	}
	return total
}
