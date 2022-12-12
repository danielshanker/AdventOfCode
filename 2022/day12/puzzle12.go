package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day12/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day12/input.txt")
		part1(input)
		part2(input)
	}

}

type point struct {
	value          int
	possiblePoints []*point
	isStart        bool
	isEnd          bool
	explored       bool
	distance       int
}

func part1(lines []string) {
	answer1 := 0
	points := [][]*point{}

	for _, line := range lines {
		chars := []*point{}
		for _, r := range line {
			p := point{}
			if string(r) == "S" {
				p.value = 0
				p.isStart = true
				chars = append(chars, &p)
				continue
			}
			if string(r) == "E" {
				p.isEnd = true
				p.value = 26
				chars = append(chars, &p)
				continue
			}
			a := int(r) - 97
			p.value = a
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	var start *point

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			if curPoint.isStart {
				start = curPoint
			}
			//up
			if i > 0 {
				if points[i-1][j].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i-1][j])
				}
			}
			//down
			if i < len(points)-1 {
				if points[i+1][j].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i+1][j])
				}
			}
			//left
			if j > 0 {
				if points[i][j-1].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j-1])
				}
			}
			//right
			if j < len(points[0])-1 {
				if points[i][j+1].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j+1])
				}
			}
		}
	}

	q := queue{}
	start.explored = true
	start.distance = 0
	q.push(*start)
	for len(q) != 0 {
		curP := q.pop()
		if curP.isEnd {
			answer1 = curP.distance
			break
		}
		for _, p := range curP.possiblePoints {
			if !p.explored {
				p.explored = true
				p.distance = curP.distance + 1
				q.push(*p)
			}
		}
	}

	fmt.Printf("Answer 1 : %d\n", answer1)
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

func part2(lines []string) {
	answer2 := 0
	points := [][]*point{}

	for _, line := range lines {
		chars := []*point{}
		for _, r := range line {
			p := point{}
			if string(r) == "S" {
				p.value = 0
				chars = append(chars, &p)
				continue
			}
			if string(r) == "E" {
				p.isEnd = true
				p.value = 26
				chars = append(chars, &p)
				continue
			}
			a := int(r) - 97
			p.value = a
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	var starts []*point

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			if curPoint.value == 0 {
				starts = append(starts, curPoint)
			}
			//up
			if i > 0 {
				if points[i-1][j].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i-1][j])
				}
			}
			//down
			if i < len(points)-1 {
				if points[i+1][j].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i+1][j])
				}
			}
			//left
			if j > 0 {
				if points[i][j-1].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j-1])
				}
			}
			//right
			if j < len(points[0])-1 {
				if points[i][j+1].value <= (curPoint.value + 1) {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[i][j+1])
				}
			}
		}
	}

	for _, start := range starts {
		q := queue{}
		start.isStart = true
		start.explored = true
		start.distance = 0
		q.push(*start)
		for len(q) != 0 {
			curP := q.pop()
			if curP.isEnd {
				if curP.distance < answer2 || answer2 == 0 {
					answer2 = curP.distance
				}
				break
			}
			for _, p := range curP.possiblePoints {
				if !p.explored {
					p.explored = true
					p.distance = curP.distance + 1
					q.push(*p)
				}
			}
		}
		reset(points, start)
	}

	fmt.Printf("Answer 2 : %d\n", answer2)
}

func reset(points [][]*point, start *point) {
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			points[i][j].distance = 0
			points[i][j].explored = false
		}
	}
	start.isStart = false
}
