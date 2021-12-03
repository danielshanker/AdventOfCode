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
	fmt.Print("sample: ")
	part1(readInputLines("../sample.txt"))
	part1(readInputLines("input.txt"))

	fmt.Print("\n\nsample: ")
	part2(readInputLines("../sample.txt"))
	part2(readInputLines("input.txt"))

}

func part1(lines []string) {
	answer1 := 0

	binLength := len(lines[0])
	listLength := len(lines)
	var count = make([]int, binLength)
	for _, line := range lines {
		chars := strings.Split(line, "")
		for i, char := range chars {
			if char == "1" {
				count[i]++
			}
		}
	}

	newBin := ""
	oppBin := ""

	for _, charCount := range count {
		if charCount > listLength/2 {
			newBin += "1"
			oppBin += "0"
		} else {
			newBin += "0"
			oppBin += "1"
		}
	}

	int1, _ := strconv.ParseInt(newBin, 2, 64)
	int2, _ := strconv.ParseInt(oppBin, 2, 64)

	answer1 = int(int1) * int(int2)

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	// O2
	binLength := len(lines[0])
	newList := make(map[string]bool)

	for _, line := range lines {
		newList[line] = true
	}

	curChar := 0

	for len(newList) > 1 {
		if curChar >= binLength {
			break
		}

		count := 0
		for line, _ := range newList {
			if string(line[curChar]) == "1" {
				count++
			}
		}
		if count*2 >= len(newList) {
			for _, line := range lines {
				if string(line[curChar]) == "0" {
					delete(newList, line)
				}
			}
		} else {
			for _, line := range lines {
				if string(line[curChar]) == "1" {
					delete(newList, line)
				}
			}
		}
		curChar++
	}

	var o2 int
	for o2Bin, _ := range newList {
		i, _ := strconv.ParseInt(o2Bin, 2, 64)
		o2 = int(i)
		delete(newList, o2Bin)
	}

	for _, line := range lines {
		newList[line] = true
	}

	// CO2

	curChar = 0

	for len(newList) > 1 {
		if curChar >= binLength {
			break
		}

		count := 0
		for line, _ := range newList {
			if string(line[curChar]) == "1" {
				count++
			}
		}
		if count*2 >= len(newList) {
			for _, line := range lines {
				if string(line[curChar]) == "1" {
					delete(newList, line)
				}
			}
		} else {
			for _, line := range lines {
				if string(line[curChar]) == "0" {
					delete(newList, line)
				}
			}
		}
		curChar++
	}

	var co2 int
	for co2Bin, _ := range newList {
		i, _ := strconv.ParseInt(co2Bin, 2, 64)
		co2 = int(i)
	}
	answer2 = o2 * co2

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
