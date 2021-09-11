package main

import "fmt"


func main() {
	// for i := 1; i <= 5; i++ {
	// 	for j := 0; j <= 5; j++ {
	// 		fmt.Printf("%d * %d = %d\n", i,j, i*j)
	// 	}
	// 	fmt.Println("------------")
	// }

	Outerloop:

		for i := 1; i <= 5; i++ {
			for j := 0; j <= 5; j++ {
				if i ==3 {
					continue Outerloop
				}
				fmt.Printf("%d * %d = %d\n", i,j, i*j)
			}
			fmt.Println("------------")
		}
}