package main

import (
	"flag"
	"fmt"
	"slices"
	"sort"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 23, part1, part2, 7, 0)

}

type connection struct {
	compA string
	compB string
}

type network struct {
	connections map[string]bool
}

func part1(lines []string) int {
	answer := 0
	connections := make([]connection, len(lines))
	for i, line := range lines {
		con := strings.Split(line, "-")
		connections[i] = connection{
			compA: con[0],
			compB: con[1],
		}
	}

	triangles := map[string]bool{}
	for i := 0; i < len(connections); i++ {
		connection1 := connections[i]
		for j := 0; j < len(connections); j++ {
			connection2 := connections[j]
			if j == i {
				continue
			}
			thirdComp := ""
			if connection1.compA == connection2.compA {
				thirdComp = connection2.compB
			} else if connection1.compA == connection2.compB {
				thirdComp = connection2.compA
			} else if connection1.compB == connection2.compA {
				thirdComp = connection2.compB
			} else if connection1.compB == connection2.compB {
				thirdComp = connection2.compA
			} else {
				continue
			}

			if string(connection1.compA[0]) != "t" && string(connection1.compB[0]) != "t" && string(connection2.compA[0]) != "t" && string(connection2.compB[0]) != "t" {
				continue
			}

			for k := 0; k < len(connections); k++ {
				if k == i || k == j {
					continue
				}
				connection3 := connections[k]
				lookFor := ""
				if connection3.compA == thirdComp {
					lookFor = connection3.compB
				} else if connection3.compB == thirdComp {
					lookFor = connection3.compA
				} else {
					continue
				}

				if connection1.compA == lookFor || connection1.compB == lookFor {
					comps := []string{connection2.compA, connection2.compB, lookFor}
					slices.Sort(comps)

					compString := strings.Join(comps, "-")
					triangles[compString] = true
				}
			}
		}
	}

	answer = len(triangles)

	return answer
}

func buildGraph(connections []connection) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)
	for _, conn := range connections {
		if _, ok := graph[conn.compA]; !ok {
			graph[conn.compA] = make(map[string]bool)
		}
		if _, ok := graph[conn.compB]; !ok {
			graph[conn.compB] = make(map[string]bool)
		}
		graph[conn.compA][conn.compB] = true
		graph[conn.compB][conn.compA] = true
	}
	return graph
}

func isClique(c []string, v string, graph map[string]map[string]bool) bool {
	for _, node := range c {
		if !graph[node][v] {
			return false
		}
	}
	return true
}

func findLargestClique(graph map[string]map[string]bool, nodes []string, clique []string, largestClique *[]string) {
	if len(clique) > len(*largestClique) {
		*largestClique = append([]string(nil), clique...)
	}

	for i, node := range nodes {
		if isClique(clique, node, graph) {
			newClique := append(clique, node)

			findLargestClique(graph, nodes[i+1:], newClique, largestClique)
		}
	}
}

func largestClique(connections []connection) []string {
	graph := buildGraph(connections)

	var nodes []string
	for node := range graph {
		nodes = append(nodes, node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return len(graph[nodes[i]]) > len(graph[nodes[j]])
	})

	var largestClique []string

	findLargestClique(graph, nodes, []string{}, &largestClique)

	return largestClique
}

func part2(lines []string) int {
	connections := make([]connection, len(lines))
	for i, line := range lines {
		con := strings.Split(line, "-")
		connections[i] = connection{
			compA: con[0],
			compB: con[1],
		}
	}

	largest := largestClique(connections)
	slices.Sort(largest)
	fmt.Println(strings.Join(largest, ","))
	return 0
}
