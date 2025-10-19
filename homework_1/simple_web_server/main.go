package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	//fmt.Println("Starting server at port 8080. Make a request on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
Напишите простой HTTP сервер, который отвечает "Hello, World!" на любой GET запрос.

Пример обращения к серверу:

curl localhost:8080/
# Hello, World!
*/