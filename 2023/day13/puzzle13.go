package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 405
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day13/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 400
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day13/input.txt")
		fmt.Printf("day13 Answer 1 : %d\n", part1(input))
		fmt.Printf("day13 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0
	yMirror := 0
	xMirror := 0

	yChunks := [][]string{}
	newChunk := []string{}
	for _, line := range lines {
		if line == "" {
			yChunks = append(yChunks, newChunk)
			newChunk = []string{}
			continue
		}
		newChunk = append(newChunk, line)
	}
	yChunks = append(yChunks, newChunk)

	xChunks := [][]string{}
	for _, chunk := range yChunks {
		newXChunk := make([]string, len(chunk[0]))
		for _, line := range chunk {
			for i, char := range line {
				newXChunk[i] += string(char)
			}
		}
		xChunks = append(xChunks, newXChunk)
	}

	for _, chunk := range yChunks {
		linesToDate := []string{}
		half := len(chunk) / 2
		for i, line := range chunk {
			count := 0
			mirrored := true
			if i <= half {
				for j := len(linesToDate) - 1; j >= 0; j-- {
					if linesToDate[j] != chunk[i+count] {
						mirrored = false
						break
					}
					count++
				}
			} else {
				for j := len(linesToDate) - 1; j > i-half; j-- {
					if i+count >= len(chunk) {
						break
					}
					if linesToDate[j] != chunk[i+count] {
						mirrored = false
						break
					}
					count++
				}
			}
			if mirrored && len(linesToDate) != 0 {
				yMirror += len(linesToDate) * 100
				break
			}
			linesToDate = append(linesToDate, line)

		}
	}
	for _, chunk := range xChunks {
		linesToDate := []string{}
		half := len(chunk) / 2
		for i, line := range chunk {
			count := 0
			mirrored := true
			if i <= half {
				for j := len(linesToDate) - 1; j >= 0; j-- {
					if linesToDate[j] != chunk[i+count] {
						mirrored = false
						break
					}
					count++
				}
			} else {
				for j := len(linesToDate) - 1; j > i-half; j-- {
					if i+count >= len(chunk) {
						break
					}
					if linesToDate[j] != chunk[i+count] {
						mirrored = false
						break
					}
					count++
				}
			}
			if mirrored && len(linesToDate) != 0 {
				xMirror += len(linesToDate)
				break
			}
			linesToDate = append(linesToDate, line)

		}
	}
	answer += xMirror + yMirror

	return answer
}

func part2(lines []string) int {
	answer := 0
	yMirror := 0
	xMirror := 0

	yChunks := [][]string{}
	newChunk := []string{}
	for _, line := range lines {
		if line == "" {
			yChunks = append(yChunks, newChunk)
			newChunk = []string{}
			continue
		}
		newChunk = append(newChunk, line)
	}
	yChunks = append(yChunks, newChunk)

	xChunks := [][]string{}
	for _, chunk := range yChunks {
		newXChunk := make([]string, len(chunk[0]))
		for _, line := range chunk {
			for i, char := range line {
				newXChunk[i] += string(char)
			}
		}
		xChunks = append(xChunks, newXChunk)
	}

	for _, chunk := range yChunks {
		linesToDate := []string{}
		half := len(chunk) / 2
		lineDiffA := ""
		lineDiffB := ""
		diffLine := 0
		for i, line := range chunk {
			count := 0
			diffs := 0
			if i <= half {
				for j := len(linesToDate) - 1; j >= 0; j-- {
					if linesToDate[j] != chunk[i+count] {
						diffs++
						lineDiffA = linesToDate[j]
						lineDiffB = chunk[i+count]
						diffLine = len(linesToDate)
						if diffs > 1 {
							break
						}
					}
					count++
				}
			} else {
				for j := len(linesToDate) - 1; j > i-half; j-- {
					if i+count >= len(chunk) {
						break
					}
					if linesToDate[j] != chunk[i+count] {
						diffs++
						lineDiffA = linesToDate[j]
						lineDiffB = chunk[i+count]
						diffLine = len(linesToDate)
						if diffs > 1 {
							break
						}
					}
					count++
				}
			}
			if diffs == 1 {
				charDiffs := 0
				for j := 0; j < len(lineDiffA); j++ {
					if lineDiffA[j] != lineDiffB[j] {
						charDiffs++
						if charDiffs > 1 {
							break
						}
					}
				}
				if charDiffs == 1 {
					yMirror += diffLine * 100
					break

				}
			}
			linesToDate = append(linesToDate, line)

		}
	}
	for _, chunk := range xChunks {
		linesToDate := []string{}
		half := len(chunk) / 2
		lineDiffA := ""
		lineDiffB := ""
		diffLine := 0
		for i, line := range chunk {
			count := 0
			diffs := 0
			if i <= half {
				for j := len(linesToDate) - 1; j >= 0; j-- {
					if linesToDate[j] != chunk[i+count] {
						diffs++
						lineDiffA = linesToDate[j]
						lineDiffB = chunk[i+count]
						diffLine = len(linesToDate)
						if diffs > 1 {
							break
						}
					}
					count++
				}
			} else {
				for j := len(linesToDate) - 1; j >= i-half; j-- {
					if i+count >= len(chunk) {
						break
					}
					a := linesToDate[j]
					b := chunk[i+count]
					if false {
						fmt.Println(a)
						fmt.Println(b)
					}
					if linesToDate[j] != chunk[i+count] {
						diffs++
						lineDiffA = linesToDate[j]
						lineDiffB = chunk[i+count]
						diffLine = len(linesToDate)
						if diffs > 1 {
							break
						}
					}
					count++
				}
			}
			if diffs == 1 {
				charDiffs := 0
				for j := 0; j < len(lineDiffA); j++ {
					if lineDiffA[j] != lineDiffB[j] {
						charDiffs++
						if charDiffs > 1 {
							break
						}
					}
				}
				if charDiffs == 1 {
					xMirror += diffLine
					break
				}
			}
			linesToDate = append(linesToDate, line)

		}
	}
	answer += xMirror + yMirror

	return answer
}
