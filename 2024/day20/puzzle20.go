package main

import (
	"flag"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 20, part1, part2, 44, 285)

}

type point struct {
	possiblePoints []*point
	isStart        bool
	isEnd          bool
	explored       bool
	distance       int
	ogDistance     int
	ogDistanceLeft int
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
	answer := 0

	maxHeight := len(lines)
	maxWidth := len(lines[0])
	points := map[Coord]*point{}
	var start *point
	walls := []Coord{}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			points[c] = &point{}
		}
	}

	for y, line := range lines {
		if line == "" {
			break
		}
		for x, r := range line {
			c := Coord{x, y}
			p := point{
				c: c,
			}
			if string(r) == "S" {
				p.isStart = true
				start = &p
			}
			if string(r) == "E" {
				p.isEnd = true
			}
			if string(r) == "#" {
				p.isWall = true
				walls = append(walls, c)
			}
			points[c] = &p
		}
	}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			curPoint := points[c]
			curPoint.possiblePoints = []*point{}

			//north
			if i > 0 {
				nc := Coord{j, i - 1}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//south
			if i < maxHeight-1 {
				nc := Coord{j, i + 1}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//west
			if j > 0 {
				nc := Coord{j - 1, i}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//east
			if j < maxWidth-1 {
				nc := Coord{j + 1, i}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
		}
	}
	q := queue{}
	start.explored = true
	start.distance = 0
	q.push(*start)

	ogLength := 0

	for len(q) != 0 {
		curP := q.pop()
		if curP.isEnd {
			ogLength = curP.distance
		}
		for _, p := range curP.possiblePoints {
			if !p.explored {
				p.explored = true
				p.distance = curP.distance + 1
				p.ogDistance = curP.ogDistance + 1
				if !p.isWall {
					q.push(*p)
				}
			}
		}
	}

	for _, p := range points {
		if !p.isWall {
			p.ogDistanceLeft = ogLength - p.distance
		}

	}

	saves := map[int]int{}
	cheats := [][]Coord{}
	for c, p := range points {
		for i := 0; i < 4; i++ {
			foundCheat := false
			cheat := make([]Coord, 2)
			if p.isWall {
				foundCheat = true
				cheat[0] = c
			}
			nc := Coord{}
			if i == 0 {
				nc = Coord{c.X, c.Y - 1}
			}
			//south
			if i == 1 {
				nc = Coord{c.X, c.Y + 1}
			}
			//west
			if i == 2 {
				nc = Coord{c.X - 1, c.Y}
			}
			//east
			if i == 3 {
				nc = Coord{c.X + 1, c.Y}
			}
			if _, ok := points[nc]; ok {
				if foundCheat && !points[nc].isWall {
					foundCheat = true
					cheat[1] = nc
				} else {
					foundCheat = false
				}
			} else {
				continue
			}

			if foundCheat {
				cheats = append(cheats, cheat)
			}
		}
	}

	for _, cheat := range cheats {
		d := points[cheat[0]].ogDistance
		nd := points[cheat[1]].ogDistanceLeft
		save := ogLength - (d + nd + 1)
		if save > 0 {
			saves[save]++
		}
	}

	for k, s := range saves {
		if k >= 100 {
			answer += s
		}
	}
	return answer
}

type cheating struct {
	start    Coord
	end      Coord
	distance int
}

func part2(lines []string) int {
	answer := 0

	maxHeight := len(lines)
	maxWidth := len(lines[0])
	points := map[Coord]*point{}
	var start *point
	walls := []Coord{}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			points[c] = &point{}
		}
	}

	for y, line := range lines {
		if line == "" {
			break
		}
		for x, r := range line {
			c := Coord{x, y}
			p := point{
				c: c,
			}
			if string(r) == "S" {
				p.isStart = true
				start = &p
			}
			if string(r) == "E" {
				p.isEnd = true
			}
			if string(r) == "#" {
				p.isWall = true
				walls = append(walls, c)
			}
			points[c] = &p
		}
	}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			c := Coord{j, i}
			curPoint := points[c]
			curPoint.possiblePoints = []*point{}

			//north
			if i > 0 {
				nc := Coord{j, i - 1}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//south
			if i < maxHeight-1 {
				nc := Coord{j, i + 1}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//west
			if j > 0 {
				nc := Coord{j - 1, i}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
			//east
			if j < maxWidth-1 {
				nc := Coord{j + 1, i}
				curPoint.possiblePoints = append(curPoint.possiblePoints, points[nc])
			}
		}
	}
	q := queue{}
	start.explored = true
	start.distance = 0
	q.push(*start)

	ogLength := 0

	for len(q) != 0 {
		curP := q.pop()
		if curP.isEnd {
			ogLength = curP.distance
			break
		}
		for _, p := range curP.possiblePoints {
			if !p.explored && !p.isWall {
				p.explored = true
				p.distance = curP.distance + 1
				p.ogDistance = curP.ogDistance + 1
				q.push(*p)
			}
		}
	}

	for _, p := range points {
		if !p.isWall {
			p.ogDistanceLeft = ogLength - p.distance
		}
	}

	saves := map[int]int{}
	cheats := map[cheating]bool{}
	for c, p := range points {
		if !p.isWall {
			for endC, endP := range points {
				if endP.isWall {
					continue
				}
				cheat := cheating{}
				d := AbsDistance(endC, c)
				if d <= 20 {
					cheat.start = c
					cheat.end = endC
					cheat.distance = d
					cheats[cheat] = true
				}
			}
		}
	}

	for cheat := range cheats {
		d := points[cheat.start].ogDistance
		nd := points[cheat.end].ogDistanceLeft
		save := ogLength - (d + nd + cheat.distance)
		if save >= 100 {
			saves[save]++
		}
	}

	for k, s := range saves {
		if k >= 100 {
			answer += s
		}
	}

	return answer
}

func checkNeighbours(endC Coord, startC Coord, points map[Coord]*point, d int) bool {
	return true
	if d == 0 {
		return true
	}

	// north
	nc := Coord{
		X: endC.X,
		Y: endC.Y - 1,
	}
	x := AbsDistance(startC, nc)
	if points[nc] != nil && points[nc].isWall && x == d-1 {
		if checkNeighbours(nc, startC, points, d-1) {
			return true
		}
	}
	// south
	nc = Coord{
		X: endC.X,
		Y: endC.Y + 1,
	}
	if points[nc] != nil && points[nc].isWall && AbsDistance(startC, nc) == d-1 {
		if checkNeighbours(nc, startC, points, d-1) {
			return true
		}
	}
	// west
	nc = Coord{
		X: endC.X - 1,
		Y: endC.Y,
	}
	if points[nc] != nil && points[nc].isWall && AbsDistance(startC, nc) == d-1 {
		if checkNeighbours(nc, startC, points, d-1) {
			return true
		}
	}
	// east
	nc = Coord{
		X: endC.X + 1,
		Y: endC.Y,
	}
	if points[nc] != nil && points[nc].isWall && AbsDistance(startC, nc) == d-1 {
		if checkNeighbours(nc, startC, points, d-1) {
			return true
		}
	}

	return false
}
