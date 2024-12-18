package main

import (
	"flag"
	"math"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 18, part1, part2, 22, 601)

}

type point struct {
	possiblePoints []*point
	isStart        bool
	isEnd          bool
	explored       bool
	distance       int
	c              Coord
	isWall         bool
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

func part1(lines []string) int {
	answer := math.MaxInt
	points := map[Coord]*point{}

	maxWidth := 6
	maxHeight := 6
	startBytes := 12

	//maxWidth = 7
	//maxHeight = 7
	//startBytes = 12
	maxWidth = 71
	maxHeight = 71
	startBytes = 1024

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			points[c] = &point{}
		}
	}

	for i, line := range lines {
		s := strings.Split(line, ",")
		x := S2i(s[0])
		y := S2i(s[1])
		c := Coord{x, y}
		points[c] = &point{
			c: c,
		}
		if i < startBytes {
			points[c].isWall = true
		}
	}
	points[Coord{X: maxWidth - 1, Y: maxHeight - 1}].isEnd = true
	var start *point
	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			curPoint := points[c]
			if curPoint.isWall {
				continue
			}
			curPoint.possiblePoints = []*point{}
			if c == (Coord{0, 0}) {
				start = curPoint
			}

			//north
			if i > 0 {
				nc := Coord{j, i - 1}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//south
			if i < maxHeight-1 {
				nc := Coord{j, i + 1}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//west
			if j > 0 {
				nc := Coord{j - 1, i}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//east
			if j < maxWidth-1 {
				nc := Coord{j + 1, i}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
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
			if curP.distance < answer {
				answer = curP.distance
			}
			continue
		}
		for _, p := range curP.possiblePoints {
			if !p.explored {
				p.explored = true
				p.distance = curP.distance + 1
				q.push(*p)
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	points := map[Coord]*point{}

	maxWidth := 0
	maxHeight := 0
	startBytes := 0

	//maxWidth = 7
	//maxHeight = 7
	//startBytes = 12
	maxWidth = 71
	maxHeight = 71
	startBytes = 1024

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			points[c] = &point{}
		}
	}

	fallPoints := []Coord{}

	for i, line := range lines {
		s := strings.Split(line, ",")
		x := S2i(s[0])
		y := S2i(s[1])
		c := Coord{x, y}
		points[c] = &point{
			c: c,
		}
		if i < startBytes {
			points[c].isWall = true
		}
		fallPoints = append(fallPoints, c)
	}
	points[Coord{X: maxWidth - 1, Y: maxHeight - 1}].isEnd = true
	var start *point
	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			curPoint := points[c]
			if curPoint.isWall {
				continue
			}
			curPoint.possiblePoints = []*point{}
			if c == (Coord{0, 0}) {
				start = curPoint
			}

			//north
			if i > 0 {
				nc := Coord{j, i - 1}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//south
			if i < maxHeight-1 {
				nc := Coord{j, i + 1}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//west
			if j > 0 {
				nc := Coord{j - 1, i}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
			//east
			if j < maxWidth-1 {
				nc := Coord{j + 1, i}
				if !points[nc].isWall {
					curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
				}
			}
		}
	}

	for i := startBytes; i < len(lines); i++ {
		points[fallPoints[i]].isWall = true
		clearExploration(points, maxHeight, maxWidth)
		found := false
		q := queue{}
		start.explored = true
		start.distance = 0
		q.push(*start)

		for len(q) != 0 {
			curP := q.pop()
			if curP.isEnd {
				found = true
				break
			}
			for _, p := range curP.possiblePoints {
				if !p.explored && !p.isWall {
					p.explored = true
					p.distance = curP.distance + 1
					q.push(*p)
				}
			}
		}
		if !found {
			answer = fallPoints[i].X * 100
			answer += fallPoints[i].Y
			break
		}
	}

	return answer
}

func clearExploration(points map[Coord]*point, h int, w int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := Coord{j, i}
			points[c].explored = false
			points[c].distance = 0
		}
	}
}
