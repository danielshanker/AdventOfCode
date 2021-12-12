package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var doubleVisited bool

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day12/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day12/input.txt")
		part1(input)
		part2(input)
	}

}

type Node struct {
	next    []*Node
	value   string
	visited bool
}

func part1(lines []string) {
	answer1 := 0
	nodes := generateNodes(lines)
	nodes = makeConnections(nodes, lines)

	traverse1(findNode(nodes, "start"))
	answer1 = count
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0
	count = 0
	nodes := generateNodes(lines)
	nodes = makeConnections(nodes, lines)

	traverse2(findNode(nodes, "start"))
	answer2 = count

	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}
func traverse2(node *Node) {
	if node.value == "end" {
		count++
		return
	}
	dv := false
	if node.visited {
		doubleVisited = true
		dv = true
	}
	if strings.ToUpper(node.value) != node.value {
		node.visited = true
	}
	for _, nextNode := range node.next {
		if nextNode.value == "start" {
			continue
		}
		if nextNode.visited {
			if doubleVisited {
				continue
			}
		}
		traverse2(nextNode)
	}
	if dv {
		doubleVisited = false
	} else {
		node.visited = false
	}
}

func traverse1(node *Node) {
	if node.value == "end" {
		count++
		return
	}
	if strings.ToUpper(node.value) != node.value {
		node.visited = true
	}
	for _, nextNode := range node.next {
		if nextNode.visited {
			continue
		}
		traverse1(nextNode)
	}
	node.visited = false
}

func makeConnections(nodes []*Node, lines []string) []*Node {
	for _, line := range lines {
		connection := strings.Split(line, "-")
		node1 := findNode(nodes, connection[0])
		node2 := findNode(nodes, connection[1])
		node1.next = append(node1.next, node2)
		node2.next = append(node2.next, node1)
	}
	return nodes
}

func generateNodes(lines []string) []*Node {
	var nodes []*Node
	for _, line := range lines {
		connection := strings.Split(line, "-")
		if !nodeExists(nodes, connection[0]) {
			n := Node{
				value: connection[0],
			}
			nodes = append(nodes, &n)
		}
		if !nodeExists(nodes, connection[1]) {
			n := Node{
				value: connection[1],
			}
			nodes = append(nodes, &n)
		}
	}
	return nodes
}
func nodeExists(nodes []*Node, node string) bool {
	for _, n := range nodes {
		if node == n.value {
			return true
		}
	}
	return false
}

func findNode(nodes []*Node, node string) *Node {
	for _, n := range nodes {
		if node == n.value {
			return n
		}
	}
	return nil
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
