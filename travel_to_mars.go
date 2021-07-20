package main

import "fmt"

func main() {
	const lightspeed = 299792// km/s
	var distance = 56000000

	fmt.Println(distance/lightspeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

	const spacex_engine = 100800
	distance = 96300000
	fmt.Println(distance/spacex_engine/24 , "Days to mars")


	// var (
	// 	distance = 56000000
	// 	speed = 100800
	// )

	// var distance, speed = 56000000, 100800

	const hours_per_day, minutes_per_hour = 24, 60

	var weight = 150.0
	weight = weight * 0.3783
	weight *= 0.3783

	var age = 41
	age = age+1
	age +=1
	age++

	weight-=2
}