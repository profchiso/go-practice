package main

import (
	"fmt"
)

func main() {
	nums := [3]int{1, 2, 3}
	for i, n := range nums {
		fmt.Println(i, n)
	}
	mod := 7 % 3

	fmt.Println("modolus", mod)
	// a := 7.0
	// b := 3.0
	// c := a / b
	// fmt.Println(c)
	sum := add(2, 6)
	fmt.Println(sum)
}

func add(a int, b int) int {

	return a + b

}
