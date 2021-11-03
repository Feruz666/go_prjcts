package main

import "fmt"

// type DumbDB struct {
// 	m map[string]string
// }

// func NewDumbDB() *DumbDB {
// 	return &DumbDB{
// 		m: make(map[string]string),
// 	}
// }

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}




type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) SetName(name string) {
	u.name = name
	
}

func (u User) getName() string {
	return u.name 
}

// Конструктор Инициализация, создание нового экземпляра структуры
func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age),
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vasya", "Male", 15, 90, 185)
	user2 := User{"Petya", 25, "Male", 80, 175}

	// user1.printUserInfo("Bok")
	// user2.printUserInfo("Deth")
	

	user1.SetName("bok")
	fmt.Println(user1.getName())
	fmt.Println(user2.getName())

	fmt.Println(user1.age.isAdult())

	// // Detail output
	// fmt.Printf("%+v\n", user1)
	// fmt.Printf("%+v\n", user2)
}
