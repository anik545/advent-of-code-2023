package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"flag"
	"fmt"
	"log"
)

func main() {
	dayPtr := flag.String("day", "", "which day's code to run. format is -day=1a")

	flag.Parse()

	fmt.Println(*dayPtr)
	switch *dayPtr {
	case "1a":
		day1.ExecuteDay1a()
	case "1b":
		day1.ExecuteDay1b()
	case "2a":
		day2.ExecuteDay2a()
	case "2b":
		day2.ExecuteDay2b()
	case "3a":
		day3.ExecuteDay3a()
	case "3b":
		day3.ExecuteDay3b()
	case "4a":
		day4.Execute4a()
	case "4b":
		day4.Execute4b()
	default:
		log.Fatalf("failed to parse args")
	}

}
