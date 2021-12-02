package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"

// Steering handles Navigation of the Submarine
type Steering struct {
	Aim        int
	Depth      int
	Horizontal int
}

// Navigate your Submarine by getting the direction and how much
// the submarine needs to move
func (s *Steering) Navigate(direction string, value int) {
	switch direction {
	case "forward":
		s.Horizontal += value
		s.Depth += s.Aim * value
	case "down":
		// s.Depth += value // Initial Task
		s.Aim += value
	case "up":
		// s.Depth -= value // Initial Task
		s.Aim -= value
	}
}

// Product of Horizontal and Depth values
func (s Steering) Product() int {
	return s.Horizontal * s.Depth
}

func main() {

	// Starting the Submarine
	steering := &Steering{Horizontal: 0, Depth: 0, Aim: 0}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// incoming data is in format: <string int>
		// e.g., forward 5, down 3, up 2
		commands := strings.Split(scanner.Text(), " ")
		units, _ := strconv.Atoi(commands[1])
		// drive the sub!!
		steering.Navigate(commands[0], units)
	}

	// Complete Course Map information
	fmt.Printf("aim: %d, depth: %d, horizontal: %d\n", steering.Aim, steering.Depth, steering.Horizontal)
	fmt.Printf("Product: %d\n", steering.Product())
}
