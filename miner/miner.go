package main

import (
	"fmt"
	"log"
	. "../models"
	"golang.org/x/net/websocket"
	. "../controller"
	"encoding/json"
	. "../crypto"
	"math/rand"
)


var origin = "http://mychain0/"
var self = "ws://10.200.208.52:8080/"
var peer = "ws://10.200.208.52:8081/"
var peers []string
var miningState bool
var urlGetBlock = peer+"getBlock"
var urlSendBlock = peer+"sendBlock"
var ownIpSendBlock = self+"sendBlock"
var addPeerIp = peer+"addPeer"
var bc Blockchain
func main() {
	peers = append(peers,self)
	peers = append(peers,"ws://10.200.208.52:8081/")
	peers = append(peers,"ws://10.200.208.52:8082/")
	bc = LoadBlockchain();
	miningState = true;
	StartMining();
	// bc.PrintChainFromRoot()
}

func StartMining(){
	generateBlock(1)
	// sendRequest(urlGetBlock,Request{2,"block"})
	bc.SaveChainUsingRoot();
}

func StopMining(){
	bc.StopBlock()
}

func generateBlock(n int){
	for i:=0;i<n;i++ {
		privKey,pubKey := GetAccount("akshay")
		value := randSeq(40)
		r,s := GetSignature(value,privKey)
		var d = Data{value,pubKey,r,s}
		bytes,_ := json.Marshal(d)

		nextBlock := bc.GenerateNextBlock(string(bytes))
		miningState = false;
		// sendBlock(urlSendBlock,nextBlock)
				//uncomment this to add it to your own chain
		// sendBlock(ownIpSendBlock,nextBlock)

		for i:=0; i<len(peers); i++{
			sendBlock(peers[i]+"sendBlock",nextBlock)
		}
		//This will send the block to the connected peers
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func sendBlock(url string, block Block){
	
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
	err = websocket.JSON.Send(ws,block)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Index,block.Hash)
	// fmt.Println("Send:", block)

	var r Request

	err = websocket.JSON.Receive(ws, &r)
	
	if r.Data == "stopMining"{
		// fmt.Println("block received from other node")
		fmt.Println("is mining",miningState)
		if miningState{
		  fmt.Println("restart mining")
		  bc.StopBlock()
		  bc.SaveChainUsingRoot()
		  bc = LoadBlockchain()
		  main()
		}else{

		}
	}else if r.Data == "success"{
		fmt.Println("block added to chain")
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Receive: %s\n", data)
}

func sendRequest(url string, message Request){
	
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
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