package main

import "fmt"

func doSomething() (int, bool) {
	return 5, false
}

func main() {
	// v, err := doSomething()
	// if err {
	// 	fmt.Println("Oh sh")
	// } else {
	// 	fmt.Println(v)
	// }

	if v, err := doSomething(); !err {
		fmt.Println(v)
	} else {
		fmt.Println("OH SHIT!")
	}
}
