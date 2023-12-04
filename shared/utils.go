package shared

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

// ReadAllLines panics if it fails to read all lines
func ReadAllLines(fileName string) []string {
	dat, err := os.Open(fileName)
	defer dat.Close()

	var lines = make([]string, 0)

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

func Max(input []int) (int, error) {
	if len(input) == 0 {

		return 0, errors.New("empty array provided")
	}
	var max = input[0]

	for _, val := range input {
		if val > max {
			max = val
		}
	}

	return max, nil
}

// IntMin wtf - why is this not in stdlib?????
func IntMin(a, b int) int {
	if a == b || a > b {
		return a
	}
	return b
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ToIntOrPanic(intable string) int {
	parsed, err := strconv.Atoi(intable)
	Check(err)
	return parsed
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
