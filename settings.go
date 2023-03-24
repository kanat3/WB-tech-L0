package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type settings struct {
	ServerHost string
	ServerPort string
	HTML       string
	Images     string
}

var cfg settings

func init() {
	//try read config
	file, err := os.Open("settings.cfg")
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
	err = json.Unmarshal(readByte, &cfg)
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't read configuratiion data")
	}
}
