package main

import (
	"fmt"
	"net/http"
)

func languageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("lang")
	if err != nil {
		cookie = &http.Cookie{Name: "lang", Value: "en"}
		http.SetCookie(w, cookie)
	}

	switch cookie.Value {
	case "ru":
		fmt.Fprintf(w, "Привет!\n")
	default:
		fmt.Fprintf(w, "Hello!\n")
	}
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(languageHandler))
}

/*
Эта задача направлена на работу с cookies и их использованием для персонализации интерфейса на основе языка, выбранного пользователем. Сервер должен проверять наличие cookie с именем lang, которое определяет язык интерфейса. Если cookie отсутствует, сервер должен установить язык по умолчанию (например, английский). Если cookie установлено и содержит значение, соответствующее поддерживаемому языку (например, ru для русского), сервер должен отвечать на этом языке.

Примечания
Если пользователь не имеет cookie lang, сервер отвечает "Hello!" и устанавливает язык по умолчанию — английский.
Если у пользователя установлено cookie lang=ru, сервер отвечает "Привет!".
Если cookie имеет другое значение, сервер также возвращает "Hello!" (английский по умолчанию).

languageHandler(w http.ResponseWriter, r *http.Request)
Примеры:


curl localhost:8080/
# Hello!


curl localhost:8080/ --cookie "lang=ru"
# Привет!


curl localhost:8080/ --cookie "lang=fr"
# Hello!
*/
