package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	go printMessage("Go is great!")
	printMessage("I will be a master!")

}
