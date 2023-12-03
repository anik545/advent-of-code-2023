package day2

import (
	"aoc/shared"
	"fmt"
	"regexp"
	"strings"
	// "slices"
)

type Draw struct {
	red, blue, green int
}

type Game struct {
	draws  []Draw
	gameId int
}

func ExecuteDay2b() {
	// lines := shared.ReadAllLines("day2/small_input.txt")
	lines := shared.ReadAllLines("day2/input.txt")
	var total int = 0
	for _, line := range lines {
		game := lineToGame(line)
		mins := minForGame(game)
		total += powerMins(mins)
	}
	fmt.Println(total)
}

func powerMins(draw Draw) int {
	return draw.blue * draw.green * draw.red
}

func minForGame(game Game) Draw {
	minR := 0
	minG := 0
	minB := 0
	for _, draw := range game.draws {
		minR = shared.IntMin(draw.red, minR)
		minG = shared.IntMin(draw.green, minG)
		minB = shared.IntMin(draw.blue, minB)
	}
	return Draw{red: minR, green: minG, blue: minB}
}

func lineToGame(line string) Game {
	onColon := strings.Split(line, ":")
	gameNumRaw := onColon[0]
	extraxctNum := regexp.MustCompile("Game ([0-9]+)")
	id := shared.ToIntOrPanic(extraxctNum.FindStringSubmatch(gameNumRaw)[1])
	draws := onColon[1]
	drawsRaws := strings.Split(draws, ";")
	drawsParsed := make([]Draw, len(drawsRaws))
	for idx, groupRaw := range drawsRaws {
		draw := groupToDraw(groupRaw)
		drawsParsed[idx] = draw
	}
	return Game{gameId: id, draws: drawsParsed}

}

func groupToDraw(group string) Draw {
	red := getForColor(group, r)
	green := getForColor(group, g)
	blue := getForColor(group, b)
	return Draw{red: red, blue: blue, green: green}
}

func getForColor(line string, regexpToUse *regexp.Regexp) int {
	matches := regexpToUse.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return 0
	}
	return shared.ToIntOrPanic(matches[0][1])
}
