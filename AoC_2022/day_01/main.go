package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
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

func nLargest(maxElements int) func(int) IntHeap {
	h := IntHeap{}
	heap.Init(&h)

	return func(val int) IntHeap {
		heap.Push(&h, val)
		for len(h) > maxElements {
			heap.Pop(&h)
		}
		return h
	}
}

func sum(slice []int) (res int) {
	res = 0
	for i := range slice {
		res += slice[i]
	}
	return
}

func TopThreePacks(brigade Elves) []int {
	var topPacks []int

	tracker := nLargest(3) // keep thrack of the three largest elements in our queue
	for _, pack := range brigade {
		calories := sum(pack)
		topPacks = tracker(calories)
	}

	sort.Ints(topPacks)
	return topPacks
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1] // remove the final element which is an empty string

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

	// calculate the three most calorically dense packs
	heaviestPacks := TopThreePacks(elfBrigade)

	fmt.Printf("Part One Solution: %v\n", heaviestPacks[len(heaviestPacks)-1])
	fmt.Printf("Part Two Solution: %v\n", sum(heaviestPacks))
}
