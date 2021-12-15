package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day15/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day15/input.txt")
		part1(input)
		part2(input)
	}

}

type Point struct {
	value    int
	distance int
}

func rowCol(row int, col int) string {
	s := strconv.Itoa(row)
	s += "-"
	s += strconv.Itoa(col)
	return s
}
func getRowCol(rc string) (int, int) {
	t := strings.Split(rc, "-")
	return s2i(t[0]), s2i(t[1])
}

func part1(lines []string) {
	answer1 := 0
	grid := map[string]Point{}
	for row, line := range lines {
		for col, c := range line {
			p := Point{
				value:    s2i(string(c)),
				distance: math.MaxInt64,
			}
			if row == 0 && col == 0 {
				p.distance = 0
			}
			grid[rowCol(row, col)] = p
		}
	}
	dist := 0
	active := map[string]bool{}
	active[rowCol(0, 0)] = true
	for len(grid) > 0 {
		if _, ok := grid[rowCol(len(lines)-1, len(lines[0])-1)]; ok {
			dist = grid[rowCol(len(lines)-1, len(lines[0])-1)].distance
		}
		minDist := math.MaxInt64
		var point Point
		var minKey string
		for i, _ := range active {
			p := grid[i]
			if p.distance < minDist {
				point = p
				minDist = p.distance
				minKey = i
			}
		}
		delete(grid, minKey)
		delete(active, minKey)

		row, col := getRowCol(minKey)
		//up
		if val, ok := grid[rowCol(row-1, col)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row-1, col)] = val
			active[rowCol(row-1, col)] = true
		}
		//down
		if val, ok := grid[rowCol(row+1, col)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row+1, col)] = val
			active[rowCol(row+1, col)] = true
		}
		//left
		if val, ok := grid[rowCol(row, col-1)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row, col-1)] = val
			active[rowCol(row, col-1)] = true
		}
		//right
		if val, ok := grid[rowCol(row, col+1)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row, col+1)] = val
			active[rowCol(row, col+1)] = true
		}
	}

	answer1 = dist
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0
	grid := map[string]Point{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for row, line := range lines {
				for col, c := range line {
					val := (s2i(string(c)) + i + j) % 9
					if val == 0 {
						val = 9
					}
					p := Point{
						value:    val,
						distance: math.MaxInt64,
					}
					if i == 0 && j == 0 && row == 0 && col == 0 {
						p.distance = 0
					}
					grid[rowCol(row+len(lines)*i, col+len(lines[0])*j)] = p
				}
			}
		}
	}

	dist := 0
	active := map[string]bool{}
	active[rowCol(0, 0)] = true
	for len(grid) > 0 {
		if _, ok := grid[rowCol(len(lines)*5-1, len(lines[0])*5-1)]; ok {
			dist = grid[rowCol(len(lines)*5-1, len(lines[0])*5-1)].distance
		}
		minDist := math.MaxInt64
		var point Point
		var minKey string
		for i, _ := range active {
			p := grid[i]
			if p.distance < minDist {
				point = p
				minDist = p.distance
				minKey = i
			}
		}
		delete(grid, minKey)
		delete(active, minKey)

		row, col := getRowCol(minKey)
		//up
		if val, ok := grid[rowCol(row-1, col)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row-1, col)] = val
			active[rowCol(row-1, col)] = true
		}
		//down
		if val, ok := grid[rowCol(row+1, col)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row+1, col)] = val
			active[rowCol(row+1, col)] = true
		}
		//left
		if val, ok := grid[rowCol(row, col-1)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row, col-1)] = val
			active[rowCol(row, col-1)] = true
		}
		//right
		if val, ok := grid[rowCol(row, col+1)]; ok {
			alt := point.distance + val.value
			if alt < val.distance {
				val.distance = alt
			}
			grid[rowCol(row, col+1)] = val
			active[rowCol(row, col+1)] = true
		}
	}
	answer2 = dist
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
