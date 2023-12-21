package storage

import (
	"fmt"

	"github.com/shaojianqing/smilebc/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDatabase struct {
	filepath string
	database *leveldb.DB
}

type LevelBatch struct {
	database  *leveldb.DB
	ldbBatch  *leveldb.Batch
	valueSize int
}

func NewLevelDatabase(config config.DBConfig) (Database, error) {
	database, err := leveldb.OpenFile(config.DBFilePath, &opt.Options{
		OpenFilesCacheCapacity: config.Handlers,
		BlockCacheCapacity:     config.CacheSize / 2 * opt.MiB,
		WriteBuffer:            config.CacheSize / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if err != nil {
		return nil, fmt.Errorf("fail to create levelDB instance, err:%w", err)
	}

	levelDatabase := &LevelDatabase{
		filepath: config.DBFilePath,
		database: database,
	}
	return levelDatabase, nil
}

func (db *LevelDatabase) Set(key Key, value Value) error {
	return db.database.Put(key, value, nil)
}

func (db *LevelDatabase) Get(key Key) (Value, error) {
	return db.database.Get(key, nil)
}

func (db *LevelDatabase) Has(key Key) (bool, error) {
	return db.database.Has(key, nil)
}

func (db *LevelDatabase) Delete(key Key) error {
	return db.database.Delete(key, nil)
}

func (db *LevelDatabase) NewBatch() Batch {
	levelBatch := &leveldb.Batch{}
	return &LevelBatch{
		database:  db.database,
		ldbBatch:  levelBatch,
		valueSize: 0,
	}
}

func (db *LevelDatabase) Close() error {
	return db.database.Close()
}

func (lb *LevelBatch) Set(key Key, value Value) error {
	lb.ldbBatch.Put(key, value)
	lb.valueSize += len(value)
	return nil
}

func (lb *LevelBatch) ValueSize() int {
	return lb.valueSize
}

func (lb *LevelBatch) SyncWrite() error {
	return lb.database.Write(lb.ldbBatch, nil)
}
