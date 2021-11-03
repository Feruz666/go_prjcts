package main

import "fmt"

func main() {
	// iterating
	// t := []int{1, 2, 3, 4, 5}
	// for i, v := range t {
	// 	fmt.Println(i, v)
	// }

	// making copies
	nums1 := [...]int{1, 2, 3, 4, 5}
	nums2 := nums1
	fmt.Println(nums1)
	fmt.Println(nums2)

	nums2[0] = 0
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println()

	// copying
	t := []int{1, 2, 3, 4, 5}
	v := make([]int, len(t))
	copy(v, t)

	fmt.Println(t)
	fmt.Println(v)
	fmt.Println()

	v2 := make([]int, 2, 5)
	copy(v2, t)
	fmt.Println(t)
	fmt.Println(len(v2))

	v3 := make([]int, 10)
	copy(v3, t)

	fmt.Println(t)
	fmt.Println(v3)

}
