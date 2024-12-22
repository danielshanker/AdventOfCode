package main

import (
	"flag"
	"strings"
	"sync"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 21, part1, part2, 126384, 0)

}

const (
	north = 0
	south = 1
	east  = 2
	west  = 3
)

var numToDir = map[string]map[string]string{
	"A": {
		"A": "A",
		"0": "<A",
		"1": "^<<A",
		"2": "<^A",
		"3": "^A",
		"4": "^^<<A",
		"5": "<^^A",
		"6": "^^A",
		"7": "^^^<<A",
		"8": "<^^^A",
		"9": "^^^A",
	},
	"0": {
		"A": ">A",
		"0": "A",
		"1": "^<A",
		"2": "^A",
		"3": "^>A",
		"4": "^^<A",
		"5": "^^A",
		"6": "^^>A",
		"7": "^^^<A",
		"8": "^^^A",
		"9": "^^^>A",
	},
	"1": {
		"A": ">>vA",
		"0": ">vA",
		"1": "A",
		"2": ">A",
		"3": ">>A",
		"4": "^A",
		"5": "^>A",
		"6": "^>>A",
		"7": "^^A",
		"8": "^^>A",
		"9": "^^>>A",
	},
	"2": {
		"A": "v>A",
		"0": "vA",
		"1": "<A",
		"2": "A",
		"3": ">A",
		"4": "<^A",
		"5": "^A",
		"6": "^>A",
		"7": "<^^A",
		"8": "^^A",
		"9": "^^>A",
	},
	"3": {
		"A": "vA",
		"0": "<vA",
		"1": "<<A",
		"2": "<A",
		"3": "A",
		"4": "<<^A",
		"5": "<^A",
		"6": "^A",
		"7": "<<^^A",
		"8": "<^^A",
		"9": "^^A",
	},
	"4": {
		"A": ">>vvA",
		"0": ">vvA",
		"1": "vA",
		"2": "v>A",
		"3": "v>>A",
		"4": "A",
		"5": ">A",
		"6": ">>A",
		"7": "^A",
		"8": "^>A",
		"9": "^>>A",
	},
	"5": {
		"A": "vv>A",
		"0": "vvA",
		"1": "<vA",
		"2": "vA",
		"3": "v>A",
		"4": "<A",
		"5": "A",
		"6": ">A",
		"7": "<^A",
		"8": "^A",
		"9": "^>A",
	},
	"6": {
		"A": "vvA",
		"0": "<vvA",
		"1": "<<vA",
		"2": "<vA",
		"3": "vA",
		"4": "<<A",
		"5": "<A",
		"6": "A",
		"7": "<<^A",
		"8": "<^A",
		"9": "^A",
	},
	"7": {
		"A": ">>vvvA",
		"0": ">vvvA",
		"1": "vvA",
		"2": "vv>A",
		"3": "vv>>A",
		"4": "vA",
		"5": "v>A",
		"6": "v>>A",
		"7": "A",
		"8": ">A",
		"9": ">>A",
	},
	"8": {
		"A": "vvv>A",
		"0": "vvvA",
		"1": "vv<A",
		"2": "vvA",
		"3": "vv>A",
		"4": "v<A",
		"5": "vA",
		"6": "v>A",
		"7": "<A",
		"8": "A",
		"9": ">A",
	},
	"9": {
		"A": "vvvA",
		"0": "<vvvA",
		"1": "<<vvA",
		"2": "<vvA",
		"3": "vvA",
		"4": "<<vvA",
		"5": "<vA",
		"6": "vA",
		"7": "<<A",
		"8": "<A",
		"9": "A",
	},
}

var dirToDir = map[string]map[string]string{
	"A": {
		"A": "A",
		"^": "<A",
		"<": "v<<A",
		">": "vA",
		"v": "<vA",
	},
	"^": {
		"A": ">A",
		"^": "A",
		"<": "v<A",
		">": "v>A",
		"v": "vA",
	},
	"<": {
		"A": ">>^A",
		"^": ">^A",
		"<": "A",
		">": ">>A",
		"v": ">A",
	},
	">": {
		"A": "^A",
		"^": "<^A",
		"<": "<<A",
		">": "A",
		"v": "<A",
	},
	"v": {
		"A": "^>A",
		"^": "^A",
		"<": "<A",
		">": ">A",
		"v": "A",
	},
}

func part1(lines []string) int {
	answer := 0
	allKeys := []string{}

	for _, line := range lines {
		curKey := "A"
		keys := ""
		for _, r := range line {
			char := string(r)
			keys += numToDir[curKey][char]
			curKey = char
		}
		allKeys = append(allKeys, keys)
	}

	c := make(chan int)
	for i, keys := range allKeys {
		go find(i, keys, lines[i], c, 2)
	}

	for range lines {
		answer += <-c
	}
	return answer
}

type memKey struct {
	val   string
	depth int
}

var cache sync.Map = sync.Map{}

func getKeys(sequence string, depth int) int {
	if depth == 0 {
		return len(sequence)
	}

	memK := memKey{sequence, depth}

	if val, ok := cache.Load(memK); ok {
		return val.(int)
	}

	curKey := "A"

	cost := 0
	for i := 0; i < len(sequence); {
		end := i + 1

		for j := i; j < len(sequence); j++ {
			if string(sequence[j]) == "A" {
				break
			}
			end++
		}

		chunk := sequence[i:end]
		var builder strings.Builder
		for _, r := range chunk {
			target := string(r)
			builder.WriteString(dirToDir[curKey][target])
			curKey = target
		}
		key := builder.String()

		cost += getKeys(key, depth-1)
		i = end

	}
	cache.Store(memK, cost)

	return cost
}

func find(i int, keys string, line string, c chan int, depth int) {
	k := getKeys(keys, depth)
	n := S2i(line[:len(line)-1])
	c <- n * k
}

func part2(lines []string) int {
	answer := 0

	allKeys := []string{}

	for _, line := range lines {
		curKey := "A"
		keys := ""
		for _, r := range line {
			char := string(r)
			keys += numToDir[curKey][char]
			curKey = char
		}
		allKeys = append(allKeys, keys)
	}

	c := make(chan int)
	for i, keys := range allKeys {
		go find(i, keys, lines[i], c, 25)
	}

	for range lines {
		answer += <-c
	}

	return answer
}
