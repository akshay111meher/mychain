package main

import(
	"fmt"
	. "../controller"
	. "../crypto"
	// "reflect"
	// "crypto/ecdsa"
)

func Test(){
	fmt.Println("test")
}

func main(){
	
	CreateAccount("prateek");
	privKey,pubKey := GetAccount("prateek")
	data :="this is the data to be signed";
	r,s := GetSignature(data,privKey)
	fmt.Println(Verify(data,pubKey,r,s))

	
}