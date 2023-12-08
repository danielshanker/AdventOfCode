package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
	. "utils"
)

const (
	high         = 0
	pair         = 1
	twoPair      = 2
	threeOfAKind = 3
	fullHouse    = 4
	fourOfAKind  = 5
	fiveOfAKind  = 6
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day7/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day7/input.txt")
		part1(input)
		part2(input)
	}

}

type handInfo struct {
	hand string
	rank int
	bet  int
}

func part1(lines []string) {
	answer := 0
	hands := []handInfo{}

	for _, line := range lines {
		info := strings.Fields(line)
		hand := handInfo{
			hand: info[0],
			bet:  S2i(info[1]),
			rank: getHandRank(info[0]),
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return sortHandRank(hands[i], hands[j])
	})

	for i, hand := range hands {
		answer += hand.bet * (i + 1)
	}
	fmt.Printf("day7 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0
	hands := []handInfo{}

	for _, line := range lines {
		info := strings.Fields(line)
		hand := handInfo{
			hand: info[0],
			bet:  S2i(info[1]),
			rank: getHandRank2(info[0]),
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return sortHandRank2(hands[i], hands[j])
	})
	for i, hand := range hands {
		answer += hand.bet * (i + 1)
	}

	fmt.Printf("day7 Answer 2 : %d\n", answer)
}

func getHandRank(hand string) int {
	handMap := map[string]int{}
	for _, cha := range hand {
		card := string(cha)
		handMap[card]++
	}

	if len(handMap) == 1 {
		return fiveOfAKind
	}
	if len(handMap) == 5 {
		return high
	}

	pairs := 0
	threes := 0

	for _, count := range handMap {
		if count == 2 {
			pairs++
		}
		if count == 3 {
			threes++
		}
	}

	if pairs == 1 {
		if threes == 1 {
			return fullHouse
		}
		return pair
	}
	if pairs == 2 {
		return twoPair
	}
	if threes == 1 {
		return threeOfAKind
	}

	return fourOfAKind
}

func sortHandRank(handA, handB handInfo) bool {
	if handA.rank < handB.rank {
		return true
	}
	if handA.rank > handB.rank {
		return false
	}

	for i := 0; i < 5; i++ {
		cardA := string(handA.hand[i])
		cardB := string(handB.hand[i])
		cardAVal := 0
		cardBVal := 0
		if cardA == "A" {
			cardAVal = 14
		} else if cardA == "K" {
			cardAVal = 13
		} else if cardA == "Q" {
			cardAVal = 12
		} else if cardA == "J" {
			cardAVal = 11
		} else if cardA == "T" {
			cardAVal = 10
		} else {
			cardAVal = S2i(cardA)
		}
		if cardB == "A" {
			cardBVal = 14
		} else if cardB == "K" {
			cardBVal = 13
		} else if cardB == "Q" {
			cardBVal = 12
		} else if cardB == "J" {
			cardBVal = 11
		} else if cardB == "T" {
			cardBVal = 10
		} else {
			cardBVal = S2i(cardB)
		}
		if cardAVal < cardBVal {
			return true
		}
		if cardAVal > cardBVal {
			return false
		}
	}
	return false
}

func getHandRank2(hand string) int {
	handMap := map[string]int{}
	for _, cha := range hand {
		card := string(cha)
		handMap[card]++
	}

	if len(handMap) == 1 {
		return fiveOfAKind
	}

	pairs := 0
	threes := 0
	fours := 0
	jokers := 0

	for card, count := range handMap {
		if card == "J" {
			jokers = count
			continue
		}
		if count == 2 {
			pairs++
		}
		if count == 3 {
			threes++
		}
		if count == 4 {
			fours++
		}
	}

	if jokers == 0 {
		if len(handMap) == 5 {
			return high
		}
		if pairs == 1 {
			if threes == 1 {
				return fullHouse
			}
			return pair
		}
		if pairs == 2 {
			return twoPair
		}
		if threes == 1 {
			return threeOfAKind
		}
		return fourOfAKind
	}
	if jokers == 1 {
		if pairs == 1 {
			return threeOfAKind
		}
		if pairs == 2 {
			return fullHouse
		}
		if threes == 1 {
			return fourOfAKind
		}
		if fours == 1 {
			return fiveOfAKind
		}
		return pair
	}
	if jokers == 2 {
		if pairs == 1 {
			return fourOfAKind
		}
		if threes == 1 {
			return fiveOfAKind
		}
		return threeOfAKind
	}
	if jokers == 3 {
		if pairs == 1 {
			return fiveOfAKind
		}
		return fourOfAKind
	}
	if jokers == 4 || jokers == 5 {
		return fiveOfAKind
	}
	return high
}

func sortHandRank2(handA, handB handInfo) bool {
	if handA.rank < handB.rank {
		return true
	}
	if handA.rank > handB.rank {
		return false
	}

	for i := 0; i < 5; i++ {
		cardA := string(handA.hand[i])
		cardB := string(handB.hand[i])
		cardAVal := 0
		cardBVal := 0
		if cardA == "A" {
			cardAVal = 14
		} else if cardA == "K" {
			cardAVal = 13
		} else if cardA == "Q" {
			cardAVal = 12
		} else if cardA == "J" {
			cardAVal = 0
		} else if cardA == "T" {
			cardAVal = 10
		} else {
			cardAVal = S2i(cardA)
		}
		if cardB == "A" {
			cardBVal = 14
		} else if cardB == "K" {
			cardBVal = 13
		} else if cardB == "Q" {
			cardBVal = 12
		} else if cardB == "J" {
			cardBVal = 0
		} else if cardB == "T" {
			cardBVal = 10
		} else {
			cardBVal = S2i(cardB)
		}
		if cardAVal < cardBVal {
			return true
		}
		if cardAVal > cardBVal {
			return false
		}
	}
	return false
}
