package main

import "fmt"

func main() {
	fmt.Println("now 2100 year. Is it a leap year?")

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("This is a leap year!")
	} else {
		fmt.Println("Unfortunately, it is not. This year is not leap")
	}

	var haveTorch = true
	var litTorch = false

	if !haveTorch || !litTorch {
		fmt.Println("Cant see anything")
	}
}
