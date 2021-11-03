package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var amount int
	fmt.Scanf("%d", &amount)
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	// reads user input until \n by default
	scanner.Scan()
	// Holds the string that was scanned
	text := scanner.Text()
	nums := []int{}
	repl_nums := strings.Split(text, " ")
	for _, v := range repl_nums {
		s := string(v)
		n, _ := strconv.Atoi(s)

		nums = append(nums, n)
	}

	pos := 0
	neg := 0
	c := 0
	z := 0

	for _, v := range nums {
		c++
		if v < 0 {
			neg++
		} else if v == 0 {
			z++
		} else if v > 0 {
			pos++
		}

	}

	var outneg float32 = float32(neg) / float32(c)
	var outpos float32 = float32(pos) / float32(c)
	var outz float32 = float32(z) / float32(c)

	// fmt.Println(outpos, "\n", outneg, "\n", outz)
	fmt.Println(fmt.Sprintf("%.6f", outpos))
	fmt.Println(fmt.Sprintf("%.6f", outneg))
	fmt.Println(fmt.Sprintf("%.6f", outz))
}
