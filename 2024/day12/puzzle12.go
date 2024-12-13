package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 12, part1, part2, 1930, 1206)

}

type plot struct {
	x          int
	y          int
	plant      string
	neighbours []*plot
	sides      int
	checked    bool
	lW         bool
	dW         bool
	rW         bool
	uW         bool
	colour     int
}

func part1(lines []string) int {
	answer := 0

	plots := make([][]*plot, len(lines))

	for y, line := range lines {
		lp := make([]*plot, len(line))
		for x, r := range line {
			p := &plot{
				x:     x,
				y:     y,
				plant: string(r),
			}
			lp[x] = p
		}
		plots[y] = lp
	}

	for y, line := range plots {
		for x, p := range line {
			// up
			if y > 0 && plots[y-1][x].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y-1][x])
			}
			// down
			if y < len(plots)-1 && plots[y+1][x].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y+1][x])
			}
			// left
			if x > 0 && plots[y][x-1].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y][x-1])
			}
			// right
			if x < len(line)-1 && plots[y][x+1].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y][x+1])
			}
			p.sides = 4 - len(p.neighbours)

		}
	}

	regions := [][]*plot{}

	colour := 41
	for _, line := range plots {
		for _, p := range line {
			if colour > 107 {
				colour = 41
			} else if colour > 47 && colour < 100 {
				colour = 100
			}
			if !p.checked {
				region := []*plot{}
				region = fillRegion(region, p, colour)
				regions = append(regions, region)
				colour++
			}
		}
	}

	//visualize(plots)

	for _, region := range regions {
		for _, plot := range region {
			answer += plot.sides * len(region)
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	plots := make([][]*plot, len(lines))

	for y, line := range lines {
		lp := make([]*plot, len(line))
		for x, r := range line {
			p := &plot{
				x:     x,
				y:     y,
				plant: string(r),
			}
			lp[x] = p
		}
		plots[y] = lp
	}

	for y, line := range plots {
		for x, p := range line {
			// up
			if y > 0 && plots[y-1][x].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y-1][x])
			} else {
				p.uW = true
			}
			// down
			if y < len(plots)-1 && plots[y+1][x].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y+1][x])
			} else {
				p.dW = true
			}
			// left
			if x > 0 && plots[y][x-1].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y][x-1])
			} else {
				p.lW = true
			}
			// right
			if x < len(line)-1 && plots[y][x+1].plant == p.plant {
				p.neighbours = append(p.neighbours, plots[y][x+1])
			} else {
				p.rW = true
			}

			p.sides = 4 - len(p.neighbours)

		}
	}

	regions := [][]*plot{}

	colour := 31
	for _, line := range plots {
		for _, p := range line {
			if colour >= 37 {
				colour = 31
			}
			if !p.checked {
				region := []*plot{}
				region = fillRegion(region, p, colour)
				regions = append(regions, region)
				colour++
			}
		}
	}

	for _, region := range regions {
		fences := walk(region)
		answer += fences * len(region)
	}

	return answer
}

type edge struct {
	dir string
	x   int
	y   int
}

func walk(region []*plot) int {
	edges := []edge{}
	for _, p := range region {
		if p.lW {
			e := edge{
				x:   p.x,
				y:   p.y,
				dir: "l",
			}
			edges = append(edges, e)
		}
		if p.rW {
			e := edge{
				x:   p.x,
				y:   p.y,
				dir: "r",
			}
			edges = append(edges, e)
		}
		if p.uW {
			e := edge{
				x:   p.x,
				y:   p.y,
				dir: "u",
			}
			edges = append(edges, e)
		}
		if p.dW {
			e := edge{
				x:   p.x,
				y:   p.y,
				dir: "d",
			}
			edges = append(edges, e)
		}
	}

	edgeMap := map[edge]struct{}{}
	for _, e := range edges {
		edgeMap[e] = struct{}{}
	}

	foundEdge := map[edge]struct{}{}

	fences := 0
	for _, e := range edges {
		if _, ok := foundEdge[e]; ok {
			continue
		}
		if e.dir == "u" || e.dir == "d" {
			i := 0
			for {
				check := edge{x: e.x - i, y: e.y, dir: e.dir}
				if _, ok := edgeMap[check]; ok {
					foundEdge[check] = struct{}{}
				} else {
					break
				}
				i++
			}
			i = 0
			for {
				check := edge{x: e.x + i, y: e.y, dir: e.dir}
				if _, ok := edgeMap[check]; ok {
					foundEdge[check] = struct{}{}
				} else {
					break
				}
				i++
			}
		} else {
			i := 0
			for {
				check := edge{x: e.x, y: e.y - i, dir: e.dir}
				if _, ok := edgeMap[check]; ok {
					foundEdge[check] = struct{}{}
				} else {
					break
				}
				i++
			}
			i = 0
			for {
				check := edge{x: e.x, y: e.y + i, dir: e.dir}
				if _, ok := edgeMap[check]; ok {
					foundEdge[check] = struct{}{}
				} else {
					break
				}
				i++
			}
		}
		fences++
	}

	return fences
}

func fillRegion(region []*plot, p *plot, colour int) []*plot {
	p.checked = true
	p.colour = colour
	region = append(region, p)

	for _, n := range p.neighbours {
		if !n.checked {
			region = fillRegion(region, n, colour)
		}
	}

	return region
}

func visualize(plots [][]*plot) {
	for _, line := range plots {
		for _, p := range line {
			fmt.Printf("\033[%dm ", p.colour)
		}
		fmt.Printf("\033[0m")
		fmt.Println()
	}
}
