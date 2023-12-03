package day1

import (
	"aoc/shared"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const FILE = "/Volumes/git/anik545/advent-of-code-2023/day1/input.txt"

func ExecuteDay1b() {
	fmt.Println("Running day1")
	linesAsString := shared.ReadAllLines(FILE)
	numbers := make([]int, len(linesAsString))
	for idx, line := range linesAsString {
		var asStrings []string = GivenLine(line)
		joined := strings.Join(asStrings, "")
		parsedInt, err := strconv.Atoi(joined)
		fmt.Println(parsedInt)
		shared.Check(err)
		numbers[idx] = parsedInt
	}
	fmt.Println(shared.Sum(numbers...))
}

var numToNum map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func GivenLine(line string) []string {
	lastNumRegex := regexp.MustCompile(".*((one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9))")
	firstNumRegex := regexp.MustCompile("((one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(1)|(2)|(3)|(4)|(5)|(6)|(7)|(8)|(9)).*")
	lastNumMatches := lastNumRegex.FindAllStringSubmatch(line, -1)
	firstNumMatches := firstNumRegex.FindAllStringSubmatch(line, -1)
	fmt.Println(lastNumMatches)
	fmt.Println(firstNumMatches)
	last := lastNumMatches[0][1]
	first := firstNumMatches[0][1]
	ret := []string{numToNum[first], numToNum[last]}
	fmt.Println(ret)
	return ret
}
