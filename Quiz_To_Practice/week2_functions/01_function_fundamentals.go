package week2_functions

import "errors"

// Challenge 1: Function Fundamentals
// Master functions, multiple returns, and variadic functions

// TODO: Uncomment and implement
// Calculator with multiple returns
// Performs operation and returns (result, remainder/decimal part, error)
// Operations: "divide" (returns quotient and remainder)
//            "power" (returns result and number of multiplications performed)
//            "factorial" (returns factorial and number of multiplications)
//            "sqrt" (returns integer part and decimal part * 100)
/*
func AdvancedCalculator(a, b int, operation string) (int, int, error) {
	// Your implementation here
	return 0, 0, nil
}
*/

// TODO: Uncomment and implement
// Variadic function that accepts any number of integers and a operation string at the end
// Operations: "sum", "product", "max", "min", "average", "median"
// For median, sort the numbers and return middle value (or average of two middle values)
/*
func VariadicProcessor(operation string, numbers ...int) (float64, error) {
	// Your implementation here
	return 0, nil
}
*/

// TODO: Uncomment and implement
// Function that returns a function (closure)
// Create a counter generator that returns a function
// The returned function increments and returns the counter
// Optional: accept initial value and step
/*
func CounterGenerator(initial, step int) func() int {
	// Your implementation here
	return nil
}
*/

// TODO: Uncomment and implement
// Higher-order function that accepts a slice and a filter function
// Returns a new slice containing only elements that pass the filter
/*
func FilterSlice(slice []int, filter func(int) bool) []int {
	// Your implementation here
	return nil
}
*/

// TODO: Uncomment and implement
// Recursive function to solve Tower of Hanoi
// Returns the sequence of moves as a slice of strings
// Format: "Move disk from A to C"
/*
func TowerOfHanoi(n int, source, auxiliary, destination string) []string {
	// Your implementation here
	return nil
}
*/