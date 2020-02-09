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
	distance := -1
	for _, intersection := range intersections(wires[0], wires[1]) {
		x, y := intersection[0], intersection[1]
		if x == 0 && y == 0 {
			continue
		}
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		d := x + y
		fmt.Println(intersection, d)
		if distance == -1 || d < distance {
			distance = d
		}
	}
	fmt.Println(distance)
}

type Wire map[int](map[int]bool)

func CreateWire(paths []string) Wire {
	wire := Wire{
		0: map[int]bool{
			0: true,
		},
	}
	x, y := 0, 0
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
				wire[x] = make(map[int]bool)
			}
			wire[x][y] = true
			itr()
		}
	}
	return wire
}

func intersections(w1, w2 Wire) [][]int {
	var intersections [][]int
	var ok bool
	for x, m := range w1 {
		if _, ok = w2[x]; !ok {
			continue
		}
		for y, _ := range m {
			if _, ok = w2[x][y]; ok {
				intersections = append(intersections, []int{x, y})
			}
		}
	}
	return intersections
}
