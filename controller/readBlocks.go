package controller

import(
	"io/ioutil"
)

func ReadFile(previousHash string) []byte{
	// re-open file
	 b, err := ioutil.ReadFile("../data/"+previousHash+".json") // just pass the file name
    if err != nil {
        return []byte("")
    }

	return b
}