package model

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	fmt.Println(GetTokenByUserId(1), Token2UserId("23308e8c5c2a175736cd55a822766c87"))
}

func TestOauth(t *testing.T) {
	// fmt.Println(SaveOauth(8, "", "", ""))
	fmt.Println(GetOauthByIp("127.0.0.1"), GetOauthByPassword("123"))
}

func TestUser(t *testing.T) {
	// fmt.Println(CreateUser(User{Name: "testing"}), UpdateUser(User{Id: 1, Name: "newName1"}))
	fmt.Println(GetUserArr([]int{1, 2, 3, 7, 8, 9}))
}
