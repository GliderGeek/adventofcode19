package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func get_input(file_name string) [2][]string {
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

    // var value int
    var wires [2][]string

    for i, line := range lines{
        wires[i] = strings.Split(line, ",")
    }

    return wires
}

func evaluate_visited(visited map[string]int, x int, y int, wire_index int, intersection_keys []string) (map[string]int, []string)  {
    //TODO: use pointers for extra performance

    visited_key := get_key(x, y)
    associated_index, present := visited[visited_key]

    if present {
        if associated_index != wire_index{
            intersection_keys = append(intersection_keys, visited_key)
        }
    } else {
        visited[visited_key] = wire_index
    }

    return visited, intersection_keys
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func get_values(key string) (int, int) {
    values := strings.Split(key, ".")
    x, _ := strconv.Atoi(values[0])
    y, _ := strconv.Atoi(values[1])
    return x, y
}

func get_key(x int, y int) string {
    return strconv.Itoa(x) + "." + strconv.Itoa(y)
}

func get_minimum_manhattan_distance(intersection_keys []string) int {

    min := 0
    var x int
    var y int

    for _, key := range intersection_keys {
        x, y = get_values(key)
        if min == 0 || (abs(x) + abs(y)) < min {
            min = abs(x) + abs(y)
        }
    }

    return min
}

func get_intersection_keys(wires [2][]string) []string{
    var value int
    var visited map[string]int
    visited = make(map[string]int)
    var intersection_keys []string
    x := 0
    y := 0

    for wire_index, wire := range wires{

        x = 0
        y = 0

        for _, instruction := range wire{
            value, _ = strconv.Atoi(string(instruction[1:]))
            for i := 0; i < value; i++ {
                switch direction := instruction[0]; direction {
                case 'R':
                    x++
                case 'L':
                    x--
                case 'U':
                    y++
                case 'D':
                    y--
                }
                visited, intersection_keys = evaluate_visited(visited, x, y, wire_index, intersection_keys)
            }
        }
    }

    return intersection_keys
}

func update_minimum_steps(minimum_steps map[string]int, x int, y int, steps int) map[string]int {
    //Check if with the latest step a new cross-section is reached with a minimum distance

    key := get_key(x, y)
    steps_in_map, present := minimum_steps[key]
    if present && (steps_in_map == 0 || steps < steps_in_map){
        minimum_steps[key] = steps
    }
    return minimum_steps
}

func get_minimum_number_of_steps(intersection_keys []string, wires [2][]string ) int {

    //map from intersection key to minimum steps needed
    //initialized with 0-s
    //re-used per wire
    var minimum_steps map[string]int  
    minimum_steps = make(map[string]int)
    
    //map from intersection key to minimum steps
    //not used per wire, global
    var steps_per_intersection map[string]int
    steps_per_intersection = make(map[string]int)
    for _, key := range intersection_keys{
        steps_per_intersection[key] = 0
    }

    var steps int
    var value int
    var x int
    var y int

    for _, wire := range wires{

        for _, key := range intersection_keys{
            minimum_steps[key] = 0
        }

        steps = 0
        x = 0
        y = 0

        for _, instruction := range wire{

            value, _ = strconv.Atoi(string(instruction[1:]))
            for i := 0; i < value; i++ {
                switch direction := instruction[0]; direction {
                case 'R': 
                    x++
                case 'L':
                    x--
                case 'U':
                    y++
                case 'D':
                    y--
                }
                steps += 1
                minimum_steps = update_minimum_steps(minimum_steps, x, y, steps)
            }
        }

        for key, value := range minimum_steps{
            steps_per_intersection[key] += value
        }
    }

    //find minimum number of steps
    min := 0
    for _, val := range steps_per_intersection{
        if min ==0 || val < min {
            min = val
        }
    }

    return min
}

func part1(){
    wires := get_input("input.txt")
    intersection_keys := get_intersection_keys(wires)
    fmt.Println("Solution part1:", get_minimum_manhattan_distance(intersection_keys))
}

func part2(){
    wires := get_input("input.txt")
    intersection_keys := get_intersection_keys(wires)
    fmt.Println("Solution part2:", get_minimum_number_of_steps(intersection_keys, wires))
}

func main() { 
    part1()
    part2()
}
