package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"sort"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day15/sample.txt")
		part1(sample, 10)
		part2(sample, 20)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day15/input.txt")
		part1(input, 2000000)
		part2(input, 4000000)
	}

}

type sensor struct {
	x        int
	y        int
	distance int
}
type r struct {
	start int
	end   int
}

func part1(lines []string, target int) {
	answer := 0

	re := regexp.MustCompile("-?[0-9]+")
	var sensors []sensor
	beacons := map[int]bool{}

	for _, line := range lines {
		var s sensor
		vals := re.FindAllString(line, 4)
		s.x = S2i(vals[0])
		s.y = S2i(vals[1])

		distX := int(math.Abs(float64(S2i(vals[2])) - float64(S2i(vals[0]))))
		distY := int(math.Abs(float64(S2i(vals[3])) - float64(S2i(vals[1]))))

		s.distance = distX + distY
		sensors = append(sensors, s)

		if S2i(vals[3]) == target {
			beacons[S2i(vals[2])] = true
		}
	}

	noBeacon := map[int]bool{}
	for _, s := range sensors {
		left := int(math.Abs(float64(target) - float64(s.y)))
		dist := s.distance - left
		if dist >= 0 {
			for i := s.x; i <= s.x+dist; i++ {
				if !beacons[i] {
					noBeacon[i] = true
				}
			}
			for i := s.x; i >= s.x-dist; i-- {
				if !beacons[i] {
					noBeacon[i] = true
				}
			}
		}
	}

	answer = len(noBeacon)
	fmt.Printf("Answer 1 : %d\n", answer)
}

func part2(lines []string, tuning int) {
	answer := 0
	re := regexp.MustCompile("-?[0-9]+")
	var sensors []sensor
	beacons := map[int]map[int]bool{}

	for _, line := range lines {
		var s sensor
		vals := re.FindAllString(line, 4)
		s.x = S2i(vals[0])
		s.y = S2i(vals[1])

		if _, ok := beacons[S2i(vals[2])]; !ok {
			beacons[S2i(vals[2])] = map[int]bool{}
		}
		beacons[S2i(vals[2])][S2i(vals[3])] = true

		distX := int(math.Abs(float64(S2i(vals[2])) - float64(S2i(vals[0]))))
		distY := int(math.Abs(float64(S2i(vals[3])) - float64(S2i(vals[1]))))

		s.distance = distX + distY

		sensors = append(sensors, s)
	}

	ranges := [][]r{}
	for j := 0; j <= tuning; j++ {
		rs := []r{}
		for _, s := range sensors {
			left := int(math.Abs(float64(j) - float64(s.y)))
			dist := s.distance - left
			if dist >= 0 {
				max := int(math.Min(float64(s.x+dist), float64(tuning)))
				min := int(math.Max(float64(s.x-dist), 0))
				var ra r
				ra.end = max
				ra.start = min
				rs = append(rs, ra)
			}
		}
		ranges = append(ranges, rs)
	}

	y := 0
	x := 0
	for a, rs := range ranges {
		if x != 0 {
			break
		}
		sort.Slice(rs, func(i, j int) bool {
			return rs[i].start < rs[j].start
		})

		end := 0
		for i := 0; i < len(rs); i++ {
			if rs[i].start > end+1 {
				y = a
				x = rs[i].start - 1
				break
			}
			if rs[i].end > end {
				end = rs[i].end
			}
		}
	}

	answer = x*4000000 + y

	fmt.Printf("Answer 2 : %d\n", answer)
}
