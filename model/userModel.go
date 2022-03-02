package model

type User struct {
	Id           int
	Name         string
	RegisterTime int
}

func CreateUser(data User) int {
	return data.Create()
}

func UpdateUser(data User) int {
	return data.Update()
}

func GetUserArr(userIdArr []int) []User {
	var res []User
	db := newDB()
	db.Find(&res, userIdArr)
	return res
}

func (user *User) Create() int {
	db := newDB()
	db.Create(user)
	return user.Id
}

func (user *User) Update() int {
	db := newDB()
	db.Model(user).Update("name", user.Name)
	return user.Id
}

func GetUser(id int) User {
	var res User
	db := newDB()
	db.Where("id = ?", id).Take(&res)

	return res
}

func UpdateUserToken(userId int, token string) {
	db := newDB()
	db.Model(&User{Id: userId}).Update("token", token)
}

func (User) TableName() string {
	return "user"
}
