package week2_functions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdvancedCalculator(t *testing.T) {
	t.Skip("Remove this line after implementing AdvancedCalculator")

	tests := []struct {
		name      string
		a, b      int
		operation string
		result    int
		extra     int
		hasError  bool
	}{
		{"Divide 10 by 3", 10, 3, "divide", 3, 1, false},
		{"Power 2^5", 2, 5, "power", 32, 5, false},
		{"Factorial of 5", 5, 0, "factorial", 120, 4, false},
		{"Square root of 50", 50, 0, "sqrt", 7, 7, false}, // 7.07...
		{"Invalid operation", 10, 5, "invalid", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result, extra, err := AdvancedCalculator(tt.a, tt.b, tt.operation)
			// if tt.hasError {
			// 	assert.Error(t, err)
			// } else {
			// 	assert.NoError(t, err)
			// 	assert.Equal(t, tt.result, result)
			// 	assert.Equal(t, tt.extra, extra)
			// }
		})
	}
}

func TestVariadicProcessor(t *testing.T) {
	t.Skip("Remove this line after implementing VariadicProcessor")

	tests := []struct {
		name      string
		operation string
		numbers   []int
		expected  float64
		hasError  bool
	}{
		{"Sum", "sum", []int{1, 2, 3, 4, 5}, 15, false},
		{"Product", "product", []int{2, 3, 4}, 24, false},
		{"Max", "max", []int{3, 7, 2, 9, 1}, 9, false},
		{"Min", "min", []int{3, 7, 2, 9, 1}, 1, false},
		{"Average", "average", []int{10, 20, 30}, 20, false},
		{"Median odd", "median", []int{3, 1, 4, 1, 5}, 3, false},
		{"Median even", "median", []int{1, 2, 3, 4}, 2.5, false},
		{"Empty slice", "sum", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result, err := VariadicProcessor(tt.operation, tt.numbers...)
			// if tt.hasError {
			// 	assert.Error(t, err)
			// } else {
			// 	assert.NoError(t, err)
			// 	assert.Equal(t, tt.expected, result)
			// }
		})
	}
}

func TestCounterGenerator(t *testing.T) {
	t.Skip("Remove this line after implementing CounterGenerator")

	// counter1 := CounterGenerator(0, 1)
	// counter2 := CounterGenerator(10, 5)

	// Test counter1
	// assert.Equal(t, 1, counter1())
	// assert.Equal(t, 2, counter1())
	// assert.Equal(t, 3, counter1())

	// Test counter2
	// assert.Equal(t, 15, counter2())
	// assert.Equal(t, 20, counter2())
	// assert.Equal(t, 25, counter2())
}

func TestFilterSlice(t *testing.T) {
	t.Skip("Remove this line after implementing FilterSlice")

	tests := []struct {
		name     string
		slice    []int
		filter   func(int) bool
		expected []int
	}{
		{
			"Even numbers",
			[]int{1, 2, 3, 4, 5, 6},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"Greater than 5",
			[]int{3, 6, 9, 2, 7, 1},
			func(n int) bool { return n > 5 },
			[]int{6, 9, 7},
		},
		{
			"None match",
			[]int{1, 3, 5},
			func(n int) bool { return n%2 == 0 },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := FilterSlice(tt.slice, tt.filter)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTowerOfHanoi(t *testing.T) {
	t.Skip("Remove this line after implementing TowerOfHanoi")

	tests := []struct {
		name     string
		n        int
		expected int // number of moves
	}{
		{"1 disk", 1, 1},
		{"2 disks", 2, 3},
		{"3 disks", 3, 7},
		{"4 disks", 4, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// moves := TowerOfHanoi(tt.n, "A", "B", "C")
			// assert.Len(t, moves, tt.expected)
			// Formula: 2^n - 1 moves for n disks
		})
	}
}