package main

import (
	"fmt"
	"time"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

func main() {
	displayDate("Mon 2006-01-02 15:04:05", "Current DateTime")
	x := 5
	y := 6
	swap(&x, &y)
	fmt.Println(x, y)
}
