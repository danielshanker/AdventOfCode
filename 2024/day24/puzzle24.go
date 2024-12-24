package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 24, part1, part2, 2024, 0)

}

type gate struct {
	output     int
	input1     string
	input2     string
	operator   string
	outputWire string
}

func part1(lines []string) int {
	answer := 0

	gates := map[string]gate{}

	next := false
	zCount := 0
	for _, line := range lines {
		if line == "" {
			next = true
			continue
		}
		if next {
			x := strings.Fields(line)
			w := gate{
				input1:     x[0],
				input2:     x[2],
				output:     -1,
				operator:   x[1],
				outputWire: x[4],
			}
			gates[x[4]] = w
			if string(x[4][0]) == "z" {
				zCount++
			}
		} else {
			x := strings.Split(line, ": ")
			w := gate{
				output: S2i(x[1]),
			}
			gates[x[0]] = w
		}
	}

	count := zCount
	for count > 0 {
		for n, g := range gates {
			if g.output != -1 {
				continue
			}
			g1 := gates[g.input1]
			g2 := gates[g.input2]
			if g1.output != -1 && g2.output != -1 {
				g.output = performOp(g, g1, g2)
				if string(n[0]) == "z" {
					count--
				}
				gates[n] = g
			}
		}
	}

	b := ""
	for i := zCount - 1; i >= 0; i-- {
		key := "z"
		key += fmt.Sprintf("%02d", i)
		b += strconv.Itoa(gates[key].output)
	}

	x, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	answer = int(x)

	return answer
}

func performOp(g gate, g1 gate, g2 gate) int {
	if g.operator == "OR" {
		return g1.output | g2.output
	} else if g.operator == "AND" {
		return g1.output & g2.output
	} else if g.operator == "XOR" {
		return g1.output ^ g2.output
	} else {
		return -1
	}
}

func part2(lines []string) int {
	answerString := ""

	gates := map[string]gate{}

	next := false
	zCount := 0
	xCount := 0
	for _, line := range lines {
		if line == "" {
			next = true
			continue
		}
		if next {
			x := strings.Fields(line)
			w := gate{
				input1:     x[0],
				input2:     x[2],
				output:     -1,
				operator:   x[1],
				outputWire: x[4],
			}
			gates[x[4]] = w
			if string(x[4][0]) == "z" {
				zCount++
			}
		} else {
			xCount++
			x := strings.Split(line, ": ")
			w := gate{
				output: S2i(x[1]),
			}
			gates[x[0]] = w
		}
	}
	xCount = xCount / 2
	fmt.Println(xCount)

	cLast := ""
	swapped := []string{}

	// looked at a full adder full adder should have xn yn pointing to both a XOR and an AND (a and b)
	// each z should have an XOR pointing to it from a and the last carry over (c1)
	// c0 is a carry over to the next group which should connect to the XOR going to the z
	// r is the AND gate that should take the carry over from the last half adder and AND it with the result of the previous XOR
	for i := 0; i < xCount; i++ {
		n := fmt.Sprintf("%02d", i)
		var a, b, r, z, c string

		a = getGateName("x"+n, "y"+n, "XOR", gates)
		b = getGateName("x"+n, "y"+n, "AND", gates)

		if cLast != "" {
			r = getGateName(cLast, a, "AND", gates)
			if r == "" {
				a, b = b, a
				swapped = append(swapped, a, b)
				r = getGateName(cLast, a, "AND", gates)
			}
			z = getGateName(cLast, a, "XOR", gates)
			if a != "" && string(a[0]) == "z" {
				a, z = z, a
				swapped = append(swapped, a, z)
			}
			if b != "" && string(b[0]) == "z" {
				b, z = z, b
				swapped = append(swapped, b, z)
			}
			if r != "" && string(r[0]) == "z" {
				r, z = z, r
				swapped = append(swapped, r, z)
			}
			c = getGateName(r, b, "OR", gates)
		}

		if c != "" && string(c[0]) == "z" && c != fmt.Sprintf("z%02d", xCount) {
			c, z = z, c
			swapped = append(swapped, c, z)
		}

		if cLast == "" {
			cLast = b
		} else {
			cLast = c
		}
	}

	sort.Strings(swapped)
	answerString = strings.Join(swapped, ",")

	fmt.Println(answerString)

	return 0
}

func getGateName(n1 string, n2 string, op string, gates map[string]gate) string {
	for name, gate := range gates {
		if (gate.input1 == n1 && gate.input2 == n2) || (gate.input2 == n1 && gate.input1 == n2) {
			if gate.operator == op {
				return name
			}
		}
	}
	return ""
}
