package concensus

import(
	"crypto/sha256"
	"math/big"
)

var stopBlockGeneration bool

func StopBlockGeneration(){
	stopBlockGeneration = true;
}
func ReturnNonce(s string) (string,bool){
	stopBlockGeneration = false;
	target:= big.NewInt(0)
	target.SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",16)
	difficulty := big.NewInt(0)
	difficulty.SetString("fffff00000000000000000000000000000000000000000000000000000000000",16)
	target.Sub(target,difficulty)
	nonce:= big.NewInt(0)
	var byteArray [32]byte
	for ;;{
		if stopBlockGeneration == true{
			return nonce.String(),false
		}
		byteArray = sha256.Sum256([]byte(s+nonce.String()))
		currentNumber:= big.NewInt(0).SetBytes(byteArray[:])
		if(currentNumber.Cmp(target) == -1){
			break;
		}else{
			nonce.Add(nonce,big.NewInt(1))
			// fmt.Println(nonce.String())
		}
	}
	return nonce.String(),true;
}