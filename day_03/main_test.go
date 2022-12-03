package main

import (
	"testing"
)

func TestPrioritization(t *testing.T) {
	testTable := []struct {
		rucksack Rucksack
		expected Priority
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
	}
	for i, test := range testTable {
		if res := FindDuplicateInPockets(test.rucksack); res != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v with input %v", i, test.expected, res, test.rucksack)
		}
	}
}

func TestCommonItems(t *testing.T) {
	testTable := []struct {
		packs    []Rucksack
		expected rune
	}{
		{[]Rucksack{"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg"}, 'r'},
		{[]Rucksack{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw"}, 'Z'},
	}
	for i, test := range testTable {
		if act := FindCommonItem(test.packs); act != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, test.expected, act)
		}
	}
}
func TestPartOne(t *testing.T) {
	lines := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
		"",
	}
	expected := 157
	if act := PartOne(lines); expected != act {
		t.Errorf("Test Failed, expected %v, got %v", expected, act)
	}
}
