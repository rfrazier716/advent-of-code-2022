package main

import (
	"testing"
)

func TestSomethingSilly(t *testing.T) {
	testTable := []struct {
		a        int
		b        int
		expected int
	}{{1, 2, 3}, {2, 3, 5}, {-4, 7, 3}}
	for i, test := range testTable {
		if res := test.a + test.b; res != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, test.expected, res)
		}
	}
}

func TestParser(t *testing.T) {
	testTable := []struct {
		input    []string
		expected []ElfPair
	}{
		{[]string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"27-36,8-1000000",
		}, []ElfPair{
			{CleaningRange{2, 4}, CleaningRange{6, 8}},
			{CleaningRange{2, 3}, CleaningRange{4, 5}},
			{CleaningRange{5, 7}, CleaningRange{7, 9}},
			{CleaningRange{27, 36}, CleaningRange{8, 1000000}},
		}},
	}
	for _, test := range testTable {
		res := ParseInput(test.input)
		if len(res) != len(test.expected) {
			t.Errorf("Test Failed, Expected %v, got %v", test.expected, res)
		} else {
			for i := range res {
				if res[i] != test.expected[i] {
					t.Errorf("Test failed at index %v. Expected %v, got %v", i, test.expected, res)
				}
			}
		}
	}
}

func TestRedundantPairs(t *testing.T) {
	testTable := []struct {
		Pair     ElfPair
		expected bool
	}{
		{ElfPair{CleaningRange{2, 4}, CleaningRange{0, 1}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{1, 3}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{3, 6}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{4, 6}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{1, 10}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{2, 4}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{2, 3}}, true},
	}
	for i, test := range testTable {
		if res := IsFullyRedundant(test.Pair); res != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, test.expected, res)
		}
	}
}

func TestPartiallyRedundantPairs(t *testing.T) {
	testTable := []struct {
		Pair     ElfPair
		expected bool
	}{
		{ElfPair{CleaningRange{2, 4}, CleaningRange{0, 1}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{5, 20}}, false},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{1, 3}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{3, 6}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{4, 6}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{1, 10}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{2, 4}}, true},
		{ElfPair{CleaningRange{2, 4}, CleaningRange{2, 3}}, true},
	}
	for i, test := range testTable {
		if res := IsPartiallyRedundant(test.Pair); res != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, test.expected, res)
		}
	}
}
