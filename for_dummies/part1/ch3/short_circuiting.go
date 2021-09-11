package main

import "fmt"

func raining() bool {
	fmt.Println("Check if it is raining now...")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now...")
	return true
}

func main() {
	// if snowing() || raining() {
	// 	fmt.Println("Stay indoors!")
	// }

	// if !raining() || snowing() {
	// 	fmt.Println("Let's SKII!!")
	// }
	
	if !raining() && !snowing() {
		fmt.Println("Let's go outdoors!")
	}
}
