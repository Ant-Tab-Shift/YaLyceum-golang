package main

import (
	"fmt"
	"net/http"
)

func ipBlockerMiddleware(blockedIP string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.Header.Get("X-Real-IP") // strings.Split(r.RemoteAddr, ":")[0]
		if remoteIP == blockedIP {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func httpServing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Access granted")
}

func main() {
	http.Handle("/", ipBlockerMiddleware("192.168.0.1", http.HandlerFunc(httpServing)))
	http.ListenAndServe(":8080", nil)
}

/*
Напишите веб-сервер, который блокирует доступ для определённого IP-адреса. Используйте middleware для проверки IP. Если IP-адрес клиента совпадает с заблокированным адресом (например, 192.168.0.1), сервер должен возвращать статус 403 Forbidden. Если IP-адрес не заблокирован, сервер возвращает ответ "Access granted".

Адрес заблокированного клиента должен передаваться в middleware по такому типу ipBlockerMiddleware(blockedIP string, next http.Handler)

Примечания
Используйте middleware для фильтрации IP.
Сервер должен блокировать запросы от заблокированных IP.
Сервер должен обрабатывать запросы от разрешённых IP, возвращая "Access granted".

curl localhost:8080/ --header "X-Real-IP: 192.168.0.1"
# 403 Forbidden

curl localhost:8080/ --header "X-Real-IP: 127.0.0.1"
# Access granted
*/
