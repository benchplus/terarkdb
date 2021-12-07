package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

var db Gorocksdb

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
