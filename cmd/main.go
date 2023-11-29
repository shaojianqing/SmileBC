package main

import (
	"fmt"
	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/node"
	"log"

	"github.com/shaojianqing/smilebc/crypto/sha3"
)

const (
	ConfigFilePath = "/Users/shaojianqing/SmileBC/config.json"
)

func main() {

	configData, err := config.LoadConfigFromFile(ConfigFilePath)
	if err != nil {
		log.Printf("load configuration from file error:%v", err)
	}
	smileNode := node.NewSmileNode(configData)
	err = smileNode.StartService()
	if err != nil {
		log.Printf("start smile blockchain node service error:%v", err)
	}
	log.Println("start smile blockchain node successfully")

	hash := sha3.NewKeccak256()
	hash.Write([]byte("I am Shaojianqing!!"))
	hashValue := hash.Sum(nil)

	fmt.Printf("Hash Value is %x", string(hashValue))
}
