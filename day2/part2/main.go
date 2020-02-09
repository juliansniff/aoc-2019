package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var opcodes []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		for _, c := range input {
			i, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			opcodes = append(opcodes, i)
		}
	}
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			t_opcodes := make([]int, len(opcodes))
			copy(t_opcodes, opcodes)
			t_opcodes[1] = noun
			t_opcodes[2] = verb
			m := &Machine{t_opcodes}
			m.Run()
			if m.opcodes[0] == 19690720 {
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}

type Machine struct {
	opcodes []int
}

func (m *Machine) Run() {
	for i := 0; i < len(m.opcodes); i += 4 {
		switch m.opcodes[i] {
		case 1:
			// addition
			pos1, pos2, resultPos := m.opcodes[i+1], m.opcodes[i+2], m.opcodes[i+3]
			m.opcodes[resultPos] = m.opcodes[pos1] + m.opcodes[pos2]
		case 2:
			// multiplication
			pos1, pos2, resultPos := m.opcodes[i+1], m.opcodes[i+2], m.opcodes[i+3]
			m.opcodes[resultPos] = m.opcodes[pos1] * m.opcodes[pos2]
		case 99:
			return
		default:
			return
		}
	}
}
