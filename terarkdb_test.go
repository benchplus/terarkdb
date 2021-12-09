package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"testing"
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

var db *Gorocksdb

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	db = NewRocksDB()
	if err := db.Open("data", false); err != nil {
		fmt.Println("open err")
	}
}

func shutdown() {
	db.Close()
}

const prefix4K = "user/"

func BenchmarkSequenceWrite4K(b *testing.B) {
	data := randomString(4096)
	var key [][]byte
	for i := 0; i < b.N; i++ {
		key = apend(key, append([]byte(prefix4K), I2b(uint64(i))...))
	}
	db.Write(key, data)
}

func BenchmarkRandWrite4K(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Set(key, data)
	}
}

func BenchmarkRandRead4K(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Get(key)
	}
}

const prefix4M = "user1/"

func BenchmarkSequenceWrite4M(b *testing.B) {
	data := randomString(4194304)
	var key [][]byte
	for i := 0; i < b.N; i++ {
		key = append(key, append([]byte(prefix4K), I2b(uint64(i))...))
	}
	db.Write(key, data)
}

func BenchmarkRandWrite4M(b *testing.B) {
	data := randomString(4194304)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4M), I2b(uint64(tmp))...)
		db.Set(key, data)
	}
}

func BenchmarkRandRead4M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4M), I2b(uint64(tmp))...)
		db.Get(key)
	}
}

func BenchmarkRandDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prefix4K), I2b(uint64(tmp))...)
		db.Del(key)
	}
}

func BenchmarkGetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db.GetAll()
	}
}
