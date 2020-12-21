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

	var deck1 []int
	var deck2 []int

	var curDeck []int
	for _, line := range lines {
		if _, ok := strconv.Atoi(line); ok != nil {
			if line == "" {
				deck1 = curDeck
				curDeck = nil
				continue
			}
			continue
		}
		curDeck = append(curDeck, s2i(line))
	}
	deck2 = curDeck
	ogDeck1 := deck1
	ogDeck2 := deck2

	var card1 int
	var card2 int
	var winningDeck []int
	for {
		if len(deck1) == 0 {
			winningDeck = deck2
			break
		}
		if len(deck2) == 0 {
			winningDeck = deck1
			break
		}
		card1, deck1 = draw(deck1)
		card2, deck2 = draw(deck2)

		if card1 > card2 {
			deck1 = placeOnBottom(deck1, card1, card2)
		} else {
			deck2 = placeOnBottom(deck2, card2, card1)
		}
	}

	answer1 := 0
	index := len(winningDeck)
	for _, card := range winningDeck {
		answer1 += index * card
		index--
	}
	deck1 = ogDeck1
	deck2 = ogDeck2

	answer2 := 0
	_, winningDeck = recursiveWar(deck1, deck2)
	index = len(winningDeck)
	for _, card := range winningDeck {
		answer2 += index * card
		index--
	}

	fmt.Println(winningDeck)

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func recursiveWar(deck1 []int, deck2 []int) (int, []int) {
	var deck1Confs [][]int
	var deck2Confs [][]int
	for {
		if len(deck1) == 0 {
			return 2, deck2
		}
		if len(deck2) == 0 {
			return 1, deck1
		}

		if checkRecurse(deck1, deck1Confs) && checkRecurse(deck2, deck2Confs) {
			return 1, deck1
		}

		deck1Confs = append(deck1Confs, deck1)
		deck2Confs = append(deck2Confs, deck2)
		var card1 int
		var card2 int

		card1, deck1 = draw(deck1)
		card2, deck2 = draw(deck2)

		if len(deck1) < card1 || len(deck2) < card2 {
			if card1 > card2 {
				deck1 = placeOnBottom(deck1, card1, card2)
			} else {
				deck2 = placeOnBottom(deck2, card2, card1)
			}
			continue
		}
		r1 := getRecursiveDeck(deck1, card1)
		r2 := getRecursiveDeck(deck2, card2)
		winner, _ := recursiveWar(r1, r2)
		if winner == 1 {
			deck1 = placeOnBottom(deck1, card1, card2)
		} else {
			deck2 = placeOnBottom(deck2, card2, card1)
		}
	}
}

func getRecursiveDeck(deck []int, cards int) []int {
	var newDeck []int
	for i := 0; i < cards; i++ {
		newDeck = append(newDeck, deck[i])
	}
	return newDeck
}

func checkRecurse(deck []int, oldDecks [][]int) bool {
	for _, arrangement := range oldDecks {
		if len(deck) != len(arrangement) {
			continue
		}
		match := true
		for i := 0; i < len(deck); i++ {
			if deck[i] != arrangement[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func draw(deck []int) (int, []int) {
	card := deck[0]
	deck = deck[1:]

	return card, deck
}

func placeOnBottom(deck []int, card1 int, card2 int) []int {

	deck = append(deck, card1)
	deck = append(deck, card2)

	return deck
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
