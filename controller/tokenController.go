package controller

import (
	"strconv"
	"time"

	"github.com/UserServer/model"
)

type Token struct {
	Id     int
	UserId int
	Token  string
}

func GetUserId(token string) int {
	return model.Token2UserId(token)
}

func NewToken(userId int) string {
	now := int(time.Now().Unix())
	token := NewMd5([]byte(strconv.Itoa(now)))
	model.SaveToken(userId, token, now)
	return token
}

func GetTokenByUserId(userId int) string {
	token := model.GetTokenByUserId(userId)
	return token
}
