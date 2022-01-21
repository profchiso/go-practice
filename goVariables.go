package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("basic go variables are")
	name := "chinedu"
	age := 20
	height := 2.5
	hobbies := []string{"sport", "programming"}
	favNumbers := []int{90, 10, 94}

	fmt.Println("the type of name is", reflect.TypeOf(name))
	fmt.Println("the type of name age", reflect.TypeOf(age))
	fmt.Println("the type of name height", reflect.TypeOf(height))
	fmt.Println("the type of name hobbies", reflect.TypeOf(hobbies))
	fmt.Println("the type of name favNumbers", reflect.TypeOf(favNumbers))

}
