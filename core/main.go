package main

import (
	"fmt"
	. "../models"
	. "../peer"
	)


func main() {
	fmt.Println("programme started");
	fmt.Println("blocks start for 0 to n-1");

	//Initiate a new blockchain with genesis block
	bc := NewBlockchain()

	//Adding 5 blocks to the chain
	secondBlock := bc.GenerateNextBlock("This is second blockdata")
	bc.AddBlock(secondBlock);

	thirdBlock := bc.GenerateNextBlock("This is third blockdata")
	bc.AddBlock(thirdBlock);

	fourthBlock := bc.GenerateNextBlock("This is fourth blockdata")
	bc.AddBlock(fourthBlock);

	fifthBlock := bc.GenerateNextBlock("This is fifth blockdata")
	bc.AddBlock(fifthBlock);

	sixthBlock := bc.GenerateNextBlock("This is sixth blockdata")
	bc.AddBlock(sixthBlock);
	//adding 5 blocks complete

	//generating 7th block
	seventhBlock := bc.GenerateNextBlock("This is seventh blockdata")
	//let us modify hash of seventh block and add this to chain
	//uncomment the below and try to add
	// seventhBlock.hash = "asjkfhskjdhkjasdhk"
	bc.AddBlock(seventhBlock);


	eightBlock := bc.GenerateNextBlock("This is eight blockdata. My name is akshay")
	bc.AddBlock(eightBlock);

	fmt.Println(bc.IsValidChain());
	fmt.Println(bc.Blocks)
	StartPeer(&bc)

}