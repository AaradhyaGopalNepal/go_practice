package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool
type personSorter struct {
	person []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.person)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.person[i], &s.person[j])
}

func (s *personSorter) Swap(i, j int) {
	s.person[i], s.person[j] = s.person[j], s.person[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		person: people,
		by:     by,
	}
	sort.Sort(ps)
}
func main() {
	person := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	byAge := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(byAge).Sort(person)
	fmt.Println("Sorted by age:", person)

}
