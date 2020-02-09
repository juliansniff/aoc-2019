package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var wires []Wire
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		paths := strings.Split(scanner.Text(), ",")
		wire := CreateWire(paths)
		wires = append(wires, wire)
	}
	fmt.Println(intersectionMinSignalDelay(wires[0], wires[1]))
}

type Wire map[int](map[int]int)

func CreateWire(paths []string) Wire {
	wire := Wire{
		0: map[int]int{
			0: 0,
		},
	}
	x, y, d := 0, 0, 0
	for _, path := range paths {
		n, err := strconv.Atoi(path[1:])
		if err != nil {
			panic(err)
		}
		var itr func()
		switch string(path[0]) {
		case "U":
			itr = func() {
				y++
			}
		case "D":
			itr = func() {
				y--
			}
		case "L":
			itr = func() {
				x--
			}
		case "R":
			itr = func() {
				x++
			}
		default:
			panic(fmt.Sprintf("Invalid path: %s", path))
		}
		for i := 0; i < n; i++ {
			_, ok := wire[x]
			if !ok {
				wire[x] = make(map[int]int)
			}
			wire[x][y] = d
			d++
			itr()
		}
	}
	return wire
}

func intersectionMinSignalDelay(w1, w2 Wire) int {
	min := -1
	var ok bool
	for x, m := range w1 {
		if _, ok = w2[x]; !ok {
			continue
		}
		for y, _ := range m {
			if x == 0 && y == 0 {
				continue
			}
			if _, ok = w2[x][y]; ok {
				delay := w1[x][y] + w2[x][y]
				if min == -1 || delay < min {
					min = delay
				}
			}
		}
	}
	return min
}
