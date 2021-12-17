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
		part1(20, 30, -10, -5)
		part2(20, 30, -10, -5)
	} else {
		part1(79, 137, -176, -117)
		part2(79, 137, -176, -117)
	}

}

func part1(xMin int, xMax int, yMin int, yMax int) {
	answer1 := 0
	var viableX []int
	for x := 1; x < xMax; x++ {
		dist := 0
		step := 0
		for i := x; i > 0; i-- {
			dist = dist + i
			step++
			if dist <= xMax && dist >= xMin {
				viableX = append(viableX, x)
				break
			}
		}
	}

	var viableY []int
	for y := -1000; y < 1000; y++ {
		dist := 0
		step := 0
		for i := y; dist > yMin; i-- {
			dist = dist + i
			step++
			if dist <= yMax && dist >= yMin {
				viableY = append(viableY, y)
				break
			}
		}
	}

	for i := len(viableY) - 1; i >= 0; i-- {
		for _, xVal := range viableX {
			x := xVal
			xD := 0
			yD := 0
			y := viableY[i]
			highest := 0
			for k := 0; k < 10000; k++ {
				xD += x
				yD += y
				if highest < yD {
					highest = yD
				}
				if x > 0 {
					x--
				}
				y--
				if xD <= xMax && xD >= xMin && yD <= yMax && yD >= yMin {
					if highest > answer1 {
						answer1 = highest
					}
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(xMin int, xMax int, yMin int, yMax int) {
	answer2 := 0

	var viableX []int
	for x := 1; x <= xMax; x++ {
		dist := 0
		step := 0
		for i := x; i > 0; i-- {
			dist = dist + i
			step++
			if dist <= xMax && dist >= xMin {
				viableX = append(viableX, x)
				break
			}
		}
	}

	var viableY []int
	for y := -1000; y < 1000; y++ {
		dist := 0
		step := 0
		for i := y; dist > yMin; i-- {
			dist = dist + i
			step++
			if dist <= yMax && dist >= yMin {
				viableY = append(viableY, y)
				break
			}
		}
	}

	found := map[string]bool{}

	for i := len(viableY) - 1; i >= 0; i-- {
		for _, xVal := range viableX {
			x := xVal
			xD := 0
			yD := 0
			y := viableY[i]
			for k := 0; k < 1000; k++ {
				xD += x
				yD += y
				if x > 0 {
					x--
				}
				y--
				if xD <= xMax && xD >= xMin && yD <= yMax && yD >= yMin {
					found[rowCol(xVal, viableY[i])] = true
					answer2++
				}
			}
		}
	}
	answer2 = len(found)
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

func rowCol(row int, col int) string {
	s := strconv.Itoa(row)
	s += "-"
	s += strconv.Itoa(col)
	return s
}
