package main

import (
	"fmt"
	"strings"
)

func main() {
	// var walkOutside = true
	// var takeTheBluePill = false

	fmt.Println("You are in the dark dungeon")
	var wearShades = true
	var command = "go out from dungeon"
	var exit = strings.Contains(command, "dungeon")
	var checkText = strings.Contains(command, "read")

	fmt.Println("You are going out from dungeon: ", exit, ". wear shades: ", wearShades)
	fmt.Println("Checking: ", checkText)
	fmt.Println("\n\n\n------\n")

	fmt.Println("On the sign outside there is a text 'Minors are not allowed to enter '")
	var age = 24
	var minor = age < 18
	fmt.Printf("in age of %v, am I adult? %v\n", age, minor)

	fmt.Println("яблоко" > "банан")

	
}
