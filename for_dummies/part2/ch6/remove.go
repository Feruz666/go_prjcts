package main

import (
	"errors"
	"fmt"
)

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

func main() {
	t := []int{1, 2, 3, 4, 5}
	t, err := delete(t, 2)

	if err==nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
}
