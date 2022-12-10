package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type visibility struct {
	northern int
	southern int
	eastern  int
	western  int
}

type coord2D struct {
	row int
	col int
}

type tree struct {
	height     int
	position   coord2D
	visibility visibility
}

type treeStack struct {
	trees          []*tree
	maxHeightIndex int
}

func tallest(t *treeStack) *tree {
	// returns a pointer to the tallest tree referenced in the stack
	return t.trees[t.maxHeightIndex]
}

func newStack() treeStack {
	return treeStack{
		make([]*tree, 0),
		-1,
	}
}

func updateStack(stack *treeStack, element *tree) (popped []*tree) {
	// handle an empty slice
	popped = make([]*tree, 0)

	if len(stack.trees) == 0 {
		stack.trees = append(stack.trees, element)
		stack.maxHeightIndex = 0
		return // we didn't pop anything so return an empty slice
	}

	// while the element is taller than the right-most element in the stack, and the right-most element is not the tallest element, pop
	for (len(stack.trees) > stack.maxHeightIndex+1) && stack.trees[len(stack.trees)-1].height <= element.height {
		popped = append(popped, stack.trees[len(stack.trees)-1])
		stack.trees = stack.trees[:len(stack.trees)-1]
	}

	// append our element onto the stack
	stack.trees = append(stack.trees, element)

	if element.height > tallest(stack).height {
		// if we've pushed an even taller element on it will end up immediately appended after the current tallest
		stack.maxHeightIndex += 1
	}

	return popped
}

func PuzzlePartOne(input []string) int {
	viewCounter := make([][]int, len(input)) // counts how many directions our tree is viewable from
	for i := range viewCounter {
		viewCounter[i] = make([]int, len(input[0]))
	}

	// need a couple of stacks keeping track of the row and column view
	// any tree remaining in the stack is visible from the perimiter of the forest
	// but we need to be aware of double counting
	rowView := make([]treeStack, len(input)) // viewing along rows
	for i := range rowView {
		rowView[i] = newStack()
	}

	colView := make([]treeStack, len(input[0])) // viewing long columns
	for i := range colView {
		colView[i] = newStack()
	}

	for r := range input {
		for c, char := range input[r] {
			// fmt.Print(r, c)
			// fmt.Printf("--%c--\n", char)
			height := int(char) - 48                              // quick conversion
			toAppend := tree{height, coord2D{r, c}, visibility{}} // create an element to insert into stacks
			viewCounter[r][c] = 2                                 // initialize the viewCounter value with 2 (we assum we see it from both horiz and vert)

			popped := updateStack(&rowView[r], &toAppend)
			for i := range popped {
				viewCounter[popped[i].position.row][popped[i].position.col] -= 1
			}

			popped = updateStack(&colView[c], &toAppend)
			for i := range popped {
				viewCounter[popped[i].position.row][popped[i].position.col] -= 1
			}

			// fmt.Println(colView)
			// fmt.Println(rowView)
			// fmt.Println()
		}
	}

	// iterate over the viewCounter array and count how many trees are visible from at least one direction
	visibleTrees := 0
	for r := range viewCounter {
		for _, val := range viewCounter[r] {
			if val > 0 {
				visibleTrees += 1
			}
		}
	}
	return visibleTrees
}

// func VisibilityScore(vis Visibility) int {
// 	return vis.North * vis.South * vis.East * vis.West
// }

// func PuzzlePartTwo(input []string) int {
// 	maxVisibility := 0

// 	visibilityMemo := make([][]*Visibility, len(input))
// 	for i := range input {
// 		visibilityMemo[i] = make([]*Visibility, len(input[i]))
// 	}

// 	var dfsHelper func(row int, col int) int
// 	dfsHelper = func(row int, col int) int {
// 		if visibilityMemo[row][col] != nil {
// 			return visibilityMemo
// 		}
// 	}

// }

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	fmt.Printf("Part One Solution: %v", PuzzlePartOne(lines))
}
