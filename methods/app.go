package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var kel kelvin = 294.0
	c := kelvinToCelsius(kel)

	var cel celsius = 127

	k := celsiusToKelvin(cel)

	fmt.Print(kel, " K is ", c, " C ")
	fmt.Print(cel, " C is ", k, " K ")
}
