package main

import (
	"errors"
)

func Add(num1, num2 int64) int64 {
	return num1 + num2
}
func Subtract(num1, num2 int64) int64 {
	return num1 - num2
}
func Multiply(num1, num2 int64) int64 {
	return num1 * num2
}
func Divide(num1, num2 int64) (int64, error) {
	quotient := num1 / num2
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return quotient, nil
}
func Modulus(num1, num2 int64) (int64, error) {
	modulus := num1 % num2
	if num2 == 0 {
		return 0, errors.New("Cannot mod by zero")
	}
	return modulus, nil
}
