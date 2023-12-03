package day2

import (
	"aoc/shared"
	"fmt"
	"regexp"

	// "slices"
	"strconv"
	"strings"
)

var r *regexp.Regexp = regexp.MustCompile("([0-9]+) red")
var b *regexp.Regexp = regexp.MustCompile("([0-9]+) blue")
var g *regexp.Regexp = regexp.MustCompile("([0-9]+) green")

func ExecuteDay2a() {
	// lines := shared.ReadAllLines("day2/small_input.txt")
	lines := shared.ReadAllLines("day2/input.txt")
	var total int = 0
	for _, line := range lines {
		gameId := getGameId(line)
		maxR := getMaxForRegexp(line, r)
		maxG := getMaxForRegexp(line, g)
		maxB := getMaxForRegexp(line, b)
		isPossible := maxR <= 12 && maxG <= 13 && maxB <= 14
		if isPossible {
			total += gameId
		}
		fmt.Println(line)
		fmt.Println(maxR, maxG, maxB)
	}
	fmt.Println(total)
}

func getMaxForRegexp(line string, regexToUse *regexp.Regexp) int {
	all := regexToUse.FindAllStringSubmatch(line, -1)
	blues := make([]int, len(all))
	for idx, v := range all {
		val, err := strconv.Atoi(v[1])
		blues[idx] = val
		shared.Check(err)
	}
	max, err := shared.Max(blues)
	shared.Check(err)
	return max
	// return slices.Max(blues)
}

func getGameId(line string) int {
	onColon := strings.Split(line, ":")
	gameNumRaw := onColon[0]
	extraxctNum := regexp.MustCompile("Game ([0-9]+)")
	id := extraxctNum.FindStringSubmatch(gameNumRaw)[1]
	ret, err := strconv.Atoi(id)
	shared.Check(err)
	return ret
}
