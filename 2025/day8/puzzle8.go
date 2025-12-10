package main

import (
	"flag"
	"math"
	"sort"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 8, part1, part2, 40, 25272)

}

type coord3 struct {
	x int
	y int
	z int
}

func (c coord3) distance(dest coord3) float64 {
	x2 := math.Pow(float64(c.x)-float64(dest.x), 2)
	y2 := math.Pow(float64(c.y)-float64(dest.y), 2)
	z2 := math.Pow(float64(c.z)-float64(dest.z), 2)

	return math.Sqrt(x2 + y2 + z2)

}

type box struct {
	visited        bool
	coord          coord3
	connectedBoxes []*box
}
type edge struct {
	a    *box
	b    *box
	dist float64
}

func traverse(b *box, count int) int {
	if b.visited {
		return count
	}
	b.visited = true
	for _, c := range b.connectedBoxes {
		count = traverse(c, count)
	}
	return count + 1
}

func part1(lines []string) int {
	answer := 0
	boxes := make([]*box, 0, len(lines))
	for _, line := range lines {
		c := strings.Split(line, ",")
		box := box{
			coord: coord3{
				x: S2i(c[0]),
				y: S2i(c[1]),
				z: S2i(c[2]),
			},
		}
		boxes = append(boxes, &box)
	}

	edges := []*edge{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			b1 := boxes[i]
			b2 := boxes[j]
			if i == j {
				continue
			}
			dist := b1.coord.distance(b2.coord)
			edges = append(edges, &edge{b1, b2, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	maxCons := 1000
	if *test {
		maxCons = 10
	}

	for i, e := range edges {
		if i >= maxCons {
			break
		}
		e.a.connectedBoxes = append(e.a.connectedBoxes, e.b)
		e.b.connectedBoxes = append(e.b.connectedBoxes, e.a)
	}

	counts := []int{}
	for _, b := range boxes {
		if b.visited {
			continue
		}
		counts = append(counts, traverse(b, 0))
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	answer = counts[0] * counts[1] * counts[2]

	return answer
}

func reset(boxes []*box) {
	for _, b := range boxes {
		b.visited = false
	}
}

func part2(lines []string) int {
	answer := 0
	boxes := make([]*box, 0, len(lines))
	for _, line := range lines {
		c := strings.Split(line, ",")
		box := box{
			coord: coord3{
				x: S2i(c[0]),
				y: S2i(c[1]),
				z: S2i(c[2]),
			},
		}
		boxes = append(boxes, &box)
	}

	edges := []*edge{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			b1 := boxes[i]
			b2 := boxes[j]
			if i == j {
				continue
			}
			dist := b1.coord.distance(b2.coord)
			edges = append(edges, &edge{b1, b2, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	for _, e := range edges {
		e.a.connectedBoxes = append(e.a.connectedBoxes, e.b)
		e.b.connectedBoxes = append(e.b.connectedBoxes, e.a)
		count := traverse(e.a, 0)
		if count == len(boxes) {
			answer = e.a.coord.x * e.b.coord.x
			break
		}
		reset(boxes)
	}

	return answer
}
