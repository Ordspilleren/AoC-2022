package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

var stack1 = Stack{
	Crates: []Crate{
		{Name: 'Z'},
		{Name: 'N'},
	},
}
var stack2 = Stack{
	Crates: []Crate{
		{Name: 'M'},
		{Name: 'C'},
		{Name: 'D'},
	},
}
var stack3 = Stack{
	Crates: []Crate{
		{Name: 'P'},
	},
}

func TestStackMovingSingleCrates(t *testing.T) {
	stack2.Move(1, &stack1, false)
	stack1.Move(3, &stack3, false)
	stack2.Move(2, &stack1, false)
	stack1.Move(1, &stack2, false)

	if !slices.Equal(stack1.Crates, []Crate{{Name: 'C'}}) {
		t.Errorf("Stack 1 contained the wrong crates")
	}
	if !slices.Equal(stack2.Crates, []Crate{{Name: 'M'}}) {
		t.Errorf("Stack 2 contained the wrong crates")
	}
	if !slices.Equal(stack3.Crates, []Crate{{Name: 'P'}, {Name: 'D'}, {Name: 'N'}, {Name: 'Z'}}) {
		t.Errorf("Stack 3 contained the wrong crates")
	}
}

func TestStackMovingMultipleCrates(t *testing.T) {
	stack2.Move(1, &stack1, true)
	stack1.Move(3, &stack3, true)
	stack2.Move(2, &stack1, true)
	stack1.Move(1, &stack2, true)

	if !slices.Equal(stack1.Crates, []Crate{{Name: 'M'}}) {
		t.Errorf("Stack 1 contained the wrong crates")
	}
	if !slices.Equal(stack2.Crates, []Crate{{Name: 'C'}}) {
		t.Errorf("Stack 2 contained the wrong crates")
	}
	if !slices.Equal(stack3.Crates, []Crate{{Name: 'P'}, {Name: 'Z'}, {Name: 'N'}, {Name: 'D'}}) {
		t.Errorf("Stack 3 contained the wrong crates")
	}
}
