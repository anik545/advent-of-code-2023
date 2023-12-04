package day4

import (
	"aoc/shared"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Card struct {
	winningNums []int
	nums        []int
	cardId      int
}

func Execute4a() {
	lines := shared.ReadAllLines("day4/input.txt")
	cards := parseCards(lines)
	total := 0
	for _, card := range cards {
		intersect := rightItemsInLeft(card.winningNums, card.nums)
		ls := len(intersect)
		fmt.Printf("%+v", card)
		fmt.Println(intersect)
		if ls != 0 {
			fmt.Println(int(math.Pow(2, float64(ls-1))))
			total += int(math.Pow(2, float64(ls-1)))
		}
	}
	fmt.Println(total)
}

func rightItemsInLeft(left, right []int) []int {
	// there is no set in go this is apparently what you're meant to do...
	inLeft := make(map[int]bool)
	result := make([]int, 0)
	for _, item := range left {
		inLeft[item] = true
	}
	for _, item := range right {
		if _, ok := inLeft[item]; ok {
			result = append(result, item)
		}
	}
	return result

}

var r = regexp.MustCompile(`Card\s*([0-9]+):(.*)\|(.*)`)

func parseCards(lines []string) []Card {
	cards := make([]Card, 0)
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		cardId := shared.ToIntOrPanic(matches[0][1])
		winningInts := matches[0][2]
		ints := matches[0][3]
		winnings := toIntArray(strings.TrimSpace(winningInts))
		nums := toIntArray(strings.TrimSpace(ints))
		cards = append(cards, Card{winningNums: winnings, nums: nums, cardId: cardId})
	}
	return cards
}

func toIntArray(spaceSeparated string) []int {
	splits := strings.Split(spaceSeparated, " ")
	ret := make([]int, 0)
	for _, val := range splits {
		if val == "" {
			continue
		}
		ret = append(ret, shared.ToIntOrPanic(val))
	}
	return ret

}
