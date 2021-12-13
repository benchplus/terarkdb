#!/bin/bash
go env

export CGO_CFLAGS="-I$PWD/third-party/terarkdb/output/output/include"
export CGO_LDFLAGS="-L$PWD/third-party/terarkdb/output/lib -Wl,-Bstatic -lrocksdb -lbz2 -ljemalloc -llz4 -lsnappy -lz -lzstd -lterark-zip-r -lboost_system -lboost_filesystem -lboost_fiber -lboost_context -Wl,-Bdynamic -lstdc++ -lm -pthread -lgomp -lrt -ldl -laio"

go get github.com/ez8-lab/gorocksdb
go test -bench 'Benchmark' -benchtime=60s | tee output.txt
