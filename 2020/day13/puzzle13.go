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

	time, _ := strconv.Atoi(lines[0])
	busses := strings.Split(lines[1], ",")

	answer1 := 0
	for i := time; i > 0; i++ {
		busFound := false
		for _, bus := range busses {
			if bus == "x" {
				continue
			}
			busTime, _ := strconv.Atoi(bus)
			if i%busTime == 0 {
				answer1 = (i - time) * busTime
				busFound = true
				break
			}
		}
		if busFound {
			break
		}
	}

	a, _ := strconv.Atoi(busses[0])

	iter := uint64(a)
	time2 := uint64(a)

	for i := 1; i < len(busses); i++ {
		if busses[i] == "x" {
			continue
		}
		c, _ := strconv.Atoi(busses[i])
		bus := uint64(c)

		for (time2+uint64(i))%bus != 0 {
			time2 += iter
		}
		iter *= bus

	}

	answer2 := time2

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
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
