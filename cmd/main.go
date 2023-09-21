package main

import (
	"log"

	"github.com/shaojianqing/smilebc/node"
)

func main() {

	smileNode := node.NewSmileNode()
	err := smileNode.StartService()
	if err != nil {
		log.Printf("start node service error:%v", err)
	}

	log.Println("start smile blockchain node successfully")
}
