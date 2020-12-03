package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readInputLines("day6.txt")

	answers := make(map[string]bool)
	var questionList []map[string]bool

	for _, line := range lines {
		if line == "" {
			questionList = append(questionList, answers)
			answers = make(map[string]bool)
		}
		for _, curLet := range line {
			answers[string(curLet)] = true
		}
	}
	questionList = append(questionList, answers)

	sum := 0
	for _, i := range questionList {
		sum += len(i)
	}

	answers2 := make(map[string]int)
	famCount := 0
	sum2 := 0
	for _, line := range lines {
		if line == "" {
			for _, i := range answers2 {
				if i == famCount {
					sum2++
				}
			}
			answers2 = make(map[string]int)
			famCount = 0
			continue
		}
		for _, curLet := range line {
			answers2[string(curLet)]++
		}
		famCount++
	}

	for _, i := range answers2 {
		if i == famCount {
			sum2++
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", sum))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", sum2))
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
