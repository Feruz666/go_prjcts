package main

import "fmt"

func kelvinToCelcius(k float64) float64 {
	k -= 273.15
	return k
}

func celciusToFarenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

func kelvinToFahrenheit(k float64) float64 {
	k = (k-273.15)*9/5 + 32
	return k
}

func main() {
	cel := 25.0
	kelvin := 233.0
	kel_0 := 0.0

	celcius := kelvinToCelcius(kelvin)
	farenheit := celciusToFarenheit(cel)
	kel_to_far := kelvinToFahrenheit(kel_0)

	fmt.Print(kelvin, " K is ", celcius, " C\n")
	fmt.Print(cel, " C is ", farenheit, " F\n")
	fmt.Print(kel_0, " K is ", kel_to_far, " F\n")
}
