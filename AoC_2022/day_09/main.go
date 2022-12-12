package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction2D int

const (
	UP    Direction2D = 0
	DOWN  Direction2D = 1
	LEFT  Direction2D = 2
	RIGHT Direction2D = 3
)

type coord2D struct {
	x int
	y int
}

func AbsInt(a int) int {
	return AbsDiffInt(a, 0)
}

func AbsDiffInt(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func IntSign(a int) int {
	if a >= 0 {
		return 1
	}

	return -1
}

type Mover interface {
	Move(Direction2D)
	TailPosition() coord2D
}

type rope struct {
	head *coord2D
	tail *coord2D
}

func NewRope() rope {
	return rope{
		&coord2D{0, 0},
		&coord2D{0, 0},
	}
}

func (t *rope) constrain() {
	// constrains the position of the tail based on the location of the head

	delX := t.head.x - t.tail.x
	delY := t.head.y - t.tail.y

	// if the rope is more than one away in any direction
	if AbsInt(delX) > 1 || AbsInt(delY) > 1 {
		if delX != 0 {
			t.tail.x += IntSign(delX)
		}
		if delY != 0 {
			t.tail.y += IntSign(delY)
		}
	}
}

func (t *rope) Move(direction Direction2D) {
	switch direction {
	case UP:
		t.head.y += 1
	case DOWN:
		t.head.y -= 1
	case LEFT:
		t.head.x -= 1
	case RIGHT:
		t.head.x += 1
	}
	t.constrain()
}

func (t *rope) TailPosition() coord2D {
	return *t.tail
}

type ropeChain struct {
	knots []coord2D
	ropes []rope
}

func (t *ropeChain) Move(direction Direction2D) {
	if t != nil && len(t.ropes) > 0 {
		t.ropes[0].Move(direction)
	}

	if len(t.ropes) > 1 {
		for _, rope := range t.ropes[1:] {
			rope.constrain()
		}
	}
}

func (t *ropeChain) TailPosition() coord2D {
	return t.knots[len(t.knots)-1]
}

func NewRopeChain(nSegments int) ropeChain {
	knots := make([]coord2D, nSegments+1)
	ropes := make([]rope, nSegments)
	for i := 0; i < nSegments; i++ {
		ropes[i].head = &knots[i]
		ropes[i].tail = &knots[i+1]
	}

	return ropeChain{
		knots,
		ropes,
	}
}

func MovementTracker(input []string, chain Mover) int {
	visitedPositions := make(map[coord2D]struct{})
	for i := range input {
		commandDirection := input[i][:1] // first character
		commandCount, _ := strconv.Atoi(input[i][2:])
		var dir Direction2D
		switch commandDirection {
		case "U":
			dir = UP
		case "D":
			dir = DOWN
		case "L":
			dir = LEFT
		case "R":
			dir = RIGHT
		}

		for i := 0; i < commandCount; i++ {
			// move the rope that many positions
			chain.Move(dir)
			visitedPositions[chain.TailPosition()] = struct{}{}
		}
	}
	return len(visitedPositions)
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]

	rope := NewRope()
	fmt.Printf("Part One Solution: %v\n", MovementTracker(lines, &rope))
	chain := NewRopeChain(9)
	fmt.Printf("Part Two Solution: %v\n", MovementTracker(lines, &chain))
}
