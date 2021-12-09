package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day9/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day9/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer1 := 0
	var lavaMap = map[int]map[int]int{}

	for i, line := range lines {
		lavaMap[i] = map[int]int{}
		for j, char := range line {
			lavaMap[i][j] = s2i(string(char))
		}
	}
	var riskLevels []int
	for i:=0; i < len(lavaMap); i++ {
		for j:= 0; j < len(lavaMap[i]); j++ {
			if (i-1 < 0 || lavaMap[i-1][j] > lavaMap[i][j]) && (i+1 >= len(lavaMap) || lavaMap[i+1][j] > lavaMap[i][j]) && (j-1 < 0 || lavaMap[i][j-1] > lavaMap[i][j]) && (j+1 >= len(lavaMap[i]) || lavaMap[i][j+1] > lavaMap[i][j]) {
				riskLevels = append(riskLevels, lavaMap[i][j])
			}
		}
	}

	for _, rl := range riskLevels {
		answer1 += rl+1
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

type point struct {
	x int
	y int
	value int
}

func part2(lines []string) {
	answer2 := 1

	var lavaMap = map[int]map[int]int{}

	for i, line := range lines {
		lavaMap[i] = map[int]int{}
		for j, char := range line {
			lavaMap[i][j] = s2i(string(char))
		}
	}
	var lowPoints []point
	var basins [][]point
	for i:=0; i < len(lavaMap); i++ {
		for j:= 0; j < len(lavaMap[i]); j++ {
			if (i-1 < 0 || lavaMap[i-1][j] > lavaMap[i][j]) && (i+1 >= len(lavaMap) || lavaMap[i+1][j] > lavaMap[i][j]) && (j-1 < 0 || lavaMap[i][j-1] > lavaMap[i][j]) && (j+1 >= len(lavaMap[i]) || lavaMap[i][j+1] > lavaMap[i][j]) {
				var lp point
				lp.x = i
				lp.y = j
				lp.value = lavaMap[i][j]
				lowPoints = append(lowPoints, lp)
				temp := []point{lp}
				basins = append(basins,temp)
			}
		}
	}
	visited := map[int]map[int]bool{}
	var finalBasins [][]point
	for _, basin := range basins {
		basin = calculateBasin(basin, lavaMap, visited)
		finalBasins = append(finalBasins, basin)
	}

	var bLen []int

	for _, basin := range finalBasins {
		bLen = append(bLen, len(basin))
	}
	sort.Ints(bLen)

	answer2 = bLen[len(bLen)-1] * bLen[len(bLen)-2] * bLen[len(bLen)-3]

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func calculateBasin(basin []point, lavaMap map[int]map[int]int, visited map[int]map[int]bool) []point {
	p := basin[len(basin)-1]
	x := p.x
	y := p.y
	if _, ok := visited[x]; ok {
		if _, ok := visited[x][y]; ok {
			return basin
		}
	}
	if _, ok := visited[x]; !ok {
		visited[x] = map[int]bool{}
	}
	visited[x][y] = true

	//up
	if y-1 >= 0 && lavaMap[x][y-1] != 9 {
		newPoint := point {
			x: x,
			y: y-1,
			value: lavaMap[x][y-1],
		}
		if _, ok := visited[x]; ok {
			if _, ok := visited[x][y-1]; !ok {
				basin = append(basin, newPoint)
				basin = calculateBasin(basin, lavaMap, visited)
			}
		} else {
			basin = append(basin, newPoint)
			basin = calculateBasin(basin, lavaMap, visited)
		}
	}
	//down
	if y+1 < len(lavaMap[x]) && lavaMap[x][y+1] != 9 {
		newPoint := point {
			x: x,
			y: y+1,
			value: lavaMap[x][y+1],
		}
		if _, ok := visited[x]; ok {
			if _, ok := visited[x][y+1]; !ok {
				basin = append(basin, newPoint)
				basin = calculateBasin(basin, lavaMap, visited)
			}
		} else {
			basin = append(basin, newPoint)
			basin = calculateBasin(basin, lavaMap, visited)
		}
	}

	//left
	if x-1 >= 0 && lavaMap[x-1][y] != 9 {
		newPoint := point {
			x: x-1,
			y: y,
			value: lavaMap[x-1][y],
		}
		if _, ok := visited[x-1]; ok {
			if _, ok := visited[x-1][y]; !ok {
				basin = append(basin, newPoint)
				basin = calculateBasin(basin, lavaMap, visited)
			}
		} else {
			basin = append(basin, newPoint)
			basin = calculateBasin(basin, lavaMap, visited)
		}
	}

	//right
	if x+1 < len(lavaMap) && lavaMap[x+1][y] != 9 {
		newPoint := point {
			x: x+1,
			y: y,
			value: lavaMap[x+1][y],
		}
		if _, ok := visited[x+1]; ok {
			if _, ok := visited[x+1][y]; !ok {
				basin = append(basin, newPoint)
				basin = calculateBasin(basin, lavaMap, visited)
			}
		} else {
			basin = append(basin, newPoint)
			basin = calculateBasin(basin, lavaMap, visited)
		}
	}

	return basin
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
