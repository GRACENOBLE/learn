package main

import (


	reader "example.com/Files/filemanagement"
)

func main(){
	//c := reader.ReadTextFile("data/text.txt")
	reader.WriteToFile("data/text2.txt","Hello")

}