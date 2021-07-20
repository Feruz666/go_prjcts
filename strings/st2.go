package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")

	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size:= utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size)

	// _ нижнее подчеркивание в цикле как переменная которая
	//  не будет использоваться внутри самого цикла
	for _, c:= range question{
		fmt.Printf("%c ", c)
	}
	fmt.Println("\n")
	word:="abcdefghijklmnopqrstuvwxyz"
	fmt.Println(len(word), "bytes")
	fmt.Println(utf8.RuneCountInString(word), "runes")
	test:="¿"
	c2, size2:= utf8.DecodeRuneInString(test)
	fmt.Printf("First rune: %c %v bytes\n", c2, size2)
}