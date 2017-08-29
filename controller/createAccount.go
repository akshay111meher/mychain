package controller

import(
	. "../crypto"
	"fmt"
	"os"
	"io/ioutil"
	// "crypto/ecdsa"
	// "reflect"
	"encoding/json"
	"encoding/hex"
)

func CreateAccount(name string) bool{
	privateKeyBytes,publicKeyBytes:= GenerateKeys()
	path:= "../accounts/"+name+".json"
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { }
		defer file.Close()
	}else{
		return false
	}

	fmt.Println("==> done creating file", path)
	return saveAccount(path,privateKeyBytes,publicKeyBytes)
}

func saveAccount(path string, privateKey, publicKey []byte) bool{

	key := Key{hex.EncodeToString(privateKey[:]),hex.EncodeToString(publicKey[:])}
	keyBytes,_ := json.Marshal(key)
	err := ioutil.WriteFile(path, keyBytes, 0644)

		if isError(err){
			fmt.Println("==> failed writing to file")
			return false
		}else{
			fmt.Println("==> done writing to file")
			return true;
		}
}