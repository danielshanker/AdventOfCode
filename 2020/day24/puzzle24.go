package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	var instructions [][]string
	for _, line := range lines {
		var inst []string
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "e" || string(line[i]) == "w" {
				inst = append(inst, string(line[i]))
				continue
			}
			inst = append(inst, string(line[i])+string(line[i+1]))
			i++
		}
		instructions = append(instructions, inst)
	}

	maxCoord := 1500
	mid := maxCoord / 2
	var hexes [][]bool
	for i := 0; i < maxCoord; i++ {
		var b []bool
		for j := 0; j < maxCoord; j++ {
			b = append(b, false)
		}
		hexes = append(hexes, b)
	}

	//ogHexes := copyGrid(hexes)

	answer1 := 0
	for _, curInstruction := range instructions {
		x := mid
		y := mid
		for _, inst := range curInstruction {
			if inst == "w" {
				x--
			}
			if inst == "e" {
				x++
			}
			if inst == "nw" {
				x--
				y++
			}
			if inst == "ne" {
				y++
			}
			if inst == "sw" {
				y--
			}
			if inst == "se" {
				x++
				y--
			}
		}
		if hexes[x][y] {
			hexes[x][y] = false
			answer1--
		} else {
			hexes[x][y] = true
			answer1++
		}
	}

	answer2 := 0

	//	hexes = copyGrid(ogHexes)
	for i := 0; i < 100; i++ {
		answer2 = 0
		hexes = move(hexes)
		for x := 0; x < len(hexes); x++ {
			for y := 0; y < len(hexes); y++ {
				if hexes[x][y] == true {
					answer2++
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func move(hexes [][]bool) [][]bool {
	new := copyGrid(hexes)

	for i := 0; i < len(hexes); i++ {
		for j := 0; j < len(hexes[i]); j++ {
			new[i][j] = checkGrid(hexes, i, j)
		}
	}

	return new
}

func checkGrid(hexes [][]bool, x int, y int) bool {

	occupyCount := 0

	// w
	if x > 0 && hexes[x-1][y] == true {
		occupyCount++
	}
	// e
	if x < len(hexes)-1 && hexes[x+1][y] == true {
		occupyCount++
	}
	// nw
	if x > 0 && y < len(hexes)-1 && hexes[x-1][y+1] == true {
		occupyCount++
	}
	// se
	if y > 0 && x < len(hexes)-1 && hexes[x+1][y-1] == true {
		occupyCount++
	}
	// ne
	if y < len(hexes)-1 && hexes[x][y+1] == true {
		occupyCount++
	}
	//sw
	if y > 0 && hexes[x][y-1] == true {
		occupyCount++
	}

	if hexes[x][y] == true {
		if occupyCount == 0 || occupyCount > 2 {
			return false
		}
		return true
	}

	if occupyCount == 2 {
		return true
	}

	return false
}

func copyGrid(old [][]bool) [][]bool {
	var new [][]bool
	for i := 0; i < len(old); i++ {
		var newLine []bool
		for j := 0; j < len(old[i]); j++ {
			char := old[i][j]
			newLine = append(newLine, char)
		}
		new = append(new, newLine)
	}
	return new
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
