package main

import (
	"fmt"
	"strconv"
)

func main(){
	age := 24
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am ", marsAge, " years old on Mars")

	countdown, err := strconv.Atoi("10")
	fmt.Println(countdown)
	fmt.Println(err)
}