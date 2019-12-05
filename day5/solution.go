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
    for _, char := range strings.Split(lines[0], ","){
        value, _ = strconv.Atoi(char)
        values = append(values, value)
    }

    return values
}

func parse_instruction(instruction int) (int, []int) {
    opcode := instruction % 100
    var parameter_modes []int
    remainder := instruction

    third_mode := instruction / 10000
    remainder -= third_mode * 10000
    second_mode := remainder / 1000
    remainder -= second_mode * 1000
    first_mode := remainder / 100

    switch opcode{
    case 3, 4: //only single parameter_mode
        parameter_modes = []int{first_mode}
    case 5, 6: //two parameter_modes
        parameter_modes = []int{first_mode, second_mode}
    case 1, 2, 7, 8: //three parameter modes
        parameter_modes = []int{first_mode, second_mode, third_mode}
    }

    return opcode, parameter_modes
}

func get_value(values []int, parameter_mode int, parameter_index int) int {
    switch parameter_mode {
    case 0:
        return values[values[parameter_index]]
    case 1:
        return values[parameter_index]
    default:
        panic("error")
    }
}

func get_last_output(filename string, input int) int{

    var value1, value2, output int
    values := get_input(filename)
    instruction_index := 0
    instruction := values[instruction_index]
    opcode, parameter_modes := parse_instruction(instruction)

    for opcode != 99 {   

        switch opcode {
        case 1:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)
            values[values[instruction_index+3]] = value1 + value2
            instruction_index += 4
        case 2:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)
            values[values[instruction_index+3]] = value1 * value2
            instruction_index += 4
        case 3:
            values[values[instruction_index+1]] = input
            instruction_index += 2
        case 4:
            if parameter_modes[0] == 0{
                output = values[values[instruction_index+1]]
            } else {
                output  = values[instruction_index+1]
            }
            fmt.Println("outputting", output)
            instruction_index += 2
        case 5:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)

            switch value1 {
            case 0:
                instruction_index += 3
            default:
                instruction_index = value2
            }
        case 6:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)

            switch value1 {
            case 0:
                instruction_index = value2
            default:
                instruction_index += 3
            }
        case 7:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)

            if value1 < value2 {
                values[values[instruction_index+3]] = 1
            } else {
                values[values[instruction_index+3]] = 0
            }
            instruction_index += 4
        case 8:
            value1 = get_value(values, parameter_modes[0], instruction_index+1)
            value2 = get_value(values, parameter_modes[1], instruction_index+2)

            if value1 == value2 {
                values[values[instruction_index+3]] = 1
            } else {
                values[values[instruction_index+3]] = 0
            }

            instruction_index += 4
        default:
            fmt.Println(opcode)
            panic("Unexpected fault")
        }
        instruction = values[instruction_index]
        opcode, parameter_modes = parse_instruction(instruction)
    } 

    return output
}

func main() { 
    fmt.Println("Result part 1 is:", get_last_output("input.txt", 1))
    fmt.Println("Result part 2 is:", get_last_output("input.txt", 5))
}
