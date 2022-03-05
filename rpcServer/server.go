package rpcserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/UserServer/controller"
)

type UserQuery struct {
	UserId    int
	UserIdArr []int
	Token     string
}

type UserRpcServer struct{}

func (server *UserRpcServer) GetUserById(query *UserQuery, reply *controller.User) error {
	*reply = *controller.GetUserById(query.UserId)
	return nil
}

func (server *UserRpcServer) GetUserArr(query *UserQuery, reply *[]controller.User) error {
	res, err := controller.GetUserArr(query.UserIdArr)
	*reply = res
	return err
}

func (server *UserRpcServer) GetUserByToken(query *UserQuery, reply *controller.User) error {
	res, err := controller.GetUserByToken(query.Token)
	*reply = *res
	return err
}

func UserRpcInit() {
	server := new(UserRpcServer)
	// 注册rpc服务
	rpc.Register(server)
	rpc.HandleHTTP()
	port := ":40001"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rpc listen to ", port)
	http.Serve(l, nil)
}
