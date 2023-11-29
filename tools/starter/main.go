package main

import (
	"fmt"
	tools "github.com/shaojianqing/smilebc/tools/database"
)

const (
	LevelDBFilepath = "/Users/shaojianqing/LevelDB/smile.ldb"
)

func main() {
	database, err := tools.CreateLevelDB(LevelDBFilepath)
	if err != nil {
		fmt.Printf("can not create levelDB, err:%v", err)
		return
	}

	nameKey := "name"
	nameValue := "Smith Shao"

	nationalityKey := "nationality"
	nationalityValue := "The People's Republic of China"

	database.Put([]byte(nameKey), []byte(nameValue), nil)
	database.Put([]byte(nationalityKey), []byte(nationalityValue), nil)

	value, err := database.Get([]byte(nameKey), nil)
	if err != nil {
		fmt.Printf("can read value from levelDB, err:%v", err)
		return
	}
	fmt.Printf("Name:%s\n", value)

	value, err = database.Get([]byte(nationalityKey), nil)
	if err != nil {
		fmt.Printf("can read value from levelDB, err:%v", err)
		return
	}
	fmt.Printf("Nationality:%s\n", value)
}
