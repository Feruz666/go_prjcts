package main

import "fmt"

func main() {
	var OS [3]string
	OS[0] = "IOS"
	OS[1] = "Android"
	OS[2] = "Windows"

	// i - indexes, v - values of indexes
	for _, v := range OS {
		fmt.Println(v)
	}
}
