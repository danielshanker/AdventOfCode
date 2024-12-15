package main

import (
	"flag"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 15, part1, part2, 10092, 9021)

}

func part1(lines []string) int {
	answer := 0

	warehouse := map[int]map[int]string{}
	instructions := ""
	flip := false
	botPos := Coord{}

	for y, line := range lines {
		if line == "" {
			flip = true
		}
		if flip {
			instructions += line
			continue
		}
		for x, r := range line {
			if _, ok := warehouse[x]; !ok {
				warehouse[x] = map[int]string{}
			}
			warehouse[x][y] = string(r)
			if string(r) == "@" {
				botPos = Coord{
					X: x,
					Y: y,
				}
			}
		}
	}

	for _, r := range instructions {
		inst := string(r)
		moved := false
		warehouse, moved = move(warehouse, botPos, inst)
		if moved {
			if inst == "^" {
				botPos.Y--
			}
			if inst == "v" {
				botPos.Y++
			}
			if inst == ">" {
				botPos.X++
			}
			if inst == "<" {
				botPos.X--
			}
		}
	}
	for y := 0; y < len(warehouse[0]); y++ {
		for x := 0; x < len(warehouse); x++ {
			if warehouse[x][y] == "O" {
				answer += 100*y + x
			}
		}
	}

	return answer
}

func move(warehouse map[int]map[int]string, pos Coord, dir string) (map[int]map[int]string, bool) {
	if warehouse[pos.X][pos.Y] == "." || warehouse[pos.X][pos.Y] == "#" {
		return warehouse, false
	}
	moved := false
	if dir == "^" {
		newPos := Coord{X: pos.X, Y: pos.Y - 1}
		warehouse, _ = move(warehouse, newPos, dir)
		if warehouse[newPos.X][newPos.Y] == "." {
			moved = true
			warehouse[pos.X][pos.Y-1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
	}
	if dir == "v" {
		newPos := Coord{X: pos.X, Y: pos.Y + 1}
		warehouse, _ = move(warehouse, newPos, dir)
		if warehouse[newPos.X][newPos.Y] == "." {
			moved = true
			warehouse[pos.X][pos.Y+1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
	}
	if dir == ">" {
		newPos := Coord{X: pos.X + 1, Y: pos.Y}
		warehouse, _ = move(warehouse, newPos, dir)
		if warehouse[newPos.X][newPos.Y] == "." {
			moved = true
			warehouse[pos.X+1][pos.Y] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
	}
	if dir == "<" {
		newPos := Coord{X: pos.X - 1, Y: pos.Y}
		warehouse, _ = move(warehouse, newPos, dir)
		if warehouse[newPos.X][newPos.Y] == "." {
			moved = true
			warehouse[pos.X-1][pos.Y] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
	}

	return warehouse, moved
}

func part2(lines []string) int {
	answer := 0

	warehouse := map[int]map[int]string{}
	instructions := ""
	flip := false
	botPos := Coord{}

	for y, line := range lines {
		if line == "" {
			flip = true
		}
		if flip {
			instructions += line
			continue
		}
		x := 0
		for _, r := range line {
			if _, ok := warehouse[x]; !ok {
				warehouse[x] = map[int]string{}
				warehouse[x+1] = map[int]string{}
			}
			if string(r) == "O" {
				warehouse[x][y] = "["
				warehouse[x+1][y] = "]"
			}
			if string(r) == "#" {
				warehouse[x][y] = "#"
				warehouse[x+1][y] = "#"
			}
			if string(r) == "." {
				warehouse[x][y] = "."
				warehouse[x+1][y] = "."
			}
			if string(r) == "@" {
				warehouse[x][y] = "@"
				warehouse[x+1][y] = "."
				botPos = Coord{
					X: x,
					Y: y,
				}
			}
			x += 2
		}
	}

	for _, r := range instructions {
		inst := string(r)
		moved := false
		if inst == ">" || inst == "<" {
			warehouse, moved = move(warehouse, botPos, inst)
		} else {
			if isMovable(warehouse, botPos, inst, &map[Coord]bool{}) {
				warehouse = move2(warehouse, botPos, inst, &map[Coord]bool{})
				moved = true
			}
		}
		if moved {
			if inst == "^" {
				botPos.Y--
			}
			if inst == "v" {
				botPos.Y++
			}
			if inst == ">" {
				botPos.X++
			}
			if inst == "<" {
				botPos.X--
			}
		}
		//for y := 0; y < len(warehouse[0]); y++ {
		//	for x := 0; x < len(warehouse); x++ {
		//		fmt.Print(warehouse[x][y])
		//	}
		//	fmt.Println()
		//}
	}

	//for y := 0; y < len(warehouse[0]); y++ {
	//	for x := 0; x < len(warehouse); x++ {
	//		fmt.Print(warehouse[x][y])
	//	}
	//	fmt.Println()
	//}

	for y := 0; y < len(warehouse[0]); y++ {
		for x := 0; x < len(warehouse); x++ {
			if warehouse[x][y] == "[" {
				answer += 100*y + x
			}
		}
	}
	return answer
}

func isMovable(warehouse map[int]map[int]string, pos Coord, dir string, seenMap *map[Coord]bool) bool {
	sm := *seenMap
	if sm[pos] {
		return true
	}
	sm[pos] = true
	seenMap = &sm

	if warehouse[pos.X][pos.Y] == "#" {
		return false
	}
	if warehouse[pos.X][pos.Y] == "." {
		return true
	}

	if warehouse[pos.X][pos.Y] == "[" {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		newPosRight := Coord{X: pos.X + 1, Y: pos.Y}
		if dir == "^" {
			if !isMovable(warehouse, newPosUp, dir, seenMap) {
				return false
			}
		} else {
			if !isMovable(warehouse, newPosDown, dir, seenMap) {
				return false
			}
		}
		if !isMovable(warehouse, newPosRight, dir, seenMap) {
			return false
		}
	} else if warehouse[pos.X][pos.Y] == "]" {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		newPosLeft := Coord{X: pos.X - 1, Y: pos.Y}
		if dir == "^" {
			if !isMovable(warehouse, newPosUp, dir, seenMap) {
				return false
			}
		} else {
			if !isMovable(warehouse, newPosDown, dir, seenMap) {
				return false
			}
		}
		if !isMovable(warehouse, newPosLeft, dir, seenMap) {
			return false
		}
	} else {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		if dir == "^" {
			if !isMovable(warehouse, newPosUp, dir, seenMap) {
				return false
			}
		} else {
			if !isMovable(warehouse, newPosDown, dir, seenMap) {
				return false
			}
		}
	}

	return true
}

func move2(warehouse map[int]map[int]string, pos Coord, dir string, seenMap *map[Coord]bool) map[int]map[int]string {
	sm := *seenMap
	if sm[pos] {
		return warehouse
	}
	sm[pos] = true
	seenMap = &sm

	if warehouse[pos.X][pos.Y] == "#" {
		return warehouse
	}
	if warehouse[pos.X][pos.Y] == "." {
		return warehouse
	}

	if warehouse[pos.X][pos.Y] == "[" {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		newPosRight := Coord{X: pos.X + 1, Y: pos.Y}
		if dir == "^" {
			move2(warehouse, newPosUp, dir, seenMap)
			warehouse[pos.X][pos.Y-1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		} else {
			move2(warehouse, newPosDown, dir, seenMap)
			warehouse[pos.X][pos.Y+1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
		move2(warehouse, newPosRight, dir, seenMap)
	} else if warehouse[pos.X][pos.Y] == "]" {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		newPosLeft := Coord{X: pos.X - 1, Y: pos.Y}
		if dir == "^" {
			move2(warehouse, newPosUp, dir, seenMap)
			warehouse[pos.X][pos.Y-1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		} else {
			move2(warehouse, newPosDown, dir, seenMap)
			warehouse[pos.X][pos.Y+1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}
		move2(warehouse, newPosLeft, dir, seenMap)
	} else {
		newPosUp := Coord{X: pos.X, Y: pos.Y - 1}
		newPosDown := Coord{X: pos.X, Y: pos.Y + 1}
		if dir == "^" {
			move2(warehouse, newPosUp, dir, seenMap)
			warehouse[pos.X][pos.Y-1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		} else {
			move2(warehouse, newPosDown, dir, seenMap)
			warehouse[pos.X][pos.Y+1] = warehouse[pos.X][pos.Y]
			warehouse[pos.X][pos.Y] = "."
		}

	}

	return warehouse
}
