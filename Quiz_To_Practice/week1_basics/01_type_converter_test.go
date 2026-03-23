package week1_basics

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvertAndProcess(t *testing.T) {
	t.Skip("Remove this line after implementing ConvertAndProcess")

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"Integer conversion", 42, "Number: 42"},
		{"String conversion", "hello", "HELLO"},
		{"Float conversion", 3.7, "Number: 4"},
		{"Float conversion down", 3.2, "Number: 3"},
		{"Bool true", true, "YES"},
		{"Bool false", false, "NO"},
		{"Slice sum", []int{1, 2, 3, 4, 5}, "Sum: 15"},
		{"Empty slice", []int{}, "Sum: 0"},
		{"Nil input", nil, "Unknown type"},
		{"Unknown type", struct{}{}, "Unknown type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := ConvertAndProcess(tt.input)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTypeAnalyzer(t *testing.T) {
	t.Skip("Remove this line after implementing TypeAnalyzer")

	tests := []struct {
		name     string
		input    []interface{}
		expected map[string]int
	}{
		{
			"Mixed types",
			[]interface{}{1, "hello", 2, "world", 3.14, true, false},
			map[string]int{"int": 2, "string": 2, "float64": 1, "bool": 2},
		},
		{
			"Only integers",
			[]interface{}{1, 2, 3, 4, 5},
			map[string]int{"int": 5},
		},
		{
			"Empty slice",
			[]interface{}{},
			map[string]int{},
		},
		{
			"Complex types",
			[]interface{}{[]int{1, 2}, map[string]int{"a": 1}, struct{}{}},
			map[string]int{"[]int": 1, "map[string]int": 1, "struct {}": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := TypeAnalyzer(tt.input)
			// assert.Equal(t, tt.expected, result)
		})
	}
}