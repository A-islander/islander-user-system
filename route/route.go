package route

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func Init() *http.ServeMux {
	userServer := http.NewServeMux()
	mid := totalMiddleware
	userServer.Handle("/user/get", mid(getUserByToken))
	userServer.Handle("/user/register", mid(registerUserByIp))

	return userServer
}
