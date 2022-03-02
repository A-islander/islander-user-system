package controller

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	fmt.Println(GetUserByToken("12322213"))
}

func TestLogin(t *testing.T) {
	// LoginByPassword("123")
	// RegisterByPassword("hedyakn", "123@email.com", "kkkk")
}

func TestCommon(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(newRandName(7))
	}
}
