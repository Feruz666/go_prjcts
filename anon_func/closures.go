package main

import "fmt"


// var f = func() {
// 	fmt.Println("Dress for msqrd")

// }

func main() {
	func() {// Объявляет анонимную функцию
		fmt.Println("Func anon")
	}()// Вызов функции


	// f := func(message string) {
	// 	fmt.Println(message)
	// }
	// f("Hey")
}