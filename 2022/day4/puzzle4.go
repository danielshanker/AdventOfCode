package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day4/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day4/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0
	for _, line := range lines {
		field := strings.Split(line, ",")
		a := strings.Split(field[0], "-")
		b := strings.Split(field[1], "-")
		minA := s2i(a[0])
		maxA := s2i(a[1])
		minB := s2i(b[0])
		maxB := s2i(b[1])

		if minA >= minB && maxA <= maxB {
			answer1++
		} else if minB >= minA && maxB <= maxA {
			answer1++
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	for _, line := range lines {
		field := strings.Split(line, ",")
		a := strings.Split(field[0], "-")
		b := strings.Split(field[1], "-")
		minA := s2i(a[0])
		maxA := s2i(a[1])
		minB := s2i(b[0])
		maxB := s2i(b[1])

		if minA >= minB && minA <= maxB {
			answer2++
		} else if minB >= minA && minB <= maxA {
			answer2++
		} else if maxA <= maxB && maxA >= minB {
			answer2++
		} else if maxB <= maxA && maxB >= minA {
			answer2++
		}
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
