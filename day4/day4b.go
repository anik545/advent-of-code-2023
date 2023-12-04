package day4

import (
	"aoc/shared"
	"fmt"
)

func Execute4b() {
	lines := shared.ReadAllLines("day4/input.txt")
	cards := parseCards(lines)

	var numberOfEachCard map[int]int = make(map[int]int)
	for _, card := range cards {
		fmt.Println(card.cardId)
		intersect := rightItemsInLeft(card.winningNums, card.nums)
		numNextCards := len(intersect)
		numberOfEachCard[card.cardId] += 1
		more, ok := numberOfEachCard[card.cardId]
		toAdd := 1
		if ok && more > 0 {
			toAdd = more
		} else {
			toAdd = 1
		}

		for i := 1; i <= numNextCards; i++ {
			if (card.cardId + i) > len(lines) {
				continue
			}
			numberOfEachCard[card.cardId+i] += toAdd
		}
	}

	total := 0
	for _, v := range numberOfEachCard {
		total += v
	}
	fmt.Println(total)
}
