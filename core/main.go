package main

import (
	"fmt"
	. "../models"
	_ "../peer"
	. "../controller"
	. "../crypto"
	"encoding/json"
	"math/rand"
	)

var bc Blockchain
func main() {
	fmt.Println("programme started");
	fmt.Println("blocks start for 0 to n-1");


//    //Init Block
//      generateInitBlock()
//    //Init Block
// 	bc.PrintChainUsingRoot();
	// //add more blocks 
	generateBlock(0);
	// //add more blocks 

	bc.PrintChainUsingRoot();

	fmt.Println(bc.GetNthBlockFromRoot(20))
	if(bc.IsValidChainFromEnd()){
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

func generateBlock(n int){
	bc = LoadBlockchain()
	for i:=0;i<n;i++ {
		privKey,pubKey := GetAccount("ayush")
		value := randSeq(40)
		r,s := GetSignature(value,privKey)
		var d = Data{value,pubKey,r,s}
		bytes,_ := json.Marshal(d)
		nextBlock := bc.GenerateNextBlock(string(bytes))
		bc.AddBlock(nextBlock);
	}
}
func generateInitBlock(){
	// //Initiate a new blockchain with genesis block
	// Create account if necessary
	// CreateAccount("akshay")
	privKey,pubKey := GetAccount("akshay")
	data := "this is genesis block data";
	r,s := GetSignature(data,privKey)
	var d = Data{data,pubKey,r,s}
	bytes,_ := json.Marshal(d)
	bc = NewBlockchain(string(bytes))
	//Initiate a new blockchain with genesis block
}