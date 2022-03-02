package controller

import "github.com/UserServer/model"

type Oauth struct {
	Id        int
	Password  string
	Email     string
	IpAddress string
}

func NewOauth(userId int, password string, email string, ipAddr string) *Oauth {
	res := &Oauth{
		Id:        userId,
		Password:  password,
		Email:     email,
		IpAddress: ipAddr,
	}
	model.SaveOauth(res.Id, res.Password, res.Email, res.IpAddress)
	return res
}

func GetOauthByIp(ipAddr string) int {
	id := model.GetOauthByIp(ipAddr)
	return id
}
