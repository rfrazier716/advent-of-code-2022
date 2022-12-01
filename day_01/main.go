package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func partOne(calories []int) {
	sum, max := 0, 0
	for i := range calories {
		if calories[i] == -1 {
			// update max and reset
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			sum += calories[i]
		}
	}
	fmt.Println(max)
}

func partTwo(calories []int) {
	h := IntHeap{}
	heap.Init(&h)
	sum := 0
	for i := range calories {
		if calories[i] == -1 {
			if sum > h[0] {
				heap.Push(&h, sum)
			}
			for len(h) > 3 {
				heap.Pop(&h)
			}
			sum = 0
		} else {
			sum += calories[i]
		}
	}

	totalSum := 0
	for i := range h {
		totalSum += h[i]
	}

	fmt.Println(totalSum)
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}
	lines := strings.Split(string(puzzleInput), "\n")
	calories := make([]int, 0)

	for i := range lines {
		if num, err := strconv.Atoi(lines[i]); err == nil {
			calories = append(calories, num)
		} else if len(lines) == 0 {
			// empty line
			calories = append(calories, -1)
		} else {
			log.Printf("Could not parse line %v: %v", i, lines[i])
		}
	}

	partOne(calories)
	partTwo(calories)

}
