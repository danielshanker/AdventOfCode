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
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	answer1 := 0
	answer2 := 0
	for _, line := range lines {
		// sanitized the input by adding a space before and after every paren
		calc := strings.Split(line, " ")
		val, _ := calculate(calc, 0)
		answer1 += val

		val, _ = calculate2(calc, 0)
		answer2 += val
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func calculate(expr []string, index int) (int, int) {
	total := 0
	lastOperation := "+"
	for i := index; i < len(expr); i++ {
		curChar := string(expr[i])
		if curChar == ")" {
			return total, i
		}
		if curChar == "(" {
			val := 0
			val, i = calculate(expr, i+1)
			if lastOperation == "+" {
				total += val
			}
			if lastOperation == "*" {
				total *= val
			}
			continue
		}

		if curChar == "+" {
			lastOperation = "+"
			continue
		}
		if curChar == "*" {
			lastOperation = "*"
			continue
		}

		if lastOperation == "+" {
			total += s2i(curChar)
		}
		if lastOperation == "*" {
			total *= s2i(curChar)
		}
	}
	return total, 0
}

func calculate2(expr []string, index int) (int, int) {
	total := 0
	lastOperation := "+"
	var multVals []int
	for i := index; i < len(expr); i++ {
		curChar := string(expr[i])
		if curChar == ")" {
			index = i
			break
		}
		if curChar == "(" {
			val := 0
			val, i = calculate2(expr, i+1)
			if lastOperation == "" {
				total = val
			}
			if lastOperation == "+" {
				total += val
			}
			continue
		}

		if curChar == "+" {
			lastOperation = "+"
			continue
		}
		if curChar == "*" {
			lastOperation = "+"
			multVals = append(multVals, total)
			total = 0
			continue
		}

		if lastOperation == "+" {
			total += s2i(curChar)
		}
	}

	if total != 0 {
		multVals = append(multVals, total)
	}

	total = 1
	for _, val := range multVals {
		total *= val
	}
	return total, index
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
