package day1

import (
	"aoc/shared"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ExecuteDay1a() {
	fmt.Println("Running day1b")
	linesAsString := shared.ReadAllLines(FILE)
	numbers := make([]int, len(linesAsString))
	for idx, line := range linesAsString {
		var firstNum string = ""
		var lastNum string = ""
		for _, char := range []rune(line) {
			isNum := unicode.IsNumber(char)
			if isNum {
				if firstNum == "" {
					firstNum = string(char)
					lastNum = string(char)
				}
				if lastNum != "" {
					lastNum = string(char)
				}
			}
		}
		joined := strings.Join([]string{firstNum, lastNum}, "")
		parsedInt, err := strconv.Atoi(joined)
		shared.Check(err)
		numbers[idx] = parsedInt
	}
	fmt.Println(shared.Sum(numbers...))
}
