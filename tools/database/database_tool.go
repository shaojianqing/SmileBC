package tools

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func CreateLevelDB(filepath string) (*leveldb.DB, error) {
	database, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to create or open levelDB, err:%w", err)
	}
	return database, nil
}
