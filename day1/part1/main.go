package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fuel := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		fuel += calculateFuel(mass)
	}
	fmt.Println(fuel)
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}
