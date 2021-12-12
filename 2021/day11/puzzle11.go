package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day11/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day11/input.txt")
		part1(input)
		part2(input)
	}

}

type octos struct {
	o     [10][10]int
	count int
}

func part1(lines []string) {
	answer1 := 0
	var octoMap octos
	y := 0
	for _, line := range lines {
		x := 0
		for _, char := range line {
			octoMap.o[y][x] = s2i(string(char))
			x++
		}
		y++
	}

	for k := 0; k < 100; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octoMap.o[i][j]++
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octoMap.o[i][j] == 10 {
					octoMap.flash(i, j)
				}
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octoMap.o[i][j] > 9 {
					octoMap.o[i][j] = 0
				}
			}
		}
	}
	answer1 = octoMap.count
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func (o *octos) flash(x int, y int) {
	o.count++
	o.o[x][y] = 100
	// up left
	if y > 0 && x > 0 {
		if o.o[x-1][y-1] < 9 {
			o.o[x-1][y-1]++
		} else if o.o[x-1][y-1] == 9 {
			o.flash(x-1, y-1)
		}
	}
	// left
	if y > 0 {
		if o.o[x][y-1] < 9 {
			o.o[x][y-1]++
		} else if o.o[x][y-1] == 9 {
			o.flash(x, y-1)
		}
	}
	// down left
	if y > 0 && x < 9 {
		if o.o[x+1][y-1] < 9 {
			o.o[x+1][y-1]++
		} else if o.o[x+1][y-1] == 9 {
			o.flash(x+1, y-1)
		}
	}
	// up
	if x > 0 {
		if o.o[x-1][y] < 9 {
			o.o[x-1][y]++
		} else if o.o[x-1][y] == 9 {
			o.flash(x-1, y)
		}
	}
	// down
	if x < 9 {
		if o.o[x+1][y] < 9 {
			o.o[x+1][y]++
		} else if o.o[x+1][y] == 9 {
			o.flash(x+1, y)
		}
	}
	// up right
	if y < 9 && x > 0 {
		if o.o[x-1][y+1] < 9 {
			o.o[x-1][y+1]++
		} else if o.o[x-1][y+1] == 9 {
			o.flash(x-1, y+1)
		}
	}
	// right
	if y < 9 {
		if o.o[x][y+1] < 9 {
			o.o[x][y+1]++
		} else if o.o[x][y+1] == 9 {
			o.flash(x, y+1)
		}
	}
	// down right
	if y < 9 && x < 9 {
		if o.o[x+1][y+1] < 9 {
			o.o[x+1][y+1]++
		} else if o.o[x+1][y+1] == 9 {
			o.flash(x+1, y+1)
		}
	}
}

func part2(lines []string) {
	answer2 := 0

	var octoMap octos
	y := 0
	for _, line := range lines {
		x := 0
		for _, char := range line {
			octoMap.o[y][x] = s2i(string(char))
			x++
		}
		y++
	}

	for k := 0; k < 1000000; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octoMap.o[i][j]++
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octoMap.o[i][j] == 10 {
					octoMap.flash(i, j)
				}
			}
		}

		flashCount := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octoMap.o[i][j] > 9 {
					octoMap.o[i][j] = 0
					flashCount++
				}
			}
		}
		if flashCount == 100 {
			answer2 = k + 1
			break
		}
	}

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
