package main

import (
	"flag"
	"fmt"
	"math"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 374
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day11/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 1030
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day11/input.txt")
		fmt.Printf("day11 Answer 1 : %d\n", part1(input))
		fmt.Printf("day11 Answer 2 : %d\n", part2(input))
	}

}

type galaxy struct {
	x int
	y int
}

func part1(lines []string) int {
	return calculateAnswer(lines, 1)
}

func part2(lines []string) int {
	space := 999999
	return calculateAnswer(lines, space)
}

func calculateAnswer(lines []string, space int) int {
	answer := 0
	xFounds := map[int]struct{}{}
	yFounds := map[int]struct{}{}

	for y, line := range lines {
		for x, char := range line {
			spot := string(char)
			if spot == "#" {
				xFounds[x] = struct{}{}
				yFounds[y] = struct{}{}
			}
		}
	}

	galaxies := []galaxy{}
	emptyYs := 0
	for y, line := range lines {
		if _, ok := yFounds[y]; !ok {
			emptyYs++
			continue
		}
		emptyXs := 0
		for x, char := range line {
			if _, ok := xFounds[x]; !ok {
				emptyXs++
				continue
			}
			spot := string(char)
			if spot == "#" {
				gal := galaxy{
					x: x + (emptyXs * space),
					y: y + (emptyYs * space),
				}
				galaxies = append(galaxies, gal)
			}
		}
	}

	for i, gal1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			gal2 := galaxies[j]
			dist := int(math.Abs(float64(gal1.x)-float64(gal2.x)) + math.Abs(float64(gal1.y)-float64(gal2.y)))
			answer += dist

		}
	}

	return answer
}
