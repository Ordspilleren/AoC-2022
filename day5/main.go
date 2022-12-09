package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Crate struct {
	Name rune
}

type Stack struct {
	Crates []Crate
}

type Stacks map[int]*Stack

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
	stacks := make(Stacks)

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
			continue
		}
		if !stacksDone {
			currentStack := 1
			for i := 1; i < len(scanner.Text()); i += 4 {
				name := scanner.Text()[i]
				if !strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", rune(name)) {
					continue
				}
				currentStack++
				log.Printf("Crate: %c", name)
				log.Printf("Stack: %d", currentStack)
				create := Crate{Name: rune(name)}
				if _, ok := stacks[currentStack]; !ok {
					stacks[currentStack] = &Stack{}
					log.Println("New stack made")
				}
				stacks[currentStack].Push(create)
				currentStack++
			}
		} else {
			log.Println(scanner.Text())
		}
	}
	log.Printf("%c", stacks[1].Crates[0].Name)
}
