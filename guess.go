package main

import (
	"fmt"
	"math/rand"
)


func main() {
	var aim = 3
	var rnd = 0
	for {
		rnd = rand.Intn(10)
		fmt.Printf("My var %v\n", rnd)
		if rnd == aim {
			fmt.Printf("I GUESSED! My var %v!!!", rnd)
			break
		}
	}	
}