package main

import (
	"flag"
	"fmt"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day18/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day18/input.txt")
		part1(input)
		part2(input)
	}

}

func part1(lines []string) {
	answer := 0
	var grid [22][22][22]bool
	for _, line := range lines {
		coord := strings.Split(line, ",")
		grid[S2i(coord[0])+1][S2i(coord[1])+1][S2i(coord[2])+1] = true
	}

	for x := 1; x < 21; x++ {
		for y := 1; y < 21; y++ {
			for z := 1; z < 21; z++ {
				if grid[x][y][z] {
					// right
					if !grid[x+1][y][z] {
						answer++
					}
					// left
					if !grid[x-1][y][z] {
						answer++
					}
					// up
					if !grid[x][y-1][z] {
						answer++
					}
					// down
					if !grid[x][y+1][z] {
						answer++
					}
					// foward
					if !grid[x][y][z+1] {
						answer++
					}
					// back
					if !grid[x][y][z-1] {
						answer++
					}
				}
			}
		}
	}

	fmt.Printf("Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0

	var grid [22][22][22]bool
	for _, line := range lines {
		coord := strings.Split(line, ",")
		grid[S2i(coord[0])+1][S2i(coord[1])+1][S2i(coord[2])+1] = true
	}

	var s Stack
	s.Push("0,0,0")
	explored := map[string]bool{
		"0,0,0": true,
	}

	for len(s) > 0 {
		curP := s.Pop()
		coord := strings.Split(curP, ",")
		x := S2i(coord[0])
		y := S2i(coord[1])
		z := S2i(coord[2])

		// right
		if x < 21 {
			key := fmt.Sprintf("%d,%d,%d", x+1, y, z)
			if !explored[key] {
				if grid[x+1][y][z] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}
		// left
		if x > 0 {
			key := fmt.Sprintf("%d,%d,%d", x-1, y, z)
			if !explored[key] {
				if grid[x-1][y][z] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}
		// down
		if y < 21 {
			key := fmt.Sprintf("%d,%d,%d", x, y+1, z)
			if !explored[key] {
				if grid[x][y+1][z] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}
		// up
		if y > 0 {
			key := fmt.Sprintf("%d,%d,%d", x, y-1, z)
			if !explored[key] {
				if grid[x][y-1][z] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}
		// toward
		if z < 21 {
			key := fmt.Sprintf("%d,%d,%d", x, y, z+1)
			if !explored[key] {
				if grid[x][y][z+1] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}
		// away
		if z > 0 {
			key := fmt.Sprintf("%d,%d,%d", x, y, z-1)
			if !explored[key] {
				if grid[x][y][z-1] {
					answer++
				} else {
					s.Push(key)
					explored[key] = true
				}
			}
		}

	}

	fmt.Printf("Answer 2 : %d\n", answer)
}
