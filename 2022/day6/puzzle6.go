package main

import (
	"flag"
	"fmt"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day6/sample.txt")
		for i := 0; i < len(sample); i++ {
			part1(sample[i])
		}
		for i := 0; i < len(sample); i++ {
			part2(sample[i])
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day6/input.txt")
		part1(input[0])
		part2(input[0])
	}
}

func part1(line string) {
	answer1 := 0

	var movingSlice []string

	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if i < 4 {
			movingSlice = append(movingSlice, char)
		} else {
			movingSlice = moveSlice(movingSlice, char)
		}

		contain := contains(movingSlice)
		if !contain && i >= 4 {
			answer1 = i + 1
			break
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func contains(slice []string) bool {
	for i := 0; i < len(slice); i++ {
		val := slice[i]
		for j := 0; j < len(slice); j++ {
			if i == j {
				continue
			}
			char := slice[j]
			if val == char {
				return true
			}
		}
	}
	return false
}

func moveSlice(slice []string, val string) []string {
	s := slice[1:]
	s = append(s, val)
	return s
}

func part2(line string) {
	answer2 := 0
	var movingSlice []string

	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if i < 14 {
			movingSlice = append(movingSlice, char)
		} else {
			movingSlice = moveSlice(movingSlice, char)
		}

		contain := contains(movingSlice)
		if !contain && i >= 14 {
			answer2 = i + 1
			break
		}
	}

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

//func s2i(val string) int {
//	num, err := strconv.Atoi(val)
//	if err != nil {
//		fmt.Println("OH NO! OH NO! NOT AN INT!")
//	}
//	return num
//}
//
//func readInputLines(fileName string) []string {
//	file, err := os.Open(fileName)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//
//	scanner.Split(bufio.ScanLines)
//	var text []string
//
//	for scanner.Scan() {
//		t := strings.TrimSpace(scanner.Text())
//		text = append(text, t)
//	}
//
//	return text
//}
