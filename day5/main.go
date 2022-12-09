package main

import (
	"bufio"
	"log"
	"os"
)

type Crate struct {
	Name rune
}

type Stack struct {
	Id     int
	Crates []Crate
}

func (s *Stack) Push(c Crate) {
	s.Crates = append(s.Crates, c)
}

func (s *Stack) Pop(quantity int) {
	n := (len(s.Crates) - 1) - quantity
	s.Crates = s.Crates[:n]
}

func (s *Stack) Move(quantity int, to *Stack) {
	n := len(s.Crates) - quantity
	to.Crates = append(to.Crates, s.Crates[n:]...)
	s.Crates = s.Crates[:n]
}

func main() {
	data, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	stacksDone := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			stacksDone = true
		}
		currentStack := 1
		for i := 1; i < len(scanner.Text()); i += 4 {
			log.Printf("Crate: %c", scanner.Text()[i])
			log.Printf("Stack: %d", currentStack)
			currentStack++
		}
	}
}
