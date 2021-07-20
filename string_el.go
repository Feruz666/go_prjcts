package main

import "fmt"

func main() {
	w:= "a b c\td\nefg hi"

	for _, e := range w{
		switch e {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c\n", e)
		}
	}
}
