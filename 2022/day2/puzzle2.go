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
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day2/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day2/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0

	var theirMoves []string
	var yourMoves []string
	for _, line := range lines {
		moves := strings.Fields(line)
		theirMoves = append(theirMoves, moves[0])
		yourMoves = append(yourMoves, moves[1])
	}

	points := 0
	for i := 0; i < len(theirMoves); i++ {
		curYou := yourMoves[i]
		curThem := theirMoves[i]

		if curYou == "X" {
			points += 1
			if curThem == "A" {
				points += 3
			}
			if curThem == "C" {
				points += 6
			}
		}
		if curYou == "Y" {
			points += 2
			if curThem == "A" {
				points += 6
			}
			if curThem == "B" {
				points += 3
			}
		}
		if curYou == "Z" {
			points += 3
			if curThem == "C" {
				points += 3
			}
			if curThem == "B" {
				points += 6
			}
		}

	}
	answer1 = points

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	var theirMoves []string
	var yourMoves []string
	for _, line := range lines {
		moves := strings.Fields(line)
		theirMoves = append(theirMoves, moves[0])
		yourMoves = append(yourMoves, moves[1])
	}

	points := 0
	for i := 0; i < len(theirMoves); i++ {
		curYou := yourMoves[i]
		curThem := theirMoves[i]

		if curYou == "X" {
			points += 0
			if curThem == "A" {
				points += 3
			}
			if curThem == "B" {
				points += 1
			}
			if curThem == "C" {
				points += 2
			}
		}
		if curYou == "Y" {
			points += 3
			if curThem == "A" {
				points += 1
			}
			if curThem == "B" {
				points += 2
			}
			if curThem == "C" {
				points += 3
			}
		}
		if curYou == "Z" {
			points += 6
			if curThem == "A" {
				points += 2
			}
			if curThem == "B" {
				points += 3
			}
			if curThem == "C" {
				points += 1
			}
		}

	}
	answer2 = points

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
