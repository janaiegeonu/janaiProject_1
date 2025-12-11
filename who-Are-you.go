package main

import (
	"fmt"
)

func main() {

	var firstname string
	var lastname string

	fmt.Println("....WELCOME TO JANAI PROGRAM....")
	fmt.Println("NAME VERIFICATION")
	fmt.Println()
	fmt.Println("what is your fullname, i wanna know")
	_, err1 := fmt.Scanln(&firstname, &lastname)
	if err1 != nil {
		fmt.Println("ERROR.. type in your lastname too!!")
		return
	}
	name := firstname + " " + lastname
	for _, char := range name {
		if (char >= '0' && char <= '9') == true {
			fmt.Println("ERROR!!..please we don't accept numbers in names(input alphabetic formats only : ", "so", name, "is not allowed")
			return
		}
	}
	fmt.Println("YOU ARE WELCOME", firstname, lastname)

	var age int
	fmt.Println("AGE VERIFICATION")
	fmt.Println()
	fmt.Println("So", firstname, lastname, "i'll like to know how old you're")
	_, err2 := fmt.Scanln(&age)
	if err2 != nil {
		fmt.Println("ERROR.. AGE NEEDS TO BE A NUMBER!!! only", age, "is allowed in your input")
		return
	}
	fmt.Println("GREAT!.. AGE :", age, "has been saved to memory!!")

	var gender string
	fmt.Println("GENDER VERIFICATION")
	fmt.Println()
	fmt.Println("what is your gender??..")
	fmt.Println("MALE or FEMALE")
	fmt.Scanln(&gender)

	switch gender {
	case "MALE", "male":
		fmt.Println("okay", firstname, lastname, "your gender is MALE..")

	case "FEMALE", "female":
		fmt.Println("okay", firstname, lastname, "your gender is FEMALE")

	default:
		fmt.Println("ERROR!!..", gender, "is not matching a gender type")
		fmt.Println("provide a gender type (MALE OR FEMALE)")
	}

}
