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

func TestSizeCalculation(t *testing.T) {
	root := ParseInput(TestInput)
	totalFileSize(&root) // run calculation once - it should calculate for all subdirectories
	testCases := []struct {
		dir          *Directory
		expectedSize int
	}{
		{&root, 48381165},
		{root.Subdirectories["d"], 24933642},
		{root.Subdirectories["a"], 94853},
		{root.Subdirectories["a"].Subdirectories["e"], 584},
	}
	for _, testCase := range testCases {
		if exp, act := testCase.expectedSize, testCase.dir.size; exp != act {
			t.Errorf("Incorrect File Size for directory %v. Expected %v, got %v", testCase.dir.Name, exp, act)
		}
	}
}

func TestPartOne(t *testing.T) {
	root := ParseInput(TestInput)
	if exp, act := 95437, PuzzlePartOne(&root); exp != act {
		t.Errorf("Part One Solution Failed. Expected %v, got %v", exp, act)
	}

}
