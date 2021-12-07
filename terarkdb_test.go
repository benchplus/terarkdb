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

const prikey = "user/"

func BenchmarkSequenceWrite(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		key := append([]byte(prikey), I2b(uint64(i))...)
		db.Set(key, data)
	}
}

func BenchmarkRandWrite(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prikey), I2b(uint64(tmp))...)
		db.Set(key, data)
	}
}

func BenchmarkRandRead(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prikey), I2b(uint64(tmp))...)
		db.Get(key)
	}
}
func BenchmarkRandDel(b *testing.B) {
	data := randomString(4096)
	for i := 0; i < b.N; i++ {
		tmp := rand.Int31n(int32(b.N))
		key := append([]byte(prikey), I2b(uint64(tmp))...)
		db.Del(key)
	}
}
