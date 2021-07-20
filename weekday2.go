package main

import (
	"fmt"
)

func main(){
	// switch time.Now().Weekday() {
	// case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
	// 	fmt.Println("Work days")
	
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekends") 
	// }

	size := "XXXL"

	switch size {
	case "XXS":
		fmt.Println("Very small")
	case "S":
		fmt.Println("Small")
	case "M":
		fmt.Println("Medium")
	case "L":
		fmt.Println("Large")

	default:
		fmt.Println("Undef")
	}


	switch num:=6; num%2==0{
	case true:
		fmt.Println("even")
	case false:
		fmt.Println("ODD")
	}
}