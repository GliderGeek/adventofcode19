package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Amplifier struct {
	values     []int
	instructionIndex int
	phaseSetting int
}

func getOpcode(amplifier *Amplifier) int {
	instruction := amplifier.values[amplifier.instructionIndex]
	opcode, _ := parseInstruction(instruction)
	return opcode
}

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

func getValue(amplifier *Amplifier, parameterMode int, parameterIndex int) int {
	switch parameterMode {
	case 0:
		return amplifier.values[amplifier.values[parameterIndex]]
	case 1:
		return amplifier.values[parameterIndex]
	default:
		panic("error")
	}
}

func inputInstruction(amplifier *Amplifier, input int) {
	index := amplifier.values[amplifier.instructionIndex+1]
	amplifier.values[index] = input
	amplifier.instructionIndex += 2
}

func executeInstruction(amplifier *Amplifier) {
	instruction := amplifier.values[amplifier.instructionIndex]
	opcode, parameterModes := parseInstruction(instruction)

	value1 := getValue(amplifier, parameterModes[0], amplifier.instructionIndex+1)
	value2 := getValue(amplifier, parameterModes[1], amplifier.instructionIndex+2)
	index := amplifier.values[amplifier.instructionIndex+3]

	switch opcode {
	case 1:
		amplifier.values[index] = value1 + value2
		amplifier.instructionIndex += 4
	case 2:
		amplifier.values[index] = value1 * value2
		amplifier.instructionIndex += 4
	case 3:
		panic("no opcode 3")
	case 4:
		panic("no opcode 4")
	case 5:
		switch value1 {
		case 0:
			amplifier.instructionIndex += 3
		default:
			amplifier.instructionIndex = value2
		}
	case 6:
		fmt.Println(6)
		switch value1 {
		case 0:
			amplifier.instructionIndex = value2
		default:
			amplifier.instructionIndex += 3
		}
	case 7:
		fmt.Println(7)
		if value1 < value2 {
			amplifier.values[index] = 1
		} else {
			amplifier.values[index] = 0
		}
		amplifier.instructionIndex += 4
	case 8:
		fmt.Println(8)
		if value1 == value2 {
			amplifier.values[index] = 1
		} else {
			amplifier.values[index] = 0
		}
		amplifier.instructionIndex += 4
	case 99:
		panic("called with opcode 99")
	default:
		fmt.Println(opcode)
		panic("Unexpected fault")
	}
}

func start(amplifier *Amplifier){

	appliedPhaseSetting := false

	for {
		instruction := amplifier.values[amplifier.instructionIndex]
		opcode, _ := parseInstruction(instruction)

		if opcode == 3{
			if appliedPhaseSetting{
				//ready for new input
				return
			} else{
				inputInstruction(amplifier, amplifier.phaseSetting)
				appliedPhaseSetting = true
			}
		} else if opcode == 4{
			panic("not expected output")
		} else if opcode == 99{
			panic("not expected end")
		} else{
			executeInstruction(amplifier)
		}
	}
}

func getOutput(amplifier *Amplifier, input int) int {

	output := input

	for {

		instruction := amplifier.values[amplifier.instructionIndex]
		opcode, parameterModes := parseInstruction(instruction)

		switch opcode {
		case 3:
			inputInstruction(amplifier, input)
		case 4:
			if parameterModes[0] == 0 {
				output = amplifier.values[amplifier.values[amplifier.instructionIndex+1]]
			} else {
				output = amplifier.values[amplifier.instructionIndex+1]
			}
			amplifier.instructionIndex += 2
			return output
		case 99:
			return output
		default:
			executeInstruction(amplifier)
		}
	}
}

func getLastOutput(amplifier *Amplifier, input int) int {
	start(amplifier)
	for {
		output := getOutput(amplifier, input)
		if getOpcode(amplifier) == 99{
			return output
		}
	}
}

func getCombinations(numbers [5]int) [][5]int {

	combinations := [][5]int{}

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
	numbers := [5]int{0, 1, 2, 3, 4}
	for _, combination := range getCombinations(numbers) {

		input := 0
		output := 0
		for _, phaseSetting := range combination {
			values := get_input(fileName)

			amplifier := Amplifier{values, 0, phaseSetting}
			output = getLastOutput(&amplifier, input)
			input = output
		}

		if output > maxOuput {
			maxOuput = output
			// maxCombination = combination
		}
	}

	return maxOuput
}

func part2(fileName string) int{

	maxOutput := 0
	numbers := [5]int{5, 6, 7, 8, 9}
	for _, combination := range getCombinations(numbers) {

		amplifiers := [5]*Amplifier{
			&Amplifier{get_input(fileName), 0, combination[0]},
			&Amplifier{get_input(fileName), 0, combination[1]},
			&Amplifier{get_input(fileName), 0, combination[2]},
			&Amplifier{get_input(fileName), 0, combination[3]},
			&Amplifier{get_input(fileName), 0, combination[4]},
		}

		input := 0
		output := 0

		outside: for {
			for _, amplifier := range amplifiers{
				input = output

				if amplifier.instructionIndex == 0{
					start(amplifier)
				} else {
					output = getOutput(amplifier, input)
				}

				if getOpcode(amplifier) == 99 {
					break outside
				}
			}
		}

		if output > maxOutput{
			maxOutput = output
		}
	}

	return maxOutput
}

func main() {
	fmt.Println("Result part 1: ", part1("example.txt"))
	// fmt.Println("Result part 2: ", part2("example.txt"))
}
