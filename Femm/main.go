package main

import (
	"fmt"
	"net/http"

	"frontendmasters.com/go/museum/api"
)
func handleHello (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go program"))
	}
func main() {
	server := http.NewServeMux() 
	server.HandleFunc("/hello", handleHello )
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8888", server)
	if err != nil {
		fmt.Printf("Error opening the server : %v", err)
	}
}
