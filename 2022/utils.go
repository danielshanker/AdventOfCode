package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
