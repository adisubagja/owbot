package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	TgId int64 `json:"id"`
	TgBotToken string `json:"token"`
}


func ConfigFromFile(path string) (Config,error){
	cfgByteArray,err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = json.Unmarshal(cfgByteArray,&config)
	if err != nil {
		return Config{}, err
	}

	return config,nil
}

