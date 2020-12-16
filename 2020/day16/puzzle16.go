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

	ranges, index := getRanges(lines)

	yourTick := strings.Split(lines[index], ",")

	index += 3
	ticks := getTicks(lines, index)
	var invalid []int
	var valid [][]string
	for _, ticket := range ticks {
		validTick := true
		for _, value := range ticket {
			val, _ := strconv.Atoi(value)
			_, ok := ranges[val]
			if !ok {
				invalid = append(invalid, val)
				validTick = false
			}
		}
		if validTick {
			valid = append(valid, ticket)
		}
	}

	answer1 := 0
	for _, i := range invalid {
		answer1 += i
	}

	disallowed := make(map[string][]int)

	ranges2 := getRanges2(lines)

	for _, ticket := range valid {
		for i, j := range ticket {
			for tickType := range ranges2 {
				val, _ := strconv.Atoi(j)
				_, ok := ranges2[tickType][val]
				if !ok {
					disallowed[tickType] = append(disallowed[tickType], i)
				}
			}
		}
	}

	found := make(map[string]bool)
	position := make(map[int]string)

	for len(found) < len(ranges2) {
		for i, _ := range ranges2 {
			_, ok := found[i]
			if ok {
				continue
			}
			if len(ranges2)-len(disallowed[i])-1 <= len(found) {
				for j := 0; j <= len(disallowed); j++ {
					_, ok := position[j]
					if ok {
						continue
					}
					foundIt := false
					for _, k := range disallowed[i] {
						if j == k {
							foundIt = true
							break
						}
					}
					if !foundIt {
						position[j] = i
						found[i] = true
						break
					}
				}
			}
		}
	}
	answer2 := 1
	for i, j := range yourTick {
		departure := strings.Split(position[i], " ")
		if departure[0] != "departure" {
			continue
		}
		value, _ := strconv.Atoi(j)
		answer2 *= value

	}

	fmt.Println(fmt.Sprintf("Answer1: %d", answer1))
	fmt.Println(fmt.Sprintf("Answer2: %d", answer2))
}

func getRanges2(lines []string) map[string]map[int]bool {

	ranges := make(map[string]map[int]bool)

	for i, line := range lines {
		if line == "" {
			i += 2
			return ranges
		}
		a := strings.Split(line, ": ")
		b := strings.Split(a[1], " or ")

		intMap := make(map[int]bool)
		for _, j := range b {
			c := strings.Split(j, "-")
			min, _ := strconv.Atoi(c[0])
			max, _ := strconv.Atoi(c[1])

			for k := min; k <= max; k++ {
				intMap[k] = true
			}
		}
		ranges[a[0]] = intMap
	}

	return ranges

}

func getTicks(lines []string, index int) [][]string {
	var ticks [][]string
	for i := index; i < len(lines); i++ {
		line := lines[i]
		ticks = append(ticks, strings.Split(line, ","))
	}

	return ticks
}

func getRanges(lines []string) (map[int]bool, int) {

	ranges := make(map[int]bool)

	for i, line := range lines {
		if line == "" {
			i += 2
			return ranges, i
		}
		a := strings.Split(line, ": ")
		b := strings.Split(a[1], " or ")

		for _, j := range b {
			c := strings.Split(j, "-")
			min, _ := strconv.Atoi(c[0])
			max, _ := strconv.Atoi(c[1])

			for k := min; k <= max; k++ {
				ranges[k] = true
			}
		}
	}

	return ranges, 0
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
