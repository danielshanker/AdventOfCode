package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct{
	value int
	called bool
}

type Card struct {
	card [][]Cell
	row []int
	column []int
	cardWon []bool
}

func main() {
	sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/sample.txt")
	input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day4/input.txt")

	fmt.Print("sample: ")
	part1(sample)
	part1(input)

	fmt.Print("\n\nsample: ")
	part2(sample)
	part2(input)

}

func part1(lines []string) {
	answer1 := 0
	lastNum := 0
	var winningCard Card

	calledNums, cards := setupBingo(lines)

	out:
	for _, numCalled := range calledNums {
	for _, card := range cards {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
					if card.card[i][j].value == numCalled {
						card.card[i][j].called = true
						card.row[i] += 1
						card.column[j] += 1
						if card.row[i] == 5 {
							lastNum = numCalled
							winningCard = card
							break out
						}
						if card.column[j] == 5 {
							lastNum = numCalled
							winningCard = card
							break out
						}
					}
				}
			}
		}
	}

	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !winningCard.card[i][j].called {
				sum += winningCard.card[i][j].value
			}
		}
	}

	answer1 = sum*lastNum

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0
	lastNum := 0
	var losingCard Card

	calledNums, cards := setupBingo(lines)

	cardsLeft := len(cards)
	out:
	for _, numCalled := range calledNums {
		for _, card := range cards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if card.card[i][j].value == numCalled {
						card.card[i][j].called = true
						card.row[i] += 1
						card.column[j] += 1
						if card.row[i] == 5 && !card.cardWon[0] {
							card.cardWon[0] = true
							cardsLeft--
							if cardsLeft == 0 {
								lastNum = numCalled
								losingCard = card
								break out
							}
						}
						if card.column[j] == 5 && !card.cardWon[0]{
							card.cardWon[0] = true
							cardsLeft--
							if cardsLeft == 0 {
								lastNum = numCalled
								losingCard = card
								break out
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !losingCard.card[i][j].called {
				sum += losingCard.card[i][j].value
			}
		}
	}

	answer2 = sum*lastNum

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func setupBingo(lines []string) ([]int, []Card){
	cardIndex := -1
	var calledNums []int
	var cards []Card

	for i, line := range lines {
		if i == 0 {
			t := strings.Split(line, ",")
			for _, val := range t {
				calledNums = append(calledNums, s2i(val))
			}
			continue
		}
		if line == "" {
			cardIndex++
			cards = append(cards, Card{})
			for i := 0; i < 5; i++ {
				cards[cardIndex].column = append(cards[cardIndex].column,0)
				cards[cardIndex].row = append(cards[cardIndex].row,0)
				cards[cardIndex].cardWon = append(cards[cardIndex].cardWon, false)
			}
			continue
		}

		t := strings.Fields(line)
		var cells []Cell
		for _, curVal := range t {
			var cell Cell
			cell.called = false
			cell.value = s2i(curVal)
			cells = append(cells, cell)
		}
		cards[cardIndex].card = append(cards[cardIndex].card, cells)
	}
	return calledNums, cards
}

func printCard(card Card) {
	fmt.Println("")
	for i:= 0; i< 5; i++ {
		fmt.Println(card.card[i])
	}
	fmt.Println("")
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
