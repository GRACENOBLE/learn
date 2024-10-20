package data

import "fmt"

type Instructor struct{
	Id 			int
	FirstName 	string
	Lastname 	string
	Score 		int
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.Lastname, i.FirstName, i.Score)
}