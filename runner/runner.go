package main

import (
	"aoc/day1"
	"aoc/day2"
	"flag"
	"fmt"
	"log"
)

func main() {
	dayPtr := flag.String("day", "1", "which day's code to run. format is 1a, 1b, 2a, etc.")

	flag.Parse()

	fmt.Println(*dayPtr)
	switch *dayPtr {
	case "1a":
		day1.ExecuteDay1a()
	case "1b":
		day1.ExecuteDay1b()
	case "2a":
		day2.ExecuteDay2a()
	default:
		log.Fatalf("failed to parse args")
	}

}
