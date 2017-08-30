package models

import(
	"strconv"
	"fmt"
	. "../crypto"
	"encoding/json"
)


type Block struct{
	Index string
	PreviousHash string
	Timestamp string
	Data string
	Hash string
	Nonce string
}

func (b Block) SHA256() string{
	return SHA256(b.Index+b.PreviousHash+b.Timestamp+b.Data)
}

func (b Block) IsNextBlockValid(nextBlock Block) bool{
	currentIndex,_ := strconv.Atoi(b.Index);
	newBlockIndex,_ := strconv.Atoi(nextBlock.Index);
	if(currentIndex + 1 != newBlockIndex){
		fmt.Println("Index mismatch. Index "+nextBlock.Index)
		return false;
	}else if(b.Hash != nextBlock.PreviousHash){
		fmt.Println("Hash mismatch with previous block. Index "+nextBlock.Index)
		return false;
	}else if(nextBlock.Hash != nextBlock.SHA256()){
		fmt.Println("Current block hash computed wrongly. Index "+nextBlock.Index)
		return false
	}else if(!nextBlock.IsThisDataValid()){
		fmt.Println("Data-Value format mismatch")
		return false
	}else{
		return true
	}
}

func (b Block) IsThisBlockValid() bool{
	if b.Hash == b.SHA256() {
		return true
	}else{
		return false;
	}
}

func (b Block) IsThisDataValid() bool{
	var d Data
	json.Unmarshal([]byte(b.Data),&d)

	return Verify(d.Value,d.PublicKey,d.R,d.S)
}

