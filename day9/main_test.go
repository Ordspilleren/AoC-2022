package main

import "testing"

func TestMoving(t *testing.T) {
	rope := CreateRope()
	rope.Move(Right, 4)
	rope.Move(Up, 4)
	rope.Move(Left, 3)
	rope.Move(Down, 1)
	rope.Move(Right, 4)
	rope.Move(Down, 1)
	rope.Move(Left, 5)
	rope.Move(Right, 2)
}
