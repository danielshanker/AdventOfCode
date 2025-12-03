package main

import (
	"flag"
	"math"
	"sync"
	"sync/atomic"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 3, part1, part2, 357, 3121910778619)

}

func part1(lines []string) int {
	answer := 0

	var atomicAnswer atomic.Int64
	wg := sync.WaitGroup{}
	for _, line := range lines {
		left := 0
		right := 0
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for i, r := range s {
				cur := S2i(string(r))
				if i < len(s)-1 {
					if cur > left {
						left = cur
						right = 0
						continue
					}
				}
				if cur > right {
					right = cur
				}
			}
			atomicAnswer.Add(int64(left*10 + right))
		}(line)
	}

	wg.Wait()
	answer = int(atomicAnswer.Load())
	return answer
}

func part2(lines []string) int {
	answer := 0
	var atomicAnswer atomic.Int64

	wg := sync.WaitGroup{}
	for _, line := range lines {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			vals := make([]int, 12)
			for i, r := range s {
				cur := S2i(string(r))
				for j := 0; j < 12; j++ {
					if checkVal(i, len(s), 11-j, cur, vals[j]) {
						vals[j] = cur
						for k := j + 1; k < 12; k++ {
							vals[k] = 0
						}
						break
					}
				}
			}
			add := 0
			for k := 0; k < 12; k++ {
				add += vals[k] * int(math.Pow10(11-k))
			}
			atomicAnswer.Add(int64(add))
		}(line)
	}
	wg.Wait()
	answer = int(atomicAnswer.Load())

	return answer
}

func checkVal(i int, lineLength int, pos int, cur int, original int) bool {
	if i < lineLength-pos {
		if cur > original {
			return true
		}
	}
	return false
}
