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
	m := &Machine{opcodes}
	m.Run()
	fmt.Println(opcodes[0])
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
			panic(fmt.Sprintf("Unknown Opcode: %d", i))
		}
	}
}
