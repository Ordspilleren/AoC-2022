package main

import "log"

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
	CurrentPosition position
	Motions         [][]visit
}

func CreateRope() *Rope {
	motions := make([][]visit, 0)
	motions[0][0].Head = true
	motions[0][0].Tail = true
	return &Rope{
		CurrentPosition: position{0, 0},
		Motions:         motions,
	}
}

func (r *Rope) Move(direction Direction, steps int) {

}

func main() {
	rope := CreateRope()
	log.Print(rope)
}
