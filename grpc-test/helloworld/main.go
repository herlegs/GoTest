package main

import (
	"github.com/herlegs/GoTest/grpc-test/helloworld/server"
	"fmt"
)

func main(){
	go server.StartRPC()
	server.StartHTTP()
}