package route

import (
	"net/http"

	"github.com/UserServer/controller"
)

func getUserByToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header["Authorization"][0]
	user, err := controller.GetUserByToken(token)
	if err != nil {
		writeError(w, 401, err.Error())
	} else {
		write(w, user)
	}
}

func registerUserByIp(w http.ResponseWriter, r *http.Request) {
	ipAddr := remoteIp(r)
	token, _ := controller.RegisterByIpAddress(ipAddr)
	write(w, struct{ Token string }{Token: token})
}
