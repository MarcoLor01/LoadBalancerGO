package main

import (
	"fmt"
	"loadBalancerGO/wordCounter"
	"log"
	"net/rpc"
	"os"
)

func main() {
	addr := "localhost:" + "8080"
	client, err := rpc.Dial("tcp", addr)
	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	if err != nil {
		log.Fatal("Error in dialing: ", err)
	}

	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal("Error in closing connection")
		}
	}(client)

	if len(os.Args) < 3 {
		fmt.Printf("No args passed in\n")
		os.Exit(1)
	}
	n1 := os.Args[1]
	n2 := os.Args[2]

	//Implementing a synchronous call
	args := &wordCounter.Args{A: n1, B: n2}
	var wordCount wordCounter.ResultCounter
	log.Printf("Synchronous call to RPC server")
	err = client.Call("WordCounter.CountLetters", args, &wordCount)
	if err != nil {
		log.Fatal("Error in CountLetters: ", err)
	}

	fmt.Printf("La prima parola ha lunghezza: %d\nLa seconda ha lunghezza: %d\n", wordCount.LenA, wordCount.LenB)
}
