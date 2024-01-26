package main

import (
	"loadBalancerGO/wordCounter"
	"log"
	"net"
	"net/rpc"
)

func main() {
	addr := "localhost:" + "2345"
	counter := new(wordCounter.Counter)
	server := rpc.NewServer()
	err := server.Register(counter)
	if err != nil {
		log.Fatal("Format of service WordCounter is not correct: ", err)
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error while starting RPC server:", err)
	}
	log.Printf("RPC server listens on port %d", 2345)
	server.Accept(lis)
}
