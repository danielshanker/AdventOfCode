package utils

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stack []string

type Coord struct {
	X int
	Y int
}

func Distance(c1, c2 Coord) Coord {
	x := int(math.Abs(float64(c1.X - c2.X)))
	y := int(math.Abs(float64(c1.Y - c2.Y)))

	return Coord{x, y}
}

func AbsDistance(c1, c2 Coord) int {
	x := int(math.Abs(float64(c1.X - c2.X)))
	y := int(math.Abs(float64(c1.Y - c2.Y)))

	return x + y
}

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
		if totalTime >= time.Second {
			fmt.Printf("%.03fs\n", totalTime.Seconds())
		} else if totalTime > time.Millisecond {
			a := float64(totalTime.Microseconds()) / 1000
			fmt.Printf("%.03fms\n", a)
		} else {
			fmt.Printf("%dμs\n", totalTime.Microseconds())
		}
		st = time.Now()
		fmt.Printf("day %d Answer 2: %d\n", day, part2(input))
		totalTime = time.Since(st)
		if totalTime >= time.Second {
			fmt.Printf("%.03fs\n", totalTime.Seconds())
		} else if totalTime >= time.Millisecond {
			a := float64(totalTime.Microseconds()) / 1000
			fmt.Printf("%.03fms\n", a)
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
		fmt.Printf("day %d part 1: %fs\n", day, totalTime.Seconds())
	} else if totalTime >= time.Millisecond {
		a := float64(totalTime.Microseconds()) / 1000
		fmt.Printf("%.03fms\n", a)
	} else {
		fmt.Printf("day %d part 1: %dμs\n", day, totalTime.Microseconds())
	}
	st = time.Now()
	part2(input)
	totalTime = time.Since(st)
	if totalTime >= time.Second {
		fmt.Printf("day %d part 2: %fs\n", day, totalTime.Seconds())
	} else if totalTime >= time.Millisecond {
		a := float64(totalTime.Microseconds()) / 1000
		fmt.Printf("%.03fms\n", a)
	} else {
		fmt.Printf("day %d part 2: %dμs\n", day, totalTime.Microseconds())
	}
}
