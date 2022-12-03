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
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day3/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day3/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0
	for _, line := range lines {
		halfway := len(line) / 2
		first := line[0:halfway]
		second := line[halfway:]
		ignore := ""
		for _, char := range first {
			if strings.Contains(second, string(char)) && !strings.Contains(ignore, string(char)) {
				answer1 += convertRune(char)
				ignore += string(char)
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0
	for i := 0; i < len(lines); i += 3 {
		line := lines[i]
		for _, char := range line {
			if strings.Contains(lines[i+1], string(char)) && strings.Contains(lines[i+2], string(char)) {
				answer2 += convertRune(char)
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func convertRune(r rune) int {
	if string(r) == strings.ToUpper(string(r)) {
		return convertUpper(int(r))
	} else {
		return convertLower(int(r))
	}
}

func convertUpper(r int) int {
	return r - 38
}
func convertLower(r int) int {
	return r - 96
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
