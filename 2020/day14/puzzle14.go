package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	mask := ""
	mem := make(map[int]int64)

	for _, line := range lines {
		arrLine := strings.Split(line, " = ")
		inst := arrLine[0]
		val := arrLine[1]
		if inst == "mask" {
			mask = val
			continue
		}
		re := regexp.MustCompile(`\d+`)
		memVal, _ := strconv.Atoi(re.FindString(inst))

		intVal, _ := strconv.Atoi(val)
		mem[memVal] = applyMask1(intVal, mask)
	}

	answer1 := int64(0)
	for _, i := range mem {
		answer1 += i
	}

	mask = ""
	mem = make(map[int]int64)

	for _, line := range lines {
		arrLine := strings.Split(line, " = ")
		inst := arrLine[0]
		val := arrLine[1]
		if inst == "mask" {
			mask = val
			continue
		}
		re := regexp.MustCompile(`\d+`)
		memVal, _ := strconv.Atoi(re.FindString(inst))

		intVal, _ := strconv.Atoi(val)
		applyMask2(memVal, int64(intVal), mem, mask)
	}
	answer2 := int64(0)
	for _, i := range mem {
		answer2 += i
	}

	fmt.Println(fmt.Sprintf("Answer1: %d", answer1))
	fmt.Println(fmt.Sprintf("Answer2: %d", answer2))

}

func applyMask1(value int, mask string) int64 {

	binaryVal := fmt.Sprintf("%036b", value)

	newBin := ""

	for i := 0; i < len(mask); i++ {
		if string(mask[i]) == "X" {
			newBin += string(binaryVal[i])
			continue
		}
		newBin += string(mask[i])
	}

	returnVal, _ := strconv.ParseInt(newBin, 2, 64)

	return returnVal
}

func applyMask2(address int, value int64, mem map[int]int64, mask string) {

	binaryVal := fmt.Sprintf("%036b", address)

	newBin := ""

	for i := 0; i < len(mask); i++ {
		if string(mask[i]) == "0" {
			newBin += string(binaryVal[i])
			continue
		}
		if string(mask[i]) == "1" {
			newBin += "1"
			continue
		}
		if string(mask[i]) == "X" {
			newBin += "X"
			continue
		}
	}

	addresses := getAllAddresses(newBin)

	for _, i := range addresses {
		mem[i] = value
	}
}

func getAllAddresses(bin string) []int {
	var addresses []int
	allXes := 0
	for _, i := range bin {
		if string(i) == "X" {
			allXes++
		}
	}
	options := int(math.Pow(2, float64(allXes)))

	for i := 0; i < options; i++ {
		newBin := bin
		binaryVal := fmt.Sprintf("%b", i)
		binIndex := len(binaryVal) - 1
		newString := ""
		for j := len(newBin) - 1; j >= 0; j-- {
			if string(newBin[j]) == "X" {
				if binIndex < 0 {
					newString = "0" + newString
				} else {
					newString = string(binaryVal[binIndex]) + newString
				}
				binIndex--
			} else {
				newString = string(newBin[j]) + newString
			}
		}
		addressInt, _ := strconv.ParseInt(newString, 2, 64)
		addresses = append(addresses, int(addressInt))
	}

	return addresses
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
