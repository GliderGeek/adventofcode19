package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "math"
)

func get_input(file_name string) []string {
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

    return lines
}


func calculate_fuel_required(mass int) int {
    return int(math.Floor(float64(mass) / 3)) - 2
}

func calculate_fuel_required_including_fuel(mass int) int {
    
    total := 0

    for mass > 0 {
        mass = calculate_fuel_required(mass)
        if mass < 0 {
            mass = 0
        } else {
            total += mass
        }
    }

    return total
}

func part1(){
    lines := get_input("input.txt")

    total := 0
    for _, line := range lines {
        mass, _ := strconv.Atoi(line)
        total += calculate_fuel_required(mass)
    }

    fmt.Print("Result part 1 is: ")
    fmt.Println(total)
}

func part2(){
    lines := get_input("input.txt")

    total := 0
    for _, line := range lines {
        mass, _ := strconv.Atoi(line)
        total += calculate_fuel_required_including_fuel(mass)
    }

    fmt.Print("Result part 2 is: ")
    fmt.Println(total)
}

func main() { 
    part1()
    part2()
}
