package main

import (
	"flag"
	"fmt"
	"math"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 16, part1, part2, 11048, 64)

}

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

var needsTurnMap = map[int]map[int]bool{
	north: {
		north: false,
		south: false,
		east:  true,
		west:  true,
	},
	south: {
		north: false,
		south: false,
		east:  true,
		west:  true,
	},
	east: {
		north: true,
		south: true,
		east:  false,
		west:  false,
	},
	west: {
		north: true,
		south: true,
		east:  false,
		west:  false,
	},
}

type point struct {
	isWall         bool
	possiblePoints []*point
	isStart        bool
	isEnd          bool
	explored       map[int]bool
	distance       int
	lastDir        int
	c              Coord
	path           []Coord
	length         int
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

	points := [][]*point{}

	for _, line := range lines {
		if line == "" {
			break
		}
		chars := []*point{}
		for _, r := range line {
			p := point{}
			p.explored = make(map[int]bool)
			if string(r) == "S" {
				p.isStart = true
				p.lastDir = east
				chars = append(chars, &p)
				continue
			}
			if string(r) == "E" {
				p.isEnd = true
				chars = append(chars, &p)
				continue
			}
			if string(r) == "#" {
				p.isWall = true
				chars = append(chars, &p)
				continue
			}
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	var start *point

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			curPoint.possiblePoints = make([]*point, 4)
			if curPoint.isStart {
				start = curPoint
			} else {
				curPoint.lastDir = -1
			}
			//north
			if i > 0 {
				if !points[i-1][j].isWall {
					curPoint.possiblePoints[north] = points[i-1][j]
				}
			}
			//south
			if i < len(points)-1 {
				if !points[i+1][j].isWall {
					curPoint.possiblePoints[south] = points[i+1][j]
				}
			}
			//west
			if j > 0 {
				if !points[i][j-1].isWall {
					curPoint.possiblePoints[west] = points[i][j-1]
				}
			}
			//east
			if j < len(points[0])-1 {
				if !points[i][j+1].isWall {
					curPoint.possiblePoints[east] = points[i][j+1]
				}
			}
		}
	}

	q := queue{}
	start.explored[east] = true
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
		for d, p := range curP.possiblePoints {
			if p == nil {
				continue
			}
			if isOpposite(curP.lastDir, d) {
				continue
			}
			if !p.explored[curP.lastDir] || p.distance > curP.distance {
				p.explored[curP.lastDir] = true
				p.lastDir = d
				cost := 1
				if needsTurn(curP.lastDir, d) {
					cost = 1001
				}
				p.distance = curP.distance + cost
				q.push(*p)
			}
		}
	}
	return answer
}

func part2(lines []string) int {
	answer := math.MaxInt

	points := [][]*point{}

	for y, line := range lines {
		if line == "" {
			break
		}
		chars := []*point{}
		for x, r := range line {
			p := point{}
			p.explored = make(map[int]bool)
			p.c = Coord{x, y}
			if string(r) == "S" {
				p.isStart = true
				p.lastDir = east
				p.path = append(p.path, p.c)
				chars = append(chars, &p)
				continue
			}
			if string(r) == "E" {
				p.isEnd = true
				chars = append(chars, &p)
				continue
			}
			if string(r) == "#" {
				p.isWall = true
				chars = append(chars, &p)
				continue
			}
			chars = append(chars, &p)
		}
		points = append(points, chars)
	}

	var start *point

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			curPoint := points[i][j]
			curPoint.possiblePoints = make([]*point, 4)
			if curPoint.isStart {
				start = curPoint
			} else {
				curPoint.lastDir = -1
			}
			//north
			if i > 0 {
				if !points[i-1][j].isWall {
					curPoint.possiblePoints[north] = points[i-1][j]
				}
			}
			//south
			if i < len(points)-1 {
				if !points[i+1][j].isWall {
					curPoint.possiblePoints[south] = points[i+1][j]
				}
			}
			//west
			if j > 0 {
				if !points[i][j-1].isWall {
					curPoint.possiblePoints[west] = points[i][j-1]
				}
			}
			//east
			if j < len(points[0])-1 {
				if !points[i][j+1].isWall {
					curPoint.possiblePoints[east] = points[i][j+1]
				}
			}
		}
	}

	q := queue{}
	start.explored[east] = true
	start.distance = 0
	q.push(*start)
	reachedEnd := map[int]map[Coord]bool{}
	shortestPath := math.MaxInt
	for len(q) != 0 {
		curP := q.pop()
		if curP.isEnd {
			if _, ok := reachedEnd[curP.distance]; !ok {
				reachedEnd[curP.distance] = map[Coord]bool{}
			}
			for _, p := range curP.path {
				reachedEnd[curP.distance][p] = true
			}
			if curP.distance < shortestPath {
				shortestPath = curP.distance
			}
			continue
		}
		for d, p := range curP.possiblePoints {
			if p == nil || isOpposite(curP.lastDir, d) || p.distance > answer {
				continue
			}
			if !p.explored[curP.lastDir] || p.distance > curP.distance {
				p.explored[curP.lastDir] = true
				p.lastDir = d
				cost := 1
				if needsTurn(curP.lastDir, d) {
					cost = 1001
				}
				p.distance = curP.distance + cost
				p.length = curP.length + 1
				if len(p.path) != 0 {
					p.path = []Coord{}
				}
				for _, a := range curP.path {
					p.path = append(p.path, a)
				}
				p.path = append(p.path, p.c)
				q.push(*p)
			}
		}
	}

	bestPath := reachedEnd[shortestPath]
	answer = len(bestPath)

	//	visualize(points, bestPath)

	// Accidentally guessed the answer (645) while output is giving me 640, go back and figure out why!

	return answer
}

func needsTurn(curDir, nextDir int) bool {
	return needsTurnMap[curDir][nextDir]
}

func isOpposite(curDir, nextDir int) bool {
	opp := map[int]map[int]bool{
		north: {
			south: true,
		},
		south: {
			north: true,
		},
		east: {
			west: true,
		},
		west: {
			east: true,
		},
	}
	return opp[curDir][nextDir]
}

func visualize(points [][]*point, bestPath map[Coord]bool) {
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[0]); j++ {
			if points[i][j].isWall {
				fmt.Print("#")
				continue
			}
			c := Coord{j, i}
			if !bestPath[c] {
				fmt.Print(".")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

}
