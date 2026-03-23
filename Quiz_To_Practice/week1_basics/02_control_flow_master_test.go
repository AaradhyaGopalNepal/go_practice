package week1_basics

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestFizzBuzzPlus(t *testing.T) {
	t.Skip("Remove this line after implementing FizzBuzzPlus")

	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			"First 20 numbers",
			20,
			[]string{
				"1²", "Prime", "FizzPrime", "2²", "BuzzPrime",
				"Fizz", "BangPrime", "8", "Fizz²", "Buzz",
				"Prime", "Fizz", "Prime", "Bang", "FizzBuzz",
				"4²", "Prime", "Fizz", "Prime", "Buzz",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := FizzBuzzPlus(tt.n)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPatternMatcher(t *testing.T) {
	t.Skip("Remove this line after implementing PatternMatcher")

	tests := []struct {
		input    string
		expected string
	}{
		{"racecar", "PALINDROME"},
		{"A man a plan a canal Panama", "PALINDROME"},
		{"abcabc", "REPEATED"},
		{"xyzxyz", "REPEATED"},
		{"abcd", "ASCENDING"},
		{"dcba", "DESCENDING"},
		{"(())", "BALANCED"},
		{"({[]})", "BALANCED"},
		{"(()", "NONE"},
		{"random", "NONE"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// result := PatternMatcher(tt.input)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSwitchCalculator(t *testing.T) {
	t.Skip("Remove this line after implementing SwitchCalculator")

	tests := []struct {
		name      string
		a, b      float64
		operation string
		expected  float64
		hasError  bool
	}{
		{"Addition", 10, 5, "+", 15, false},
		{"Subtraction", 10, 5, "-", 5, false},
		{"Multiplication", 10, 5, "*", 50, false},
		{"Division", 10, 5, "/", 2, false},
		{"Power", 2, 3, "^", 8, false},
		{"Modulo", 10, 3, "%", 1, false},
		{"Square root", 16, 0, "sqrt", 4, false},
		{"Invalid operation", 10, 5, "invalid", 0, true},
		{"Division by zero", 10, 0, "/", math.Inf(1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result, err := SwitchCalculator(tt.a, tt.b, tt.operation)
			// if tt.hasError {
			// 	assert.Error(t, err)
			// } else {
			// 	assert.NoError(t, err)
			// 	if math.IsInf(tt.expected, 1) {
			// 		assert.True(t, math.IsInf(result, 1))
			// 	} else {
			// 		assert.Equal(t, tt.expected, result)
			// 	}
			// }
		})
	}
}