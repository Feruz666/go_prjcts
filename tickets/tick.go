package main

import (
	"math/rand"
)

var company = ""
var vel = 16

func main() {
	distance := 62100000

	for count := 0; count <= 10; count++ {
		switch num := rand.Intn(3) + 1; num {
		case 1:
			company = "Space Adventures"
		case 2:
			company = "SpaceX"
		case 3:
			company = "Virgin Galactic"
		}

		
	}

}
