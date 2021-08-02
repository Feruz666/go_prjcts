package main

import (
	"fmt"
)

type fahrenheit float64
type celsius float64


var cels celsius = 40.0


func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit(c * 9.0/5.0) + 32
}

func fahrenheitToCelsius(f fahrenheit) celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func drawTable() {
	fmt.Println("====================")
	fmt.Println("|   C     |   F    |")
	fmt.Println("====================")

	for i := int(cels); i<= 100; i++ {
		if i%5 ==0  {
			fmt.Printf("|   %vC    |   F   |\n", i)
		}
	}
}


func main() {
	drawTable()
}