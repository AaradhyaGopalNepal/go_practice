package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &person{}
		},
	}

	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18
	fmt.Println("Got Person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person again:", person2)
	person3 := pool.Get().(*person)
	fmt.Println("Got another person", person3)
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Returned persons to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got person4", person4)
	person5 := pool.Get().(*person)
	fmt.Println("Got person5", person5)
}
