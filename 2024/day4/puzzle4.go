package main

import (
	"flag"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 4, part1, part2, 18, 9)

}

func part1(lines []string) int {
	answer := 0

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			if char == "X" {
				if i >= 3 {
					if string(lines[i-1][j]) == "M" && string(lines[i-2][j]) == "A" && string(lines[i-3][j]) == "S" {
						answer++
					}
				}
				if i < len(lines)-3 {
					if string(lines[i+1][j]) == "M" && string(lines[i+2][j]) == "A" && string(lines[i+3][j]) == "S" {
						answer++
					}
				}
				if j >= 3 {
					if string(lines[i][j-1]) == "M" && string(lines[i][j-2]) == "A" && string(lines[i][j-3]) == "S" {
						answer++
					}
				}
				if j < len(lines[0])-3 {
					if string(lines[i][j+1]) == "M" && string(lines[i][j+2]) == "A" && string(lines[i][j+3]) == "S" {
						answer++
					}
				}
				if i >= 3 && j >= 3 {
					if string(lines[i-1][j-1]) == "M" && string(lines[i-2][j-2]) == "A" && string(lines[i-3][j-3]) == "S" {
						answer++
					}
				}
				if i >= 3 && j < len(lines[0])-3 {
					if string(lines[i-1][j+1]) == "M" && string(lines[i-2][j+2]) == "A" && string(lines[i-3][j+3]) == "S" {
						answer++
					}
				}
				if i < len(lines)-3 && j < len(lines[0])-3 {
					if string(lines[i+1][j+1]) == "M" && string(lines[i+2][j+2]) == "A" && string(lines[i+3][j+3]) == "S" {
						answer++
					}
				}
				if i < len(lines)-3 && j >= 3 {
					if string(lines[i+1][j-1]) == "M" && string(lines[i+2][j-2]) == "A" && string(lines[i+3][j-3]) == "S" {
						answer++
					}
				}
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	for i, line := range lines {
		for j, r := range line {
			char := string(r)
			if i == 0 || i == len(lines)-1 || j == 0 || j == len(line)-1 {
				continue
			}
			if char == "A" {
				if (string(lines[i-1][j-1]) == "M" && string(lines[i+1][j+1]) == "S") || (string(lines[i-1][j-1]) == "S" && string(lines[i+1][j+1]) == "M") {
					if (string(lines[i-1][j+1]) == "M" && string(lines[i+1][j-1]) == "S") || (string(lines[i-1][j+1]) == "S" && string(lines[i+1][j-1]) == "M") {
						answer++
					}

				}
			}
		}
	}

	return answer
}
