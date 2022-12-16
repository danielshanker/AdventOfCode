package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	. "utils"

	"github.com/ernestosuarez/itertools"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day16/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day16/input.txt")
		part1(input)
		part2(input)
	}

}

type valve struct {
	name       string
	rate       int
	leads      []*valve
	costs      map[string]int
	leadString string
	visited    bool
	cost       int
}

func part1(lines []string) {
	answer := 0
	rName := regexp.MustCompile("Valve ([A-Z][A-Z])")
	rRate := regexp.MustCompile("rate=([0-9]+)")
	rLeads := regexp.MustCompile("to valves? (.*)")
	valves := map[string]*valve{}

	for _, line := range lines {
		name := rName.FindStringSubmatch(line)[1]
		rate := rRate.FindStringSubmatch(line)[1]
		leads := rLeads.FindStringSubmatch(line)[1]
		v := valve{
			name:       name,
			rate:       S2i(rate),
			leadString: leads,
		}
		valves[name] = &v
	}

	for _, v := range valves {
		s := strings.Split(v.leadString, ", ")
		for _, lValve := range s {
			v.leads = append(v.leads, valves[lValve])
		}
	}

	for _, start := range valves {
		start.costs = map[string]int{}
		q := queue{}
		start.visited = true
		q.push(*start)
		for len(q) != 0 {
			curP := q.pop()
			for _, p := range curP.leads {
				if !p.visited {
					p.visited = true
					p.cost = curP.cost + 1
					q.push(*p)
					if p.rate != 0 {
						start.costs[p.name] = p.cost
					}
				}
			}
		}
		for _, v := range valves {
			v.visited = false
			v.cost = 0
		}
	}

	names := []string{}
	for _, v := range valves {
		if v.rate != 0 {
			names = append(names, v.name)
		}
	}

	answer = dfs(valves, valves["AA"], 0, 0, 0, names, 30)

	fmt.Printf("Answer 1 : %d\n", answer)
}

func dfs(valves map[string]*valve, current *valve, currentTime int, currentFlow int, currentPressure int, remaining []string, totalTime int) int {

	score := currentPressure + (totalTime-currentTime)*currentFlow
	max := score

	for _, name := range remaining {
		d := current.costs[name] + 1
		if currentTime+d < totalTime {
			newTime := currentTime + d
			newFlow := currentFlow + valves[name].rate
			newPressure := currentPressure + d*currentFlow
			new := []string{}
			for _, i := range remaining {
				if i != name {
					new = append(new, i)
				}
			}

			possibeScore := dfs(valves, valves[name], newTime, newFlow, newPressure, new, totalTime)
			if possibeScore > max {
				max = possibeScore
			}
		}
	}

	return max

}

func part2(lines []string) {
	answer := 0
	rName := regexp.MustCompile("Valve ([A-Z][A-Z])")
	rRate := regexp.MustCompile("rate=([0-9]+)")
	rLeads := regexp.MustCompile("to valves? (.*)")
	valves := map[string]*valve{}

	for _, line := range lines {
		name := rName.FindStringSubmatch(line)[1]
		rate := rRate.FindStringSubmatch(line)[1]
		leads := rLeads.FindStringSubmatch(line)[1]
		v := valve{
			name:       name,
			rate:       S2i(rate),
			leadString: leads,
		}
		valves[name] = &v
	}

	for _, v := range valves {
		s := strings.Split(v.leadString, ", ")
		for _, lValve := range s {
			v.leads = append(v.leads, valves[lValve])
		}
	}

	for _, start := range valves {
		start.costs = map[string]int{}
		q := queue{}
		start.visited = true
		q.push(*start)
		for len(q) != 0 {
			curP := q.pop()
			for _, p := range curP.leads {
				if !p.visited {
					p.visited = true
					p.cost = curP.cost + 1
					q.push(*p)
					if p.rate != 0 {
						start.costs[p.name] = p.cost
					}
				}
			}
		}
		for _, v := range valves {
			v.visited = false
			v.cost = 0
		}
	}

	names := []string{}
	for _, v := range valves {
		if v.rate != 0 {
			names = append(names, v.name)
		}
	}

	for i := 1; i < len(names); i++ {
		for n := range itertools.CombinationsStr(names, i) {
			left := []string{}
			for _, j := range names {
				if !contains(n, j) {
					left = append(left, j)
				}
			}

			you := dfs(valves, valves["AA"], 0, 0, 0, n, 26)
			elephant := dfs(valves, valves["AA"], 0, 0, 0, left, 26)
			if you+elephant > answer {
				answer = you + elephant
			}
		}
	}

	fmt.Printf("Answer 2 : %d\n", answer)
}

func contains(a []string, b string) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}

type queue []valve

func (q *queue) pop() valve {
	if len(*q) == 0 {
		return valve{}
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q *queue) push(v valve) {
	*q = append(*q, v)
}
