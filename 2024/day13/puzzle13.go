package main

import (
	"flag"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 13, part1, part2, 480, 0)

}

type clawMachine struct {
	aX int
	bX int
	aY int
	bY int
	pX int
	pY int
}

func part1(lines []string) int {
	answer := 0

	cms := []clawMachine{}

	cm := clawMachine{}
	i := 0
	for _, line := range lines {
		if line == "" {
			cms = append(cms, cm)
			cm = clawMachine{}
			i = 0
			continue
		}
		if i == 0 {
			f := strings.Fields(line)
			x := strings.Split(f[2], "+")
			cm.aX = S2i(x[1][:len(x[1])-1])
			y := strings.Split(f[3], "+")
			cm.aY = S2i(y[1])
		}
		if i == 1 {
			f := strings.Fields(line)
			x := strings.Split(f[2], "+")
			cm.bX = S2i(x[1][:len(x[1])-1])
			y := strings.Split(f[3], "+")
			cm.bY = S2i(y[1])
		}
		if i == 2 {
			f := strings.Fields(line)
			x := strings.Split(f[1], "=")
			cm.pX = S2i(x[1][:len(x[1])-1])
			y := strings.Split(f[2], "=")
			cm.pY = S2i(y[1])
		}
		i++
	}
	cms = append(cms, cm)

	for _, c := range cms {
		startX := int(c.pX / c.bX)

		for i := startX; i >= 0; i-- {
			for j := 0; j < 100; j++ {
				if c.pX == c.bX*i+c.aX*j && c.pY == c.bY*i+c.aY*j {
					answer += i + j*3
				}
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	cms := []clawMachine{}

	cm := clawMachine{}
	i := 0
	for _, line := range lines {
		if line == "" {
			cms = append(cms, cm)
			cm = clawMachine{}
			i = 0
			continue
		}
		if i == 0 {
			f := strings.Fields(line)
			x := strings.Split(f[2], "+")
			cm.aX = S2i(x[1][:len(x[1])-1])
			y := strings.Split(f[3], "+")
			cm.aY = S2i(y[1])
		}
		if i == 1 {
			f := strings.Fields(line)
			x := strings.Split(f[2], "+")
			cm.bX = S2i(x[1][:len(x[1])-1])
			y := strings.Split(f[3], "+")
			cm.bY = S2i(y[1])
		}
		if i == 2 {
			f := strings.Fields(line)
			x := strings.Split(f[1], "=")
			cm.pX = S2i(x[1][:len(x[1])-1]) + 10000000000000
			y := strings.Split(f[2], "=")
			cm.pY = S2i(y[1]) + 10000000000000
		}
		i++
	}
	cms = append(cms, cm)

	for _, claw := range cms {
		a := claw.aX
		b := claw.bX
		c := claw.pX
		d := claw.aY
		e := claw.bY
		f := claw.pY

		top := (a*f - d*c)
		bottom := (a*e - d*b)

		i := top / bottom
		if i*bottom != top || i < 0 {
			continue
		}

		j := (c - b*int(i)) / a
		if j*a != (c-b*int(i)) || j < 0 {
			continue
		}
		answer += int(i + 3*j)
	}

	return answer
}
