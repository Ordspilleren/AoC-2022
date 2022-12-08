package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

type Tournament struct {
	Rounds     []Round
	TotalScore int
}

type Round struct {
	OpponentMove Move
	OwnMove      Move
	Score        int
}

func MapMove(move string) (Move, error) {
	switch move {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return 0, nil
	}
}

func MapMoveOnRoundOutcome(opponentMove Move, outcome string) (Move, error) {
	if outcome == "X" {
		switch opponentMove {
		case Rock:
			return Scissors, nil
		case Paper:
			return Rock, nil
		case Scissors:
			return Paper, nil
		}
	} else if outcome == "Y" {
		return opponentMove, nil
	} else if outcome == "Z" {
		switch opponentMove {
		case Rock:
			return Paper, nil
		case Paper:
			return Scissors, nil
		case Scissors:
			return Rock, nil
		}
	}

	return 0, nil
}

func (r *Round) Play() {
	score := 0

	switch r.OwnMove {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	switch {
	case r.OpponentMove == Rock && r.OwnMove == Paper:
		score += 6
	case r.OpponentMove == Paper && r.OwnMove == Scissors:
		score += 6
	case r.OpponentMove == Scissors && r.OwnMove == Rock:
		score += 6
	case r.OpponentMove == r.OwnMove:
		score += 3
	}

	r.Score = score
}

func main() {
	tournament := Tournament{}
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		round := Round{}
		moves := strings.Split(scanner.Text(), " ")
		round.OpponentMove, err = MapMove(moves[0])
		if err != nil {
			log.Panic("Error mapping opponent move")
		}
		//round.OwnMove, err = MapMove(moves[1])
		round.OwnMove, err = MapMoveOnRoundOutcome(round.OpponentMove, moves[1])
		if err != nil {
			log.Panic("Error mapping own move")
		}
		round.Play()
		tournament.Rounds = append(tournament.Rounds, round)
	}

	for _, round := range tournament.Rounds {
		tournament.TotalScore += round.Score
	}

	log.Print(tournament)

	log.Printf("Tournament Score: %d\n", tournament.TotalScore)
}
