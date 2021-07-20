package main 

type kelvin float64
type celsius float64
type farenheit float64


// Функция

// func kelvinToCelsius(k kelvin) celsius{
// 	return celsius(k - 273.15)
// }

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}



//  Метод|ПРИЕМНИК(КАК параметр)|название метода|возвращаемый тип!
	func (f farenheit)  celsius() celsius {
		return celsius((f - 32.0) * 5.0 / 9.0)
	}



var k kelvin = 294.0
var c celsius
var f farenheit

c = k.celsius()
c = f.celsius()