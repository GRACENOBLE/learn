package main

import (
	"fmt"
)

func main() {
	scanned := 0
	array := []int{}
	fmt.Print("Enter 3 numbers\n")
	for i := 0; i < 3; i++ {
		fmt.Scan(&scanned)
		array = append(array, scanned)
	}
	
	for _, value := range array {
		fmt.Print(value)
	} 
}
