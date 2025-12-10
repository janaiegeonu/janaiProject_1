package main

import "fmt"

func main() {

	var name string

	fmt.Println("....WELCOME TO JANAI PROGRAM....")
	fmt.Println("what is your name, i wanna know")
	fmt.Scanln(&name)
	fmt.Println("YOU ARE WELCOME", name)

	var age int
	fmt.Println("So", name, "i'll like to know how old you're")
	fmt.Scanln(&age)
	fmt.Println("GREAT!.. AGE :", age, "is good to paticipate!!!")
}
