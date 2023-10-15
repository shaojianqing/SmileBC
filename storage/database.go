package storage

import "github.com/shaojianqing/smilebc/config"

type (
	Key   []byte
	Value []byte
)

type Batch interface {
	Set(key Key, value Value) error
	ValueSize() int
	SyncWrite() error
}

type Database interface {
	Set(key Key, value Value) error
	Get(key Key) (Value, error)
	Has(key Key) (bool, error)
	Delete(key Key) error
	NewBatch() Batch
	Close() error
}

func NewDatabase(config config.DBConfig) (Database, error) {
	// Currently we only support the levelDB as the underlying key-value storage infrastructure.
	// So it returns the levelDB instance by default
	return NewLevelDatabase(config)
}
