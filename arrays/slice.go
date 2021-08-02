package main

import "fmt"

func main() {

	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	terrestial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestial, gasGiants, iceGiants)
}
