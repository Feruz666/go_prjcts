package main

import "fmt"

func main() {
	// s := make([]int, 5)

	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))


	// r := make([]int, 2, 5)
	// fmt.Println(r)
	// fmt.Println(len(r))
	// fmt.Println(cap(r))

	t := []int{1,2,3,4,5}
	t = append(t, 6,7,8, 9, 10)
	// fmt.Println(t)
	// fmt.Println(len(t))
	// fmt.Println(cap(t))

	u := t
	fmt.Println(u)
	fmt.Println(t)
	fmt.Println()

	u[9] = 44
	fmt.Println(u)
	fmt.Println(t)
	fmt.Println()

	t = append(t, 15)
	fmt.Println(u)
	fmt.Println(t)

	

}