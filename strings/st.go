package main

import "fmt"

func main() {
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(`strings can be \n`)

	fmt.Printf("%v is a %[1]T\n", "literal string")

	// var star byte = '*'

	// c := message[5]
	// fmt.Printf("%c\n", c)
	message := "shalom"
	for i := 0; i < 6; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}

	fmt.Println(len(message))
	fmt.Println("\n\n\n")
	message2 := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message2); i++ {
		c:=message2[i]
		if c>='a' && c<='z'{
			c=c+13
			if c>'z'{
				c=c-26
			}
		}
		fmt.Printf("%c", c)
	}
}
