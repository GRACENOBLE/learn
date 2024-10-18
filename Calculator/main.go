package main

import (
	"fmt"
)

func main() {
	var num1, num2 int64
	var operator string
	fmt.Println("Enter two numbers")
	fmt.Scanf("%d\n%d", &num1, &num2)

	fmt.Println("Enter an operator")
	fmt.Scanf("%s", &operator)

	switch operator {
	case "+":
		fmt.Printf("the sum is %d ", Add(num1, num2))
	case "-":
		fmt.Printf("the difference is %d ", Subtract(num1, num2))

	case "*":
		fmt.Printf("the product is %d ", Multiply(num1, num2))

	case "/":

		quotient, error := Divide(num1, num2)
		if error != nil {

			fmt.Printf("the quotient is %d ", quotient)
		} else {
			fmt.Println(error)
		}

	case "%":
		modulus, error := Divide(num1, num2)
		if error != nil {

			fmt.Printf("the modulus is %d ", modulus)
		} else {
			fmt.Println(error)
		}

	}
}
