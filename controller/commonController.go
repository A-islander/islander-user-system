package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func NewMd5(str []byte) string {
	m := md5.New()
	m.Write(str)
	return hex.EncodeToString(m.Sum(nil))
}

func newRandName(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
