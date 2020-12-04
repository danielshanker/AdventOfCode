package tools

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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
