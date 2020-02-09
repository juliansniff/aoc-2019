package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := make([]int, 2)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var err error
		input := strings.Split(scanner.Text(), "-")
		r[0], err = strconv.Atoi(input[0])
		if err != nil {
			panic(err)
		}
		r[1], err = strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(numPasswords(r[0], r[1]))
}

func numPasswords(min, max int) int {
	count := 0
	for p := min; p <= max; p++ {
		if validPassword(p, min, max) {
			count++
		}
	}
	return count
}

func validPassword(p, min, max int) bool {
	repeats := false
	lastDigit := -1
	digitCount := 0
	for p >= 1 {
		digit := p % 10
		if digit > lastDigit && lastDigit != -1 {
			return false
		}
		if digit == lastDigit {
			repeats = true
		}
		p /= 10
		lastDigit = digit
		digitCount++
	}
	return digitCount == 6 && repeats
}
