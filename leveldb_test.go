package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

var defaultLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func randomString(n int) []byte {

	letters := defaultLetters

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return b
}

// I2b returns an 8-byte big endian representation of v
// v uint64(123456) -> 8-byte big endian.
func I2b(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

var db *leveldb.DB

func init() {
	var err error
	if db, err = leveldb.OpenFile("data", nil); err != nil {
		fmt.Println("open err")
	}
}

func shutdown() {
	db.Close()
}

const prefix4K = "user/"

func BenchmarkSequenceWrite4K(b *testing.B) {
	data := randomString(4096)
	batch := new(leveldb.Batch)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		batch.Put(key, data)
	}
	db.Write(batch, nil)
}

func BenchmarkRandWrite4K(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Put(key, data, nil)
	}
}

func BenchmarkRandRead4K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Get(key, nil)
	}
}

const prefix4M = "user1/"

func BenchmarkSequenceWrite4M(b *testing.B) {
	data := randomString(4194304)
	batch := new(leveldb.Batch)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4M), I2b(uint64(tmp))...)
		batch.Put(key, data)
	}
	db.Write(batch, nil)
}

func BenchmarkRandWrite4M(b *testing.B) {
	data := randomString(4194304)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4M), I2b(uint64(tmp))...)
		db.Put(key, data, nil)
	}
}

func BenchmarkRandRead4M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4M), I2b(uint64(tmp))...)
		db.Get(key, nil)
	}
}

func BenchmarkRandDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Delete(key, nil)
	}
}

func BenchmarkGetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter := db.NewIterator(nil, nil)
		for iter.Next() {
		}
		iter.Release()
		if err := iter.Error(); err != nil {
			fmt.Println("iter err")
		}
	}
}
