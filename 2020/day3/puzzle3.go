package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readInputLines("day3.txt")
	var grid []string

	for _, line := range lines {
		grid = append(grid, line)
	}

	var xPat [5]int
	var yPat [5]int
	xPat[0] = 1
	xPat[1] = 3
	xPat[2] = 5
	xPat[3] = 7
	xPat[4] = 1

	yPat[0] = 1
	yPat[1] = 1
	yPat[2] = 1
	yPat[3] = 1
	yPat[4] = 2

	trees := 0
	treesMult := 1

	for i := 0; i < 5; i++ {
		x := 0
		y := 0
		trees2 := 0
		patternx := xPat[i]
		patterny := yPat[i]
		for {
			x += patternx
			y += patterny
			if y >= len(lines) {
				break
			}
			x = x % len(lines[0])
			if string(grid[y][x]) == "#" {
				if i == 1 {
					trees++
				}
				trees2++
			}
		}
		treesMult *= trees2
	}

	fmt.Println(fmt.Sprintf("Answer 1: %d", trees))
	fmt.Println(fmt.Sprintf("Answer 2: %d", treesMult))

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
