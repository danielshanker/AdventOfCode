package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sample := readInputLines("../sample.txt")
	input := readInputLines("input.txt")

	fmt.Print("sample: ")
	part1(sample)
	part1(input)

	fmt.Print("\n\nsample: ")
	part2(sample)
	part2(input)

}

func part1(lines []string) {
	answer1 := 0

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

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
