package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
	
	var distance = rand.Intn(401000000 - 56000000) + 56000000
	fmt.Println(distance)
}
