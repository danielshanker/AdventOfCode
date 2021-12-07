package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day6/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("input.txt")
		part1(input)
		part2(input)
	}
}

func part1(lines []string) {
	answer1 := spawn(lines[0], 80)

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {

	answer2 := spawn(lines[0], 256)

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func spawn (line string, days int) int{
	var fish [9]int

	fishInput := strings.Split(line, ",")
	for _, i := range fishInput {
		fish[s2i(i)]++
	}

	totalFish := len(fishInput)

	for i:=0; i< days; i++ {
		var tempFish [9]int
		for j:=8; j>=0; j-- {
			if j == 0 {
				tempFish[6] += fish[0]
				tempFish[8] = fish[0]
				totalFish += fish[0]
				continue
			}
			tempFish[j-1] = fish[j]
		}
		for j:=0; j<9; j++ {
			fish[j] = tempFish[j]
		}
	}

	return totalFish
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
