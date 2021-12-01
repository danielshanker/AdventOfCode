package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	lastDepth := -1
	increases := 0

	for _, line := range lines {
		if s2i(line) > lastDepth && lastDepth >= 0 {
			increases++
		}
		lastDepth = s2i(line)
	}

	answer1 := increases
	answer2 := 0

	lastSum := -1
	for i, _ := range lines {
		if i >= len(lines) - 2 {
			break
		}
		sum := s2i(lines[i]) + s2i(lines[i+1]) + s2i(lines[i+2])
		if sum > lastSum && lastSum >= 0 {
			answer2++
		}
		lastSum = sum
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
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
