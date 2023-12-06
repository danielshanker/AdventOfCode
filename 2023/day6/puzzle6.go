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

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day6/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day6/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 1

	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	distances = distances[1:]
	times = times[1:]

	for i, t := range times {
		time := S2i(t)
		winCount := 0
		up := true
		distance := S2i(distances[i])
		for j := 0; j < time; j++ {
			a := getTotalDist(j, time)
			if a > distance {
				winCount++
				up = false
			} else if !up {
				break
			}
		}
		answer *= winCount
	}

	fmt.Printf("day6 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 1

	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	distances = distances[1:]
	times = times[1:]
	distance := S2i(strings.Join(distances, ""))
	time := S2i(strings.Join(times, ""))

	winCount := 0
	up := true
	for j := 0; j < time; j++ {
		a := getTotalDist(j, time)
		if a > distance {
			winCount++
			up = false
		} else if !up {
			break
		}
	}
	answer *= winCount

	fmt.Printf("day6 Answer 2 : %d\n", answer)
}

func getTotalDist(heldTime int, totalTime int) int {
	if heldTime >= totalTime {
		return 0
	}
	td := heldTime * (totalTime - heldTime)

	return td

}
