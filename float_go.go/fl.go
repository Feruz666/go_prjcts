package main

import "fmt"

func main() {
	// third := 1.0 / 3

	// fmt.Println(third)
	// fmt.Printf("%v\n", third)
	// fmt.Printf("%f\n", third)
	// fmt.Printf("%.3f\n", third)
	// fmt.Printf("%07.5f\n", third)

	// var pi64 = math.Pi
	// var pi32 float32 = math.Pi

	// fmt.Println(pi64)
	// fmt.Println(pi32)

	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")

	cop := 0.0

	for cop < 11 {
		cop+=0.10
	}
	fmt.Println(cop)
}
