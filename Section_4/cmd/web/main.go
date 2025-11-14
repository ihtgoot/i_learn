package main

import (
	"fmt"
	"net/http"
	"section_4/pkg/handler"
)

const portNumber string = ":8080"

func hi_there(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, fmt.Sprint("<html><h1>Hi There</h1></html>"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Print("server started")
	http.HandleFunc("/", hi_there)
	http.HandleFunc("/home", handler.Home)
}
