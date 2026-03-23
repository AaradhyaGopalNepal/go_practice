package week2_functions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestResourceManager(t *testing.T) {
	t.Skip("Remove this line after implementing ResourceManager")

	resources := []string{"database", "file", "network"}
	expected := []string{
		"Open: database",
		"Open: file",
		"Open: network",
		"Process: database",
		"Process: file",
		"Process: network",
		"Close: network",  // Note: reverse order
		"Close: file",
		"Close: database",
	}

	// result := ResourceManager(resources)
	// assert.Equal(t, expected, result)
}

func TestSafeDivision(t *testing.T) {
	t.Skip("Remove this line after implementing SafeDivision")

	tests := []struct {
		name     string
		a, b     int
		expected int
		hasError bool
	}{
		{"Normal division", 10, 2, 5, false},
		{"Division by zero", 10, 0, -1, true},
		{"Negative division", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result, err := SafeDivision(tt.a, tt.b)
			// if tt.hasError {
			// 	assert.Error(t, err)
			// 	assert.Equal(t, -1, result)
			// } else {
			// 	assert.NoError(t, err)
			// 	assert.Equal(t, tt.expected, result)
			// }
		})
	}
}

func TestTransactionSimulator(t *testing.T) {
	t.Skip("Remove this line after implementing TransactionSimulator")

	tests := []struct {
		name           string
		operations     []string
		expectedState  string
		expectedOps    []string
	}{
		{
			"Successful transaction",
			[]string{"begin", "insert", "update", "commit"},
			"committed",
			[]string{"begin", "insert", "update", "commit"},
		},
		{
			"Failed transaction",
			[]string{"begin", "insert", "fail-update", "commit"},
			"rolled-back",
			[]string{"begin", "insert", "rollback"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// state, ops := TransactionSimulator(tt.operations)
			// assert.Equal(t, tt.expectedState, state)
			// assert.Equal(t, tt.expectedOps, ops)
		})
	}
}

func TestChainedDefers(t *testing.T) {
	t.Skip("Remove this line after implementing ChainedDefers")

	tests := []struct {
		n        int
		expected []int
	}{
		{3, []int{3, 2, 1}},        // LIFO order
		{5, []int{5, 4, 3, 2, 1}},
		{1, []int{1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			// result := ChainedDefers(tt.n)
			// assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPanicPropagation(t *testing.T) {
	t.Skip("Remove this line after implementing PanicPropagation")

	tests := []struct {
		name          string
		shouldPanic   bool
		recoveryLevel int
		expectedMsg   string
		hasError      bool
	}{
		{"No panic", false, 1, "completed successfully", false},
		{"Panic with immediate recovery", true, 1, "recovered at level 1", false},
		{"Panic with parent recovery", true, 2, "recovered at level 2", false},
		{"Panic with grandparent recovery", true, 3, "recovered at level 3", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// msg, err := PanicPropagation(tt.shouldPanic, tt.recoveryLevel)
			// if tt.hasError {
			// 	assert.Error(t, err)
			// } else {
			// 	assert.NoError(t, err)
			// 	assert.Contains(t, msg, tt.expectedMsg)
			// }
		})
	}
}