package main

import (
	"flag"
	"math"
	"sort"
	"strings"
	"utils"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 1, part1, part2, 0, 31)
}

func part1(lines []string) int {
	answer := 0
	left := []int{}
	right := []int{}

	for _, line := range lines {
		vals := strings.Fields(line)
		left = append(left, utils.S2i(vals[0]))
		right = append(right, utils.S2i(vals[1]))
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		answer += int(math.Abs(float64(left[i] - right[i])))
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	left := []int{}
	right := []int{}
	for _, line := range lines {
		vals := strings.Fields(line)
		left = append(left, utils.S2i(vals[0]))
		right = append(right, utils.S2i(vals[1]))
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		target := left[i]
		count := 0
		for j := 0; j < len(right); j++ {
			if right[j] > target {
				break
			}
			if right[j] == target {
				count++
			}
		}
		answer += target * count
	}

	return answer
}
