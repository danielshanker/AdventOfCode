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
	numList = append(numList, 0)

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numList = append(numList, num)
	}
	sortArray(numList)

	numList = append(numList, numList[len(numList)-1]+3)

	one := 0
	three := 0

	for i := 0; i < len(numList)-1; i++ {
		diff := numList[i+1] - numList[i]
		if diff == 1 {
			one++
		} else if diff == 2 {

		} else if diff == 3 {
			three++
		} else {
			fmt.Println("error")
		}
	}

	answer1 := three * one
	var subSeqs [][]int

	var subSeq []int
	var diffArr []int
	for i := 0; i < len(numList)-1; i++ {
		diff := numList[i+1] - numList[i]
		diffArr = append(diffArr, diff)
	}
	for _, val := range diffArr {
		if val != 3 {
			subSeq = append(subSeq, val)
		} else {
			if len(subSeq) > 0 {
				subSeqs = append(subSeqs, subSeq)
				subSeq = nil
			}
		}
	}
	paths := 1
	for _, i := range subSeqs {
		paths *= opts(len(i))
	}
	answer2 := paths

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func opts(length int) int {

	if length == 1 {
		return 1
	}
	if length == 2 {
		return 2
	}
	if length == 3 {
		return 4
	}
	if length == 4 {
		return 7
	}
	if length == 5 {
		return 13
	}
	if length == 6 {
		return 24
	}
	fmt.Println(fmt.Sprintf("need more! %d", length))
	return 0
}

func sortArray(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[j] < array[i] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
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
