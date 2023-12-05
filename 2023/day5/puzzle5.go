package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day5/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := ReadInputLines("/home/daniel.shanker/Pers/AdventOfCode/2023/day5/input.txt")
		part1(input)
		part2(input)
	}

}

type numRange struct {
	start int
	end   int
	moved bool
}

func part1(lines []string) {
	answer := 0
	seeds := []int{}
	for _, val := range strings.Fields(lines[0]) {
		seed, err := strconv.Atoi(val)
		if err == nil {
			seeds = append(seeds, seed)
		}
	}

	seedToSoil := map[numRange]numRange{}
	soilToFert := map[numRange]numRange{}
	fertToWater := map[numRange]numRange{}
	waterToLight := map[numRange]numRange{}
	lightToTemp := map[numRange]numRange{}
	tempToHumid := map[numRange]numRange{}
	humidToLocation := map[numRange]numRange{}
	mapIndex := -1
	skip := true

	for _, line := range lines {
		if skip {
			skip = false
			continue
		}
		if line == "" {
			skip = true
			mapIndex++
			continue
		}

		rules := strings.Fields(line)
		destStart := S2i(rules[0])
		sourceStart := S2i(rules[1])
		rangeLength := S2i(rules[2])
		source := numRange{
			start: sourceStart,
			end:   sourceStart + rangeLength - 1,
		}
		dest := numRange{
			start: destStart,
			end:   destStart + rangeLength - 1,
		}

		if mapIndex == 0 {
			seedToSoil[source] = dest
		}
		if mapIndex == 1 {
			soilToFert[source] = dest
		}
		if mapIndex == 2 {
			fertToWater[source] = dest
		}
		if mapIndex == 3 {
			waterToLight[source] = dest
		}
		if mapIndex == 4 {
			lightToTemp[source] = dest
		}
		if mapIndex == 5 {
			tempToHumid[source] = dest
		}
		if mapIndex == 6 {
			humidToLocation[source] = dest
		}
		if mapIndex > 6 {
			fmt.Println("something has gone horribly wrong")
		}

	}

	answer = 99999999999
	for _, seed := range seeds {
		soil := lookup(seed, seedToSoil)
		fert := lookup(soil, soilToFert)
		water := lookup(fert, fertToWater)
		light := lookup(water, waterToLight)
		temp := lookup(light, lightToTemp)
		humid := lookup(temp, tempToHumid)
		location := lookup(humid, humidToLocation)
		if location < answer {
			answer = location
		}
	}

	fmt.Printf("day5 Answer 1 : %d\n", answer)
}

func part2(lines []string) {
	answer := 0
	seeds := []numRange{}
	startSeed := true
	seedIndex := 0

	for _, val := range strings.Fields(lines[0]) {
		seed, err := strconv.Atoi(val)
		if err == nil {
			if startSeed {
				startSeed = false
				seedRange := numRange{start: seed}
				seeds = append(seeds, seedRange)
			} else {
				startSeed = true
				seeds[seedIndex].end = seeds[seedIndex].start + seed - 1
				seedIndex++
			}
		}
	}

	seedToSoil := map[numRange]numRange{}
	soilToFert := map[numRange]numRange{}
	fertToWater := map[numRange]numRange{}
	waterToLight := map[numRange]numRange{}
	lightToTemp := map[numRange]numRange{}
	tempToHumid := map[numRange]numRange{}
	humidToLocation := map[numRange]numRange{}
	mapIndex := -1
	skip := true

	for _, line := range lines {
		if skip {
			skip = false
			continue
		}
		if line == "" {
			skip = true
			mapIndex++
			continue
		}

		rules := strings.Fields(line)
		destStart := S2i(rules[0])
		sourceStart := S2i(rules[1])
		rangeLength := S2i(rules[2])
		source := numRange{
			start: sourceStart,
			end:   sourceStart + rangeLength - 1,
		}
		dest := numRange{
			start: destStart,
			end:   destStart + rangeLength - 1,
		}

		if mapIndex == 0 {
			seedToSoil[source] = dest
		}
		if mapIndex == 1 {
			soilToFert[source] = dest
		}
		if mapIndex == 2 {
			fertToWater[source] = dest
		}
		if mapIndex == 3 {
			waterToLight[source] = dest
		}
		if mapIndex == 4 {
			lightToTemp[source] = dest
		}
		if mapIndex == 5 {
			tempToHumid[source] = dest
		}
		if mapIndex == 6 {
			humidToLocation[source] = dest
		}
		if mapIndex > 6 {
			fmt.Println("something has gone horribly wrong")
		}

	}

	answer = 99999999999

	soils := []numRange{}
	ferts := []numRange{}
	waters := []numRange{}
	lights := []numRange{}
	temps := []numRange{}
	humids := []numRange{}
	locations := []numRange{}
	for _, seed := range seeds {
		soils = append(soils, lookupPart2(seed, seedToSoil)...)
	}
	for _, soil := range soils {
		ferts = append(ferts, lookupPart2(soil, soilToFert)...)
	}
	for _, fert := range ferts {
		waters = append(waters, lookupPart2(fert, fertToWater)...)
	}
	for _, water := range waters {
		lights = append(lights, lookupPart2(water, waterToLight)...)
	}
	for _, light := range lights {
		temps = append(temps, lookupPart2(light, lightToTemp)...)
	}
	for _, temp := range temps {
		humids = append(humids, lookupPart2(temp, tempToHumid)...)
	}
	for _, humid := range humids {
		locations = append(locations, lookupPart2(humid, humidToLocation)...)
	}
	sort.Slice(locations, func(i, j int) bool {
		return locations[i].start < locations[j].start
	})
	answer = locations[0].start

	fmt.Printf("day5 Answer 2 : %d\n", answer)
}

func lookup(val int, m map[numRange]numRange) int {
	returnVal := val
	for source, dest := range m {
		if val >= source.start && val <= source.end {
			dist := val - source.start
			returnVal = dest.start + dist
			break
		}
	}
	return returnVal
}

func lookupPart2(val numRange, m map[numRange]numRange) []numRange {
	returnVal := []numRange{}
	for source, dest := range m {
		// entirely within
		if source.start <= val.start && source.end >= val.end {
			startDist := val.start - source.start
			endDist := source.end - val.end
			rv := numRange{
				start: dest.start + startDist,
				end:   dest.end - endDist,
			}
			returnVal = append(returnVal, rv)
			val.moved = true
		} else if source.start <= val.start && source.end <= val.end && source.end >= val.start {
			// halfway in
			newRange1 := numRange{
				start: val.start,
				end:   source.end,
			}
			newRange2 := numRange{
				start: source.end + 1,
				end:   val.end,
			}
			returnVal = append(returnVal, lookupPart2(newRange1, m)...)
			returnVal = append(returnVal, lookupPart2(newRange2, m)...)
			val.moved = true
		} else if source.end >= val.end && source.start >= val.start && source.start <= val.end {
			// halfway out
			newRange1 := numRange{
				start: val.start,
				end:   source.start - 1,
			}
			newRange2 := numRange{
				start: source.start,
				end:   val.end,
			}
			returnVal = append(returnVal, lookupPart2(newRange1, m)...)
			returnVal = append(returnVal, lookupPart2(newRange2, m)...)
			val.moved = true
		}
	}

	if !val.moved {
		rv := numRange{
			start: val.start,
			end:   val.end,
		}
		returnVal = append(returnVal, rv)
	}

	return returnVal
}
