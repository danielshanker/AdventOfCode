package main

import (
	"flag"
	"strconv"
	"strings"
	. "utils"
)

func main() {
	test := flag.Bool("t", false, "use sample")
	flag.Parse()

	Start(test, 22, part1, part2, 37327623, 23)

}

func part1(lines []string) int {
	answer := 0

	for _, line := range lines {
		secret := S2i(line)
		for i := 0; i < 2000; i++ {
			secret = step1(secret)
			secret = step2(secret)
			secret = step3(secret)
		}
		answer += secret
	}

	return answer
}

func step1(secret int) int {
	a := secret * 64
	secret = mix(secret, a)
	return prune(secret)
}

func step2(secret int) int {
	a := secret / 32
	secret = mix(secret, a)
	return prune(secret)
}

func step3(secret int) int {
	a := secret * 2048
	secret = mix(secret, a)
	return prune(secret)
}

func mix(val, mixVal int) int {
	return val ^ mixVal
}

func prune(val int) int {
	return val % 16777216
}

func part2(lines []string) int {
	answer := 0

	sequences := map[string]int{}

	for _, line := range lines {
		singleSequence := map[string]int{}
		four := 0
		three := 0
		two := 0
		one := 0
		secret := S2i(line)
		lastSecret := secret % 10
		for i := 0; i < 2000; i++ {
			secret = step1(secret)
			secret = step2(secret)
			secret = step3(secret)
			four = three
			three = two
			two = one
			one = (secret % 10) - lastSecret
			if i > 3 {
				s := make([]string, 4)
				s[0] = strconv.Itoa(four)
				s[1] = strconv.Itoa(three)
				s[2] = strconv.Itoa(two)
				s[3] = strconv.Itoa(one)
				seq := strings.Join(s, ",")
				if _, ok := singleSequence[seq]; !ok {
					singleSequence[seq] = secret % 10
				}
			}
			lastSecret = secret % 10
		}
		for k, v := range singleSequence {
			sequences[k] += v
		}
	}

	for _, v := range sequences {
		if v >= answer {
			answer = v
		}
	}

	return answer
}
