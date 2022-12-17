package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type position struct {
	Row    int
	Column int
}

type visit struct {
	Head bool
	Tail bool
}

type Rope struct {
	CurrentHeadPosition position
	CurrentTailPosition position
	Motions             [][]visit
}

func CreateRope() *Rope {
	motions := make([][]visit, 1000)
	for i := range motions {
		motions[i] = make([]visit, 1000)
	}
	currentPosition := position{500, 500}
	motions[currentPosition.Row][currentPosition.Column].Head = true
	motions[currentPosition.Row][currentPosition.Column].Tail = true
	return &Rope{
		CurrentHeadPosition: currentPosition,
		CurrentTailPosition: currentPosition,
		Motions:             motions,
	}
}

func (r *Rope) SetHeadPosition() {
	r.Motions[r.CurrentHeadPosition.Row][r.CurrentHeadPosition.Column].Head = true
}

func (r *Rope) SetTailPosition(direction Direction) {
	r.CalculateTailKinematics(direction)
	r.Motions[r.CurrentTailPosition.Row][r.CurrentTailPosition.Column].Tail = true
}

func (r *Rope) CalculateTailKinematics(direction Direction) {
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row-1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column+1 {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row-1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column-1 {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row+1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column-1 {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row+1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column+1 {
		return
	}

	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row-1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row+1 && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column+1 {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column-1 {
		return
	}
	if r.CurrentHeadPosition.Row == r.CurrentTailPosition.Row && r.CurrentHeadPosition.Column == r.CurrentTailPosition.Column {
		return
	}

	switch direction {
	case Up:
		r.CurrentTailPosition.Row = r.CurrentHeadPosition.Row + 1
		r.CurrentTailPosition.Column = r.CurrentHeadPosition.Column
	case Down:
		r.CurrentTailPosition.Row = r.CurrentHeadPosition.Row - 1
		r.CurrentTailPosition.Column = r.CurrentHeadPosition.Column
	case Left:
		r.CurrentTailPosition.Row = r.CurrentHeadPosition.Row
		r.CurrentTailPosition.Column = r.CurrentHeadPosition.Column + 1
	case Right:
		r.CurrentTailPosition.Row = r.CurrentHeadPosition.Row
		r.CurrentTailPosition.Column = r.CurrentHeadPosition.Column - 1
	}
}

func (r *Rope) Move(direction Direction, steps int) {
	for i := 0; i < steps; i++ {
		switch direction {
		case Up:
			r.CurrentHeadPosition.Row = r.CurrentHeadPosition.Row - 1
			r.SetHeadPosition()
		case Down:
			r.CurrentHeadPosition.Row = r.CurrentHeadPosition.Row + 1
			r.SetHeadPosition()
		case Left:
			r.CurrentHeadPosition.Column = r.CurrentHeadPosition.Column - 1
			r.SetHeadPosition()
		case Right:
			r.CurrentHeadPosition.Column = r.CurrentHeadPosition.Column + 1
			r.SetHeadPosition()
		}
		r.SetTailPosition(direction)
	}
}

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	rope := CreateRope()
	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(move[1])
		switch move[0] {
		case "U":
			rope.Move(Up, steps)
		case "D":
			rope.Move(Down, steps)
		case "L":
			rope.Move(Left, steps)
		case "R":
			rope.Move(Right, steps)
		}
	}

	tailVisits := 0
	for row := range rope.Motions {
		for column := range rope.Motions[row] {
			if !rope.Motions[row][column].Tail {
				fmt.Print(".")
			}
			if rope.Motions[row][column].Tail {
				tailVisits++
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	log.Printf("Total tail visits: %d", tailVisits)
}
