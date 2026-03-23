package parallel

import (
	"fmt"
	"time"
)

func printNumberees() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(time.Now())
		fmt.Println(string(letter))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	//Parallel
	go printNumberees()
	go printLetters()

	time.Sleep(4 * time.Second)
}
