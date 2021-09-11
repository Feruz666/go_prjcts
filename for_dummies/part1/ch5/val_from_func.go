package main

import "fmt"

func addNum(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return sum
}

func countOddEven(s string) (int, int) {
	odds, evens := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

// Variable number of arguments (...), fixed param
func addNums2(total int ,nums ... int) int {
	total = 0
	for _, n := range nums {
		total += n
	}
	return total
}



func main() {
	s := addNum(6, 8)
	fmt.Println(s)

	o,e := countOddEven("qwertyui")
	fmt.Println(o, e)

	fmt.Println(addNums2(1,2,3,4,5,6))
}
