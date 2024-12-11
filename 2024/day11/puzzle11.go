package main

import (
	"flag"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 11, part1, part2, 55312, 0)

}

/*
type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
}
*/

func part1(lines []string) int {
	return run(lines[0], 25)

	/*
		answer := 0
			A TESTAMENT TO MY HUBRIS!!!

			list := linkedList{}
			s := strings.Fields(lines[0])

			list.head = &node{
				val:  S2i(s[0]),
				next: nil,
			}

			current := list.head
			for i, st := range s {
				if i == 0 {
					continue
				}
				current.next = &node{val: S2i(st)}
				current = current.next
			}

			for i := 0; i < 25; i++ {
				current = list.head
				for current != nil {
					if current.val == 0 {
						current.val = 1
						current = current.next
						continue
					}
					stringStone := strconv.Itoa(current.val)
					if len(stringStone)%2 == 0 {
						a := stringStone[len(stringStone)/2:]
						b := stringStone[:len(stringStone)/2]
						current.val = S2i(a)
						newNode := &node{
							val:  S2i(b),
							next: current.next,
						}
						current.next = newNode
						current = current.next.next
						continue
					}
					current.val *= 2024
					current = current.next
				}
			}

			current = list.head
			for current != nil {
				current = current.next
				answer++
			}

			return answer
	*/
}

func part2(lines []string) int {
	return run(lines[0], 75)
}

func run(line string, length int) int {
	stoneMap := map[int]int{}
	s := strings.Fields(line)
	for _, st := range s {
		stoneMap[S2i(st)]++
	}

	answer := len(s)

	for i := 0; i < length; i++ {
		newStoneMap := make(map[int]int, len(stoneMap))
		for stone, count := range stoneMap {
			if count <= 0 {
				continue
			}
			stringStone := strconv.Itoa(stone)
			if stone == 0 {
				newStoneMap[1] += count
			} else if len(stringStone)%2 == 0 {
				a := stringStone[len(stringStone)/2:]
				b := stringStone[:len(stringStone)/2]
				newStoneMap[S2i(a)] += count
				newStoneMap[S2i(b)] += count
				answer += count
			} else {
				newStoneMap[stone*2024] += count
			}
		}
		stoneMap = newStoneMap
	}

	return answer
}
