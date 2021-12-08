package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day8/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day8/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")
		outputs := strings.Fields(splitLine[1])
		for _, output := range outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7  {
				answer1++
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")

		allVals := strings.Fields(splitLine[0])
		var one string
		var two string
		var three string
		var four string
		var five string
		var six string
		var seven string
		var eight string
		var nine string
		var zero string
		for _, val := range allVals {
			if len(val) == 2 {
				one = val
			}
			if len(val) == 3 {
				seven = val
			}
			if len(val) == 4 {
				four = val
			}
			if len(val) == 7 {
				eight = val
			}
		}
		var top string
		var topLeft string
		var topRight string
		var bottomRight string
		var middle string
		var bottom string
		var bottomLeft string
		for _, seg := range seven {
			if !contains(string(seg), one) {
				top = string(seg)
			}
		}

		for _, val := range allVals {
			if len(val) == 5 {
				//3
				if contains(string(one[0]), val) && contains(string(one[1]), val) {
					three = val
					for _, seg := range four {
						if !contains(string(seg), three) {
							topLeft = string(seg)
						} else if !contains(string(seg), one) {
							middle = string(seg)
						}
					}
				}
			}
		}
		for _, val := range allVals {
			if len(val) == 5 {
				//5
				if contains(top, val) && contains(topLeft, val) && contains(middle, val) {
					five = val
					if contains(string(one[0]), val) {
						bottomRight = string(one[0])
						topRight = string(one[1])
					} else {
						bottomRight = string(one[1])
						topRight = string(one[0])
					}
					for _, seg := range five {
						stringSeg := string(seg)
						if stringSeg != top && stringSeg != topLeft && stringSeg != middle && stringSeg != bottomRight {
							bottom = stringSeg
						}
					}

				}
			}
		}
		for _, char := range "abcdefg" {
			stringSeg := string(char)
			if stringSeg != top && stringSeg != topLeft && stringSeg != topRight && stringSeg != middle && stringSeg != bottomRight && stringSeg != bottom {
				bottomLeft = stringSeg
			}
		}

		zero = sortString(top+topLeft+topRight+bottomLeft+bottomRight+bottom)
		one = sortString(one)
		two = sortString(top+topRight+middle+bottomLeft+bottom)
		three = sortString(three)
		four = sortString(four)
		five = sortString(five)
		six = sortString(top+topLeft+middle+bottomLeft+bottomRight+bottom)
		seven = sortString(seven)
		eight = sortString(eight)
		nine = sortString(top+topLeft+middle+topRight+bottomRight+bottom)

		outputs := strings.Fields(splitLine[1])
		outputVal := ""
		for _, output := range outputs {
			val := 0
			output = sortString(output)
			switch output {
			case zero:
				val = 0
			case one:
				val = 1
			case two:
				val = 2
			case three:
				val = 3
			case four:
				val = 4
			case five:
				val = 5
			case six:
				val = 6
			case seven:
				val = 7
			case eight:
				val = 8
			case nine:
				val = 9
			}
			outputVal += strconv.Itoa(val)
		}
		answer2 += s2i(outputVal)
	}

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

func contains (c string, s string) bool {
	for _, char := range s {
		if string(char) == c {
			return true
		}
	}
	return false
}
func sortString(s string) string {
	splitString := strings.Split(s, "")
	sort.Strings(splitString)
	return strings.Join(splitString, "")
}
