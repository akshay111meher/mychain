package controller

import(
	"io/ioutil"
)

func ReadFile(blockNumber string) []byte{
	// re-open file
	 b, err := ioutil.ReadFile("../data/"+blockNumber+".json") // just pass the file name
    if err != nil {
        return []byte("")
    }

	return b
}