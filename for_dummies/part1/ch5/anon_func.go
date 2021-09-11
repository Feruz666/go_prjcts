package main

import "fmt"

func doSmth() int {
	return 5
}

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, (f1 + f2)
		return f1
	}
}

func main() {
	var i func() int
	i = func() int {
		return 5
	}
	fmt.Println(i())

	gen := fib()
	fmt.Println(gen())
	for i := 0; i < 8; i++ {
		fmt.Println(gen())
	}
}
