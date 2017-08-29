package controller

import(
	"io/ioutil"
	"encoding/json"
)

func GetAccount(name string) (string, string){
	// re-open file
	var key Key
	 b, err := ioutil.ReadFile("../accounts/"+name+".json") // just pass the file name
	if err != nil {
        return (""),("")
    }
	
	err = json.Unmarshal(b,&key)
	if err!=nil {
		return (""),("")
	}
	return key.Private,key.Public
}