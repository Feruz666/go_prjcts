package main

import "fmt"

func main() {
	// b, err := strconv.ParseBool("t")
	// fmt.Println(b)
	// fmt.Println(err)
	// fmt.Printf("%T\n", b)

	// f, err := strconv.ParseFloat("3.1415", 64)
	// fmt.Println(f)
	// fmt.Println(err)
	// fmt.Printf("%T\n", f)

	num1 := 5
	num2 := float32(num1)
	num3 := float64(num2)
	num4 := float32(num3)
	num5 := int(num4)
	
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num2)
	fmt.Printf("%T\n", num3)
	fmt.Printf("%T\n", num4)
	fmt.Printf("%T\n", num5)

}
