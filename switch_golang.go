package main

import "fmt"

func main() {
	// fmt.Println("there is an enter to the cave and the way to the East")

	// var command = "go in"

	// switch command {
	// case "go to the East":
	// 	fmt.Println("You are going to the mountain")
	// case "go into the cave", "go in":
	// 	fmt.Println("you are in bad lighted cave ")
	// case "read the sign":
	// 	fmt.Println("On the sign outside there is a text 'Minors are not allowed to enter'.")
	// default:
	// 	fmt.Println("Its not clear so far")
	// }


	var room = "sea"

	switch {
	case room == "cave":
		fmt.Println("you are in bad lighted cave")
	case room == "sea":
		fmt.Println("The ice seems strong enough")
		fallthrough
	case room == "depth":
		fmt.Println("The water so cold")
	}
}