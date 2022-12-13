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

func newInt2d(rows int, cols int, defaultValue int) [][]int {
	masterSlice := make([]int, rows*cols) // allocate a continguous chunk of memory for our slice
	res := make([][]int, rows)
	for i := range res {
		res[i] = masterSlice[cols*i : cols*(i+1)] // make subslices
		for j := range res[i] {
			res[i][j] = defaultValue // fill with the default value
		}
	}
	return res
}

func parseInput(input []string) (start coord2D, end coord2D, heights [][]int) {
	rows, cols := len(input), len(input[0])

	//make the heights array
	heights = make([][]int, rows)
	for i := 0; i < rows; i++ {
		heights[i] = make([]int, cols)
	}

	for r := range input {
		for c, char := range input[r] {
			switch char {
			case 'S':
				start = coord2D{r, c}
				heights[r][c] = 0
			case 'E':
				end = coord2D{r, c}
				heights[r][c] = 'z' - 'a'
			default:
				heights[r][c] = int(char - 'a')
			}
		}
	}

	return
}

func minDistanceToEnd(heights [][]int, end coord2D) [][]int {
	/*
		To make part two faster, this calculates the shortest distance from any node on the map to the end.
		If the node is not reachable from the end, it will have a default value of -1.
		Constructed with an Iterative BFS
	*/

	// convenient to have
	rows, cols := len(heights), len(heights[0])

	// make our result structure
	minDistance := newInt2d(rows, cols, -1)

	// helpers for our iterative BFS
	exists := struct{}{}                  // empty struct saves space
	visited := make(map[coord2D]struct{}) // map to track which nodes we've visited
	toVisit := make([]coord2D, 0)         // deque for the nodes we need to visit

	// initialize our stack and visited structure with the end node
	toVisit = append(toVisit, end)
	visited[end] = exists

	// convenience function to return if an item is in the visited set
	isVisited := func(node coord2D) bool {
		_, ok := visited[node]
		return ok
	}

	// loop as long as we have spaces to visit
	for distance := 0; len(toVisit) > 0; distance += 1 {

		// pull all values off the stack and log the distance
		for range toVisit {

			// pop from the stack and update the slice
			node := toVisit[0]
			toVisit = toVisit[1:]

			// check if we have accessable neigbors we haven't visited
			toCheck := []coord2D{
				{node.row + 1, node.col}, // up
				{node.row - 1, node.col}, // down
				{node.row, node.col - 1}, // left
				{node.row, node.col + 1}, // right
			}

			for _, coord := range toCheck {
				if (coord.col >= 0 && coord.col < cols) && // in cols
					(coord.row >= 0 && coord.row < rows) && // in rows
					(!isVisited(coord)) && // haven't visited this yet
					(heights[coord.row][coord.col] >= heights[node.row][node.col]-1) {
					minDistance[coord.row][coord.col] = distance + 1 //	update the distance
					visited[coord] = exists                          //	mark them as visited
					toVisit = append(toVisit, coord)                 //  push them onto the stack
				}
			}
		}
	}
	return minDistance
}

func puzzlePartOne(distances [][]int, start coord2D) int {
	return distances[start.row][start.col]
}

func puzzlePartTwo(distances [][]int, heights [][]int) int {
	rows, cols := len(distances), len(distances[0])
	shortestPath := rows * cols // initialize it to the largest possible path
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if heights[r][c] == 0 &&
				distances[r][c] < shortestPath &&
				distances[r][c] > -1 {
				shortestPath = distances[r][c]
			}
		}
	}
	return shortestPath
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	start, end, heights := parseInput(lines)
	distances := minDistanceToEnd(heights, end)

	partOneSoln := puzzlePartOne(distances, start)
	partTwoSoln := puzzlePartTwo(distances, heights)

	fmt.Printf("Part One Solution: %v\n", partOneSoln)
	fmt.Printf("Part Two Solution: %v\n", partTwoSoln)

}
