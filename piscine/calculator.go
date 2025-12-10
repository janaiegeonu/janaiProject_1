package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Println("WELCOME TO JANAI CALCULATOR PROGRAM 🧮")
	fmt.Println("INPUT ONLY NUMBERS 🔢")

	// First number
	fmt.Print("Enter the first number : ")
	_, err1 := fmt.Scanln(&number1)
	if err1 != nil {
		fmt.Println("⚠️ERROR!!! NOT A NUMBER (please input numbers only)⚠️")
		return
	}

	// Second number
	fmt.Print("Enter the second number : ")
	_, err2 := fmt.Scanln(&number2)
	if err2 != nil {
		fmt.Println("⚠️ERROR!!! NOT A NUMBER (please input numbers only)⚠️")
		return
	}

	// Operator
	fmt.Print("PICK YOUR OPERATOR ( + - * / ) : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0 {
			fmt.Println("Division by zero.")
		} else {
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("SYNTAX ERROR (Not an Operator)")
	}
}
