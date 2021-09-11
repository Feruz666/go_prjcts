package main

import "fmt"

func main() {
	// var c [3]string
	// c[0] = "IOS"
	// c[1] = "Android"
	// c[2] = "Wind"
	// fmt.Println(c[0:2]) // IOS android
	// fmt.Println(c[:2]) //  IOS android
	// fmt.Println(c[1:]) //  android wind
	// fmt.Println(c[:]) //  ios android wind


	t := []int {1,2,3,4,5}
	// fmt.Println(len(t)) //5
	// fmt.Println(cap(t)) //5

	// t = t[2:4]
	// fmt.Println(t)
	// fmt.Println(len(t)) 
	// fmt.Println(cap(t)) 

	t = t[1:3] // 2 3
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))
}