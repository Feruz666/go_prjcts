package main

import "fmt"



func main() {
	// var planets [8]string

	// planets[0] = "Mercury"
	// planets[1] = "Venera"
	// planets[2] = "Earth"

	// earth := planets[2]
    // fmt.Println(earth)

	// fmt.Println(len(planets))
	// fmt.Println(planets[3]== "") 

	dwarfs := [5]string{"Cerera", "Xaumea", "Pluto", "Makemake", "Erid"}

	planets := [...]string{
		"Mercury",
		"Venera",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptun",
	}


	for i, dwarf:= range dwarfs{
		fmt.Println(i, dwarf)
	}

	planetsMarkII := planets

	planets[2] = "oops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
