package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	lines := readInputLines("input.txt")
	//	lines := readInputLines("../sample.txt")

	rotationOrders := [4][2]int{
		//E
		{0, 1},
		//S
		{-1, 0},
		//W
		{0, -1},
		//N
		{1, 0},
	}
	curRot := 0

	// N, E
	position := [2]int{0, 0}

	// N, E
	rotation := [2]int{0, 1}

	for _, line := range lines {
		move := [2]int{0, 0}
		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		if direction == "N" {
			move[0] = distance
			move[1] = 0
		}
		if direction == "S" {
			move[0] = distance * -1
			move[1] = 0
		}
		if direction == "E" {
			move[1] = distance
			move[0] = 0
		}
		if direction == "W" {
			move[1] = distance * -1
			move[0] = 0
		}

		if direction == "F" {
			move[0] = rotation[0] * distance
			move[1] = rotation[1] * distance
		}

		if direction == "R" {
			turn := curRot
			if distance == 90 {
				turn = (curRot + 1) % 4
			}
			if distance == 180 {
				turn = (curRot + 2) % 4
			}
			if distance == 270 {
				turn = (curRot + 3) % 4
			}
			curRot = turn
			rotation = rotationOrders[turn]
		}

		if direction == "L" {
			turn := curRot
			if distance == 90 {
				turn = (curRot + 3) % 4
			}
			if distance == 180 {
				turn = (curRot + 2) % 4
			}
			if distance == 270 {
				turn = (curRot + 1) % 4
			}
			curRot = turn
			rotation = rotationOrders[turn]
		}
		position[0] += move[0]
		position[1] += move[1]
	}
	answer1 := int(math.Abs(float64(position[0])) + math.Abs(float64(position[1])))

	waypointPos := [2]int{1, 10}
	position[0] = 0
	position[1] = 0
	for _, line := range lines {
		move := [2]int{0, 0}
		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		if direction == "N" {
			waypointPos[0] += distance
		}
		if direction == "S" {
			waypointPos[0] -= distance
		}
		if direction == "E" {
			waypointPos[1] += distance
		}
		if direction == "W" {
			waypointPos[1] -= distance
		}

		if direction == "F" {
			position[0] += waypointPos[0] * distance
			position[1] += waypointPos[1] * distance
		}

		if direction == "R" {
			newWP := [2]int{0, 0}
			if distance == 90 {
				newWP[0] = waypointPos[1] * -1
				newWP[1] = waypointPos[0]
			}
			if distance == 180 {
				newWP[0] = waypointPos[0] * -1
				newWP[1] = waypointPos[1] * -1
			}
			if distance == 270 {
				newWP[0] = waypointPos[1]
				newWP[1] = waypointPos[0] * -1
			}
			waypointPos = newWP
		}

		if direction == "L" {
			newWP := [2]int{0, 0}
			if distance == 90 {
				newWP[0] = waypointPos[1]
				newWP[1] = waypointPos[0] * -1
			}
			if distance == 180 {
				newWP[0] = waypointPos[0] * -1
				newWP[1] = waypointPos[1] * -1
			}
			if distance == 270 {
				newWP[0] = waypointPos[1] * -1
				newWP[1] = waypointPos[0]
			}
			waypointPos = newWP
		}
		position[0] += move[0]
		position[1] += move[1]
	}

	answer2 := int(math.Abs(float64(position[0])) + math.Abs(float64(position[1])))

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
