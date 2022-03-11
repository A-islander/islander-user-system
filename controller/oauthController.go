package controller

import "github.com/user_server/model"

type Oauth struct {
	Id        int    `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	IpAddress string `json:"ipAddress"`
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

func GetOauthByIp(ipAddr string) (int, error) {
	id, err := model.GetOauthByIp(ipAddr)
	return id, err
}
