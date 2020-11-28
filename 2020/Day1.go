package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readInputLines("day1.txt")

	found := false
	answer1 := 0
	answer2 := 0
	for _, i := range lines {
		for _, j := range lines {
			if i == j {
				continue
			}
			if i+j == 2020 {
				found = true
				answer1 = i * j
				break
			}
		}
		if found == true {
			break
		}
	}

	found = false
	for _, i := range lines {
		for _, j := range lines {
			for _, k := range lines {
				if i+j+k == 2020 {
					found = true
					answer2 = i * j * k
					break
				}
			}
			if found == true {
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("answer 1: %d", answer1))
	fmt.Println(fmt.Sprintf("answer 2: %d", answer2))
}

func readInputLines(fileName string) []int {
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

	var data []int

	for _, ln := range text {
		num, _ := strconv.Atoi(ln)
		data = append(data, num)
	}

	return data
}
