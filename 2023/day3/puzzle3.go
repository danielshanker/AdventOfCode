package main

import (
	"flag"
	"fmt"
	"strconv"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		expectedAnswer := 4361
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day3/sample.txt")
		answer := part1(sample)
		fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer))
		expectedAnswer = 467835
		answer2 := part2(sample)
		fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day3/input.txt")
		fmt.Printf("day3 Answer 1 : %d\n", part1(input))
		fmt.Printf("day3 Answer 2 : %d\n", part2(input))
	}

}

type coord struct {
	value string
	x1    int
	x2    int
	x3    int
	y     int
}

func part1(lines []string) int {
	answer := 0
	numberCoords := []coord{}
	symbolCoords := []coord{}
	partNumbers := []int{}
	for y, line := range lines {
		if line == "" {
			break
		}
		num := ""
		x1 := 0
		firstDigit := true
		for x, cha := range line {
			_, isNum := strconv.Atoi(string(cha))
			if isNum == nil {
				num += string(cha)
				if firstDigit {
					x1 = x
					firstDigit = false
				}
			} else if string(cha) == "." {
				if num != "" {
					nc := coord{
						value: num,
						x1:    x1,
						y:     y,
					}
					numberCoords = append(numberCoords, nc)
					num = ""
					x1 = 0
					firstDigit = true
				}
			} else {
				if num != "" {
					nc := coord{
						value: num,
						x1:    x1,
						y:     y,
					}
					numberCoords = append(numberCoords, nc)
					num = ""
					x1 = 0
					firstDigit = true
				}
				sc := coord{
					value: string(cha),
					x1:    x,
					y:     y,
				}
				symbolCoords = append(symbolCoords, sc)
			}
			if !firstDigit && x == len(line)-1 {
				nc := coord{
					value: num,
					x1:    x1,
					y:     y,
				}
				numberCoords = append(numberCoords, nc)
				num = ""
				x1 = 0
				firstDigit = true
			}
		}
	}

	for _, numCoord := range numberCoords {
		for _, symbolCoord := range symbolCoords {
			numLength := len(numCoord.value)
			//left
			if (symbolCoord.x1 == numCoord.x1-1) && (symbolCoord.y == numCoord.y) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//left up
			if (symbolCoord.x1 == numCoord.x1-1) && (symbolCoord.y == numCoord.y-1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//left down
			if (symbolCoord.x1 == numCoord.x1-1) && (symbolCoord.y == numCoord.y+1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//down
			if (symbolCoord.x1 == numCoord.x1) && (symbolCoord.y == numCoord.y-1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//up
			if (symbolCoord.x1 == numCoord.x1) && (symbolCoord.y == numCoord.y+1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//right
			if (symbolCoord.x1 == numCoord.x1+numLength) && (symbolCoord.y == numCoord.y) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//right up
			if (symbolCoord.x1 == numCoord.x1+numLength) && (symbolCoord.y == numCoord.y-1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			//right down
			if (symbolCoord.x1 == numCoord.x1+numLength) && (symbolCoord.y == numCoord.y+1) {
				partNumbers = append(partNumbers, S2i(numCoord.value))
				break
			}
			if numLength == 2 {
				//right up
				if (symbolCoord.x1 == numCoord.x1+1) && (symbolCoord.y == numCoord.y-1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
				//right down
				if (symbolCoord.x1 == numCoord.x1+1) && (symbolCoord.y == numCoord.y+1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
			}
			if numLength == 3 {
				//right up
				if (symbolCoord.x1 == numCoord.x1+1) && (symbolCoord.y == numCoord.y-1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
				//right down
				if (symbolCoord.x1 == numCoord.x1+1) && (symbolCoord.y == numCoord.y+1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
				//right up
				if (symbolCoord.x1 == numCoord.x1+2) && (symbolCoord.y == numCoord.y-1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
				//right down
				if (symbolCoord.x1 == numCoord.x1+2) && (symbolCoord.y == numCoord.y+1) {
					partNumbers = append(partNumbers, S2i(numCoord.value))
					break
				}
			}
		}
	}
	for _, num := range partNumbers {
		answer += num
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	numberCoords := []coord{}
	starCoords := []coord{}
	for y, line := range lines {
		num := ""
		x1 := 0
		x2 := -10
		x3 := -10
		firstDigit := true
		for x, cha := range line {
			_, isNum := strconv.Atoi(string(cha))
			if isNum == nil {
				num += string(cha)
				if len(num) == 1 {
					x1 = x
					firstDigit = false
				} else if len(num) == 2 {
					x2 = x
				} else if len(num) == 3 {
					x3 = x
				}
			} else if string(cha) == "." {
				if num != "" {
					nc := coord{
						value: num,
						x1:    x1,
						x2:    x2,
						x3:    x3,
						y:     y,
					}
					numberCoords = append(numberCoords, nc)
					num = ""
					x1 = 0
					x2 = -10
					x3 = -10
					firstDigit = true
				}
			} else {
				if num != "" {
					nc := coord{
						value: num,
						x1:    x1,
						x2:    x2,
						x3:    x3,
						y:     y,
					}
					numberCoords = append(numberCoords, nc)
					num = ""
					x1 = 0
					x2 = -10
					x3 = -10
					firstDigit = true
				}
				if string(cha) == "*" {
					sc := coord{
						value: string(cha),
						x1:    x,
						y:     y,
					}
					starCoords = append(starCoords, sc)
				}
			}
			if !firstDigit && x == len(line)-1 {
				nc := coord{
					value: num,
					x1:    x1,
					x2:    x2,
					x3:    x3,
					y:     y,
				}
				numberCoords = append(numberCoords, nc)
				num = ""
				x1 = 0
				x2 = -10
				x3 = -10
				firstDigit = true
			}
		}
	}

	for _, star := range starCoords {
		found := 0
		num1 := 0
		num2 := 0
		for _, num := range numberCoords {
			l := star.x1 - 1
			r := star.x1 + 1
			u := star.y - 1
			d := star.y + 1
			//left
			if star.y == num.y && (l == num.x1 || l == num.x2 || l == num.x3) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//left up
			if u == num.y && (l == num.x1 || l == num.x2 || l == num.x3) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//left down
			if d == num.y && (l == num.x1 || l == num.x2 || l == num.x3) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//right
			if star.y == num.y && (r == num.x1) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//right up
			if u == num.y && (r == num.x1) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//right down
			if d == num.y && (r == num.x1) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			//up
			if u == num.y && (star.x1 == num.x1 || star.x1 == num.x2 || star.x1 == num.x3) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
			// down
			if d == num.y && (star.x1 == num.x1 || star.x1 == num.x2 || star.x1 == num.x3) {
				if num1 == 0 {
					num1 = S2i(num.value)
				} else {
					num2 = S2i(num.value)
				}
				found++
				continue
			}
		}
		if found == 2 {
			answer += num1 * num2
		}
	}

	return answer
}
