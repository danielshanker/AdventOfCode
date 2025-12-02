package main

import (
	"flag"
	"strconv"
	"strings"
	. "utils"
)

var test *bool

func main() {
	//	test = flag.Bool("t", false, "use sample")
	test = flag.Bool("t", false, "use sample")
	flag.Parse()
	Start(test, 2, part1, part2, 1227775554, 4174379265)

}

type ranges struct {
	min int
	max int
}

func part1(lines []string) int {
	answer := 0

	rangeList := []ranges{}

	rSplit := strings.Split(lines[0], ",")
	for _, r := range rSplit {
		if r == "" {
			continue
		}
		s := strings.Split(r, "-")
		min := S2i(s[0])
		max := S2i(s[1])

		rangeList = append(rangeList, ranges{min: min, max: max})
	}

	for _, r := range rangeList {
		for i := r.min; i <= r.max; i++ {
			if hasDupe(i) {
				answer += i
			}
		}
	}

	return answer
}

func hasDupe(val int) bool {
	valString := strconv.Itoa(val)
	return valString[:len(valString)/2] == valString[len(valString)/2:]
}

func hasDupes(val int) bool {
	valString := strconv.Itoa(val)
	for i := 1; i <= len(valString)/2; i++ {
		subString := valString[:i]
		if len(valString)%i != 0 {
			continue
		}
		j := 0
		found := true
		for j < len(valString)/i {
			compString := valString[j*i : j*i+i]
			if compString != subString {
				found = false
				break
			}
			j++
		}
		if found {
			return true
		}
	}
	return false
}

func part2(lines []string) int {
	answer := 0

	rangeList := []ranges{}

	rSplit := strings.Split(lines[0], ",")
	for _, r := range rSplit {
		if r == "" {
			continue
		}
		s := strings.Split(r, "-")
		min := S2i(s[0])
		max := S2i(s[1])

		rangeList = append(rangeList, ranges{min: min, max: max})
	}

	for _, r := range rangeList {
		for i := r.min; i <= r.max; i++ {
			if hasDupes(i) {
				answer += i
			}
		}
	}

	return answer
}
