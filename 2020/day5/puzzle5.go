package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := readInputLines("day5.txt")
	var seatList []string
	for _, line := range lines {
		seatList = append(seatList, line)
	}

	highestSeat := 0
	var takenSeats [128*8 + 5]bool

	for _, seat := range seatList {
		seatID := findSeatID(seat)
		takenSeats[seatID] = true
		if seatID > highestSeat {
			highestSeat = seatID
		}
	}

	var yourSeat int
	for i := 1; i < len(takenSeats)-1; i++ {
		if !takenSeats[i] {
			if takenSeats[i+1] && takenSeats[i-1] {
				yourSeat = i
			}
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", highestSeat))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", yourSeat))
}

func findSeatID(seat string) int {
	rowMin := 0
	rowMax := 127
	rowMid := 128 / 2
	colMin := 0
	colMax := 7
	colMid := 4

	for _, curLetter := range seat {
		if string(curLetter) == "F" {
			rowMax = rowMid
			rowMid = (((rowMax + 1) - rowMin) / 2) + rowMin
		}
		if string(curLetter) == "B" {
			rowMin = rowMid
			rowMid = (((rowMax + 1) - rowMin) / 2) + rowMin
		}
		if string(curLetter) == "L" {
			colMax = colMid
			colMid = (((colMax + 1) - colMin) / 2) + colMin
		}
		if string(curLetter) == "R" {
			colMin = colMid
			colMid = (((colMax + 1) - colMin) / 2) + colMin
		}
	}

	seatID := rowMin*8 + colMin

	return seatID
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
