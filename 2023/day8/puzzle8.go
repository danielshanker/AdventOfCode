package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 6
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day8/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 6
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day8/input.txt")
		fmt.Printf("day8 Answer 1 : %d\n", part1(input))
		fmt.Printf("day8 Answer 2 : %d\n", part2(input))
	}

}

type connection struct {
	l string
	r string
}

func part1(lines []string) int {
	answer := 0
	connections := map[string]connection{}
	conReg := regexp.MustCompile("[A-Z0-9]{3}")

	instructions := strings.Split(lines[0], "")
	start := false
	for _, line := range lines {
		if !start {
			if line == "" {
				start = true
			}
			continue
		}
		rule := conReg.FindAllStringSubmatch(line, -1)
		con := connection{
			l: rule[1][0],
			r: rule[2][0],
		}
		connections[rule[0][0]] = con
	}

	curNode := "AAA"

	for {
		found := false
		for _, cha := range instructions {
			answer++
			inst := string(cha)
			if inst == "L" {
				curNode = connections[curNode].l
			} else {
				curNode = connections[curNode].r
			}
			if curNode == "ZZZ" {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	connections := map[string]connection{}
	conReg := regexp.MustCompile("[A-Z0-9]{3}")
	aNodes := []string{}

	instructions := strings.Split(lines[0], "")
	start := false
	for _, line := range lines {
		if !start {
			if line == "" {
				start = true
			}
			continue
		}
		rule := conReg.FindAllStringSubmatch(line, -1)
		con := connection{
			l: rule[1][0],
			r: rule[2][0],
		}
		connections[rule[0][0]] = con
		if string(rule[0][0][2]) == "A" {
			aNodes = append(aNodes, rule[0][0])
		}
	}

	zFounds := []int{}

	for i, curNode := range aNodes {
		count := 0
		for {
			found := false
			for _, cha := range instructions {
				count++
				inst := string(cha)
				if inst == "L" {
					curNode = connections[curNode].l
				} else {
					curNode = connections[curNode].r
				}
				if string(curNode[2]) == "Z" {
					zFounds = append(zFounds, count)
					found = true
					break
				}
				aNodes[i] = curNode
			}
			if found {
				break
			}
		}
	}

	answer = LCM(zFounds[0], zFounds[1], zFounds[2:]...)

	return answer
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
