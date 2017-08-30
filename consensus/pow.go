package concensus

import(
	"crypto/sha256"
	"math/big"
)

func ReturnNonce(s string) string{
	target:= big.NewInt(0)
	target.SetString("00000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",16)
	nonce:= big.NewInt(0)
	var byteArray [32]byte
	for ;;{
		byteArray = sha256.Sum256([]byte(s+nonce.String()))
		currentNumber:= big.NewInt(0).SetBytes(byteArray[:])
		if(currentNumber.Cmp(target) == -1){
			break;
		}else{
			nonce.Add(nonce,big.NewInt(1))
			// fmt.Println(nonce.String())
		}
	}
	return nonce.String();
}