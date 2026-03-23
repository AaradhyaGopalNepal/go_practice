package week1_basics

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// Challenge 1: Type Converter
// Create a function that takes various types and performs intelligent conversions
// This tests understanding of Go's type system and type assertions

// ConvertAndProcess should:
// 1. Accept any type (interface{})
// 2. Detect the type and convert it appropriately:
//    - int to string with "Number: " prefix
//    - string to uppercase
//    - float64 to int (rounded) then to string
//    - []int to sum of all elements as string
//    - bool to "YES" (true) or "NO" (false) - Done
//    - For unknown types, return "Unknown type" - Done

func ConvertAndProcess(input interface{}) string {
	switch v := input.(type) {
	case int:
		return fmt.Sprint("Number: ", v)
	case string:
		return strings.ToUpper(v)
	case float64:
		return fmt.Sprint(int(math.Round(v)))
	case []int:
		sum := 0
		for _, e := range v {
			sum = sum + e
		}
		return fmt.Sprint("Sum: ", sum)
	case bool:
		if input == true {
			return "YES"
		} else {
			return "NO"
		}
	default:
		return "Unknown type"
	}
}

// TypeAnalyzer should analyze a slice of interface{} and return a map
// with type names as keys and count of occurrences as values
// Example: []interface{}{1, "hello", 2, "world"} -> map[string]int{"int": 2, "string": 2}

func TypeAnalyzer(items []interface{}) map[string]int {
	value := map[string]int{}
	for _, e := range items {
		typeOf := reflect.TypeOf(e).String()
		oldValue, ok := value[typeOf]
		if ok {
			value[typeOf] = oldValue + 1
		} else {
			value[typeOf] = 1
		}
	}
	return value
}
