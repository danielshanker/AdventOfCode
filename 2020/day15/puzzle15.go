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

	memory := make(map[int]int)

	start := strings.Split(lines[0], ",")

	lastNum := 0
	for i, val := range start {
		intVal, _ := strconv.Atoi(val)
		if i == len(start)-1 {
			lastNum = intVal
			break
		}
		memory[intVal] = i + 1
		lastNum = intVal
	}

	answer1 := 0
	for i := len(start); i < 30000000; i++ {
		if i == 2020 {
			answer1 = lastNum
		}
		if _, ok := memory[lastNum]; !ok {
			memory[lastNum] = i
			lastNum = 0
			continue
		}

		lastVal := i - memory[lastNum]
		memory[lastNum] = i
		lastNum = lastVal
	}
	answer2 := lastNum

	fmt.Println(fmt.Sprintf("Answer1: %d", answer1))
	fmt.Println(fmt.Sprintf("Answer2: %d", answer2))

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
