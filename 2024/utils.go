package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stack []string

func S2i(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Print("OH NO! OH NO! NOT AN INT! - ")
		fmt.Println(val)
	}
	return num
}

func ReadInputLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		text = append(text, t)
	}

	return text
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

type Queue []string

func (q *Queue) Pop() string {
	if len(*q) == 0 {
		return ""
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q *Queue) Push(str string) {
	*q = append(*q, str)
}

type P func([]string) int

func Start(test *bool, day int, part1 P, part2 P, a1 int, a2 int) {
	//	justTime(day, part1, part2)
	//	return
	if *test {
		expectedAnswer := a1
		sample := ReadInputLines(fmt.Sprintf("/home/daniel.shanker/Pers/AdventOfCode/2024/day%d/sample.txt", day))
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = a2
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines(fmt.Sprintf("/home/daniel.shanker/Pers/AdventOfCode/2024/day%d/input.txt", day))
		st := time.Now()
		fmt.Printf("day %d Answer 1: %d\n", day, part1(input))
		totalTime := time.Since(st)
		if totalTime > time.Millisecond {
			fmt.Printf("%dms\n", totalTime.Milliseconds())
		} else {
			fmt.Printf("%dμs\n", totalTime.Microseconds())
		}
		st = time.Now()
		fmt.Printf("day %d Answer 2: %d\n", day, part2(input))
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

func justTime(day int, part1 P, part2 P) {
	input := ReadInputLines(fmt.Sprintf("/home/daniel.shanker/Pers/AdventOfCode/2024/day%d/input.txt", day))
	st := time.Now()
	part1(input)
	totalTime := time.Since(st)
	if totalTime >= time.Second {
		fmt.Printf("day %d part 1: %ds\n", day, int(totalTime.Seconds()))
	} else if totalTime >= time.Millisecond {
		fmt.Printf("day %d part 1: %dms\n", day, totalTime.Milliseconds())
	} else {
		fmt.Printf("day %d part 1: %dμs\n", day, totalTime.Microseconds())
	}
	st = time.Now()
	part2(input)
	totalTime = time.Since(st)
	if totalTime >= time.Second {
		fmt.Printf("day %d part 2: %ds\n", day, int(totalTime.Seconds()))
	} else if totalTime >= time.Millisecond {
		fmt.Printf("day %d part 2: %dms\n", day, totalTime.Milliseconds())
	} else {
		fmt.Printf("day %d part 2: %dμs\n", day, totalTime.Microseconds())
	}
}
