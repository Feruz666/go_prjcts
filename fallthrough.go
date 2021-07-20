package main

import "fmt"

func main() {
	nextstop :="B"

	fmt.Println("Stops ahead of us: ")

	switch nextstop{
	
	case "A":
		fmt.Println("A")
		fallthrough

	case "B":
		fmt.Println("B")
		fallthrough

	case "C":
		fmt.Println("C")
		fallthrough

	case "D":
		fmt.Println("D")
		fallthrough

	case "E":
		fmt.Println("E")
	}
}