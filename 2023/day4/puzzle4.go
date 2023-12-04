package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day4/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day4/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0
	for _, line := range lines {
		pointCount := 0
		cardSplit := strings.Split(line, " | ")
		winners := strings.Split(cardSplit[0], ": ")
		winningNumbers := strings.Fields(winners[1])
		winMap := make(map[string]struct{})
		for _, num := range winningNumbers {
			winMap[num] = struct{}{}

		}

		yourNumbers := strings.Fields(cardSplit[1])
		for _, num := range yourNumbers {
			if _, ok := winMap[num]; ok {
				pointCount++
			}
		}
		if pointCount != 0 {
			answer += int(math.Pow(2, float64(pointCount-1)))
		}
	}

	fmt.Printf("day4 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0
	lineMap := map[int]string{}
	winCount := map[int]int{}
	for i, line := range lines {
		lineMap[i+1] = line
	}
	for index := 0; index > -1; index++ {
		if index >= len(lines)-1 {
			break
		}
		line := lines[index]
		pointCount := 0
		cardNumber := S2i(strings.Split(strings.Fields(line)[1], ":")[0])
		if _, ok := winCount[cardNumber]; !ok {
			cardSplit := strings.Split(line, " | ")
			winners := strings.Split(cardSplit[0], ": ")
			winningNumbers := strings.Fields(winners[1])
			winMap := make(map[string]struct{})
			for _, num := range winningNumbers {
				winMap[num] = struct{}{}
			}

			yourNumbers := strings.Fields(cardSplit[1])
			for _, num := range yourNumbers {
				if _, ok := winMap[num]; ok {
					pointCount++
				}
			}
			winCount[cardNumber] = pointCount
		}
		pointCount = winCount[cardNumber]
		for i := 1; i <= pointCount; i++ {
			lines = insert(lines, index+i, lineMap[cardNumber+i])
		}
	}
	answer = len(lines)
	fmt.Printf("day4 Answer 2 : %d\n", answer)
}

func insert(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
