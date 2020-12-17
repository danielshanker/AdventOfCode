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
	var grid [][][]string
	for i := 0; i < 30; i++ {
		var box [][]string
		for j := 0; j < 30; j++ {
			var l []string
			for k := 0; k < 30; k++ {
				l = append(l, ".")
			}
			box = append(box, l)
		}
		grid = append(grid, box)
	}
	mid := 15

	y := 0
	for _, line := range lines {
		charArray := strings.Split(line, "")
		for i, j := range charArray {
			grid[mid+i][mid+y][mid] = j
		}
		y++
	}

	answer1 := 0
	for i := 0; i < 6; i++ {
		grid = move(grid)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for k := 0; k < len(grid[i][j]); k++ {
				if grid[i][j][k] == "#" {
					answer1++
				}
			}
		}
	}

	answer2 := problem2(lines)

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func problem2(lines []string) int {
	var grid [][][][]string
	for l := 0; l < 30; l++ {
		var cube [][][]string
		for i := 0; i < 30; i++ {
			var box [][]string
			for j := 0; j < 30; j++ {
				var l []string
				for k := 0; k < 30; k++ {
					l = append(l, ".")
				}
				box = append(box, l)
			}
			cube = append(cube, box)
		}
		grid = append(grid, cube)
	}
	mid := 15

	y := 0
	for _, line := range lines {
		charArray := strings.Split(line, "")
		for i, j := range charArray {
			grid[mid+i][mid+y][mid][mid] = j
		}
		y++
	}

	for i := 0; i < 6; i++ {
		grid = move2(grid)
	}

	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for k := 0; k < len(grid[i][j]); k++ {
				for l := 0; l < len(grid[i][j][k]); l++ {
					if grid[i][j][k][l] == "#" {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func move(grid [][][]string) [][][]string {

	newGrid := copyGrid(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for k := 0; k < len(grid[i][j]); k++ {
				newGrid[i][j][k] = checkGrid(grid, i, j, k)
			}
		}
	}
	grid = copyGrid(newGrid)
	return grid
}

func checkGrid(grid [][][]string, x int, y int, z int) string {

	occupyCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				if x+i < 0 || x+i >= len(grid) || y+j < 0 || y+j >= len(grid[0]) || z+k < 0 || z+k >= len(grid[0][0]) {
					continue
				}
				if grid[x+i][y+j][z+k] == "#" {
					occupyCount++
				}
			}
		}
	}

	if grid[x][y][z] == "#" {
		if occupyCount == 3 || occupyCount == 2 {
			return "#"
		} else {
			return "."
		}
	}

	if grid[x][y][z] == "." {
		if occupyCount == 3 {
			return "#"
		} else {
			return "."
		}
	}

	return "."
}

func copyGrid(old [][][]string) [][][]string {
	var new [][][]string
	for i := 0; i < len(old); i++ {
		var newBox [][]string
		for j := 0; j < len(old[i]); j++ {
			var newLine []string
			for k := 0; k < len(old[i][j]); k++ {
				char := old[i][j][k]
				newLine = append(newLine, char)
			}
			newBox = append(newBox, newLine)
		}
		new = append(new, newBox)
	}
	return new
}

func move2(grid [][][][]string) [][][][]string {

	newGrid := copyGrid2(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for k := 0; k < len(grid[i][j]); k++ {
				for l := 0; l < len(grid[i][j][k]); l++ {
					newGrid[i][j][k][l] = checkGrid2(grid, i, j, k, l)
				}
			}
		}
	}
	grid = copyGrid2(newGrid)
	return grid
}
func copyGrid2(old [][][][]string) [][][][]string {
	var new [][][][]string
	for i := 0; i < len(old); i++ {
		var newCube [][][]string
		for j := 0; j < len(old[i]); j++ {
			var newBox [][]string
			for k := 0; k < len(old[i][j]); k++ {
				var newLine []string
				for l := 0; l < len(old[i][j][k]); l++ {
					char := old[i][j][k][l]
					newLine = append(newLine, char)
				}
				newBox = append(newBox, newLine)
			}
			newCube = append(newCube, newBox)
		}
		new = append(new, newCube)
	}
	return new
}
func checkGrid2(grid [][][][]string, x int, y int, z int, w int) string {

	occupyCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					if x+i < 0 || x+i >= len(grid) || y+j < 0 || y+j >= len(grid[0]) || z+k < 0 || z+k >= len(grid[0][0]) || w+l < 0 || w+l >= len(grid[0][0][0]) {
						continue
					}
					if grid[x+i][y+j][z+k][w+l] == "#" {
						occupyCount++
					}
				}
			}
		}
	}

	if grid[x][y][z][w] == "#" {
		if occupyCount == 3 || occupyCount == 2 {
			return "#"
		} else {
			return "."
		}
	}

	if grid[x][y][z][w] == "." {
		if occupyCount == 3 {
			return "#"
		} else {
			return "."
		}
	}

	return "."
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
