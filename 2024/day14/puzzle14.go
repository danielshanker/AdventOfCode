package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 14, part1, part2, 12, 0)

}

type bot struct {
	startPos  coord
	curPos    coord
	xVelocity int
	yVelocity int
	id        int
}

type coord struct {
	x int
	y int
}

func part1(lines []string) int {
	answer := 0

	wh := strings.Fields(lines[0])
	width := S2i(wh[0])
	height := S2i(wh[1])

	bots := make([]bot, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			continue
		}

		pv := strings.Fields(line)
		p := pv[0][2:]
		pos := strings.Split(p, ",")
		start := coord{
			x: S2i(pos[0]),
			y: S2i(pos[1]),
		}
		v := pv[1][2:]
		vel := strings.Split(v, ",")

		b := bot{
			startPos:  start,
			curPos:    start,
			xVelocity: S2i(vel[0]),
			yVelocity: S2i(vel[1]),
			id:        i,
		}
		bots[i-1] = b
	}

	quads := make([]int, 4)
	c := map[coord]int{}

	for _, b := range bots {
		b.curPos.x = (b.startPos.x + b.xVelocity*100) % width
		b.curPos.y = (b.startPos.y + b.yVelocity*100) % height
		if b.curPos.x < 0 {
			b.curPos.x = width + b.curPos.x
		}
		if b.curPos.y < 0 {
			b.curPos.y = height + b.curPos.y
		}
		c[coord{x: b.curPos.x, y: b.curPos.y}]++

		if b.curPos.x <= width/2-1 && b.curPos.y <= height/2-1 {
			quads[0]++
		} else if b.curPos.x > width/2 && b.curPos.y <= height/2-1 {
			quads[1]++
		} else if b.curPos.x <= width/2-1 && b.curPos.y > height/2 {
			quads[2]++
		} else if b.curPos.x > width/2 && b.curPos.y > height/2 {
			quads[3]++
		}
	}

	answer = 1
	for _, q := range quads {
		answer *= q
	}

	return answer
}

func part2(lines []string) int {
	answer := 0

	wh := strings.Fields(lines[0])
	width := S2i(wh[0])
	height := S2i(wh[1])

	bots := make([]bot, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			continue
		}

		pv := strings.Fields(line)
		p := pv[0][2:]
		pos := strings.Split(p, ",")
		start := coord{
			x: S2i(pos[0]),
			y: S2i(pos[1]),
		}
		v := pv[1][2:]
		vel := strings.Split(v, ",")

		b := bot{
			startPos:  start,
			curPos:    start,
			xVelocity: S2i(vel[0]),
			yVelocity: S2i(vel[1]),
			id:        i,
		}
		bots[i-1] = b
	}

	lastDanger := math.MaxInt
	for a := 0; a < 10000; a++ {
		quads := make([]int, 4)
		c := map[coord]int{}
		for _, b := range bots {
			b.curPos.x = (b.startPos.x + b.xVelocity*a) % width
			b.curPos.y = (b.startPos.y + b.yVelocity*a) % height
			if b.curPos.x < 0 {
				b.curPos.x = width + b.curPos.x
			}
			if b.curPos.y < 0 {
				b.curPos.y = height + b.curPos.y
			}
			c[coord{x: b.curPos.x, y: b.curPos.y}]++

			if b.curPos.x <= width/2-1 && b.curPos.y <= height/2-1 {
				quads[0]++
			} else if b.curPos.x > width/2 && b.curPos.y <= height/2-1 {
				quads[1]++
			} else if b.curPos.x <= width/2-1 && b.curPos.y > height/2 {
				quads[2]++
			} else if b.curPos.x > width/2 && b.curPos.y > height/2 {
				quads[3]++
			}
		}
		danger := 1
		for _, q := range quads {
			danger *= q
		}

		if danger < lastDanger {
			lastDanger = danger
			fmt.Println(a)
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					if _, ok := c[coord{x: j, y: i}]; ok {
						fmt.Print(".")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}
	return answer
}
