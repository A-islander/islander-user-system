package controller

import (
	"errors"
	"time"

	"github.com/user_server/model"
)

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	RegisterTime int    `json:"registerTime"`
}

type UserAuth struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser(name string) *User {
	res := User{
		Name:         name,
		RegisterTime: int(time.Now().Unix()),
	}
	res.Id = model.CreateUser(TransferUserModel(res))
	return &res
}

func GetUserById(id int) *User {
	res := TransferUser(model.GetUser(id))
	return &res
}

func GetUserByToken(token string) (*User, error) {
	if err := checkToken(token); err != nil {
		return nil, err
	}
	id, err := model.Token2UserId(token)
	if err != nil {
		return nil, err
	}
	return GetUserById(id), nil
}

func GetUserArr(userIdArr []int) ([]User, error) {
	return TransferUserArr(model.GetUserArr(userIdArr)), nil
}

func (user *User) update() bool {
	model.UpdateUser(TransferUserModel(*user))
	return true
}

func checkToken(token string) error {
	if len(token) != 32 {
		return errors.New("token is field")
	}
	return nil
}
