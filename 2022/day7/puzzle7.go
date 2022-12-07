package main

import (
	"flag"
	"fmt"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day7/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2022/day7/input.txt")
		part1(input)
		part2(input)
	}

}

type file struct {
	name string
	size int
}

type directory struct {
	dirs     map[string]*directory
	files    []file
	upperDir *directory
	name     string
	dirSize  int
}

func part1(lines []string) {
	answer1 := 0
	rootDir := directory{
		name: "/",
		dirs: make(map[string]*directory),
	}

	var curDir *directory
	inList := false

	for _, line := range lines {
		command := strings.Split(line, " ")

		if inList {
			if command[0] == "$" {
				inList = false
			} else {
				if command[0] == "dir" {
					newDir := directory{
						name:     command[1],
						upperDir: curDir,
						dirs:     make(map[string]*directory),
					}
					curDir.dirs[command[1]] = &newDir
				} else {
					newFile := file{
						size: S2i(command[0]),
						name: command[1],
					}
					curDir.files = append(curDir.files, newFile)
				}
			}
		}

		if command[0] == "$" {
			if command[1] == "cd" {
				if command[2] == "/" {
					curDir = &rootDir
				} else if command[2] == ".." {
					curDir = curDir.upperDir
				} else {
					if _, ok := curDir.dirs[command[2]]; ok {
						curDir = curDir.dirs[command[2]]
					} else {
						newDir := directory{
							name:     command[2],
							upperDir: curDir,
							dirs:     make(map[string]*directory),
						}
						curDir.dirs[command[2]] = &newDir
						curDir = curDir.dirs[command[2]]
					}
				}
			}
			if command[1] == "ls" {
				inList = true
				continue
			}
		}
	}

	getDirSize(&rootDir, &answer1)

	fmt.Printf("Answer 1 : %d\n", answer1)
}

func getDirSize(dir *directory, answer *int) int {
	size := 0
	for _, f := range dir.files {
		size += f.size
	}

	for _, d := range dir.dirs {
		size += getDirSize(d, answer)
	}

	if size <= 100000 {
		*answer += size
	}
	dir.dirSize = size
	return size
}

func part2(lines []string) {
	answer2 := 0

	rootDir := directory{
		name: "/",
		dirs: make(map[string]*directory),
	}

	var curDir *directory
	inList := false

	for _, line := range lines {
		command := strings.Split(line, " ")

		if inList {
			if command[0] == "$" {
				inList = false
			} else {
				if command[0] == "dir" {
					newDir := directory{
						name:     command[1],
						upperDir: curDir,
						dirs:     make(map[string]*directory),
					}
					curDir.dirs[command[1]] = &newDir
				} else {
					newFile := file{
						size: S2i(command[0]),
						name: command[1],
					}
					curDir.files = append(curDir.files, newFile)
				}
			}
		}

		if command[0] == "$" {
			if command[1] == "cd" {
				if command[2] == "/" {
					curDir = &rootDir
				} else if command[2] == ".." {
					curDir = curDir.upperDir
				} else {
					if _, ok := curDir.dirs[command[2]]; ok {
						curDir = curDir.dirs[command[2]]
					} else {
						newDir := directory{
							name:     command[2],
							upperDir: curDir,
							dirs:     make(map[string]*directory),
						}
						curDir.dirs[command[2]] = &newDir
						curDir = curDir.dirs[command[2]]
					}
				}
			}
			if command[1] == "ls" {
				inList = true
				continue
			}
		}
	}

	temp := 0
	getDirSize(&rootDir, &temp)

	freeSpace := 70000000 - rootDir.dirSize
	neededSpace := 30000000 - freeSpace
	var dirSizes []int
	traverse(&rootDir, neededSpace, &dirSizes)
	answer2 = 30000000

	for _, s := range dirSizes {
		if s < answer2 {
			answer2 = s
		}
	}

	fmt.Printf("Answer 2 : %d\n", answer2)
}

func traverse(dir *directory, neededSpace int, answer *[]int) {
	size := dir.dirSize

	for _, d := range dir.dirs {
		traverse(d, neededSpace, answer)
	}

	if size > neededSpace {
		*answer = append(*answer, size)
	}
}
