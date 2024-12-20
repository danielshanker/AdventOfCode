package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	reg := regexp.MustCompile(`"([0-9]+)"`)
	query := `select(.[70]=="301" or .[70]=="303" or .[70]=="322" or .[70]=="330" or .[70]=="4410" or .[70]=="4480" or .[70]=="4520" or .[70]=="1850" or .[70]=="675" or .[70]=="5103" or .[70]=="5104" or .[70]=="1180" or .[70]=="5150" or .[70]=="3281" or .[70]=="3282" or .[70]=="3283" or .[70]=="3284" or .[70]=="3285" or .[70]=="3300" or .[70]=="2160" or .[70]=="3360" or .[70]=="3250" or .[70]=="5190" or .[70]=="053" or .[70]=="059" or .[70]=="1068" or .[70]=="1069" or .[70]=="2022" or .[70]=="2584" or .[70]=="2585" or .[70]=="2333" or .[70]=="088" or .[70]=="091" or .[70]=="150" or .[70]=="502" or .[70]=="532" or .[70]=="1341" or .[70]=="1342" or .[70]=="3402" or .[70]=="3466" or .[70]=="3632" or .[70]=="3730" or .[70]=="4784" or .[70]=="4732" or .[70]=="4924" or .[70]=="120" or .[70]=="167" or .[70]=="1025" or .[70]=="1027" or .[70]=="1725" or .[70]=="4165" or .[70]=="4193" or .[70]=="3919" or .[70]=="5367" or .[70]=="5404" or .[70]=="5487" or .[70]=="5524" or .[70]=="078" or .[70]=="081" or .[70]=="632" or .[70]=="634" or .[70]=="636" or .[70]=="638" or .[70]=="1680" or .[70]=="2845" or .[70]=="2865" or .[70]=="2885" or .[70]=="2761" or .[70]=="2762" or .[70]=="2770")`
	rule := reg.FindAllStringSubmatch(query, -1)
	servers := []string{}

	for _, a := range rule {
		intVal := S2i(a[1])
		hexVal := strconv.FormatInt(int64(intVal), 16)
		servers = append(servers, hexVal)
	}
	newQuery := strings.Join(servers, `" or .[70]=="`)
	fmt.Println(newQuery)

	return
	if *test {
		expectedAnswer := 6
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day8/sample.txt")
		answer1 := part1(sample)
		if expectedAnswer == answer1 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer1))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer1))
		}
		expectedAnswer = 6
		answer2 := part2(sample)
		if expectedAnswer == answer2 {
			fmt.Println(fmt.Sprintf("Correct Answer %d", answer2))
		} else {
			fmt.Println(fmt.Sprintf("expected %d, got %d", expectedAnswer, answer2))
		}
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day8/input.txt")
		fmt.Printf("day8 Answer 1 : %d\n", part1(input))
		fmt.Printf("day8 Answer 2 : %d\n", part2(input))
	}

}

type connection struct {
	l string
	r string
}

func part1(lines []string) int {
	answer := 0
	connections := map[string]connection{}
	conReg := regexp.MustCompile("[A-Z0-9]{3}")

	instructions := strings.Split(lines[0], "")
	start := false
	for _, line := range lines {
		if !start {
			if line == "" {
				start = true
			}
			continue
		}
		rule := conReg.FindAllStringSubmatch(line, -1)
		con := connection{
			l: rule[1][0],
			r: rule[2][0],
		}
		connections[rule[0][0]] = con
	}

	curNode := "AAA"

	for {
		found := false
		for _, cha := range instructions {
			answer++
			inst := string(cha)
			if inst == "L" {
				curNode = connections[curNode].l
			} else {
				curNode = connections[curNode].r
			}
			if curNode == "ZZZ" {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return answer
}

func part2(lines []string) int {
	answer := 0
	connections := map[string]connection{}
	conReg := regexp.MustCompile("[A-Z0-9]{3}")
	aNodes := []string{}

	instructions := strings.Split(lines[0], "")
	start := false
	for _, line := range lines {
		if !start {
			if line == "" {
				start = true
			}
			continue
		}
		rule := conReg.FindAllStringSubmatch(line, -1)
		con := connection{
			l: rule[1][0],
			r: rule[2][0],
		}
		connections[rule[0][0]] = con
		if string(rule[0][0][2]) == "A" {
			aNodes = append(aNodes, rule[0][0])
		}
	}

	zFounds := []int{}

	for i, curNode := range aNodes {
		count := 0
		for {
			found := false
			for _, cha := range instructions {
				count++
				inst := string(cha)
				if inst == "L" {
					curNode = connections[curNode].l
				} else {
					curNode = connections[curNode].r
				}
				if string(curNode[2]) == "Z" {
					zFounds = append(zFounds, count)
					found = true
					break
				}
				aNodes[i] = curNode
			}
			if found {
				break
			}
		}
	}

	answer = LCM(zFounds[0], zFounds[1], zFounds[2:]...)

	return answer
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
