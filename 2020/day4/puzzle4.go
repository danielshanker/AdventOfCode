package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"encoding/hex"
)

func main() {

	lines := readInputLines("day4.txt")
	//lines := readInputLines("../sample.txt")

	var passport []string
	var passports [][]string
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = nil
		}
		entries := strings.Split(line, " ")
		for _, entry := range entries {
			passport = append(passport, entry)
		}
	}
	if passports != nil {
		passports = append(passports, passport)
	}

	requiredEntries := strings.Split("byr iyr eyr hgt hcl ecl pid", " ")

	validPassports := 0
	validPassports2 := 0
	for _, curPass := range passports {
		validFields := 0
		validFields2 := 0
		for _, curVal := range curPass {
			entry := strings.Split(curVal, ":")
			if len(entry) != 2 {
				continue
			}
			if isValueInArray(entry[0], requiredEntries) {
				validFields++
			}
			if isValueInArray(entry[0], requiredEntries) && isValidEntry(entry[0], entry[1]) {
				validFields2++
			}
		}
		if validFields >= len(requiredEntries) {
			validPassports++
		}
		if validFields2 >= len(requiredEntries) {
			validPassports2++
		}
	}

	fmt.Println(fmt.Sprintf("Answer 1: %d", validPassports))
	fmt.Println(fmt.Sprintf("Answer 2: %d", validPassports2))

}

func isValueInArray(value string, array []string) bool {
	for _, curVal := range array {
		if curVal == value {
			return true
		}
	}
	return false
}

func isValidEntry(entry string, value string) bool {

	switch entry {
	case "byr":
		intVal, err := strconv.Atoi(value)
		if err == nil {
			if intVal >= 1920 && intVal <= 2002 {
				return true
			}
		}
	case "iyr":
		intVal, err := strconv.Atoi(value)
		if err == nil {
			if intVal >= 2010 && intVal <= 2020 {
				return true
			}
		}
	case "eyr":
		intVal, err := strconv.Atoi(value)
		if err == nil {
			if intVal >= 2020 && intVal <= 2030 {
				return true
			}
		}
	case "hgt":
		unit := value[len(value)-2:]
		heightString := value[0 : len(value)-2]
		if unit == "cm" || unit == "in" {
			height, err := strconv.Atoi(heightString)
			if err == nil {
				if unit == "cm" {
					if height > 193 || height < 150 {
						return false
					}
				} else {
					if height > 76 || height < 59 {
						return false
					}
				}
				return true
			}
		}
	case "hcl":
		if string(value[0]) != "#" {
			return false
		}
		colour := value[1:]
		_, err := hex.DecodeString(colour)
		if err == nil {
			return true
		}

	case "ecl":
		validColours := strings.Split("amb blu brn gry grn hzl oth", " ")
		if isValueInArray(value, validColours) {
			return true
		}
	case "pid":
		if len(value) != 9 {
			return false
		}
		_, err := strconv.Atoi(value)
		if err == nil {
			return true
		}
	case "cid":
		return true
	}

	return false
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
