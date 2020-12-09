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
	var numList []int
	preamble := 25

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numList = append(numList, num)
	}
	answer1 := 0
	for i := preamble; i < len(numList); i++ {
		found := false
		for j := i - preamble; j < i; j++ {
			for k := i - preamble; k < i; k++ {
				if numList[j] == numList[k] {
					continue
				}
				if numList[j]+numList[k] == numList[i] {
					found = true
				}
			}
		}
		if !found {
			answer1 = numList[i]
			break
		}
	}
	answer2 := 0
	for i := 0; i < len(numList); i++ {
		sum := 0
		smallest := 1000000000000000000
		largest := 0
		found := false
		for j := i; j < len(numList)-i; j++ {
			sum += numList[j]
			if sum > answer1 {
				break
			}
			if numList[j] > largest {
				largest = numList[j]
			}
			if numList[j] < smallest {
				smallest = numList[j]
			}
			if j == 0 {
				continue
			}
			if sum == answer1 {
				answer2 = smallest + largest
				found = false
				break
			}
		}
		if found {
			break
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
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
