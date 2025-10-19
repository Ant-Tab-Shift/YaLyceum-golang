package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func generateId() string {
	symbols := make([]byte, 10)
	for i := range symbols {
		switch rand.Intn(3) {
		case 0:
			symbols[i] = byte(rand.Intn(26)) + 'a' // 26 букв от 'a' до 'z'
		case 1:
			symbols[i] = byte(rand.Intn(26)) + 'A' // 26 букв от 'A' до 'Z'
		case 2:
			symbols[i] = byte(rand.Intn(10)) + '0' // 10 цифр от '0' до '9'
		}
	}
	return string(symbols)
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "user_id",
		Value:    generateId(),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Please log in")
}

func httpServing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Middleware checke cookie status")
}

func redirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		if c, err := r.Cookie("user_id"); err != nil || c.Valid() != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", redirectMiddleware(http.HandlerFunc(httpServing)))
	mux.HandleFunc("/login", login)

	http.ListenAndServe(":8080", mux)
}

/*
Напишите HTTP сервер с middleware, которое проверяет наличие cookie user_id. Если cookie отсутствует, middleware должно перенаправлять запрос на /login, где сервер устанавливает cookie и отвечает "Please log in". При наличии cookie сервер возвращает "Access granted".

Пример:

curl localhost:8080/
# Redirects to /login

curl --cookie "user_id=123" localhost:8080/
# Access granted
*/
