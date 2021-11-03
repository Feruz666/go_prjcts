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
func (p *point) length() float64 {
	return math.Sqrt(
		(math.Pow(float64(p.x), 2) +
			math.Pow(float64(p.y), 2) +
			math.Pow(float64(p.z), 2)))
}

func (p *point) move(deltax,deltay,deltaz float32) {
	p.x += deltax
	p.y += deltay
	p.z += deltaz
}

func (p point) output() {
	fmt.Println("This is" , p)
}


func newPoint(x, y, z float32) *point {
	p := point{x: x, y: y, z: z}
	return &p
}


func main() {
	pt1 := newPoint(7.8, 9.1, 2.3)
	fmt.Println(pt1.length())
	pt1.output()

	pt1.move(0.1, 0.1, 0.1)
	fmt.Println(*pt1)
}