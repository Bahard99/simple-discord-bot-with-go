package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token		string //store Token from config.json
	BotPrefix	string //store BotPrefix from config.json

	config *configStruct //store value extracted from config.json
)

type configStruct struct {
	Token		string	`json:	"Token"`
	BotPrefix	string	`json:	"BotPrefix"`
}

func ReadConfig() error {
	
	fmt.Println("Reading config file ....")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	} 

	fmt.Println(string(file))
	
	//copy file variable value into var config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token 	  = config.Token
	BotPrefix = config.BotPrefix

	return nil
}