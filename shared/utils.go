package shared

import (
	"bufio"
	"os"
)

// panics if it fails to read all lines
func ReadAllLines(fileName string) []string {
	dat, err := os.Open(fileName)
	defer dat.Close()

	var lines []string = make([]string, 0)

	Check(err)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Sum(input ...int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
