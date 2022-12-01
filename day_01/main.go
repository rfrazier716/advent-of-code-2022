package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Backpack = []int
type Elves = []Backpack

// Create a Heap for Part 2, since we only need the top three backpacks
type IntHeap []int

// required for the sort.Interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func partOne(elfBrigade Elves) int {
	// this is just keeping track of the max
	max := 0
	for _, pack := range elfBrigade {
		if calories := sum(pack); calories > max {
			max = calories
		}
	}
	return max
}

func partTwo(elfBrigade Elves) int {
	h := IntHeap{}
	heap.Init(&h)

	for _, pack := range elfBrigade {
		calories := sum(pack)
		heap.Push(&h, calories)
		for len(h) > 3 {
			heap.Pop(&h)
		}
	}

	return sum(h)
}

func sum(slice []int) (res int) {
	res = 0
	for i := range slice {
		res += slice[i]
	}
	return
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")

	// Create our Elves on a Hike
	elfBrigade := Elves{make(Backpack, 0)}
	for i := range lines {
		if num, err := strconv.Atoi(lines[i]); err == nil {
			elfBrigade[len(elfBrigade)-1] = append(elfBrigade[len(elfBrigade)-1], num)
		} else if len(lines[i]) == 0 {
			// empty line - make a new struct
			elfBrigade = append(elfBrigade, make(Backpack, 0))
		} else {
			// Shouldn't hit this...
			log.Printf("Could not parse line %v: %v", i, lines[i])
		}
	}

	fmt.Printf("Part One Solution: %v\n", partOne(elfBrigade))
	fmt.Printf("Part Two Solution: %v\n", partTwo(elfBrigade))
}
