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
