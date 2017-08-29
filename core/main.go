package main

import (
	"fmt"
	. "../models"
	_ "../peer"
	// . "../controller"
	// . "../crypto"
	// "encoding/json"
	"math/rand"
	)


func main() {
	fmt.Println("programme started");
	fmt.Println("blocks start for 0 to n-1");

	// // //Initiate a new blockchain with genesis block
	// CreateAccount("akshay")
	// privKey,pubKey := GetAccount("akshay")
	// data := "this is the data to be hashed"
	// r,s := GetSignature(data,privKey)
	// var d = Data{data,pubKey,r,s}
	// bytes,_ := json.Marshal(d)
	// bc := NewBlockchain(string(bytes))
	// //Initiate a new blockchain with genesis block



	//add more blocks 
	bc := LoadBlockchain()
	// // create Account if necessary
	// CreateAccount("deepak")
	
	// for i:=0;i<100;i++ {
	// 	privKey,pubKey := GetAccount("saurabh")
	// 	value := randSeq(150)
	// 	r,s := GetSignature(value,privKey)
	// 	var d = Data{value,pubKey,r,s}
	// 	bytes,_ := json.Marshal(d)
	// 	nextBlock := bc.GenerateNextBlock(string(bytes))
	// 	bc.AddBlock(nextBlock);
	// }

	//add more blocks 

	bc.PrintChain();
	if(bc.IsValidChain()){
		fmt.Println("This is a valid Chain")
	}else{
		fmt.Println("This is an invalid Chain")
	}
	// StartPeer(&bc)

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}