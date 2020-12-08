package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = [][]string{}

type exitResult struct {
	ExitCode    int
	ReturnValue int
}

func main() {
	loadData()
	part1()
	part2()
}

func part1() {
	returnValue := runProgram(input)
	fmt.Println("Part 1:", returnValue.ReturnValue)
}

func part2() {
	i := 0
	retVal := exitResult{}
	for i < len(input) {
		instructions := copySwap(i)
		retVal = runProgram(instructions)
		if retVal.ExitCode == 0 {
			fmt.Println("Part 2:", retVal.ReturnValue)
			return
		}
		i++
	}
}

func copySwap(idx int) [][]string {
	instructions := [][]string{}
	for i := range input {
		if input[i][0] == "jmp" && i == idx {
			instruction := []string{"nop", input[i][1]}
			instructions = append(instructions, instruction)
		} else if input[i][0] == "nop" && i == idx {
			instruction := []string{"jmp", input[i][1]}
			instructions = append(instructions, instruction)
		} else {
			instructions = append(instructions, input[i])
		}
	}
	return instructions
}

func runProgram(instructions [][]string) exitResult {
	acc := 0
	pointer := 0
	prev := map[int]bool{}
	code := 1
	for pointer < len(instructions) {
		if _, e := prev[pointer]; e {
			return exitResult{ExitCode: code, ReturnValue: acc}
		}
		prev[pointer] = true
		val, _ := strconv.Atoi(instructions[pointer][1])

		switch instructions[pointer][0] {
		case "jmp":
			pointer += val
		case "acc":
			acc += val
			pointer++
		default:
			pointer++
		}
	}
	if pointer == len(instructions) {
		code = 0
	}
	return exitResult{ExitCode: code, ReturnValue: acc}
}

func loadData() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line2 := strings.Split(line, " ")
		input = append(input, line2)
	}
}
