package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_input(file_name string) []int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var value int
	var values []int
	for _, char := range strings.Split(lines[0], ",") {
		value, _ = strconv.Atoi(char)
		values = append(values, value)
	}

	return values
}

func parseInstruction(instruction int) (int, []int) {
	opcode := instruction % 100
	var parameterModes []int
	remainder := instruction

	thirdMode := instruction / 10000
	remainder -= thirdMode * 10000
	secondMode := remainder / 1000
	remainder -= secondMode * 1000
	firstMode := remainder / 100

	switch opcode {
	case 3, 4: //only single parameter_mode
		parameterModes = []int{firstMode}
	case 5, 6: //two parameterModes
		parameterModes = []int{firstMode, secondMode}
	case 1, 2, 7, 8: //three parameter modes
		parameterModes = []int{firstMode, secondMode, thirdMode}
	}

	return opcode, parameterModes
}

func getValue(values []int, parameterMode int, parameterIndex int) int {
	switch parameterMode {
	case 0:
		return values[values[parameterIndex]]
	case 1:
		return values[parameterIndex]
	default:
		panic("error")
	}
}

func getLastOutput(filename string, phaseSetting int, input int) int {

	var value1, value2, output int
	values := get_input(filename)
	instructionIndex := 0
	instruction := values[instructionIndex]
	opcode, parameterModes := parseInstruction(instruction)

	usedInputs := 0

	for opcode != 99 {

		switch opcode {
		case 1:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)
			values[values[instructionIndex+3]] = value1 + value2
			instructionIndex += 4
		case 2:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)
			values[values[instructionIndex+3]] = value1 * value2
			instructionIndex += 4
		case 3:

			switch usedInputs {
			case 0:
				values[values[instructionIndex+1]] = phaseSetting
			case 1:
				values[values[instructionIndex+1]] = input
			default:
				panic("unsupported number of inputs")
			}

			instructionIndex += 2
			usedInputs++
		case 4:
			if parameterModes[0] == 0 {
				output = values[values[instructionIndex+1]]
			} else {
				output = values[instructionIndex+1]
			}
			// fmt.Println("outputting", output)
			instructionIndex += 2
		case 5:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)

			switch value1 {
			case 0:
				instructionIndex += 3
			default:
				instructionIndex = value2
			}
		case 6:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)

			switch value1 {
			case 0:
				instructionIndex = value2
			default:
				instructionIndex += 3
			}
		case 7:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)

			if value1 < value2 {
				values[values[instructionIndex+3]] = 1
			} else {
				values[values[instructionIndex+3]] = 0
			}
			instructionIndex += 4
		case 8:
			value1 = getValue(values, parameterModes[0], instructionIndex+1)
			value2 = getValue(values, parameterModes[1], instructionIndex+2)

			if value1 == value2 {
				values[values[instructionIndex+3]] = 1
			} else {
				values[values[instructionIndex+3]] = 0
			}

			instructionIndex += 4
		default:
			fmt.Println(opcode)
			panic("Unexpected fault")
		}
		instruction = values[instructionIndex]
		opcode, parameterModes = parseInstruction(instruction)
	}

	return output
}

func getCombinations() [][5]int {

	combinations := [][5]int{}
	numbers := []int{0, 1, 2, 3, 4}

	for _, number0 := range numbers {
		for _, number1 := range numbers {
			if number1 == number0 {
				continue
			}

			for _, number2 := range numbers {
				if (number2 == number1) || (number2 == number0) {
					continue
				}

				for _, number3 := range numbers {
					if number3 == number2 || number3 == number1 || number3 == number0 {
						continue
					}

					for _, number4 := range numbers {
						if number4 == number3 || number4 == number2 || number4 == number1 || number4 == number0 {
							continue
						} else {
							combinations = append(combinations, [5]int{number0, number1, number2, number3, number4})
						}
					}
				}

			}

		}
	}

	return combinations
}

func part1(fileName string) int {
	maxOuput := 0
	// maxCombination := [5]int{0, 0, 0, 0, 0}

	for _, combination := range getCombinations() {

		input := 0
		output := 0
		for _, phaseSetting := range combination {
			output = getLastOutput(fileName, phaseSetting, input)
			input = output
		}

		if output > maxOuput {
			maxOuput = output
			// maxCombination = combination
		}
	}

	return maxOuput
}

func main() {
	fmt.Println("Result part 1: ", part1("input.txt"))
}
