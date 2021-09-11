package main

import (
	"fmt"
	"strings"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// max := 100
	// a,b := 0, 1
	// for ;b<= max; {
	// 	fmt.Println(b)
	// 	a,b = b, a+b
	// }

	// Endless LOOP
	// for{}

	
	for {
		fmt.Println("Enter QUIT to exit")
		var input string
		fmt.Print("Please enter a string: ")
		fmt.Scanln(&input)
		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}
}