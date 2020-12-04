package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")

	bagContents := make(map[string][]string)
	bagContents2 := make(map[string]map[string]int)

	for _, line := range lines {
		words := strings.Split(line, " bags contain ")
		initBag := words[0]
		iBags, iBags2 := getInternalBags(words[1])

		bagContents[initBag] = iBags
		bagContents2[initBag] = iBags2

	}

	foundBags := 0
	for k := range bagContents {
		if findBags(bagContents, k) {
			foundBags++
		}
	}

	answer2 := findBagAmount(bagContents2, "shiny gold")

	fmt.Println(fmt.Sprintf("Answer 1 : %d", foundBags))
	fmt.Println(fmt.Sprintf("Answer 2 : %d", answer2))
}

func findBags(bagContents map[string][]string, bag string) bool {
	if bagContents[bag][0] == "no other" {
		return false
	}
	for _, curBag := range bagContents[bag] {
		if curBag == "shiny gold" {
			return true
		}
		found := findBags(bagContents, curBag)
		if found {
			return true
		}
	}

	return false
}

func findBagAmount(bagContents map[string]map[string]int, bag string) int {
	val := 0
	for curBag, amount := range bagContents[bag] {
		val += amount
		val += amount * findBagAmount(bagContents, curBag)
	}
	return val
}

func getInternalBags(input string) ([]string, map[string]int) {
	re := regexp.MustCompile(`\d+ [a-zA-Z]+\s[a-zA-Z]+`)
	reString := regexp.MustCompile(`[a-zA-Z]+\s[a-zA-Z]+`)
	bagContents := reString.FindAllString(input, -1)

	bc := re.FindAllString(input, -1)
	bagContents2 := make(map[string]int)
	for _, i := range bc {
		reS := regexp.MustCompile(`[a-zA-Z]+\s[a-zA-Z]+`)
		reI := regexp.MustCompile(`\d+`)
		numba2 := reS.FindAllString(i, -1)
		bagContents2[numba2[0]], _ = strconv.Atoi(reI.FindAllString(i, -1)[0])

	}
	return bagContents, bagContents2
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
