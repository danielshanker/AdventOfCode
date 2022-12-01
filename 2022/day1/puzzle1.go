package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day1/sample.txt")
		cals := part1(sample)
		part2(cals)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day1/input.txt")
		cals := part1(input)
		part2(cals)
	}

}

func part1(lines []string) []int {
	answer1 := 0
	var calories []int
	curCount := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, curCount)
			curCount = 0
			continue
		}
		curCount += s2i(line)
	}
	calories = append(calories, curCount)
	curCount = 0

	sort.Ints(calories[:])
	answer1 = calories[len(calories)-1]
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	return calories
}

func part2(calories []int) {
	size := len(calories) - 1

	answer2 := calories[size] + calories[size-1] + calories[size-2]

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
