package main

import (
	"fmt"
)

func main(){
	var command = "go to the East"

	if command == "go to the East" {
		fmt.Println("You are going to the mountain")
	} else if command == "Go in" {
		fmt.Println("You are in the dungeon")
	} else {
		fmt.Println("It is not clear so far")
	}
}