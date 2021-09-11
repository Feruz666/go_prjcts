package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var nums [5]int // an array of int (5 elements)
	// fmt.Println(nums)

	// nums2 := [5]int{1,2,3,4,5}
	// fmt.Println(nums2)

	nums := [...]int{1,2,3,4,5}
	fmt.Println(len(nums))

	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] = strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}
	fmt.Println(table)
}
