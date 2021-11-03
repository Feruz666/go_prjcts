package main

import "fmt"

func main() {
	message := "Soon I will become a Ninja"

	printMessage(message)

	fmt.Println(message)
}

func printMessage(message string) {
	message += " (message from the printMessage())"
	fmt.Println(message)
}