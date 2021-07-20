package main

import "fmt"


func main() {
	text_2 := "L fdph, L vdz, L frqtxhuhg"

	for i:=0; i <len(text_2); i++{
		c:=text_2[i]
		if c>='a' && c<='z'{
			c=c-3
			if c>'z'{
				c=c+6
			}
		} else if c>='A' && c<='Z'{
			c=c-3
			if c>'Z'{
				c=c+6
			}
		}
		fmt.Printf("%c", c)
	} 
}