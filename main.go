package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/UserServer/route"
	rpcserver "github.com/UserServer/rpcServer"
)

func init() {
	file := "./" + "message" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[UserServer]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	port := ":50001"
	userServer := route.Init()
	fmt.Printf("http listen to port %s \n", port)
	go http.ListenAndServe(port, userServer)
	go rpcserver.UserRpcInit()
	select {}
}
