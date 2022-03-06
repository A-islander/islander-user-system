package controller

import "github.com/user_server/model"

func TransferUserArr(data []model.User) []User {
	var res []User
	for i := 0; i < len(data); i++ {
		res = append(res, TransferUser(data[i]))
	}
	return res
}

func TransferUser(data model.User) User {
	return User{
		Id:           data.Id,
		Name:         data.Name,
		RegisterTime: data.RegisterTime,
	}
}

func TransferUserModel(data User) model.User {
	return model.User{
		Id:           data.Id,
		Name:         data.Name,
		RegisterTime: data.RegisterTime,
	}
}
