package main

import (
	"flag"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 9, part1, part2, 1928, 2858)

}

type point struct {
	id  int
	val int
}

type file struct {
	size    int
	startID int
	val     int
}

func part1(lines []string) int {
	answer := 0

	line := lines[0]
	id := 0
	pos := 0
	charPoints := []point{}
	freePoints := []point{}

	for i, r := range line {
		rep := S2i(string(r))
		if i%2 == 0 {
			for j := 0; j < rep; j++ {
				p := point{
					id:  pos,
					val: id,
				}
				charPoints = append(charPoints, p)
				pos++
			}
			id++
		} else {
			for j := 0; j < rep; j++ {
				p := point{
					id:  pos,
					val: 0,
				}
				freePoints = append(freePoints, p)
				pos++
			}
		}
	}

	for i, p := range freePoints {
		var last point
		if charPoints[len(charPoints)-1].id < p.id {
			continue
		}
		charPoints, last = pop(charPoints)
		last.id = p.id
		freePoints[i] = last
	}

	charPoints = append(charPoints, freePoints...)

	for _, val := range charPoints {
		answer += val.id * val.val
	}

	return answer
}

func pop(p []point) ([]point, point) {
	if len(p) == 0 {
		return p, point{}
	}
	last := (p)[len(p)-1]
	p = (p)[:len(p)-1]
	return p, last
}

func part2(lines []string) int {
	answer := 0

	line := lines[0]
	id := 0
	pos := 0
	fullFiles := []file{}
	emptyFiles := []file{}

	for i, r := range line {
		rep := S2i(string(r))
		if i%2 == 0 {
			f := file{
				size:    rep,
				startID: pos,
				val:     id,
			}
			pos += rep
			fullFiles = append(fullFiles, f)
			id++
		} else {
			f := file{
				size:    rep,
				startID: pos,
				val:     0,
			}
			pos += rep
			emptyFiles = append(emptyFiles, f)
		}
	}

	for j := len(fullFiles) - 1; j >= 0; j-- {
		for i, e := range emptyFiles {
			f := fullFiles[j]
			if fullFiles[j].startID < e.startID {
				continue
			}
			if f.size <= e.size && e.val == 0 {
				sizeDiff := e.size - f.size
				e.val = f.val
				e.size = f.size
				emptyFiles[i] = e
				fullFiles = append(fullFiles[:j], fullFiles[j+1:]...)
				if sizeDiff > 0 {
					nf := file{
						val:     0,
						size:    sizeDiff,
						startID: e.startID + e.size,
					}
					emptyFiles = append(emptyFiles[:i+1], emptyFiles[i:]...)
					emptyFiles[i] = nf
				}
				break
			}
		}
	}

	fullFiles = append(fullFiles, emptyFiles...)
	for _, f := range fullFiles {
		for i := f.startID; i < f.startID+f.size; i++ {
			answer += i * f.val
		}
	}

	return answer
}
