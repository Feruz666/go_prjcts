package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0{
			break
		}
		count--
	}
	if count==0{
		fmt.Println("Start!")
	} else {
		fmt.Println("the launch is canseled")
	}

	
	var degrees = 0

	for {
		fmt.Println(degrees)

		degrees++
		if degrees>=360{
			degrees=0
			if rand.Intn(100) == 0 {
				break
			}
		}
	}

	
}