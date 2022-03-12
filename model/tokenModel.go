package model

type Token struct {
	Id     int
	UserId int
	Token  string
	time   int
}

func Token2UserId(token string) (int, error) {
	var res Token
	db := newDB()
	err := db.Where("token = ?", token).Take(&res).Error
	return res.UserId, err
}

func GetTokenByUserId(userId int) string {
	var res Token
	db := newDB()
	db.Where("user_id = ?", userId).Take(&res)
	return res.Token
}

func SaveToken(userId int, token string, time int) int {
	res := Token{Token: token, UserId: userId, time: time}
	db := newDB()
	db.Create(&res)
	UpdateUserToken(userId, token)
	return res.Id
}

func (Token) TableName() string {
	return "user_token"
}
