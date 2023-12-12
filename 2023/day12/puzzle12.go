package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 21
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day12/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 525152
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day12/input.txt")
		fmt.Printf("day12 Answer 1 : %d\n", part1(input))
		fmt.Printf("day12 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0
	for _, line := range lines {
		a := strings.Fields(line)
		groupString := strings.Split(a[1], ",")
		groups := []int{}
		for _, val := range groupString {
			groups = append(groups, S2i(val))
		}
		broken, _ := checkString(a[0], 0, groups, 0, "", map[string]int{})
		answer += broken
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	for _, line := range lines {
		a := strings.Fields(line)
		springs := a[0]
		springs2 := a[1]
		for i := 0; i < 4; i++ {
			springs += "?"
			springs += a[0]
			springs2 += ","
			springs2 += a[1]
		}
		groupString := strings.Split(springs2, ",")
		groups := []int{}
		for _, val := range groupString {
			groups = append(groups, S2i(val))
		}
		broken, _ := checkString(springs, 0, groups, 0, "", map[string]int{})
		answer += broken
	}

	return answer
}

func checkString(line string, i int, groups []int, curLength int, fullLine string, savedStrings map[string]int) (int, map[string]int) {
	goodOnes := 0

	if i > len(groups) {
		return 0, savedStrings
	}

	if line == "" {
		if i == len(groups) {
			return 1, savedStrings
		}
		return 0, savedStrings
	}

	curGroupVal := 0
	if i < len(groups) {
		curGroupVal = groups[i]
	}
	firstChar := string(line[0])

	val := 0
	if firstChar == "." {
		if curLength != 0 && curLength < curGroupVal {
			return 0, savedStrings
		}
		if curLength != 0 && curLength == curGroupVal {
			i++
		}

		groupString := ""
		for j := i; j < len(groups); j++ {
			groupString += strconv.Itoa(groups[i]) + ","
		}
		leftKey := line + " - " + groupString
		if _, ok := savedStrings[leftKey]; ok {
			return savedStrings[leftKey], savedStrings
		}
		newLine := line[1:]
		val, savedStrings = checkString(newLine, i, groups, 0, fullLine+".", savedStrings)
		goodOnes += val
		savedStrings[leftKey] = goodOnes
	}
	if firstChar == "#" {
		curLength++
		if curLength > curGroupVal {
			return 0, savedStrings
		}
		newLine := line[1:]
		if newLine == "" && curLength == curGroupVal {
			i++
		}
		val, savedStrings = checkString(newLine, i, groups, curLength, fullLine+"#", savedStrings)
		goodOnes += val
	}
	if firstChar == "?" {
		newLine := line[1:]
		val, savedStrings = checkString("#"+newLine, i, groups, curLength, fullLine, savedStrings)
		goodOnes += val
		val, savedStrings = checkString("."+newLine, i, groups, curLength, fullLine, savedStrings)
		goodOnes += val
	}

	return goodOnes, savedStrings
}
