package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	if *test {
		sample := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day16/sample.txt")
		part1(sample)
		part2(sample)
	} else {
		input := readInputLines("/home/daniel.shanker/Pers/AdventOfCode/2021/day16/input.txt")
		part1(input)
		part2(input)
	}

}

var versionTotal = 0

func part1(lines []string) {
	answer1 := 0
	binaryInput := convertHexToBinary(lines[0])
	index := 0
	parse(binaryInput, index)

	answer1 = versionTotal
	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
}

func part2(lines []string) {
	answer2 := 0

	binaryInput := convertHexToBinary(lines[0])
	index := 0
	_, val := parse2(binaryInput, index)

	answer2 = val
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func parse2(binaryInput string, index int) (int, int) {
	typeID := ""
	version := ""
	val := 0

	index, version, typeID = parseHeader(binaryInput, index)
	versionTotal += bin2Int(version)
	if bin2Int(typeID) == 4 {
		index, val = parseLiteral(binaryInput, index)
	} else {
		var values []int
		if string(binaryInput[index]) == "0" {
			index++
			l := ""
			for i := 0; i < 15; i++ {
				l += string(binaryInput[index])
				index++
			}
			length := bin2Int(l)
			subPacketEnd := index + length
			for index < subPacketEnd {
				v := 0
				index, v = parse2(binaryInput, index)
				values = append(values, v)
			}
		} else {
			index++
			l := ""
			for i := 0; i < 11; i++ {
				l += string(binaryInput[index])
				index++
			}
			length := bin2Int(l)
			for i := 0; i < length; i++ {
				v := 0
				index, v = parse2(binaryInput, index)
				values = append(values, v)
			}
		}
		switch bin2Int(typeID) {
		case 0:
			for _, i := range values {
				val += i
			}
		case 1:
			val = 1
			for _, i := range values {
				val *= i
			}
		case 2:
			val = math.MaxInt64
			for _, i := range values {
				if i < val {
					val = i
				}
			}
		case 3:
			for _, i := range values {
				if i > val {
					val = i
				}
			}
		case 5:
			if values[0] > values[1] {
				val = 1
			} else {
				val = 0
			}
		case 6:
			if values[0] < values[1] {
				val = 1
			} else {
				val = 0
			}
		case 7:
			if values[0] == values[1] {
				val = 1
			} else {
				val = 0
			}
		}

	}
	return index, val
}

func parse(binaryInput string, index int) int {
	typeID := ""
	version := ""

	index, version, typeID = parseHeader(binaryInput, index)
	versionTotal += bin2Int(version)
	if bin2Int(typeID) == 4 {
		index, _ = parseLiteral(binaryInput, index)
	} else {
		if string(binaryInput[index]) == "0" {
			index++
			l := ""
			for i := 0; i < 15; i++ {
				l += string(binaryInput[index])
				index++
			}
			length := bin2Int(l)
			subPacketEnd := index + length
			for index < subPacketEnd {
				index = parse(binaryInput, index)
			}
		} else {
			index++
			l := ""
			for i := 0; i < 11; i++ {
				l += string(binaryInput[index])
				index++
			}
			length := bin2Int(l)
			for i := 0; i < length; i++ {
				index = parse(binaryInput, index)
			}

		}
	}
	return index
}

func parseHeader(b string, index int) (int, string, string) {
	version := string(b[index]) + string(b[index+1]) + string(b[index+2])
	index += 3
	typeID := string(b[index]) + string(b[index+1]) + string(b[index+2])
	index += 3

	return index, version, typeID
}

func parseLiteral(b string, index int) (int, int) {
	val := ""
	for true {
		breakOut := false
		if string(b[index]) == "0" {
			breakOut = true
		}
		index++
		for i := index; i < index+4; i++ {
			val += string(b[i])
		}
		index += 4
		if breakOut {
			break
		}
	}

	return index, bin2Int(val)
}
func bin2Int(val string) int {
	valInt, _ := strconv.ParseInt(val, 2, 64)
	return int(valInt)
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

func convertHexToBinary(val string) string {
	bin := ""
	for _, i := range val {
		char := string(i)
		switch char {
		case "0":
			bin += "0000"
		case "1":
			bin += "0001"
		case "2":
			bin += "0010"
		case "3":
			bin += "0011"
		case "4":
			bin += "0100"
		case "5":
			bin += "0101"
		case "6":
			bin += "0110"
		case "7":
			bin += "0111"
		case "8":
			bin += "1000"
		case "9":
			bin += "1001"
		case "A":
			bin += "1010"
		case "B":
			bin += "1011"
		case "C":
			bin += "1100"
		case "D":
			bin += "1101"
		case "E":
			bin += "1110"
		case "F":
			bin += "1111"
		}
	}
	return bin
}
