package storage

import "github.com/shaojianqing/smilebc/config"

type Database interface {
}

func NewDatabase(config config.DBConfig) (Database, error) {
	return nil, nil
}
