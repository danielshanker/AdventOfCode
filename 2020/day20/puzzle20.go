package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

//THIS CODE IS DISGUSTING, IF YOU CAN READ THIS, PLEASE DON'T JUDGE ME
func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")
	var tiles []tile

	i := 0
	for _, line := range lines {
		if line == "" {
			i++
			continue
		}
		if len(tiles) <= i {
			var emptyTile tile
			tiles = append(tiles, emptyTile)
		}
		re := regexp.MustCompile("Tile (\\d+)")
		if re.MatchString(line) {
			curID := re.FindStringSubmatch(line)
			tiles[i].id = s2i(curID[1])
			continue
		}
		tiles[i].body = append(tiles[i].body, line)
	}

	for i, curTile := range tiles {
		var newBody []string
		tiles[i].top = curTile.body[0]
		tiles[i].bottom = curTile.body[len(curTile.body)-1]
		for index, j := range curTile.body {
			tiles[i].left += string(j[0])
			tiles[i].right += string(j[len(curTile.body)-1])
			if index != 0 && index != len(curTile.body)-1 {
				line := j[1:]
				line = line[0 : len(line)-1]
				newBody = append(newBody, line)
			}
		}
		tiles[i].body = newBody
	}

	var grid [12][12]tile
	//var grid [3][3]tile
	var corners []int
	for _, curTile := range tiles {
		sideMatches := 0
		for _, testTile := range tiles {
			if testTile.id == curTile.id {
				continue
			}
			checkTile := testTile
			matchFound := false
			for i := 0; i < 4; i++ {
				checkTile = rotate(checkTile)
				for j := 0; j < 2; j++ {
					checkTile = flip(checkTile, true)
					side := checkFit(curTile, checkTile)
					if side != NONE {
						curTile.sides = append(curTile.sides, side)
						matchFound = true
						break
					}
					checkTile = flip(checkTile, false)
					side = checkFit(curTile, checkTile)
					if side != NONE {
						curTile.sides = append(curTile.sides, side)
						matchFound = true
						break
					}
				}
				if matchFound {
					break
				}
			}
			if matchFound {
				sideMatches++
				continue
			}
		}

		if sideMatches == 2 {
			corners = append(corners, curTile.id)
			if len(corners) == 1 {
				newTile := curTile
				if curTile.sides[0] == T || curTile.sides[1] == T {
					newTile = rotate(newTile)
					if curTile.sides[0] == L || curTile.sides[1] == L {
						newTile = rotate(newTile)
					}
				} else if curTile.sides[0] == L || curTile.sides[1] == L {
					newTile = flip(newTile, true)
				}
				newTile.id = curTile.id
				grid[0][0] = newTile
			}
		}
	}

	usedTiles := make(map[int]bool)
	usedTiles[grid[0][0].id] = true
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			var curTile tile
			if j == len(grid)-1 {
				curTile = grid[i][0]
			} else {
				curTile = grid[i][j]
			}
			for _, testTile := range tiles {
				if _, ok := usedTiles[testTile.id]; ok {
					continue
				}
				checkTile := testTile
				matchFound := false
				for a := 0; a < 2; a++ {
					for k := 0; k < 4; k++ {
						checkTile = rotate(checkTile)
						if j != len(grid)-1 {
							if curTile.right == checkTile.left {
								matchFound = true
								grid[i][j+1] = checkTile
								break
							}
						} else {
							if curTile.bottom == checkTile.top {
								matchFound = true
								grid[i+1][0] = checkTile
								break
							}
						}
					}
					checkTile = flip(checkTile, true)
					if matchFound {
						break
					}
				}
				if matchFound {
					usedTiles[testTile.id] = true
					break
				}
			}
		}
	}

	var seaMap [96]string
	k := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			for l := 0; l < 8; l++ {
				seaMap[k+l] += grid[i][j].body[l]
			}
		}
		k += 8
	}
	regex := "#.{4}##.{4}##.{4}###"
	regexBottom := ".#..#..#..#..#..#"
	regexT := "#.{4}##.{4}##.{4}###"
	re := regexp.MustCompile(regex)
	re2 := regexp.MustCompile(regexBottom)

	seaMonsters := 0
	for {
		for i := 1; i < len(seaMap)-1; i++ {
			matches := re.FindAllStringIndex(seaMap[i], -1)
			for _, match := range matches {
				if string(seaMap[i-1][match[1]-1]) == "#" {
					matches2 := re2.FindAllStringIndex(seaMap[i+1], -1)
					for _, match2 := range matches2 {
						if match2[0] == match[0] {
							seaMonsters++
							break
						}
					}
				}
			}
		}
		if seaMonsters != 0 {
			break
		}
		seaMap = flipt(seaMap, true)
		for i := 1; i < len(seaMap)-1; i++ {
			matches := re.FindAllStringIndex(seaMap[i], -1)
			for _, match := range matches {
				if string(seaMap[i-1][match[1]-1]) == "#" {
					matches2 := re2.FindAllStringIndex(seaMap[i+1], -1)
					for _, match2 := range matches2 {
						if match2[0] == match[0] {
							seaMonsters++
							break
						}
					}
				}
			}
		}
		if seaMonsters != 0 {
			break
		}
		seaMap = flipt(seaMap, false)
		for i := 1; i < len(seaMap)-1; i++ {
			matches := re.FindAllStringIndex(seaMap[i], -1)
			for _, match := range matches {
				if string(seaMap[i-1][match[1]-1]) == "#" {
					matches2 := re2.FindAllStringIndex(seaMap[i+1], -1)
					for _, match2 := range matches2 {
						if match2[0] == match[0] {
							seaMonsters++
							break
						}
					}
				}
			}
		}
		if seaMonsters != 0 {
			break
		}
		seaMap = flipt(seaMap, true)
		for i := 1; i < len(seaMap)-1; i++ {
			matches := re.FindAllStringIndex(seaMap[i], -1)
			for _, match := range matches {
				if string(seaMap[i-1][match[1]-1]) == "#" {
					matches2 := re2.FindAllStringIndex(seaMap[i+1], -1)
					for _, match2 := range matches2 {
						if match2[0] == match[0] {
							seaMonsters++
							break
						}
					}
				}
			}
		}
		if seaMonsters != 0 {
			break
		}
		seaMap = flipt(seaMap, true)
		seaMap = rotatet(seaMap)
	}

	reT := regexp.MustCompile(regexT)

	seaMapText := ""
	for _, i := range seaMap {
		seaMapText += i + "\n"

	}

	a := reT.FindAllStringIndex(seaMapText, -1)

	seaMonsters = 0
	for range a {
		seaMonsters++
	}

	hashCount := 0
	for _, i := range seaMap {
		for _, j := range i {
			if string(j) == "#" {
				hashCount++
			}
		}
	}
	answer2 := hashCount - (seaMonsters * 15)

	answer1 := 1
	for _, i := range corners {
		answer1 *= i
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func checkFit(curTile tile, testTile tile) string {
	if (curTile.top) == testTile.bottom {
		return T
	}
	if (curTile.bottom) == testTile.top {
		return B
	}
	if (curTile.left) == testTile.right {
		return L
	}
	if (curTile.right) == testTile.left {
		return R
	}
	return NONE
}

func rotate(curTile tile) tile {
	var newTile tile

	top := curTile.top
	bottom := curTile.bottom
	left := curTile.left
	right := curTile.right

	newTile.right = top
	newTile.bottom = reverse(right)
	newTile.left = bottom
	newTile.top = reverse(left)
	newTile.id = curTile.id
	var newBody []string
	for i := 0; i < len(curTile.body); i++ {
		newLine := ""
		for j := 0; j < len(curTile.body); j++ {
			newLine += string(curTile.body[len(curTile.body)-j-1][i])
		}
		newBody = append(newBody, newLine)
	}
	newTile.body = newBody

	return newTile
}

func flip(curTile tile, horizontal bool) tile {
	top := curTile.top
	bottom := curTile.bottom
	left := curTile.left
	right := curTile.right
	var newTile tile
	if horizontal {
		newTile.right = left
		newTile.left = right
		newTile.bottom = reverse(bottom)
		newTile.top = reverse(top)
		for _, i := range curTile.body {
			line := reverse(i)
			newTile.body = append(newTile.body, line)
		}
	} else {
		newTile.right = reverse(right)
		newTile.left = reverse(left)
		newTile.bottom = top
		newTile.top = bottom
		for i := len(curTile.body) - 1; i >= 0; i-- {
			newTile.body = append(newTile.body, curTile.body[i])
		}
	}
	return newTile
}

func flipt(curTile [96]string, horizontal bool) [96]string {
	var newTile [96]string
	if horizontal {
		for j, i := range curTile {
			line := reverse(i)
			newTile[j] = line
		}
	} else {
		for i := len(curTile) - 1; i >= 0; i-- {
			newTile[len(curTile)-i-1] = curTile[i]
		}
	}
	return newTile
}
func rotatet(curTile [96]string) [96]string {
	var newBody [96]string
	for i := 0; i < len(curTile); i++ {
		newLine := ""
		for j := 0; j < len(curTile); j++ {
			newLine += string(curTile[len(curTile)-j-1][i])
		}
		newBody[i] = newLine
	}

	return newBody
}

func reverse(r string) string {
	newString := ""
	for i := len(r) - 1; i >= 0; i-- {
		newString += string(r[i])
	}
	return newString
}

type tile struct {
	left   string
	right  string
	bottom string
	top    string
	id     int
	body   []string
	sides  []string
}

const (
	T    = "t"
	L    = "l"
	B    = "b"
	R    = "r"
	NONE = "n"
)

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
