package simple

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 3, 4, 1, 2}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)
	sortingSlice := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	sort.Strings(sortingSlice)
	fmt.Println("Sorted strings", sortingSlice)
}
