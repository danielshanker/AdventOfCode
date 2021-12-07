package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"flag"
	"strings"
	"sort"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day7/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day7/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0

	crabPosString := strings.Split(lines[0], ",")
	var crabPos	[]int
	for _, crab := range crabPosString {
		crabPos = append(crabPos, s2i(crab))
	}

	sort.Ints(crabPos)
	medianIndex := len(crabPos)/2
	target := crabPos[medianIndex]

	fuel := 0

	for _, pos := range crabPos {
		if pos - target < 0 {
			fuel -= pos-target
		} else {
			fuel += pos-target
		}
	}
	answer1 = fuel
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	crabPosString := strings.Split(lines[0], ",")
	var crabPos	[]int
	total := 0
	for _, crab := range crabPosString {
		crabPos = append(crabPos, s2i(crab))
		total += s2i(crab)
	}

	// the target is either the int mean rounded down, or the int mean rounded up
	target := int(float32(total)/float32(len(crabPos)))
	target2 := int(float32(total)/float32(len(crabPos)) + .5)

	fuel := 0
	fuel2 := 0

	for _, pos := range crabPos {
		steps := 0
		if pos - target < 0 {
			steps = target - pos
		} else {
			steps = pos-target
		}

		steps2 := 0
		if pos - target2 < 0 {
			steps2 = target2 - pos
		} else {
			steps2 = pos-target2
		}

		fuel += (steps*(steps+1))/2
		fuel2 += (steps2*(steps2+1))/2
	}

	if fuel > fuel2 {
		answer2 = fuel2
	} else {
		answer2 = fuel
	}
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}


func s2i(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("OH NO! OH NO! NOT AN INT!")
	}
	return num
}

func readInputLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}
