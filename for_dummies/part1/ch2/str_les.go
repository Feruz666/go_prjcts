package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("Your age: ", age)

	var input string
	fmt.Print("Enter age: ")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Age is: ", age)
	}
}
