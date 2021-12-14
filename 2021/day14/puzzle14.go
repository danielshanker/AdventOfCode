package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day14/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day14/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0
	rules := map[string]string{}
	input := ""
	for i, line := range lines {
		if i == 0 {
			input = line
			continue
		}
		if i == 1 {
			continue
		}
		l := strings.Split(line, " -> ")
		rules[l[0]] = l[1]
	}
	for i := 0; i < 10; i++ {
		newInput := ""
		for j := 0; j < len(input)-1; j++ {
			cur := string(input[j]) + string(input[j+1])
			mid := rules[cur]
			newInput += string(input[j]) + mid
		}
		input = newInput + string(input[len(input)-1])
	}
	count := map[string]int{}
	for _, char := range input {
		count[string(char)]++
	}
	biggest := 0
	smallest := math.MaxInt64
	for _, val := range count {
		if val > biggest {
			biggest = val
		}
		if val < smallest {
			smallest = val
		}
	}
	answer1 = biggest - smallest
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	rules := map[string]string{}
	pairs := map[string]int{}
	for i, line := range lines {
		if i == 0 {
			for j := 0; j < len(line)-1; j++ {
				pairs[string(line[j])+string(line[j+1])]++
			}
			continue
		}
		if i == 1 {
			continue
		}
		l := strings.Split(line, " -> ")
		rules[l[0]] = l[1]
	}
	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}
		for pair, val := range pairs {
			mid := rules[pair]
			newPairs[string(pair[0])+mid] += val
			newPairs[mid+string(pair[1])] += val
		}
		pairs = map[string]int{}
		for k, v := range newPairs {
			pairs[k] = v
		}
	}

	count := map[string]int{}
	for pair, val := range pairs {
		count[string(pair[0])] += val
	}
	// the last character will never change and will not have its value added from the pairs sequence
	count[string(lines[0][len(lines[0])-1])]++

	biggest := 0
	smallest := math.MaxInt64
	for _, val := range count {
		if val > biggest {
			biggest = val
		}
		if val < smallest {
			smallest = val
		}
	}
	answer2 = biggest - smallest
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
