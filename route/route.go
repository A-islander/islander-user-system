package route

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func UserHttpInit() {
	port := ":50001"
	userServer := http.NewServeMux()
	mid := totalMiddleware
	userServer.Handle("/user/get", mid(getUserByToken))
	userServer.Handle("/user/register", mid(registerUserByIp))
	fmt.Printf("http listen to port %s \n", port)
	http.ListenAndServe(port, userServer)
}
