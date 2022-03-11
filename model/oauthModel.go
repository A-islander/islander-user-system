package model

type Oauth struct {
	Id        int
	Password  string
	Email     string
	IpAddress string
}

func GetOauthByPassword(password string) int {
	var res Oauth
	db := newDB()
	db.Where("password = ?", password).Take(&res)
	return res.Id
}

func GetOauthByIp(ipAddr string) (int, error) {
	var res Oauth
	db := newDB()
	err := db.Where("ip_address = ?", ipAddr).Take(&res).Error
	return res.Id, err
}

func SaveOauth(userId int, password string, email string, ipAddr string) int {
	res := Oauth{Id: userId, Password: password, Email: email, IpAddress: ipAddr}
	db := newDB()
	db.Create(&res)
	return res.Id
}

func (Oauth) TableName() string {
	return "user_oauth"
}
