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

	cardPk := s2i(lines[0])
	doorPk := s2i(lines[1])

	doorLS := 0
	cardLS := 0

	i := 0

	key := 1

	for {
		if key == doorPk {
			doorLS = i
		}
		if key == cardPk {
			cardLS = i
		}
		if cardLS != 0 && doorLS != 0 {
			break
		}
		i++
		key *= 7
		key = key % 20201227
	}

	encryption := s2i(lines[1])

	key = 1
	for i = 0; i < cardLS; i++ {
		key *= encryption
		key = key % 20201227
	}

	answer1 := key
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
