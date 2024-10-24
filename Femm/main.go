package main

import (
	"fmt"
	"net/http"
)

func main(){
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error opening the server")
	}
}