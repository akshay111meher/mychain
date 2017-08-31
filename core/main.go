package main

import (
	"fmt"
	. "../models"
	peer "../peer"
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
	generateAdditionalBlock(0);
	// //add more blocks 

	// // bc.PrintChainFromRoot();
	// fmt.Println("##############")
	// fmt.Println(bc.GetNthBlockFromRoot(2).Hash)
	// fmt.Println(bc.GetNthBlockFromRoot(3).Hash)
	// fmt.Println("##############")


	// 	privKey,pubKey := GetAccount("ayush")
	// 	value := randSeq(40)
	// 	r,s := GetSignature(value,privKey)
	// 	var d = Data{value,pubKey,r,s}
	// 	bytes,_ := json.Marshal(d)
		
	// firstForkBlock := bc.GetForkBlock(string(bytes),bc.GetNthBlockFromRoot(1).Hash,2)
	// firstForkHash := firstForkBlock.SHA256();
	// bc.AddBlock(firstForkBlock)
	// // bc.PrintChainFromRoot()
	// secondForkBlock := bc.GetForkBlock(string(bytes),firstForkHash,3)
	// secondForkHash := secondForkBlock.SHA256();
	// bc.AddBlock(secondForkBlock)
	// thirdForkBlock := bc.GetForkBlock(string(bytes),secondForkHash,4)
	// thirdForkHash := thirdForkBlock.SHA256();
	// bc.AddBlock(thirdForkBlock)
	// fourthForkBlock := bc.GetForkBlock(string(bytes),thirdForkHash,5)
	// _ = fourthForkBlock.SHA256();
	// bc.AddBlock(fourthForkBlock)
	

	// fmt.Println("************")
	// fmt.Println(bc.GetNthBlockFromRoot(2).Hash)
	// fmt.Println(bc.GetNthBlockFromRoot(3).Hash)
	// fmt.Println(bc.GetNthBlockFromRoot(4).Hash)
	// fmt.Println(bc.GetNthBlockFromRoot(5).Hash)
	// fmt.Println("************")
	// // bc.AddBlock(forkBlock)
	// bc.PrintChainFromRoot();
	// if(bc.IsValidChainFromEnd()){
	// 	fmt.Println("This is a valid Chain")
	// }else{
	// 	fmt.Println("This is an invalid Chain")
	// }

	// bc.SaveChainUsingRoot()
	fmt.Println("latest Block:",bc.GetLatestBlock().Index,bc.GetLatestBlock().Hash)
	peer.StartPeer(&bc)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func generateAdditionalBlock(n int){
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