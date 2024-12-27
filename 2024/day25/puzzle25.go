package main

import (
	"flag"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 25, part1, part2, 3, 0)

}

func part1(lines []string) int {
	answer := 0

	keys := [][]int{}
	locks := [][]int{}

	curLock := make([]int, 5)
	curKey := make([]int, 5)
	isKey := false
	for i, line := range lines {
		if i%8 == 7 {
			if isKey {
				keys = append(keys, curKey)
			} else {
				locks = append(locks, curLock)
			}
			curLock = make([]int, 5)
			curKey = make([]int, 5)
			continue
		}

		if i%8 == 6 {
			continue
		}

		if i%8 == 0 {
			if line == "....." {
				isKey = true
			} else {
				isKey = false
			}
			continue
		}
		if isKey {
			for i, r := range line {
				char := string(r)
				if char == "#" {
					curKey[i]++
				}
			}
		} else {
			for i, r := range line {
				char := string(r)
				if char == "#" {
					curLock[i]++
				}
			}
		}
	}
	if isKey {
		keys = append(keys, curKey)
	} else {
		locks = append(locks, curLock)
	}

	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(locks); j++ {
			count := 0
			for k := 0; k < 5; k++ {
				if keys[i][k]+locks[j][k] <= 5 {
					count++
				}
				if count == 5 {
					answer++
				}
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	return answer
}
