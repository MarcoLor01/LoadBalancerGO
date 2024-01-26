package main

import (
	"encoding/json"
	"fmt"
	"io"
	"loadBalancerGO/wordCounter"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

type Count struct {
	lbMu sync.Mutex
	lb   *LoadBal
}

type Address struct {
	Addr string `json:"addr"`
}

type LoadBal struct {
	Address []Address `json:"address"`
	current int
}

func (lb *LoadBal) ForwardRPC(args wordCounter.Args, replyResultCounter *wordCounter.ResultCounter) (*wordCounter.ResultCounter, error) {
	actualServer := lb.Address[lb.current].Addr
	lb.current = (lb.current + 1) % len(lb.Address)
	client, err := rpc.Dial("tcp", actualServer)

	if err != nil {
		log.Fatal("Error in dialing to server: ", err)
	}
	defer client.Close()

	err = client.Call("Counter.CountLettersReal", args, replyResultCounter)
	if err != nil {
		return &wordCounter.ResultCounter{LenA: -1, LenB: -1}, err
	}
	return replyResultCounter, nil
}

func (wc *Count) CountLetters(args wordCounter.Args, reply *wordCounter.ResultCounter) error {
	wc.lbMu.Lock()
	defer wc.lbMu.Unlock()

	if wc.lb == nil {
		var err error
		wc.lb, err = CreateLoadBalancer()
		if err != nil {
			return err
		}
	}

	reply, err := wc.lb.ForwardRPC(args, reply)

	if err != nil {
		return err
	}
	return nil
}

func CreateLoadBalancer() (*LoadBal, error) {
	jsonFile, err := os.Open("serversAddr.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	loadBalancer := new(LoadBal)
	err = json.Unmarshal(byteValue, &loadBalancer)
	loadBalancer.current = 0

	if err != nil {
		log.Fatal("error unmarshalling the address")
	}

	return loadBalancer, nil
}

func main() {

	addr := "localhost:" + "8080"
	counter := new(Count)
	server := rpc.NewServer()
	err := server.RegisterName("WordCounter", counter)
	if err != nil {
		log.Fatal("error registering the loadBalancer")
	}
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Error while starting RPC load balancer:", err)
	}

	log.Printf("RPC load balancer listening on port %d", 8080)
	for {
		conn, _ := lis.Accept()
		go server.ServeConn(conn)
	}
}
