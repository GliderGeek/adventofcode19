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

func get_index0_value(noun int, verb int) int {
    values := get_input("input.txt")

    values[1] = noun
    values[2] = verb

    opcode_index := 0
    opcode := values[opcode_index]
    for opcode != 99 {

        if opcode == 1{
            values[values[opcode_index+3]] = values[values[opcode_index+1]] + values[values[opcode_index+2]]
        } else if opcode == 2 {
            values[values[opcode_index+3]] = values[values[opcode_index+1]] * values[values[opcode_index+2]]
        } else{
            fmt.Println("Something went wrong")
        }

        opcode_index += 4
        opcode = values[opcode_index]

    } 

    return values[0]
}

func part1(){
    fmt.Println("Result part 1 is:", get_index0_value(12, 2))
}

func part2(){
    found := false
    target := 19690720
    var result int

    for noun := 0; noun < 100; noun ++ {
        for verb := 0; verb < 100; verb ++ {
            if target == get_index0_value(noun, verb){
                found = true
                result = 100 * noun + verb
                break
            }
        }

        if found {
            break
        }
    }

    fmt.Println("Result part 2 is:", result)
}

func main() { 
    part1()
    part2()
}
