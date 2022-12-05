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
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day5/sample.txt")
		part1(sample, *test)
		part2(sample, *test)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day5/input.txt")
		part1(input, *test)
		part2(input, *test)
	}

}

type Stack []string

func part1(lines []string, test bool) {
	answer1 := ""
	var boxes []Stack
	boxes = getBoxes(test)
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		num := s2i(lineSplit[1])
		from := s2i(lineSplit[3]) - 1
		to := s2i(lineSplit[5]) - 1

		for i := 0; i < num; i++ {
			box := boxes[from].Pop()
			boxes[to].Push(box)
		}
	}
	for _, box := range boxes {
		answer1 += box.Pop()
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %s", answer1))
}

func part2(lines []string, test bool) {
	answer2 := ""
	var boxes []Stack
	boxes = getBoxes(test)
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		num := s2i(lineSplit[1])
		from := s2i(lineSplit[3]) - 1
		to := s2i(lineSplit[5]) - 1

		var temp Stack
		for i := 0; i < num; i++ {
			box := boxes[from].Pop()
			temp.Push(box)
		}
		for i := 0; i < num; i++ {
			box := temp.Pop()
			boxes[to].Push(box)
		}
	}

	for _, box := range boxes {
		answer2 += box.Pop()
	}

	fmt.Println(fmt.Sprintf("Answer 2 : %s", answer2))
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
		t := strings.TrimSpace(scanner.Text())
		text = append(text, t)
	}

	return text
}

func getBoxes(test bool) []Stack {
	var boxes []Stack
	if test {
		boxes = []Stack{
			{
				"Z",
				"N",
			},
			{
				"M",
				"C",
				"D",
			},
			{
				"P",
			},
		}
	} else {
		boxes = []Stack{
			{
				"S",
				"T",
				"H",
				"F",
				"W",
				"R",
			},
			{
				"S",
				"G",
				"D",
				"Q",
				"W",
			},
			{
				"B",
				"T",
				"W",
			},
			{
				"D",
				"R",
				"W",
				"T",
				"N",
				"Q",
				"Z",
				"J",
			},
			{
				"F",
				"B",
				"H",
				"G",
				"L",
				"V",
				"T",
				"Z",
			},
			{
				"L",
				"P",
				"T",
				"C",
				"V",
				"B",
				"S",
				"G",
			},
			{
				"Z",
				"B",
				"R",
				"T",
				"W",
				"G",
				"P",
			},
			{
				"N",
				"G",
				"M",
				"T",
				"C",
				"J",
				"R",
			},
			{
				"L",
				"G",
				"B",
				"W",
			},
		}
	}
	return boxes
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
