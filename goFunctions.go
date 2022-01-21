package main

import "fmt"

func main() {
	sum := add(1, 9)
	fmt.Println("sum =", sum)
	isNumEven := isEven(6)
	fmt.Println("is even", isNumEven)
	avg := findAvg([]int{1, 2, 3})
	fmt.Println("avg", avg)
}

func add(a int, b int) int {
	return a + b
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func findAvg(a []int) float32 {
	arrLen := len(a)
	fmt.Println("slice length", arrLen)
	sum := 0
	for i, e := range a {
		fmt.Println(i)
		sum = sum + e
	}
	return float32(sum) / float32(arrLen)
}
