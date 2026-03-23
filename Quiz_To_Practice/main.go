package main

import (
	"fmt"
	"go-quiz/week1_basics"
)

func main() {
	fmt.Println(week1_basics.ConvertAndProcess(5))
	fmt.Println(week1_basics.ConvertAndProcess("hello"))
	fmt.Println(week1_basics.ConvertAndProcess(5.5))
	fmt.Println(week1_basics.ConvertAndProcess([]int{1, 2, 3}))
	fmt.Println(week1_basics.ConvertAndProcess(true))
	fmt.Println(week1_basics.ConvertAndProcess(false))
	fmt.Println(week1_basics.TypeAnalyzer([]interface{}{1, "hello", 2, "world", 3.14, true, false}))

}
