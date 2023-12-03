package day3

import (
	"aoc/shared"
	"fmt"
)

type Position struct {
	x, y int
}
type PotentialGear struct {
	pos                      Position
	adjacentNumbersPositions []Position
}

func ExecuteDay3b() {
	//lines := shared.ReadAllLines("day3/small_input.txt")
	lines := shared.ReadAllLines("day3/input.txt")
	mapRet := make(map[Position][]NumberAtPosition)
	for idx, line := range lines {
		var poss = getNumbersAtPosition(line)
		fmt.Println("poss", poss)
		for _, num := range poss {
			potentialGearPositions := getAdjacentPotentialGears(lines, idx, num)
			fmt.Println("potentialGearPositions", potentialGearPositions)
			for _, g := range potentialGearPositions {
				curr, ok := mapRet[g]
				if ok {
					mapRet[g] = append(curr, num)
				} else {
					newSlice := []NumberAtPosition{num}
					mapRet[g] = newSlice
				}

			}
		}
	}
	total := 0
	for _, adjacentNumbers := range mapRet {
		uniq := shared.Unique(adjacentNumbers)
		if len(uniq) == 2 {
			total += adjacentNumbers[0].num * adjacentNumbers[1].num
		}
	}
	fmt.Printf("total %d\n", total)
}

//func parseAdjacentNumbers(lines []string, gearCandidate PotentialGear) {
//	for idx, intPositions := range gearCandidate.adjacentNumbersPositions {
//
//	}
//}

func getAdjacentPotentialGears(lines []string, idx int, numberAtPosition NumberAtPosition) []Position {
	positions := make([]Position, 0)
	aboveBelowEnd := numberAtPosition.endIdx
	aboveBelowStart := numberAtPosition.startIdx - 1

	if idx != len(lines)-1 {
		for x := aboveBelowStart; x <= aboveBelowEnd; x++ {
			nextLine := lines[idx+1]
			if (len(nextLine)-1) < x || x < 0 {
				continue
			}
			if isSymbol(nextLine[x]) {
				positions = append(positions, Position{x: x, y: idx + 1})
			}
		}
	}
	if idx != 0 {
		for x := aboveBelowStart; x <= aboveBelowEnd; x++ {
			previousLine := lines[idx-1]
			if (len(previousLine)-1) < x || x < 0 {
				continue
			}
			if isSymbol(previousLine[x]) {
				positions = append(positions, Position{x: x, y: idx - 1})
			}
		}
	}
	thisLine := lines[idx]
	if aboveBelowEnd < len(thisLine)-1 {
		if isSymbol(thisLine[aboveBelowEnd]) {
			positions = append(positions, Position{x: aboveBelowEnd, y: idx})
		}
	}
	if aboveBelowStart > 0 {
		if isSymbol(thisLine[aboveBelowStart]) {
			positions = append(positions, Position{x: aboveBelowStart, y: idx})
		}
	}
	return positions
}
