package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	sample := ReadInputLines(fmt.Sprintf("/home/daniel.shanker/Pers/AdventOfCode/2024/day%d/sample.txt", 17))
	input := ReadInputLines(fmt.Sprintf("/home/daniel.shanker/Pers/AdventOfCode/2024/day%d/input.txt", 17))

	if *test {
		part1(sample)
		part2(sample)
	} else {
		st := time.Now()
		part1(input)
		totalTime := time.Since(st)
		if totalTime >= time.Second {
			fmt.Printf("%ds\n", int(totalTime.Seconds()))
		} else if totalTime >= time.Millisecond {
			fmt.Printf("%dms\n", totalTime.Milliseconds())
		} else {
			fmt.Printf("%dμs\n", totalTime.Microseconds())
		}
		st = time.Now()
		part2(input)
		totalTime = time.Since(st)
		if totalTime >= time.Second {
			fmt.Printf("%ds\n", int(totalTime.Seconds()))
		} else if totalTime >= time.Millisecond {
			fmt.Printf("%dms\n", totalTime.Milliseconds())
		} else {
			fmt.Printf("%dμs\n", totalTime.Microseconds())
		}
	}

}

type register struct {
	a      int
	b      int
	c      int
	output string
}

func part1(lines []string) int {

	reg := register{}
	reg.a = S2i(strings.Split(lines[0], ": ")[1])
	reg.b = S2i(strings.Split(lines[1], ": ")[1])
	reg.c = S2i(strings.Split(lines[2], ": ")[1])
	p := strings.Fields(lines[4])[1]
	program := []int{}

	for _, inst := range strings.Split(p, ",") {
		program = append(program, S2i(inst))
	}
	reg = run(reg, program)
	fmt.Println(reg.output)

	return 0
}

func part2(lines []string) int {

	reg := register{}
	reg.a = S2i(strings.Split(lines[0], ": ")[1])
	reg.b = S2i(strings.Split(lines[1], ": ")[1])
	reg.c = S2i(strings.Split(lines[2], ": ")[1])
	p := strings.Fields(lines[4])[1]
	program := []int{}

	for _, inst := range strings.Split(p, ",") {
		program = append(program, S2i(inst))
	}
	expected := "2,4,1,1,7,5,0,3,1,4,4,4,5,5,3,0"

	pows := make([]int, 16)
	e := strings.Split(expected, ",")
	search(pows, 15, e, program)
	return 0
}

func search(pows []int, index int, e []string, program []int) {
	if index < 0 {
		return
	}
	for i := 0; i < 8; i++ {
		answer := 0
		for j := 0; j < 16; j++ {
			if j == index {
				answer += int(math.Pow(8, float64(j))) * i
			} else {
				answer += int(math.Pow(8, float64(j))) * pows[j]
			}
		}
		reg := register{
			a: answer,
		}
		reg = run(reg, program)
		sp := strings.Split(reg.output, ",")
		if len(sp) > index && sp[index] == e[index] {
			pows[index] = i
			if index == 0 {
				fmt.Println(answer)
				return
			}
			search(pows, index-1, e, program)
		}
	}
}

func run(reg register, program []int) register {
	reg.output = ""
	i := 0
	for i < len(program) {
		inst := program[i]
		switch inst {
		case 0:
			reg = adv(reg, program[i+1])
			i += 2
		case 1:
			reg = bxl(reg, program[i+1])
			i += 2
		case 2:
			reg = bst(reg, program[i+1])
			i += 2
		case 3:
			a := jnz(reg, program[i+1])
			if a >= 0 {
				i = a
			} else {
				i += 2
			}

		case 4:
			reg = bxc(reg)
			i += 2
		case 5:
			reg = outCom(reg, program[i+1])
			i += 2
		case 6:
			reg = bdv(reg, program[i+1])
			i += 2
		case 7:
			reg = cdv(reg, program[i+1])
			i += 2
		}
	}
	return reg
}

func adv(reg register, val int) register {
	a := reg.a
	v := getCombo(reg, val)
	out := a / (int(math.Pow(2, float64(v))))
	reg.a = out
	return reg
}
func bxl(reg register, val int) register {
	b := reg.b
	out := b ^ val
	reg.b = out
	return reg
}
func bst(reg register, val int) register {
	v := getCombo(reg, val)
	out := v % 8
	reg.b = out
	return reg
}
func jnz(reg register, val int) int {
	if reg.a == 0 {
		return -1
	} else {
		return val
	}
}
func bxc(reg register) register {
	b := reg.b
	c := reg.c
	out := b ^ c
	reg.b = out
	return reg
}

func outCom(reg register, val int) register {
	v := getCombo(reg, val)
	out := v % 8
	if reg.output == "" {
		reg.output = strconv.Itoa(out)
	} else {
		reg.output += "," + strconv.Itoa(out)
	}
	return reg
}
func bdv(reg register, val int) register {
	a := reg.a
	v := getCombo(reg, val)
	out := a / (int(math.Pow(2, float64(v))))
	reg.b = out
	return reg
}
func cdv(reg register, val int) register {
	a := reg.a
	v := getCombo(reg, val)
	out := a / (int(math.Pow(2, float64(v))))
	reg.c = out
	return reg
}

func getCombo(reg register, val int) int {
	if val <= 3 && val >= 0 {
		return val
	}
	if val == 4 {
		return reg.a
	}
	if val == 5 {
		return reg.b
	}
	if val == 6 {
		return reg.c
	}
	fmt.Println("invalid")
	return 0
}
