package main

import "fmt"

const (
	ConfigFilePath = "/Users/shaojianqing/SmileBC/config.json"
)

func main() {

	/*configData, err := config.LoadConfigFromFile(ConfigFilePath)
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

	fmt.Printf("Hash Value is %x", string(hashValue))*/

	/*aesEncryptionKey, err := ecies.GenerateAESKey(p2p.AESKeyLength)
	if err != nil {
		log.Fatalf("generate AES encryption key error:%v", err)
	}

	secret, err := p2p.NewSecretTest(aesEncryptionKey)
	if err != nil {
		log.Fatalf("new secret object error:%v", err)
	}

	plain := "I am Smith Shao, I am testing AES encryption^!^"
	cipher := secret.Encrypt([]byte(plain))
	fmt.Printf("cipher from plain:%s\n", cipher)

	newPlain := secret.Decrypt(cipher)
	fmt.Printf("new plain from cipher:%s\n", newPlain)*/

	maxUint24 := ^uint32(0) >> 8
	fmt.Printf("maxUint24:%d", maxUint24)
}
