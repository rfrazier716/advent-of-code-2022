package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TerminalCommand int

const (
	CD TerminalCommand = 0
	LS TerminalCommand = 1
)

type SystemFile struct {
	Name string
	Size int
}

type Directory struct {
	Name           string
	Parent         *Directory
	Subdirectories map[string]*Directory
	Files          []SystemFile
	size           int
}

func NewDirectory(name string, parent *Directory) Directory {
	return Directory{
		name,
		parent,
		map[string]*Directory{"..": parent},
		make([]SystemFile, 0),
		-1, // initially unsized
	}
}

func ExecuteTerminalCommand(commandString string, pwd *Directory) (new *Directory, err error) {
	switch commandString[:2] {
	case "cd":
		// change directory
		if new, ok := pwd.Subdirectories[commandString[3:]]; ok {
			return new, nil
		} else {
			return new, fmt.Errorf("subdirectory %v does not exist in %v", commandString[3:], pwd.Name)
		}
	case "ls":
		new = pwd
		return
	}
	err = fmt.Errorf("could not process command %v", commandString[:2])
	return
}

func ParseInput(input []string) Directory {
	anchor := NewDirectory("/", nil)
	pwd := &anchor

	for _, val := range input {
		switch val[0] {
		case '$':
			new, err := ExecuteTerminalCommand(val[2:], pwd)
			if err != nil {
				log.Println(err)
			} else {
				pwd = new
			}
		case 'd':
			dirName := val[4:]
			dir := NewDirectory(dirName, pwd)
			pwd.Subdirectories[dirName] = &dir
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			splitFileInfo := strings.Split(val, " ")
			fileSize, _ := strconv.Atoi(splitFileInfo[0])
			fileName := splitFileInfo[1]
			pwd.Files = append(pwd.Files, SystemFile{fileName, fileSize})
		}
	}

	return anchor
}

func totalFileSize(d *Directory) int {
	// Returns the total size of a directory - including the size of it's files and the files in all subdirectories
	// This is effectively memoized, since running it once will recursively calculate the size of all subdirectories
	var dfs func(dir *Directory) int
	dfs = func(dir *Directory) int {
		// if size has been calculated, do it
		if dir.size != -1 {
			return dir.size
		}

		// Otherwise recursively calculate
		dirSize := 0

		// sum all files
		for i := range dir.Files {
			dirSize += dir.Files[i].Size
		}

		// sum the directories
		for key, val := range dir.Subdirectories {
			if val != nil && key != ".." {
				dirSize += dfs(val)
			}
		}
		dir.size = dirSize // update dirSize
		return dirSize
	}

	return dfs(d)
}

func PuzzlePartOne(root *Directory) int {
	// This will be an iterative Breadth First Search
	// Pull a file off the queue, check if it's sized appropriately 
	deque := make([]*Directory, 0) // make a deque for our iterative BFS
	deque = append(deque, root)    // append our root directory

	maxFileSize := 100000
	sum := 0
	// Loop as long as we have files to explore
	for len(deque) > 0 {

		// empty the deque
		for range deque {
			// pop our element
			dir := deque[0]
			deque = deque[1:]

			// if the size is less than our max - update the sum
			if size := totalFileSize(dir); size <= maxFileSize {
				sum += size
			}

			// put any subdirectories onto the deque
			for key, sub := range dir.Subdirectories {
				if key != ".." {
					deque = append(deque, sub)
				}
			}
		}
	}

	return sum
}

func PuzzlePartTwo(root *Directory) int {
	// This will be an iterative Breadth First Search
	// Pull a file off the queue, check if it's sized appropriately 
	deque := make([]*Directory, 0) // make a deque for our iterative BFS
	deque = append(deque, root)    // append our root directory

	spaceToClear := totalFileSize(root) - 40000000
	smallestClearableFileSize := totalFileSize(root)
	// Loop as long as we have files to explore
	for len(deque) > 0 {

		// empty the deque
		for range deque {
			// pop our element
			dir := deque[0]
			deque = deque[1:]

			// if the size is less than our max - update the sum
			if size := totalFileSize(dir); size >= spaceToClear && size < smallestClearableFileSize {
				smallestClearableFileSize = size
			}

			// put any subdirectories onto the deque
			for key, sub := range dir.Subdirectories {
				if key != ".." {
					deque = append(deque, sub)
				}
			}
		}
	}

	return smallestClearableFileSize
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}
	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	root := ParseInput(lines)

	fmt.Printf("Part One Solution: %v\n", PuzzlePartOne(&root))
	fmt.Printf("Part Two Solution: %v\n", PuzzlePartTwo(&root))
}
