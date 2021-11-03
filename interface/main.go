package main

import (
	"fmt"
	"math"
)



func main() {
	square := Square{5}
	circle := Circle{3}

	printShapeArea(square)
	printShapeArea(circle)


	
	// // Вызов пустого интерфейса
	// printInterface(square)
	// printInterface(circle)
	// printInterface("askjchksjd")
	// printInterface(222)
	// printInterface(true)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

// Пустой интерфейс
func printInterface(i interface{}) {
	// Приведение интерфейса к конкретному значению
	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case string:
	// 	fmt.Println("string", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// default:
	// 	fmt.Println("Unknown type", value)
	// }

	str, ok := i.(string)
	if !ok {
		fmt.Println("Interface is not string")
		return
	}
	fmt.Println(len(str))

	// fmt.Printf("%+v\n", i)
}
