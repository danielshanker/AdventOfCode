package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//lines := readInputLines("input.txt")
	lines := readInputLines("../sample.txt")

	for _, line := range lines {
		fmt.Println(line)
	}

	answer1 := 0
	answer2 := 0

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
