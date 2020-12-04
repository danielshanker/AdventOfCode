package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	answer1, _ := readOpCode(lines)
	var answer2 int

	i := 0
	for _, line := range lines {
		var valid bool
		inst := strings.Split(line, " ")[0]
		val := strings.Split(line, " ")[1]
		ogLine := line
		if inst == "acc" {
			i++
			continue
		}
		if inst == "jmp" {
			line = "nop " + val
		}
		if inst == "nop" {
			line = "jmp " + val
		}
		lines[i] = line
		answer2, valid = readOpCode(lines)
		lines[i] = ogLine
		if valid {
			break
		}
		i++
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func readOpCode(instructions []string) (int, bool) {
	curLine := 0
	value := 0
	readLines := make([]bool, len(instructions))
	for {
		if curLine >= len(readLines) {
			return value, true
		}
		if readLines[curLine] {
			return value, false
		}
		readLines[curLine] = true
		split := strings.Split(instructions[curLine], " ")
		inst := split[0]
		intVal, _ := strconv.Atoi(split[1])

		switch inst {
		case "nop":
			curLine++
		case "acc":
			value += intVal
			curLine++
		case "jmp":
			curLine += intVal
		}
	}
}

func readInputLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}
