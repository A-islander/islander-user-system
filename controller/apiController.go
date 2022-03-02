package controller

func Login(email string, password string) (string, error) {
	token, err := LoginByPassword(password)
	if err != nil {
		return "", err
	}
	return token, err
}

func GetUserInfo(token string) (User, error) {
	user, err := GetUserByToken(token)
	if err != nil {
		return User{}, err
	}
	return *user, err
}

func GetUserInfoArr(idArr []int) ([]User, error) {
	return GetUserArr(idArr)
}
