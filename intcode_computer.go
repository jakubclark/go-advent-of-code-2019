package go_advent_of_code_2019

import "fmt"

type Machine struct {
	Memory []int
	i      int
}

func NewMachine(program []int) Machine {
	machine := Machine{Memory: program}
	return machine
}

func (m Machine) GetArgument(argPosition int) int {
	index := m.i + argPosition + 1
	return m.Memory[m.Memory[index]]
}

func (m Machine) GetAddress(argPosition int) int {
	return m.Memory[m.i+argPosition+1]
}

func (m Machine) RunProgram() {
	for {
		opcode := m.Memory[m.i]

		switch opcode {
		case 1:
			left := m.GetArgument(0)
			right := m.GetArgument(1)
			dest := m.GetAddress(2)
			m.Memory[dest] = left + right
			m.i += 4
		case 2:
			left := m.GetArgument(0)
			right := m.GetArgument(1)
			dest := m.GetAddress(2)
			m.Memory[dest] = left * right
			m.i += 4
		case 99:
			return
		default:
			panic(fmt.Sprintf("Unexpected opcode: %d", opcode))
		}
	}
}
