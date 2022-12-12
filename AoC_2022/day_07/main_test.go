package main

import (
	"testing"
)

var TestInput = []string{
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}


func TestParser(t *testing.T) {
	result := ParseInput(TestInput)
	expectedFiles := []SystemFile{
		{"b.txt", 14848514},
		{"c.dat", 8504156},
	}
	if len(result.Files) != len(expectedFiles) {
		t.Errorf("Files in root directory do not match, expected length %v, got %v", len(expectedFiles), len(result.Files))
	} else {
		for i := range expectedFiles {
			if expectedFiles[i] != result.Files[i] {
				t.Errorf("Files in root directory do not match at index %v. Expected %v, got %v", i, expectedFiles[i], result.Files[i])
			}
		}
	}
	// checking if the subdirectories work
	if subdir, ok := result.Subdirectories["a"].Subdirectories["e"]; ok {
		expectedFiles := []SystemFile{
			{"i", 584},
		}
		if len(subdir.Files) != len(expectedFiles) {
			t.Errorf("Files in root directory do not match, expected length %v, got %v", len(expectedFiles), len(result.Files))
		} else {
			for i := range expectedFiles {
				if expectedFiles[i] != subdir.Files[i] {
					t.Errorf("Files in root directory do not match at index %v. Expected %v, got %v", i, expectedFiles[i], subdir.Files[i])
				}
			}
		}
	} else {
		t.Errorf("Root Directory does not contain expected subdirectory")
	}
}

func TestPartOne(t *testing.T) {
	result := ParseInput(TestInput)
	sum := Solver(&result, 100000)
	expected := 95437
	if sum != expected {
		t.Errorf("Part one Failed, expected %v, got %v", expected, sum)
	}
}
