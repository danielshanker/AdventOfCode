package main

import (
	"flag"
	"math"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 5, part1, part2, 3, 14)

}

type iRange struct {
	min  int
	max  int
	skip bool
}

func part1(lines []string) int {
	answer := 0

	fresh := make([]iRange, 0)
	firstChunk := true

	for _, line := range lines {
		if line == "" {
			firstChunk = false
			continue
		}
		if firstChunk {
			fRange := strings.Split(line, "-")
			ir := iRange{
				min: S2i(fRange[0]),
				max: S2i(fRange[1]),
			}
			fresh = append(fresh, ir)
		} else {
			ingredient := S2i(line)
			for _, f := range fresh {
				if ingredient >= f.min && ingredient <= f.max {
					answer++
					break
				}
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	fresh := make([]*iRange, 0)

	for _, line := range lines {
		if line == "" {
			break
		}
		fRange := strings.Split(line, "-")
		ir := &iRange{
			min: S2i(fRange[0]),
			max: S2i(fRange[1]),
		}
		fresh = append(fresh, ir)
	}

	newFresh := []*iRange{}

	for {
		hadOverlap := false
		for i, r1 := range fresh {
			overlapped := false
			for j, r2 := range fresh {
				if i == j || r2.skip || r1.skip {
					continue
				}
				// no overlap
				if (r1.min < r2.min && r1.max < r2.min) || (r1.max > r2.max && r1.min > r2.max) {
					continue
				}
				overlapped = true
				min := int(math.Min(float64(r1.min), float64(r2.min)))
				max := int(math.Max(float64(r1.max), float64(r2.max)))
				newFresh = append(newFresh, &iRange{min, max, false})
				r1.skip = true
				r2.skip = true
			}
			if !overlapped {
				newFresh = append(newFresh, r1)
			} else {
				hadOverlap = true
			}
		}
		if !hadOverlap {
			break
		}
		fresh = make([]*iRange, 0)
		for _, f := range newFresh {
			if f.skip {
				continue
			}
			nf := &iRange{
				min: f.min,
				max: f.max,
			}
			fresh = append(fresh, nf)
		}

		newFresh = make([]*iRange, 0)
	}

	for _, f := range fresh {
		rVal := f.max - f.min
		answer += rVal + 1
	}

	return answer
}
