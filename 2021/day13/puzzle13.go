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
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day13/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day13/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0

	var dots = map[int]map[int]bool{}
	var foldsX []int
	var foldsY []int
	var foldOrder []string
	maxX := 0
	maxY := 0

	endDots := false
	for _, line := range lines {
		if line == "" {
			endDots = true
			continue
		}
		if endDots {
			a := strings.Fields(line)
			b := strings.Split(a[2], "=")
			if b[0] == "x" {
				foldsX = append(foldsX, s2i(b[1]))
				foldOrder = append(foldOrder, "x")
			} else {
				foldsY = append(foldsY, s2i(b[1]))
				foldOrder = append(foldOrder, "y")
			}
			continue
		}
		coord := strings.Split(line, ",")
		if s2i(coord[1]) > maxY {
			maxY = s2i(coord[1])
		}
		if s2i(coord[0]) > maxX {
			maxX = s2i(coord[0])
		}

		if _, ok := dots[s2i(coord[1])]; !ok {
			dots[s2i(coord[1])] = map[int]bool{}
		}
		dots[s2i(coord[1])][s2i(coord[0])] = true

	}

	if foldOrder[0] == "x" {
		dots = foldX(dots, foldsX[0], maxY, maxX)
		maxX = foldsX[0]
	} else {
		dots = foldY(dots, foldsY[0], maxY, maxX)
		maxY = foldsY[0]
	}

	for i := 0; i < maxY+1; i++ {
		for j := 0; j < maxX+1; j++ {
			if val, ok := dots[i][j]; ok {
				if val {
					answer1++
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func foldX(dots map[int]map[int]bool, curFold int, maxY int, maxX int) map[int]map[int]bool {
	for i := 0; i < maxY+1; i++ {
		for j := curFold; j < maxX+1; j++ {
			if val, ok := dots[i][j]; ok && val {
				dots[i][j] = false
				dist := j - curFold
				dots[i][curFold-dist] = true
			}
		}
	}
	return dots
}
func foldY(dots map[int]map[int]bool, curFold int, maxY int, maxX int) map[int]map[int]bool {
	for i := curFold; i < maxY+1; i++ {
		for j := 0; j < maxX+1; j++ {
			if val, ok := dots[i][j]; ok && val {
				dots[i][j] = false
				dist := i - curFold
				if _, ok := dots[curFold-dist]; !ok {
					dots[curFold-dist] = map[int]bool{}
				}
				dots[curFold-dist][j] = true
			}
		}
	}
	return dots
}

func part2(lines []string) {
	var dots = map[int]map[int]bool{}
	var foldsX []int
	var foldsY []int
	var foldOrder []string
	maxX := 0
	maxY := 0

	endDots := false
	for _, line := range lines {
		if line == "" {
			endDots = true
			continue
		}
		if endDots {
			a := strings.Fields(line)
			b := strings.Split(a[2], "=")
			if b[0] == "x" {
				foldsX = append(foldsX, s2i(b[1]))
				foldOrder = append(foldOrder, "x")
			} else {
				foldsY = append(foldsY, s2i(b[1]))
				foldOrder = append(foldOrder, "y")
			}
			continue
		}
		coord := strings.Split(line, ",")
		if s2i(coord[1]) > maxY {
			maxY = s2i(coord[1])
		}
		if s2i(coord[0]) > maxX {
			maxX = s2i(coord[0])
		}

		if _, ok := dots[s2i(coord[1])]; !ok {
			dots[s2i(coord[1])] = map[int]bool{}
		}
		dots[s2i(coord[1])][s2i(coord[0])] = true

	}

	x := 0
	y := 0
	for i := 0; i < len(foldOrder); i++ {
		if foldOrder[i] == "x" {
			dots = foldX(dots, foldsX[x], maxY, maxX)
			maxX = foldsX[x]
			x++
		} else {
			dots = foldY(dots, foldsY[y], maxY, maxX)
			maxY = foldsY[y]
			y++
		}
	}

	for i := 0; i < maxY+1; i++ {
		for j := 0; j < maxX+1; j++ {
			if val, ok := dots[i][j]; ok {
				if val {
					fmt.Print("█")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
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
