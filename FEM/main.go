package main

import "frontendmasters.com/go/server/data"

func main() {
	max := data.Instructor {Id: 3, Lastname: "Firtman", } //new instructor
	max.FirstName = "Maximilliano"
	print(max.Print())
}
