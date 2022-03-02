package controller

import (
	"errors"

	"github.com/UserServer/model"
)

func LoginByPassword(password string) (string, error) {
	userId := model.GetOauthByPassword(password)
	if userId == -1 {
		return "", errors.New("cant search this password")
	}
	token := NewToken(userId)
	return token, nil
}

func RegisterByPassword(username string, email string, password string, ipAddr string) (string, error) {
	user := NewUser(username)
	NewOauth(user.Id, password, "", ipAddr)
	token := NewToken(user.Id)
	return token, nil
}

// 根据ip注册
func RegisterByIpAddress(ipAddr string) (string, error) {
	if userId := GetOauthByIp(ipAddr); userId != -1 {
		return GetTokenByUserId(userId), nil
	}
	user := NewUser(newRandName(7))
	NewOauth(user.Id, "", "", ipAddr)
	token := NewToken(user.Id)
	return token, nil
}
