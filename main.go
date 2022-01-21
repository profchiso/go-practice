package main

import "fmt"

func main() {
	fmt.Println("its a go")
	sum := Add(8, 8)
	fmt.Println("sum =", sum)
	result := printOddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	fmt.Println(result)
}

func Add(a int, b int) int {
	return a + b
}
func printOddNumbers(a []int) []int {
	var result []int
	for i, e := range a {
		if e%2 != 0 {

		} else {
			fmt.Println(i)
			result = append(result, e)
		}
	}
	return result
}
