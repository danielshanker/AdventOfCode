package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")
	var grid [][]string

	allDot := strings.Repeat(".", len(lines[0])+2)
	border := strings.Split(allDot, "")

	grid = append(grid, border)

	for _, line := range lines {
		line = "." + line + "."
		charArray := strings.Split(line, "")
		grid = append(grid, charArray)
	}
	grid = append(grid, border)
	ogGrid := copyGrid(grid)

	grid = move(grid)
	answer1 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" {
				answer1++
			}
		}
	}

	answer2 := 0
	grid = move2(ogGrid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" {
				answer2++
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func move(grid [][]string) [][]string {
	hadChange := true

	newGrid := copyGrid(grid)
	for hadChange {
		localHadChange := false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "." {
					continue
				}
				changed := false
				newGrid[i][j], changed = checkGrid(grid, i, j)
				if changed {
					localHadChange = true
				}
			}
		}
		grid = copyGrid(newGrid)
		if !localHadChange {
			hadChange = false
		}
	}
	return grid
}

func move2(grid [][]string) [][]string {
	hadChange := true

	newGrid := copyGrid(grid)

	for hadChange {
		localHadChange := false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "." {
					continue
				}
				changed := false
				newGrid[i][j], changed = checkGrid2(grid, i, j)
				if changed {
					localHadChange = true
				}
			}
		}
		grid = copyGrid(newGrid)
		if !localHadChange {
			hadChange = false
		}
	}
	return grid
}

func copyGrid(old [][]string) [][]string {
	var new [][]string
	for i := 0; i < len(old); i++ {
		var newLine []string
		for j := 0; j < len(old[i]); j++ {
			char := old[i][j]
			newLine = append(newLine, char)
		}
		new = append(new, newLine)
	}
	return new
}

func checkGrid(grid [][]string, x int, y int) (string, bool) {

	occupyCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid[x+i][y+j] == "#" {
				occupyCount++
			}

		}
	}

	if grid[x][y] == "#" {
		if occupyCount >= 4 {
			return "L", true
		} else {
			return "#", false
		}
	}

	if grid[x][y] == "L" {
		if occupyCount == 0 {
			return "#", true
		} else {
			return "L", false
		}
	}

	return ".", false
}

func checkGrid2(grid [][]string, x int, y int) (string, bool) {

	occupyCount := 0
	//u
	for i := 1; i <= x; i++ {
		if grid[x-i][y] == "#" {
			occupyCount++
			break
		}
		if grid[x-i][y] == "L" {
			break
		}
	}
	//ul
	for i := 1; i <= x; i++ {
		if y-i < 0 {
			break
		}
		if grid[x-i][y-i] == "#" {
			occupyCount++
			break
		}
		if grid[x-i][y-i] == "L" {
			break
		}
	}
	//ur
	for i := 1; i <= x; i++ {
		if y+i >= len(grid[x]) {
			break
		}
		if grid[x-i][y+i] == "#" {
			occupyCount++
			break
		}
		if grid[x-i][y+i] == "L" {
			break
		}
	}
	//d
	for i := 1; i < len(grid)-x; i++ {
		if grid[x+i][y] == "#" {
			occupyCount++
			break
		}
		if grid[x+i][y] == "L" {
			break
		}
	}
	//dl
	for i := 1; i < len(grid)-x; i++ {
		if y-i < 0 {
			break
		}
		if grid[x+i][y-i] == "#" {
			occupyCount++
			break
		}
		if grid[x+i][y-i] == "L" {
			break
		}
	}
	//dr
	for i := 1; i < len(grid)-x; i++ {
		if y+i >= len(grid[x]) {
			break
		}
		if grid[x+i][y+i] == "#" {
			occupyCount++
			break
		}
		if grid[x+i][y+i] == "L" {
			break
		}
	}
	//l
	for i := 1; i <= y; i++ {
		if grid[x][y-i] == "#" {
			occupyCount++
			break
		}
		if grid[x][y-i] == "L" {
			break
		}
	}
	//r
	for i := 1; i < len(grid[x])-y; i++ {
		if grid[x][y+i] == "#" {
			occupyCount++
			break
		}
		if grid[x][y+i] == "L" {
			break
		}
	}

	if grid[x][y] == "#" {
		if occupyCount >= 5 {
			return "L", true
		} else {
			return "#", false
		}
	}

	if grid[x][y] == "L" {
		if occupyCount == 0 {
			return "#", true
		} else {
			return "L", false
		}
	}

	return ".", false
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
