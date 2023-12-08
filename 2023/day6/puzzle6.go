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

	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]

	for i, t := range times {
		time := S2i(t)
		distance := S2i(distances[i])
		a, b := quadraticForm(time, distance)
		answer *= (a - b + 1)
	}

	fmt.Printf("day6 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 1

	distance := S2i(strings.Join(strings.Fields(lines[1])[1:], ""))
	time := S2i(strings.Join(strings.Fields(lines[0])[1:], ""))
	a, b := quadraticForm(time, distance)
	answer *= (a - b + 1)

	fmt.Printf("day6 Answer 2 : %d\n", answer)
}

func quadraticForm(totalTime, distance int) (int, int) {
	b := float64(-1 * totalTime)
	a := float64(1)
	c := float64(distance) + 0.00000000001

	inside := math.Pow(float64(b), 2) - 4*a*c
	quadA := (-1*b + math.Sqrt(float64(inside))) / 2 * a
	quadB := (-1*b - math.Sqrt(float64(inside))) / 2 * a

	return int(math.Floor(quadA)), int(math.Ceil(quadB))

}
