package day3

import (
	"aoc/shared"
	"fmt"
	"regexp"
)

type NumberAtPosition struct {
	// this is `[startIdx, endIdx)`
	startIdx, endIdx int
	num              int
}

var numRegex = regexp.MustCompile("[0-9]+")

func ExecuteDay3a() {
	// lines := shared.ReadAllLines("day3/small_input.txt")
	lines := shared.ReadAllLines("day3/input.txt")
	total := 0
	for idx, line := range lines {
		var poss = getNumbersAtPosition(line)
		for _, num := range poss {
			var hasAdjacentSymbol = checkAdjacentSymbol(lines, idx, num)
			if hasAdjacentSymbol {
				total += num.num
			}
		}

	}
	fmt.Printf("total, %d", total)
}

func checkAdjacentSymbol(lines []string, idx int, numberAtPosition NumberAtPosition) bool {
	aboveBelowEnd := numberAtPosition.endIdx
	aboveBelowStart := numberAtPosition.startIdx - 1

	var nextLineHasSymbol bool
	if idx != len(lines)-1 {
		for x := aboveBelowStart; x <= aboveBelowEnd; x++ {
			nextLine := lines[idx+1]
			if (len(nextLine)-1) < x || x < 0 {
				continue
			}
			if isSymbol(nextLine[x]) {
				nextLineHasSymbol = true
			}
		}
	}
	var prevLineHasSymbol bool
	if idx != 0 {
		for x := aboveBelowStart; x <= aboveBelowEnd; x++ {
			previousLine := lines[idx-1]
			if (len(previousLine)-1) < x || x < 0 {
				continue
			}
			if isSymbol(previousLine[x]) {
				nextLineHasSymbol = true
			}
		}
	}
	var thisLineHasSymbol bool
	thisLine := lines[idx]
	if aboveBelowEnd < len(thisLine)-1 {
		if isSymbol(thisLine[aboveBelowEnd]) {
			thisLineHasSymbol = true
		}
	}
	if aboveBelowStart > 0 {
		if isSymbol(thisLine[aboveBelowStart]) {
			thisLineHasSymbol = true
		}
	}
	return thisLineHasSymbol || prevLineHasSymbol || nextLineHasSymbol
}

var isSymbolRegex = regexp.MustCompile("[^0-9\\.]")

func isSymbol(char byte) bool {
	matches := isSymbolRegex.Match([]byte{char})
	return matches
}

func getNumbersAtPosition(line string) []NumberAtPosition {
	matches := numRegex.FindAllStringSubmatchIndex(line, -1)
	nums := make([]NumberAtPosition, len(matches))
	for idx, match := range matches {
		startIdx := match[0]
		endIdx := match[1]
		num := line[startIdx:endIdx]
		nums[idx] = NumberAtPosition{startIdx: startIdx, endIdx: endIdx, num: shared.ToIntOrPanic(num)}
	}
	return nums
}
