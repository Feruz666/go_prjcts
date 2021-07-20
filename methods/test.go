package main

import "fmt"



type celsius float64

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) farenheit() farenheit{
	return farenheit((c * 9.0 / 5.0) + 32.0)
}



type farenheit float64

func (f farenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0) 
}

func (f farenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}



type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) farenheit() farenheit {
	return k.celsius().farenheit()
}




func main() {
	var kel kelvin = 280.0
	c := kel.celsius()
	fmt.Print(c)
}
