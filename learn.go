/*Reading from and printing out to the console.*/
package main

import (
	"fmt"
)

func main() {
	scanned := 0// a temp variable for the scanned value
	array := []int{} //an array to store the scanned values

	fmt.Print("Enter 3 numbers\n")

	for i := 0; i < 3; i++ {//scanning and appending
		fmt.Scan(&scanned)
		array = append(array, scanned)
	}
	
	for _, value := range array {//printing output
		fmt.Printf("%d\n",value)
	} 
}
