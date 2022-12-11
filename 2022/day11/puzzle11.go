package main

import (
	"flag"
	"fmt"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day11/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day11/input.txt")
		part1(input)
		part2(input)
	}

}

type monkey struct {
	items     []int64
	operation string
	divisible int64
	trueMonk  int64
	falseMonk int64
	inspected int64
}

func part1(lines []string) {
	answer1 := int64(0)

	var monkeys []monkey

	curMonkey := monkey{}
	monkeyLine := 0

	for i, line := range lines {
		monkeyLine = i % 7
		if line == "" {
			monkeys = append(monkeys, curMonkey)
			curMonkey = monkey{}
			continue
		}
		if monkeyLine == 0 {
			continue
		}
		if monkeyLine == 1 {
			a := strings.Split(line, ":")
			b := strings.ReplaceAll(a[1], " ", "")
			c := strings.Split(b, ",")
			for _, num := range c {
				curMonkey.items = append(curMonkey.items, int64(S2i(num)))
			}
		}

		if monkeyLine == 2 {
			a := strings.Split(line, "new = ")
			curMonkey.operation = a[1]
		}

		if monkeyLine == 3 {
			a := strings.Split(line, " by ")
			curMonkey.divisible = int64(S2i(a[1]))
		}
		if monkeyLine == 4 {
			a := strings.Split(line, " monkey ")
			curMonkey.trueMonk = int64(S2i(a[1]))
		}
		if monkeyLine == 5 {
			a := strings.Split(line, " monkey ")
			curMonkey.falseMonk = int64(S2i(a[1]))
		}
	}
	monkeys = append(monkeys, curMonkey)

	for i := 0; i < 20; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeys[j].inspected++
				operation := strings.Split(monkey.operation, " ")
				if operation[1] == "+" {
					if operation[2] == "old" {
						item += item
					} else {
						item += int64(S2i(operation[2]))
					}
				} else if operation[1] == "*" {
					if operation[2] == "old" {
						item *= item
					} else {
						item *= int64(S2i(operation[2]))
					}
				}

				item /= 3
				if item%monkey.divisible == 0 {
					monkeys[monkey.trueMonk].items = append(monkeys[monkey.trueMonk].items, item)
				} else {
					monkeys[monkey.falseMonk].items = append(monkeys[monkey.falseMonk].items, item)
				}
			}
			monkeys[j].items = []int64{}
		}
	}

	inspected := []int64{}
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}

	max := int64(0)
	max2 := int64(0)

	for _, i := range inspected {
		if i > max {
			max2 = max
			max = i
			continue
		}
		if i > max2 {
			max2 = i
		}
	}
	answer1 = max * max2

	fmt.Printf("Answer 1 : %d\n", answer1)
}

func part2(lines []string) {
	answer2 := int64(0)
	var monkeys []monkey

	curMonkey := monkey{}
	monkeyLine := 0

	for i, line := range lines {
		monkeyLine = i % 7
		if line == "" {
			monkeys = append(monkeys, curMonkey)
			curMonkey = monkey{}
			continue
		}
		if monkeyLine == 0 {
			continue
		}
		if monkeyLine == 1 {
			a := strings.Split(line, ":")
			b := strings.ReplaceAll(a[1], " ", "")
			c := strings.Split(b, ",")
			for _, num := range c {
				curMonkey.items = append(curMonkey.items, int64(S2i(num)))
			}
		}

		if monkeyLine == 2 {
			a := strings.Split(line, "new = ")
			curMonkey.operation = a[1]
		}

		if monkeyLine == 3 {
			a := strings.Split(line, " by ")
			curMonkey.divisible = int64(S2i(a[1]))
		}
		if monkeyLine == 4 {
			a := strings.Split(line, " monkey ")
			curMonkey.trueMonk = int64(S2i(a[1]))
		}
		if monkeyLine == 5 {
			a := strings.Split(line, " monkey ")
			curMonkey.falseMonk = int64(S2i(a[1]))
		}
	}
	monkeys = append(monkeys, curMonkey)
	lcm := LCM(monkeys)
	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeys[j].inspected++
				operation := strings.Split(monkey.operation, " ")
				if operation[1] == "+" {
					if operation[2] == "old" {
						item += item
					} else {
						item += int64(S2i(operation[2]))
					}
				} else if operation[1] == "*" {
					if operation[2] == "old" {
						item *= item
					} else {
						item *= int64(S2i(operation[2]))
					}
				}
				item := item % lcm
				if item%monkey.divisible == 0 {
					monkeys[monkey.trueMonk].items = append(monkeys[monkey.trueMonk].items, item)
				} else {
					monkeys[monkey.falseMonk].items = append(monkeys[monkey.falseMonk].items, item)
				}
			}
			monkeys[j].items = []int64{}
		}
	}

	inspected := []int64{}
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}

	max := int64(0)
	max2 := int64(0)

	for _, i := range inspected {
		if i > max {
			max2 = max
			max = i
			continue
		}
		if i > max2 {
			max2 = i
		}
	}

	answer2 = max * max2

	fmt.Printf("Answer 2 : %d\n", answer2)
}

func LCM(monkeys []monkey) int64 {
	lcm := int64(1)

	for _, m := range monkeys {
		lcm *= m.divisible
	}
	return lcm
}
