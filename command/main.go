package main

import (
	"fmt"
	"log"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/crypto/sha3"
	"github.com/shaojianqing/smilebc/system"
)

const (
	ConfigFilePath = "/Users/shaojianqing/SmileBC/config.json"
)

func main() {

	configData, err := config.LoadConfigFromFile(ConfigFilePath)
	if err != nil {
		log.Fatalf("load configuration from file error:%v", err)
	}

	err = config.ValidateConfiguration(configData)
	if err != nil {
		log.Fatalf("validate configuration error:%v", err)
	}

	smile := system.NewSmile(configData)
	err = smile.StartService()
	if err != nil {
		log.Fatalf("start smile blockchain system service error:%v", err)
	}
	log.Println("start smile blockchain system successfully")

	hash := sha3.NewKeccak256()
	hash.Write([]byte("I am Shaojianqing!!"))
	hashValue := hash.Sum(nil)

	fmt.Printf("Hash Value is %x", string(hashValue))
}
