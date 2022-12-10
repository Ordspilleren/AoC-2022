package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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
	s.Crates = append([]Crate{c}, s.Crates...)
}

func (s *Stack) Pop(quantity int) {
	n := (len(s.Crates) - 1) - quantity
	s.Crates = s.Crates[:n]
}

func (s *Stack) Move(quantity int, to *Stack, crateMover9001 bool) {
	if crateMover9001 {
		n := len(s.Crates) - quantity
		to.Crates = append(to.Crates, s.Crates[n:]...)
		s.Crates = s.Crates[:n]
	} else {
		for i := 0; i < quantity; i++ {
			n := (len(s.Crates) - 1)
			to.Crates = append(to.Crates, s.Crates[n])
			s.Crates = s.Crates[:n]
		}
	}
}

func main() {
	stacks := make(Stacks)

	data, err := os.Open("input.txt")
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
					currentStack++
					continue
				}
				log.Printf("Crate: %c", name)
				log.Printf("Stack: %d", currentStack)
				create := Crate{Name: rune(name)}
				if _, ok := stacks[currentStack]; !ok {
					stacks[currentStack] = &Stack{}
					log.Println("New stack made")
				}
				log.Println()
				stacks[currentStack].Push(create)
				currentStack++
			}
		} else {
			var quantity int
			var from int
			var to int
			r := regexp.MustCompile("move ([0-9]*?) from ([0-9]*?) to ([0-9]*?)$")
			match := r.FindStringSubmatch(scanner.Text())
			quantity, _ = strconv.Atoi(match[1])
			from, _ = strconv.Atoi(match[2])
			to, _ = strconv.Atoi(match[3])
			stacks[from].Move(quantity, stacks[to], true)
		}
	}

	var sortedStacksIds []int
	for k := range stacks {
		sortedStacksIds = append(sortedStacksIds, k)
	}
	sort.Ints(sortedStacksIds)

	for _, stackId := range sortedStacksIds {
		n := len(stacks[stackId].Crates) - 1
		topCrate := stacks[stackId].Crates[n]

		log.Printf("Stack: %d, Top Crate: %c", stackId, topCrate.Name)
	}
}
