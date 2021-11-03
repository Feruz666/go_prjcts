package main

import (
	"fmt"
)


func main() {
	// users := map[string]int {
	// 	"Vasya": 12,
	// 	"Petr" : 22,
	// 	"Jack" : 15,
	// }

	// users["Reyn"] = 55

	var users map[string]int
	users = make(map[string]int)

	users["Petr"] = 22

	// delete(users, "Petr")

	for key, value := range users{
		fmt.Println(key, value)
	}



	// age, exists := users["Petr"]
	// fmt.Println(age, exists)

	// if exists {
	// 	fmt.Println("Jack", age)
	// } else {
	// 	fmt.Println("нет в списке")
	// }

}