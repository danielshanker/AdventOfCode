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
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day10/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day10/input.txt")
		part1(input)
		part2(input)
	}

}
type Stack []string

func part1(lines []string) {
	match := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
		"<": ">",
	}
	points := map[string]int{
		"}": 1197,
		"]": 57,
		")": 3,
		">": 25137,
	}
	answer1 := 0

	for _, line := range lines {
		var stack Stack
		for _, char := range line {
			curChar := string(char)
			if curChar == "(" || curChar == "{" || curChar == "[" || curChar == "<" {
				stack.Push(curChar)
			} else {
				lastChar := stack.Pop()
				if match[lastChar] != curChar {
					answer1 += points[curChar]
					continue
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	pointsArray := map[string]int{
		"{": 3,
		"[": 2,
		"(": 1,
		"<": 4,
	}
	match := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
		"<": ">",
	}
	var legalList []string

	for _, line := range lines {
		var stack Stack
		legal := true
		for _, char := range line {
			curChar := string(char)
			if curChar == "(" || curChar == "{" || curChar == "[" || curChar == "<" {
				stack.Push(curChar)
			} else {
				lastChar := stack.Pop()
				if match[lastChar] != curChar {
					legal = false
					continue
				}
			}
		}
		if legal {
			legalList = append(legalList, line)
		}
	}
	var points []int
	for _, line := range legalList {
		var stack Stack
		for _, char := range line {
			curChar := string(char)
			if curChar == "(" || curChar == "{" || curChar == "[" || curChar == "<" {
				stack.Push(curChar)
			} else {
				stack.Pop()
			}
		}
		pointVal := 0
		length := len(stack)
		for i:=	0; i < length; i++ {
			pointVal *= 5
			pointVal += pointsArray[stack.Pop()]
		}
		points = append(points, pointVal)
	}

	sort.Ints(points)
	answer2 = points[len(points)/2]

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

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}