package main

import (
	"log"

	"github.com/shaojianqing/smilebc/config"
	"github.com/shaojianqing/smilebc/node"
)

const (
	ConfigFilePath = "~/SmileBC/config.json"
)

func main() {

	configData := config.LoadConfigFromFile(ConfigFilePath)
	smileNode := node.NewSmileNode(configData)
	err := smileNode.StartService()
	if err != nil {
		log.Printf("start smile blockchain node service error:%v", err)
	}

	log.Println("start smile blockchain node successfully")
}
