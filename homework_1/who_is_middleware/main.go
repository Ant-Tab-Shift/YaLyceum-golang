package main

import (
	"fmt"
	"net/http"
    "log"
)

func httpServing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Middleware Test")
}

func loggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.String())
        next.ServeHTTP(w, r)
    })
}

func main() {
    server := http.Server{
		Addr:   ":8080",
		Handler: loggerMiddleware(http.HandlerFunc(httpServing)),
	}

	//fmt.Println("Starting server at port 8080. Make a request on http://localhost:8080/")
	server.ListenAndServe()
}

/*
Реализуйте HTTP сервер с middleware, который запускается на порту 8080, которое логирует все входящие запросы (метод и URL) в консоль. Основная функция сервера — возвращать "Middleware Test" на любой GET запрос.
*/
