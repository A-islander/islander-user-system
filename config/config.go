package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	UserName string
	PassWord string
	Ip       string
	Database string
	SageNum  int
}

func GetConfig() Config {
	file, err := ioutil.ReadFile("./config.json")
	// test中
	// file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal(err)
	}
	var res Config
	json.Unmarshal(file, &res)
	return res
}
