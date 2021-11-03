package main

import "fmt"




func main() {
	//defer 
	defer handlePanic()

	// fmt.Println("Main")
	// fmt.Println("Main 2")
	// fmt.Println("Main 3")



	// panic()
	//
	messages := []string{
		"message1",
		"message2",
		"message3",
		"message4",
	}
	fmt.Println(messages)

	messages[4] = "message5"
	// panic("AAAAAAA HELP") 
}

func handlePanic() {
	if r:=recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlePanic() HAS DONE")
}