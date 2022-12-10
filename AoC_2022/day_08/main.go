package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type coord2D struct {
	row int
	col int
}

type tree struct {
	height   int
	position coord2D
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

func PuzzlePartOne(forest [][]int) int {
	viewCounter := make([][]int, len(forest)) // counts how many directions our tree is viewable from
	for i := range viewCounter {
		viewCounter[i] = make([]int, len(forest[0]))
	}

	// need a couple of stacks keeping track of the row and column view
	// any tree remaining in the stack is visible from the perimiter of the forest
	// but we need to be aware of double counting
	rowView := make([]treeStack, len(forest)) // viewing along rows
	for i := range rowView {
		rowView[i] = newStack()
	}

	colView := make([]treeStack, len(forest[0])) // viewing long columns
	for i := range colView {
		colView[i] = newStack()
	}

	for r := range forest {
		for c, char := range forest[r] {
			// fmt.Print(r, c)
			// fmt.Printf("--%c--\n", char)
			height := int(char) - 48                // quick conversion
			toAppend := tree{height, coord2D{r, c}} // create an element to insert into stacks
			viewCounter[r][c] = 2                   // initialize the viewCounter value with 2 (we assum we see it from both horiz and vert)

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

type forestViews struct {
	up    [][]int
	down  [][]int
	left  [][]int
	right [][]int
}

type stackHelper struct {
	val   int
	index int
}

func furthestViewLookingLeft(forest [][]int) (view [][]int) {
	view = make([][]int, len(forest))
	for r := range forest {
		view[r] = make([]int, len(forest[r]))
		stack := make([]stackHelper, 0)
		for c := range forest[r] {
			for len(stack) > 0 && forest[r][c] > stack[len(stack)-1].val {
				// pop our stack
				stack = stack[:len(stack)-1]
			}
			// now we can update the view and push onto the stack
			view[r][c] = c
			if len(stack) > 0 {
				view[r][c] -= stack[len(stack)-1].index
			}
			stack = append(stack, stackHelper{forest[r][c], c})
		}
	}
	return
}

func furthestViewLookingRight(forest [][]int) (view [][]int) {
	view = make([][]int, len(forest))
	for r := range forest {
		view[r] = make([]int, len(forest[r]))
		stack := make([]stackHelper, 0)
		for offset := range forest[r] {
			c := len(forest[r]) - offset - 1
			for len(stack) > 0 && forest[r][c] > stack[len(stack)-1].val {
				// pop our stack
				stack = stack[:len(stack)-1]
			}
			// now we can update the view and push onto the stack
			view[r][c] = offset
			if len(stack) > 0 {
				view[r][c] -= stack[len(stack)-1].index
			}
			stack = append(stack, stackHelper{forest[r][c], offset})
		}
	}
	return
}

func furthestViewLookingUp(forest [][]int) (view [][]int) {
	view = make([][]int, len(forest))
	for i := range forest {
		view[i] = make([]int, len(forest[i]))
	}

	for c := range forest[0] {
		stack := make([]stackHelper, 0)
		for r := range forest {
			for len(stack) > 0 && forest[r][c] > stack[len(stack)-1].val {
				// pop our stack
				stack = stack[:len(stack)-1]
			}
			// now we can update the view and push onto the stack
			view[r][c] = r
			if len(stack) > 0 {
				view[r][c] -= stack[len(stack)-1].index
			}
			stack = append(stack, stackHelper{forest[r][c], r})
		}
	}
	return
}

func furthestViewLookingDown(forest [][]int) (view [][]int) {
	view = make([][]int, len(forest))
	for i := range forest {
		view[i] = make([]int, len(forest[i]))
	}

	for c := range forest[0] {
		stack := make([]stackHelper, 0)
		for offset := range forest {
			r := len(forest) - offset - 1
			for len(stack) > 0 && forest[r][c] > stack[len(stack)-1].val {
				// pop our stack
				stack = stack[:len(stack)-1]
			}
			// now we can update the view and push onto the stack
			view[r][c] = offset
			if len(stack) > 0 {
				view[r][c] -= stack[len(stack)-1].index
			}
			stack = append(stack, stackHelper{forest[r][c], offset})
		}
	}
	return
}

func calculateViews(forest [][]int) forestViews {
	return forestViews{
		furthestViewLookingUp(forest),
		furthestViewLookingDown(forest),
		furthestViewLookingLeft(forest),
		furthestViewLookingRight(forest),
	}
}

func parseInput(input []string) [][]int {
	parsed := make([][]int, len(input))
	for r := range input {
		parsed[r] = make([]int, len(input[r]))
		for c := range input[r] {
			parsed[r][c] = int(input[r][c]) - 48 // simple since we know it's ascii 0-9
		}
	}
	return parsed
}

func puzzlePartTwo(views forestViews) int {
	maxViewScore := 0
	for r := range views.up {
		for c := range views.up[r] {
			viewScore := views.up[r][c] * views.down[r][c] * views.left[r][c] * views.right[r][c]
			if viewScore > maxViewScore {
				maxViewScore = viewScore
			}
		}
	}

	return maxViewScore
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	forest := parseInput(lines)

	fmt.Printf("Part One Solution: %v\n", PuzzlePartOne(forest))
	views := calculateViews(forest)
	fmt.Printf("Part Two Solution: %v\n", puzzlePartTwo(views))
}
