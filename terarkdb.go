package main

import (
	"errors"

	"github.com/ez8-lab/gorocksdb"
)

type Gorocksdb struct {
	DB    *gorocksdb.DB
	Bsync bool
}

func NewRocksDB() *Gorocksdb {
	return &Gorocksdb{}
}

func (db *Gorocksdb) Open(path string, sync bool) error {
	opt := gorocksdb.NewDefaultOptions()
	env := gorocksdb.NewDefaultEnv()
	env.SetBackgroundThreads(3)
	opt.SetEnv(env)
	opt.SetBytesPerSync(32768)
	opt.SetMaxBackgroundCompactions(5)
	opt.SetCreateIfMissing(true)

	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockSize(8192)

	tzto := gorocksdb.NewDefaultTerarkZipTableOptions()
	tzto.SetLocalTempDir("/tmp")
	tzto.SetIndexNestLevel(3)
	tzto.SetSampleRatio(0.01)
	tzto.SetTerarkZipMinLevel(2)

	opt.SetTerarkZipTableFactory(tzto, bbto)

	database, err := gorocksdb.OpenDb(opt, path)
	if err != nil {
		return err
	}
	db.DB = database
	db.Bsync = sync
	return nil
}

func (db *Gorocksdb) Close() error {
	db.DB.Close()
	return nil
}

func (db *Gorocksdb) Get(key []byte) ([]byte, error) {
	opt := gorocksdb.NewDefaultReadOptions()

	v, err := db.DB.Get(opt, key)

	if v == nil {
		err = errors.New("keyNotFound")
	}
	return v.Data(), err
}

func (db *Gorocksdb) Set(key, val []byte) error {
	opts := gorocksdb.NewDefaultWriteOptions()
	opts.SetSync(db.Bsync)
	return db.DB.Put(opts, key, val)
}

func (db *Gorocksdb) Write(key [][]byte, val []byte) error {
	wb := gorocksdb.NewWriteBatch()
	for i := 0; i < len(key); i++ {
		wb.Put(key[i], val)
	}
	opts := gorocksdb.NewDefaultWriteOptions()
	opts.SetSync(db.Bsync)
	return db.DB.Write(opts, wb)
}

func (db *Gorocksdb) Del(key []byte) error {
	opts := gorocksdb.NewDefaultWriteOptions()
	opts.SetSync(db.Bsync)
	return db.DB.Delete(opts, key)
}

func (db *Gorocksdb) GetAll() (int, error) {
	var cout int
	opt := gorocksdb.NewDefaultReadOptions()
	iter := db.DB.NewIterator(opt)
	defer iter.Close()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		cout++
	}
	return cout, nil
}
