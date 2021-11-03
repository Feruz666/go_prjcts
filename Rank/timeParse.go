package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// reads user input until \n by default
	scanner.Scan()
	// Holds the string that was scanned
	text := scanner.Text()
	layout1 := "03:04:05PM"
    layout2 := "15:04:05"
    t, _ := time.Parse(layout1, text)
    fmt.Println(t.Format(layout2))
}