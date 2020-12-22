package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readInputLines("input.txt")
	//	lines := readInputLines("../sample.txt")

	var cups []int
	for _, i := range lines[0] {
		cups = append(cups, s2i(string(i)))
	}

	for i := 0; i < 100; i++ {
		var heldCups [3]int
		cups, heldCups = pickupCups(cups)
		dest := cups[0] - 1

		for {
			safeCup := true
			for _, i := range heldCups {
				if dest == 0 {
					dest = 9
				}
				if dest == i {
					dest--
					safeCup = false
					break
				}
			}
			if safeCup {
				break
			}
		}

		cups = addCupsBack(cups, heldCups, dest)
		cups = shiftCurrentCup(cups)
	}

	answer1 := ""
	index := 0
	for i, j := range cups {
		if j == 1 {
			index = i
			break
		}
	}

	for i := 1; i < len(cups); i++ {
		cupI := (i + index) % len(cups)
		answer1 += strconv.Itoa(cups[cupI])

	}

	answer2 := part2(lines[0])

	fmt.Println(fmt.Sprintf("Answer 1 : %s", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func part2(line string) int {
	var l cupRing
	for _, i := range line {
		l.AddToRing(s2i(string(i)))
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			break
		}
		ptr = ptr.next
	}
	for i := 10; i <= 1000000; i++ {
		var curCup cup
		curCup.label = i
		ptr.next = &curCup
		l.len++
		ptr = ptr.next
	}
	l.loopList()

	l.lookUp = make(map[int]*cup)
	ptr = l.head
	for i := 0; i < l.len; i++ {
		l.lookUp[ptr.label] = ptr
		ptr = ptr.next
	}

	//for k := 0; k < 10; k++ {
	for k := 0; k < 10000000; k++ {
		pickedCups := l.pickupCups2()
		destination := l.head.label - 1
		for {
			if destination == 0 {
				destination = 1000000
				//			destination = 9
			}
			safeCup := true
			if pickedCups.label == destination {
				safeCup = false
				destination--
				continue
			}
			if pickedCups.next.label == destination {
				safeCup = false
				destination--
				continue
			}
			if pickedCups.next.next.label == destination {
				safeCup = false
				destination--
				continue
			}
			if safeCup {
				break
			}
		}

		l.addCupsBack2(pickedCups, destination)
		l.head = l.head.next
	}
	return l.getAnswer()
}

func (l *cupRing) getAnswer() int {
	ptr := l.head
	answer := 1
	for i := 0; i < l.len; i++ {
		if ptr.label == 1 {
			answer = ptr.next.label
			answer *= ptr.next.next.label
			break
		}
		ptr = ptr.next
	}

	return answer
}

func (l *cupRing) addCupsBack2(pickedCups cup, destIndex int) {
	destination := l.lookUp[destIndex]
	destNext := destination.next
	destination.next = &pickedCups
	l.lookUp[destination.next.label] = destination.next
	destination.next.next.next.next = destNext
	l.lookUp[destination.next.next.next.label] = destination.next.next.next
}

func (l *cupRing) pickupCups2() cup {
	curCup := l.head

	pickedCups := curCup.next

	curCup.next = curCup.next.next.next.next
	l.lookUp[curCup.label] = curCup
	pickedCups.next.next.next = nil
	l.lookUp[pickedCups.next.next.label] = pickedCups.next.next

	return *pickedCups
}

func (l *cupRing) loopList() {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = l.head
			return
		}
		ptr = ptr.next
	}
}

func (l *cupRing) AddToRing(val int) {
	curCup := cup{}
	curCup.label = val
	if l.len == 0 {
		l.head = &curCup
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &curCup
			l.len++
			return
		}
		ptr = ptr.next
	}
}

type cup struct {
	next  *cup
	label int
}

type cupRing struct {
	head   *cup
	len    int
	lookUp map[int]*cup
}

func addCupsBack(cups []int, heldCups [3]int, dest int) []int {
	var newCups []int

	for _, i := range cups {
		newCups = append(newCups, i)
		if i == dest {
			newCups = append(newCups, heldCups[0])
			newCups = append(newCups, heldCups[1])
			newCups = append(newCups, heldCups[2])
		}
	}

	return newCups
}

func shiftCurrentCup(cups []int) []int {
	var newCups []int

	for i := 1; i < len(cups); i++ {
		newCups = append(newCups, cups[i])
	}
	newCups = append(newCups, cups[0])

	return newCups
}

func pickupCups(cups []int) ([]int, [3]int) {
	var heldCups [3]int
	heldCups[0] = cups[1]
	heldCups[1] = cups[2]
	heldCups[2] = cups[3]
	var newCups []int
	newCups = append(newCups, cups[0])
	for i := 4; i < len(cups); i++ {
		newCups = append(newCups, cups[i])
	}

	return newCups, heldCups
}

func s2i(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("OH NO! OH NO! NOT AN INT!")
	}
	return num
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
