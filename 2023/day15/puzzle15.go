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
		expectedAnswer := 1320
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day15/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 145
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day15/input.txt")
		fmt.Printf("day15 Answer 1 : %d\n", part1(input))
		fmt.Printf("day15 Answer 2 : %d\n", part2(input))
	}

}

func part1(lines []string) int {
	answer := 0
	line := lines[0]
	commands := strings.Split(line, ",")
	for _, command := range commands {
		answer += getHashAlgorithmValue(command)
	}

	return answer
}

type lens struct {
	register string
	value    int
}

func part2(lines []string) int {
	answer := 0
	line := lines[0]
	hashmap := map[int][]lens{}
	commands := strings.Split(line, ",")
	for _, command := range commands {
		register := ""
		getReg := true
		useEqual := false
		useDash := false
		equalVal := ""
		for _, cur := range command {
			char := string(cur)
			if getReg && char != "=" && char != "-" {
				register += char
				continue
			}
			getReg = false
			if char == "=" {
				useEqual = true
				continue
			}
			if useEqual {
				equalVal += char
				continue
			}
			if char == "-" {
				useDash = true
			}
		}
		box := getHashAlgorithmValue(register)
		if useEqual {
			if _, ok := hashmap[box]; !ok {
				newLens := lens{
					register: register,
					value:    S2i(equalVal),
				}
				hashmap[box] = append(hashmap[box], newLens)
			} else {
				lensFound := false
				for i, l := range hashmap[box] {
					if l.register == register {
						l.value = S2i(equalVal)
						lensFound = true
						hashmap[box][i] = l
					}
				}
				if !lensFound {
					newLens := lens{
						register: register,
						value:    S2i(equalVal),
					}
					hashmap[box] = append(hashmap[box], newLens)

				}
			}
		}
		if useDash {
			for i, l := range hashmap[box] {
				if l.register == register {
					hashmap[box] = append(hashmap[box][:i], hashmap[box][i+1:]...)
				}
			}
		}
	}

	for box, lenses := range hashmap {
		for i, lens := range lenses {
			answer += (1 + box) * (i + 1) * lens.value
		}
	}

	return answer
}

func getHashAlgorithmValue(line string) int {
	curVal := 0
	for _, cur := range line {
		asciiVal := int(cur)
		curVal += asciiVal
		curVal *= 17
		curVal = curVal % 256
	}
	return curVal
}
