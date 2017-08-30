package controller

import(
	"os"
	"fmt"
	"io/ioutil"
)
func CreateFile(previousHash string, blockData []byte) {
	// detect if file exists
	path:= "../data/"+previousHash+".json"
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { return }
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
	writeFile(path,blockData)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func writeFile(path string,data []byte) {
	// open file using READ & WRITE permission
	err := ioutil.WriteFile(path, data, 0644)

	if isError(err){
		fmt.Println("==> failed writing to file")
	}else{

		fmt.Println("==> done writing to file")
	}
}