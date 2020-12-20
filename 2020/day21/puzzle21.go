package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("input.txt")
	//lines := readInputLines("../sample.txt")

	possibleAllergenWords := make(map[string][]string)
	regex := "\\(contains (.*)\\)"
	re := regexp.MustCompile(regex)
	allWords := make(map[string]bool)

	for _, line := range lines {
		a := re.FindStringSubmatch(line)
		allergens := strings.Split(a[1], ", ")

		words := strings.Split(line, " (")
		wordList := strings.Fields(words[0])
		for _, i := range wordList {
			allWords[i] = true
		}

		for _, i := range allergens {
			if _, ok := possibleAllergenWords[i]; !ok {
				possibleAllergenWords[i] = wordList
			} else {
				var newAllergenList []string
				for _, allergenWord := range possibleAllergenWords[i] {
					foundMatch := false
					for _, word := range wordList {
						if allergenWord == word {
							foundMatch = true
							break
						}
					}
					if foundMatch {
						newAllergenList = append(newAllergenList, allergenWord)
					}
				}
				possibleAllergenWords[i] = newAllergenList
			}
		}
	}

	for _, i := range possibleAllergenWords {
		for _, j := range i {
			allWords[j] = false
		}
	}

	answer1 := 0
	for i, j := range allWords {
		if j {
			for _, line := range lines {
				for _, k := range strings.Fields(line) {
					if k == i {
						answer1++
					}
				}
			}
		}
	}

	translations := make(map[string]string)
	translations2 := make(map[string]string)
	for len(translations) < len(possibleAllergenWords) {
		for allergen, ingredients := range possibleAllergenWords {
			wordsTranslated := 0
			lastWord := ""
			for _, ingredient := range ingredients {
				if _, ok := translations[ingredient]; ok {
					wordsTranslated++
				} else {
					lastWord = ingredient
				}
			}
			if wordsTranslated+1 == len(ingredients) {
				translations[lastWord] = allergen
				translations2[allergen] = lastWord
			}
		}
	}

	var safeWords []string
	for _, i := range translations {
		safeWords = append(safeWords, i)
	}
	sort.Strings(safeWords)
	answer2 := ""
	for _, i := range safeWords {
		if answer2 != "" {
			answer2 += ","
		}
		answer2 += translations2[i]
	}

	fmt.Println(fmt.Sprintf("Answer 1 : %d", answer1))
	fmt.Println("Answer 2 : " + answer2)
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
