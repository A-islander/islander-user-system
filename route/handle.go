package route

import (
	"errors"
	"log"
	"net/http"

	"github.com/user_server/controller"
)

func getUserByToken(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Header["Authorization"]
	if !ok {
		writeError(w, 403, errors.New("without Authorization").Error())
		return
	}
	user, err := controller.GetUserByToken(token[0])
	if err != nil {
		writeError(w, 403, err.Error())
	} else {
		write(w, user)
	}
}

func registerUserByIp(w http.ResponseWriter, r *http.Request) {
	ipAddr := remoteIp(r)
	log.Println(ipAddr)
	token, _ := controller.RegisterByIpAddress(ipAddr)
	write(w, struct {
		Token string `json:"token"`
	}{Token: token})
}
