package main

import (
	"fmt"
	. "github.com/jakubclark/go_advent_of_code_2019"
)

func runParameterized(program []int, noun int, verb int) int {
	program[1] = noun
	program[2] = verb
	m := NewMachine(program)
	m.RunProgram()
	return m.Memory[0]
}
func computeNounAndVerb(program []int) []int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			programCopy := make([]int, len(program))
			copy(programCopy, program)
			res := runParameterized(programCopy, noun, verb)
			if res == 19690720 {
				return []int{noun, verb}
			}
		}
	}
	panic("Failed to compute noun and verb")
}

func part1() {
	program := LoadProgram("input.txt")
	res := runParameterized(program, 12, 2)
	fmt.Printf("Solution for part 1: %d\n", res)
}

func part2() {
	program := LoadProgram("input.txt")
	res := computeNounAndVerb(program)
	noun := res[0]
	verb := res[1]

	fmt.Printf("Solution for part 2: %d\n", 100*noun+verb)
}

func main() {
	part1()
	part2()
}
