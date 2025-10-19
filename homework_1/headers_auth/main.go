package main

import (
	"fmt"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func httpServing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Authorized access")
}

func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(httpServing)))
	http.ListenAndServe(":8080", nil)
}

/*
Напишите веб-сервер, который проверяет наличие заголовка Authorization для каждого запроса. Если заголовок отсутствует или содержит неверное значение, сервер возвращает 401 Unauthorized. Если заголовок правильный (например, "Authorization: Bearer valid_token"), сервер возвращает "Authorized access".

Примечания
Реализуйте проверку заголовка Authorization.
Сервер должен возвращать 401 Unauthorized, если заголовок неверен.
Сервер должен возвращать "Authorized access" при корректном заголовке.

Cигнатура middleware: authMiddleware(next http.Handler) http.Handler

Пример:

curl localhost:8080/ --header "Authorization: Bearer valid_token"
# Authorized access

curl localhost:8080/
# 401 Unauthorized
*/