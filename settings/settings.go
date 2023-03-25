package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	ServerHost string
	ServerPort string
	Static     string
	Images     string
}

func NewConfig(any interface{}, fName string) {

	//try read config
	file, err := os.Open(fName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't open the configuration")
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		panic("Cant't read the structure of the configuration")
	}

	readByte := make([]byte, stat.Size())

	_, err = file.Read(readByte)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't read the configuration")
	}
	err = json.Unmarshal(readByte, &any)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't read configuratiion data")
	}
}
