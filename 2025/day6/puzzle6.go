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

	Start(test, 6, part1, part2, 4277556, 3263827)

}

func part1(lines []string) int {
	answer := 0
	row := make([][]int, 4)
	op := strings.Fields(lines[4])

	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}
		vals := strings.Fields(line)
		for _, v := range vals {
			row[i] = append(row[i], S2i(v))
		}
	}
	for i := 0; i < len(op); i++ {
		if op[i] == "+" {
			ans := row[0][i] + row[1][i] + row[2][i] + row[3][i]
			answer += ans
		} else {
			ans := row[0][i] * row[1][i] * row[2][i] * row[3][i]
			answer += ans
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	op := strings.Fields(lines[4])
	cols := make([][]int, len(op))

	col := []int{}
	max := math.Max(float64(len(lines[0])), float64(len(lines[1])))
	max = math.Max(max, float64(len(lines[2])))
	max = math.Max(max, float64(len(lines[3])))
	colIndex := len(op) - 1
	for i := int(max) - 1; i >= 0; i-- {
		a, b, c, d := " ", " ", " ", " "
		if len(lines[0]) > i {
			a = string(lines[0][i])
		}
		if len(lines[1]) > i {
			b = string(lines[1][i])
		}
		if len(lines[2]) > i {
			c = string(lines[2][i])
		}
		if len(lines[3]) > i {
			d = string(lines[3][i])
		}
		if a == " " && b == " " && c == " " && d == " " {
			cols[colIndex] = col
			colIndex--
			col = []int{}
			continue
		}
		val := 0
		mult := 1

		if d != " " && d != "." {
			val += S2i(d) * mult
			mult *= 10
		}
		if c != " " && c != "." {
			val += S2i(c) * mult
			mult *= 10
		}
		if b != " " && b != "." {
			val += S2i(b) * mult
			mult *= 10
		}
		if a != " " && a != "." {
			val += S2i(a) * mult
			mult *= 10
		}
		col = append(col, val)
	}
	cols[0] = col

	for i, o := range op {
		if o == "+" {
			ans := 0
			for _, v := range cols[i] {
				ans += v
			}
			answer += ans
		} else {
			ans := 1
			for _, v := range cols[i] {
				ans *= v
			}
			answer += ans
		}
	}

	return answer
}
