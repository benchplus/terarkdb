#!/bin/bash
go env

export CGO_CFLAGS="-I$PWD/third-party/rocksdb/output/output/include"
export CGO_LDFLAGS="-L$PWD/third-party/rocksdb/output/output/lib -Wl,-Bstatic -lrocksdb -lbz2 -ljemalloc -llz4 -lsnappy -lz -lzstd -Wl,-Bdynamic -lstdc++ -lm -pthread -lgomp -lrt -ldl"

go get github.com/huxinhuxin/gorocksdb
go test -bench 'Benchmark' | tee output.txt
