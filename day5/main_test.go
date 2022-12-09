package main_test

import "log"

func main() {
	stack1 := Stack{
		Id: 1,
		Crates: []Crate{
			{Name: "Z"},
			{Name: "N"},
		},
	}
	stack2 := Stack{
		Id: 2,
		Crates: []Crate{
			{Name: "M"},
			{Name: "C"},
			{Name: "D"},
		},
	}

	log.Print(stack1)
	log.Print(stack2)
	stack2.Move(1, &stack1)
	log.Print(stack1)
	log.Print(stack2)
}
