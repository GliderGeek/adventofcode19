package main

import (
    "fmt"
    "strconv"
    "strings"
)

func get_individual_numbers(number int) [6]int {
    stringed_number := strconv.Itoa(number)
    var individual_number int
    var individual_numbers [6]int

    for i, char := range strings.Split(stringed_number, ""){
        individual_number, _ = strconv.Atoi(string(char))
        individual_numbers[i] = individual_number
    }

    return individual_numbers
}

func password_meets_criteria_part1(password [6]int) bool{
    double_adjacent_numbers := false
    for index, number := range password {
        if index != 0{
            if number < password[index-1]{
                return false
            } else {
                if number == password[index-1]{
                    double_adjacent_numbers = true
                }
            }
        }
    }

    return double_adjacent_numbers
}

func password_meets_criteria_part2(password [6]int) bool{

    var repeating_numbers map[int]int
    repeating_numbers = make(map[int]int)
    var occurences int
    var present bool

    for index, number := range password {
        if index != 0{
            if number < password[index-1]{
                return false
            } else {
                if number == password[index-1]{

                    occurences, present = repeating_numbers[number]
                    if present{
                        repeating_numbers[number] = repeating_numbers[number] + 1
                    } else {
                        repeating_numbers[number] = 2
                    }
                }
            }
        }
    }

    for _, occurences = range repeating_numbers{
        if occurences == 2{
            return true
        }
    }
    return false
}

func part1(start_number int, end_number int){
    number_of_possible_passwords := 0
    var individual_numbers [6]int

    for number := start_number; number < end_number; number++ {
        individual_numbers = get_individual_numbers(number)
        if password_meets_criteria_part1(individual_numbers) {
            number_of_possible_passwords++
        }
    }

    fmt.Println("Solution part1:", number_of_possible_passwords)
}

func part2(start_number int, end_number int){
    number_of_possible_passwords := 0
    var individual_numbers [6]int

    for number := start_number; number < end_number; number++ {
        individual_numbers = get_individual_numbers(number)
        if password_meets_criteria_part2(individual_numbers) {
            number_of_possible_passwords++
        }
    }

    fmt.Println("Solution part2:", number_of_possible_passwords)
}

func main() {
    part1(156218, 652527)
    part2(156218, 652527)
}
