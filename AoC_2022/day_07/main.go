package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
}

func NewDirectory(name string, parent *Directory) Directory {
	return Directory{
		name,
		parent,
		map[string]*Directory{"..": parent},
		make([]SystemFile, 0),
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

func Solver(rootDir *Directory, maxFileSize int, spaceToClear int) (sum int, deletedDirectorySize int) {
	deletedDirectorySize = math.MaxInt // highest integer size

	// helper function to traverse the directory tree4
	var dfs func(dir *Directory) int
	dfs = func(dir *Directory) int {
		dirSize := 0
		for i := range dir.Files {
			// sum the files
			dirSize += dir.Files[i].Size
		}
		for key, val := range dir.Subdirectories {
			// sum the directories
			if val != nil && key != ".." {
				dirSize += dfs(val)
			}
		}
		if dirSize <= maxFileSize {
			// add to sum if not oversized
			sum += dirSize
		}
		if dirSize < deletedDirectorySize && dirSize >= spaceToClear{
			// tracking for part2
			deletedDirectorySize = dirSize
		}
		return dirSize
	}

	dfs(rootDir)

	return 
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}
	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	root := ParseInput(lines)
	sizeCount, dirSizes := Solver(&root, 100000)
	sort.Ints(dirSizes)
	fmt.Println(sizeCount)
	fmt.Println(dirSizes)
}
