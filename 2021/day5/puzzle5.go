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
	sample := readInputLines("sample.txt")
	input := readInputLines("input.txt")

	fmt.Print("sample: ")
	part1(sample)
	part1(input)

	fmt.Print("\n\nsample: ")
	part2(sample)
	part2(input)

}

func part1(lines []string) {
	answer1 := 0

	var seaMap [1000][1000]int
	overlap := 0

	for _,  line := range lines {
		coords := strings.Split(line, " -> ")
		coords1 := strings.Split(coords[0], ",")
		coords2 := strings.Split(coords[1], ",")
		x1 := s2i(coords1[0])
		x2 := s2i(coords2[0])
		y1 := s2i(coords1[1])
		y2 := s2i(coords2[1])

		if x1 != x2 && y1 != y2 {
			continue
		}

		if x1 == x2 {
			first := y1
			second := y2
			if y1 > y2 {
				first = y2
				second = y1
			}
			for i:=first; i<=second; i++ {
				seaMap[x1][i]++
				if seaMap[x1][i] == 2 {
					overlap++
				}
			}
		}
		if y1 == y2 {
			first := x1
			second := x2
			if x1 > x2 {
				first = x2
				second = x1
			}
			for i:=first; i<=second; i++ {
				seaMap[i][y1]++
				if seaMap[i][y1] == 2 {
					overlap++
				}
			}
		}
	}
	answer1 = overlap
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	var seaMap [1000][1000]int
	overlap := 0

	for _,  line := range lines {
		coords := strings.Split(line, " -> ")
		coords1 := strings.Split(coords[0], ",")
		coords2 := strings.Split(coords[1], ",")
		x1 := s2i(coords1[0])
		x2 := s2i(coords2[0])
		y1 := s2i(coords1[1])
		y2 := s2i(coords2[1])

		if x1 == x2 {
			first := y1
			second := y2
			if y1 > y2 {
				first = y2
				second = y1
			}
			for i:=first; i<=second; i++ {
				seaMap[x1][i]++
				if seaMap[x1][i] == 2 {
					overlap++
				}
			}
		} else if y1 == y2 {
			first := x1
			second := x2
			if x1 > x2 {
				first = x2
				second = x1
			}
			for i:=first; i<=second; i++ {
				seaMap[i][y1]++
				if seaMap[i][y1] == 2 {
					overlap++
				}
			}
		} else {
			if x1 > x2 && y1 > y2 {
				lineLength := x1-x2
				for i := 0; i <= lineLength; i++ {
					seaMap[x1-i][y1-i]++
					if seaMap[x1-i][y1-i] == 2 {
						overlap++
					}
				}
			} else if x1 > x2 && y2 > y1 {
				lineLength := x1-x2
				for i := 0; i <= lineLength; i++ {
					seaMap[x1-i][y1+i]++
					if seaMap[x1-i][y1+i] == 2 {
						overlap++
					}
				}

			} else if x2 > x1 && y2 > y1 {
				lineLength := x2-x1
				for i := 0; i <= lineLength; i++ {
					seaMap[x1+i][y1+i]++
					if seaMap[x1+i][y1+i] == 2 {
						overlap++
					}
				}
			} else {
				lineLength := x2-x1
				for i := 0; i <= lineLength; i++ {
					seaMap[x1+i][y1-i]++
					if seaMap[x1+i][y1-i] == 2 {
						overlap++
					}
				}
			}
		}
	}
	answer2 = overlap

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func printMap(seaMap [1000][1000]int) {

	fmt.Println("")
	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {
			fmt.Print(seaMap[j][i])
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
