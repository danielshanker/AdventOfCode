package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("day2.txt")

	validPasswords := 0
	validPasswords2 := 0
	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		password := splitLine[1]
		splitRule := strings.Split(splitLine[0], " ")
		letter := splitRule[1]
		minMax := strings.Split(splitRule[0], "-")

		letterCount := 0

		for _, curLetter := range password {
			if string(curLetter) == letter {
				letterCount++
			}
		}

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		if letterCount >= min && letterCount <= max {
			validPasswords++
		}

		foundSpot := 0
		if string(password[min]) == letter {
			foundSpot++
		}
		if string(password[max]) == letter {
			foundSpot++
		}
		if foundSpot == 1 {
			validPasswords2++
		}
	}
	fmt.Println(fmt.Sprintf("answer1: %d", validPasswords))
	fmt.Println(fmt.Sprintf("answer2: %d", validPasswords2))
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
