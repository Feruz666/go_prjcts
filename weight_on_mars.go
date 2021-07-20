package main

import "fmt"

func main() {
	fmt.Print("My weight on the Mars equal ")
	fmt.Print(70.0 * 0.3783)
	fmt.Print(" kg and my age equal ")
	fmt.Print(24 * 365 / 687)
	fmt.Print(" years\n")

	fmt.Printf("my weight on Mars = %v kg, ", 70 * 0.3783)
	fmt.Printf("and my age equals %v years. \n\n", 24*365/687 )

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
