package main

import (
	"flag"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 10, part1, part2, 36, 81)

}

type point struct {
	value          int
	possiblePoints []*point
	explored       bool
	x              int
	y              int
}

func part1(lines []string) int {
	answer := 0

	points := [][]*point{}
	starts := []*point{}
	ends := []*point{}

	for y, line := range lines {
		chars := []*point{}
		for x, r := range line {
			p := point{}
			p.y = y
			p.x = x
			if string(r) == "0" {
				starts = append(starts, &p)
			}
			if string(r) == "9" {
				ends = append(ends, &p)
			}
			p.value = S2i(string(r))
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			//up
			if i > 0 {
				if points[i-1][j].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i-1][j])
				}
			}
			//down
			if i < len(points)-1 {
				if points[i+1][j].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i+1][j])
				}
			}
			//left
			if j > 0 {
				if points[i][j-1].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j-1])
				}
			}
			//right
			if j < len(points[0])-1 {
				if points[i][j+1].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j+1])
				}
			}
		}
	}

	for _, start := range starts {
		for _, end := range ends {
			q := queue{}
			start.explored = true
			q.push(*points[start.y][start.x])
			for len(q) != 0 {
				curP := q.pop()
				if curP.x == end.x && curP.y == end.y {
					answer++
					break
				}
				for _, p := range curP.possiblePoints {
					if !p.explored {
						p.explored = true
						q.push(*p)
					}
				}
			}
			clearExplored(points)
		}
	}

	return answer
}

func clearExplored(p [][]*point) {

	for _, a := range p {
		for _, b := range a {
			b.explored = false
		}
	}

}

func part2(lines []string) int {
	answer := 0

	points := [][]*point{}
	starts := []*point{}
	ends := []*point{}

	for y, line := range lines {
		chars := []*point{}
		for x, r := range line {
			p := point{}
			p.y = y
			p.x = x
			if string(r) == "0" {
				starts = append(starts, &p)
			}
			if string(r) == "9" {
				ends = append(ends, &p)
			}
			p.value = S2i(string(r))
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			//up
			if i > 0 {
				if points[i-1][j].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i-1][j])
				}
			}
			//down
			if i < len(points)-1 {
				if points[i+1][j].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i+1][j])
				}
			}
			//left
			if j > 0 {
				if points[i][j-1].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j-1])
				}
			}
			//right
			if j < len(points[0])-1 {
				if points[i][j+1].value == (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j+1])
				}
			}
		}
	}

	for _, start := range starts {
		for _, end := range ends {
			q := queue{}
			q.push(*points[start.y][start.x])
			for len(q) != 0 {
				curP := q.pop()
				if curP.x == end.x && curP.y == end.y {
					answer++
				}
				for _, p := range curP.possiblePoints {
					if !p.explored {
						q.push(*p)
					}
				}
			}
			clearExplored(points)
		}
	}

	return answer
}

type queue []point

func (q *queue) pop() point {
	if len(*q) == 0 {
		return point{}
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q *queue) push(p point) {
	*q = append(*q, p)
}
