package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")
	//	lines := readInputLines("../sample.txt")
	length := 132
	//	length = 31

	var rule []string
	var text []string
	// List was sorted
	for i, line := range lines {
		if i >= length {
			text = append(text, line)
		} else {
			rule = append(rule, line)
		}
	}

	regex := findRegex(rule)
	re := regexp.MustCompile("^" + regex + "$")
	answer1 := 0
	for _, i := range text {
		matched := re.MatchString(i)
		if matched {
			answer1++
		}
	}
	answer2 := 0

	for loop := 0; loop < 10; loop++ {
		newRule := loopWithDepth(rule, loop)
		regex = findRegex(newRule)
		re = regexp.MustCompile("^" + regex + "$")
		answer2 = 0
		for _, i := range text {
			matched := re.MatchString(i)
			if matched {
				answer2++
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func loopWithDepth(rules []string, loop int) []string {
	var newRule []string
	for _, rule := range rules {
		splitNum := strings.Split(rule, ":")
		if splitNum[0] == "8" {
			newRule = append(newRule, rule+" + ")
			continue
		}
		if splitNum[0] == "11" {
			eleven := "11: 42 31 "
			for j := 2; j < loop; j++ {
				eleven += " | "
				for i := 0; i < j; i++ {
					eleven += " 42 "
				}
				for i := 0; i < j; i++ {
					eleven += " 31 "
				}
			}
			newRule = append(newRule, eleven)
			continue
		}
		newRule = append(newRule, rule)
	}

	return newRule

}

func findRegex(input []string) string {
	rules := make(map[int]string)
	for {
		if _, ok := rules[0]; ok {
			break
		}
		for _, rule := range input {
			goNext := false
			ruleSplit := strings.Split(rule, ": ")
			i := s2i(ruleSplit[0])
			if _, ok := rules[i]; ok {
				continue
			}
			parts := strings.Fields(rule)
			parts = parts[1:]

			if string(parts[0][0]) == "\"" {
				rules[i] = string(parts[0][1])
				continue
			}
			regex := "(?:"
			for _, part := range parts {
				if part == "|" {
					regex += "|"
				} else if part == "+" {
					regex += "+"
				} else {
					token := s2i(part)
					a, ok := rules[token]
					if !ok {
						goNext = true
						break
					}
					regex += a
				}
			}
			if goNext {
				continue
			}
			regex += ")"
			rules[i] = regex
		}
	}
	return rules[0]
}

func s2i(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("OH NO! OH NO! NOT AN INT!")
		fmt.Println(err)
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
