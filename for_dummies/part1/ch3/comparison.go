package main

import "fmt"

func main() {
	num := 5
	fmt.Println(num == 0) // || && !
	fmt.Println(num != 0)

	num2 := 6
	condition := num2 > 2 && num2 < 9 // ||
	fmt.Println(condition)

	//making decision if else

	condition2 := num%2 == 1
	if condition2 {
		fmt.Println("num is odd")
	} else {
		fmt.Println("even num")
	}

	if condition == true {
		fmt.Println("Number is odd")
	}

	if num%2 == 1 {
		fmt.Println("Odd")
	}
}
