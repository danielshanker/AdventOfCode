package main

import (
	"flag"
	"strings"
	. "utils"
)

var test *bool

func main() {
	test = flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 11, part1, part2, 5, 2)

}

type server struct {
	id          string
	connections []*server
	path        map[string]int
	downs       int
}

func part1(lines []string) int {
	answer := 0

	serverMap := make(map[string]*server)
	var start *server
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		id := splitLine[0]
		cons := strings.Fields(splitLine[1])
		var s *server
		if _, ok := serverMap[id]; ok {
			s = serverMap[id]
		} else {
			s = &server{
				id: id,
			}
		}
		for _, cs := range cons {
			if curServer, ok := serverMap[cs]; ok {
				s.connections = append(s.connections, curServer)
			} else {
				curServer := &server{id: cs}
				s.connections = append(s.connections, curServer)
				serverMap[cs] = curServer
			}
		}
		serverMap[id] = s
		if id == "you" {
			start = s
		}
	}
	if start == nil {
		return 0
	}

	answer = travel(start)

	return answer
}

func part2(lines []string) int {
	answer := 0
	serverMap := make(map[string]*server)
	var start, fft, dac *server
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		id := splitLine[0]
		cons := strings.Fields(splitLine[1])
		var s *server
		if _, ok := serverMap[id]; ok {
			s = serverMap[id]
		} else {
			s = &server{
				id:    id,
				path:  make(map[string]int),
				downs: -1,
			}
		}
		for _, cs := range cons {
			if curServer, ok := serverMap[cs]; ok {
				s.connections = append(s.connections, curServer)
			} else {
				curServer := &server{id: cs, path: make(map[string]int), downs: -1}
				s.connections = append(s.connections, curServer)
				serverMap[cs] = curServer
			}
		}
		serverMap[id] = s
		if id == "svr" {
			start = s
		}
		if id == "fft" {
			fft = s
		}
		if id == "dac" {
			dac = s
		}
	}

	answer = travel2(dac, "out")
	clearDowns(serverMap)
	answer *= travel2(fft, "dac")
	clearDowns(serverMap)
	answer *= travel2(start, "fft")

	return answer
}

func travel(curServer *server) int {
	if curServer.id == "out" {
		return 1
	}
	if curServer.downs > 0 {
		return curServer.downs
	}
	world := 0
	for _, con := range curServer.connections {
		world += travel(con)
	}
	curServer.downs = world
	return world
}

func travel2(curServer *server, target string) int {
	if curServer.id == target {
		return 1
	}
	if curServer.downs >= 0 {
		return curServer.downs
	}
	world := 0
	for _, con := range curServer.connections {
		world += travel2(con, target)
	}
	curServer.downs = world
	return world
}

func clearDowns(servers map[string]*server) {
	for _, s := range servers {
		s.downs = -1
	}
}
