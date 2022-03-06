package route

import (
	"errors"
	"net/http"

	"github.com/user_server/controller"
)

func getUserByToken(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Header["Authorization"]
	if !ok {
		writeError(w, 403, errors.New("without Authorization").Error())
	}
	user, err := controller.GetUserByToken(token[0])
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
