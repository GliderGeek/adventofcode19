package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	name     string
	children []*node
}

func getInput(fileName string) [][2]string {
	file, err := os.Open(fileName)
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
	input := [][2]string{}

	for _, line := range lines {
		parentChild := strings.Split(line, ")")
		parent := parentChild[0]
		child := parentChild[1]
		input = append(input, [2]string{parent, child})
	}

	return input
}

func getOrbitMap(entries [][2]string) map[string]*node {
	nodes := map[string]*node{
		"COM": &node{"COM", []*node{}},
	}

	for _, entry := range entries {
		parentNode, present := nodes[entry[0]]
		if !present {
			parentNode = &node{entry[0], []*node{}}
			nodes[parentNode.name] = parentNode
		}

		childNode, present := nodes[entry[1]]
		if !present {
			childNode = &node{entry[1], []*node{}}
			nodes[childNode.name] = childNode

		}

		parentNode.children = append(parentNode.children, childNode)
	}

	return nodes
}

func calculateOrbits(children []*node, level int) int {
	orbits := 0
	for _, child := range children {
		orbits += level
		orbits += calculateOrbits(child.children, level+1)
	}
	return orbits
}

func main() {

	input := getInput("input.txt")
	orbitMap := getOrbitMap(input)
	fmt.Println("orbits", calculateOrbits(orbitMap["COM"].children, 1))

	// fmt.Println("Result part 1 is:", get_last_output("input.txt", 1))
	// fmt.Println("Result part 2 is:", get_last_output("input.txt", 5))
}
