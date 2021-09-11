package main

import "fmt"

func main() {
	grade := "C"
	switch grade {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "C":
		fallthrough
	case "D":
		fmt.Println("Passed")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Absent")
	}

	score := 79
	gr := ""
	switch {
	case score < 50:
		gr = "F"
	case score < 60:
		gr = "D"
	case score < 70:
		gr = "C"
	case score < 80:
		gr = "B"
	default:
		gr = "A"

	}
	fmt.Println(gr)
}
