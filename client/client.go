package main

import (
	"fmt"
	"log"
	. "../models"
	"golang.org/x/net/websocket"
)


var origin = "http://localhost/"
var url = "ws://localhost:8080/addPeer"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
	// message := []byte("4321")
	message := Request{400,"Block"}
	err = websocket.JSON.Send(ws,message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send:", message)

	var data Block
	err = websocket.JSON.Receive(ws, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", data)
}