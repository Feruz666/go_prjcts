package main

import (
	"fmt"
	"math"
)

type point struct {
	x float32
	y float32
	z float32
}

// defining method of struct
func (p point) length() float64 {
	return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.x), 2)))
}

func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}

func main() {

	// Struct basics
	var pt1 point
	pt1.x = 3.1
	pt1.y = 5.7
	pt1.z = 4.2
	fmt.Println(pt1.x)
	fmt.Println(pt1.y)
	fmt.Println(pt1.z)

	pt2 := point{
		x: 5.6,
		y: 3.3,
		z: 1.1,
	}
	pt3 := point{x: 2.2, y: 4.4}
	fmt.Println(pt2, pt3)
	fmt.Println()

	// Create a new point struct using a func
	pt4 := newPoint(7.8, 6.7, 2.3)
	fmt.Println(pt4)
	fmt.Println()

	//Making a copy
	// pt5 := pt4

	// pt5.x = 0
	// fmt.Println(pt4)
	// fmt.Print(pt5 , "\n\n")

	// Independent copy of pt4 sign of asterisk *
	pt6 := *pt4
	pt6.z = 0
	fmt.Println(pt4)
	fmt.Print(pt6, "\n\n")

	pt7 := pt6
	pt7.x = 1
	fmt.Println(pt7)
	fmt.Print("\n\n")

	pt8 := &pt7
	pt8.x = 8
	fmt.Println(pt7)
	fmt.Println(pt8)

}
