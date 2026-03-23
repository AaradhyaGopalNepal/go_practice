package week1_basics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBitwiseOperations(t *testing.T) {
	t.Skip("Remove this line after implementing BitwiseOperations")

	tests := []struct {
		name      string
		number    int
		n         int
		operation string
		expected  int
	}{
		{"Set 3rd bit of 5", 5, 3, "set", 13},       // 0101 -> 1101
		{"Clear 2nd bit of 7", 7, 2, "clear", 3},    // 0111 -> 0011
		{"Toggle 1st bit of 10", 10, 1, "toggle", 8}, // 1010 -> 1000
		{"Check 2nd bit of 6", 6, 2, "check", 1},    // 0110, bit 2 is set
		{"Count bits in 15", 15, 0, "count", 4},     // 1111 has 4 bits
		{"Rightmost bit of 12", 12, 0, "rightmost", 3}, // 1100, rightmost at position 3
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := BitwiseOperations(tt.number, tt.n, tt.operation)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLogicalPuzzle(t *testing.T) {
	t.Skip("Remove this line after implementing LogicalPuzzle")

	tests := []struct {
		name       string
		a, b, c    bool
		expression string
		expected   bool
	}{
		{"AND all true", true, true, true, "AND", true},
		{"AND with false", true, false, true, "AND", false},
		{"OR with one true", false, true, false, "OR", true},
		{"XOR odd trues", true, false, false, "XOR", true},
		{"XOR even trues", true, true, false, "XOR", false},
		{"NAND all true", true, true, true, "NAND", false},
		{"NOR all false", false, false, false, "NOR", true},
		{"XNOR even trues", true, true, false, "XNOR", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := LogicalPuzzle(tt.a, tt.b, tt.c, tt.expression)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestComparisonChain(t *testing.T) {
	t.Skip("Remove this line after implementing ComparisonChain")

	tests := []struct {
		name     string
		numbers  []int
		pattern  string
		expected bool
	}{
		{"Ascending", []int{1, 2, 3, 4, 5}, "ascending", true},
		{"Not ascending", []int{1, 3, 2, 4}, "ascending", false},
		{"Descending", []int{5, 4, 3, 2, 1}, "descending", true},
		{"Equal", []int{3, 3, 3, 3}, "equal", true},
		{"Not equal", []int{3, 3, 4, 3}, "equal", false},
		{"Alternating", []int{1, 3, 2, 5, 3, 7}, "alternating", true},
		{"Non-decreasing", []int{1, 2, 2, 3, 3, 4}, "non-decreasing", true},
		{"Non-increasing", []int{5, 5, 4, 3, 3, 1}, "non-increasing", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := ComparisonChain(tt.numbers, tt.pattern)
			// assert.Equal(t, tt.expected, result)
		})
	}
}