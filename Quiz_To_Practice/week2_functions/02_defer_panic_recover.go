package week2_functions

import (
	"fmt"
	"errors"
)

// Challenge 2: Defer, Panic, and Recover
// Master Go's unique error handling and cleanup mechanisms

// TODO: Uncomment and implement
// ResourceManager simulates opening multiple resources with proper cleanup
// Uses defer to ensure resources are closed in reverse order
// Returns a slice of strings showing the order of operations
/*
func ResourceManager(resources []string) []string {
	var operations []string

	// Your implementation here
	// 1. Open each resource (append "Open: resourceName" to operations)
	// 2. Use defer to close each resource (append "Close: resourceName")
	// 3. Process resources (append "Process: resourceName")
	// 4. Ensure close happens in reverse order of open

	return operations
}
*/

// TODO: Uncomment and implement
// SafeDivision performs division with panic recovery
// If b is 0, panic and recover, returning -1 and error
// Use defer and recover to handle the panic gracefully
/*
func SafeDivision(a, b int) (result int, err error) {
	// Your implementation here
	// Use defer to recover from panic
	// If panic occurs, set result to -1 and err to appropriate error
	return
}
*/

// TODO: Uncomment and implement
// TransactionSimulator simulates a database transaction
// Operations: "begin", "insert", "update", "delete", "commit", "rollback"
// If any operation fails (contains "fail" in operation name), panic and rollback
// Use defer to ensure cleanup happens
// Return final state and operations performed
/*
func TransactionSimulator(operations []string) (string, []string) {
	var performed []string
	var state string = "initial"

	// Your implementation here
	// Use defer to handle rollback on panic
	// Track all performed operations

	return state, performed
}
*/

// TODO: Uncomment and implement
// ChainedDefers demonstrates defer execution order
// Accept a number n and create n deferred functions
// Each deferred function appends its number to a result slice
// Return the final slice showing defer LIFO order
/*
func ChainedDefers(n int) []int {
	var result []int

	// Your implementation here
	// Create n deferred functions
	// Each should append its number to result

	return result
}
*/

// TODO: Uncomment and implement
// PanicPropagation demonstrates panic propagation and recovery at different levels
// Create nested function calls that might panic
// Recover at different levels based on recovery level parameter
// Levels: 1 (immediate), 2 (parent), 3 (grandparent)
/*
func PanicPropagation(shouldPanic bool, recoveryLevel int) (string, error) {
	// Your implementation here
	// Create nested functions
	// Panic if shouldPanic is true
	// Recover at the specified level

	return "", nil
}
*/