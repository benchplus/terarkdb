#!/bin/bash
go env

go get github.com/syndtr/goleveldb/leveldb
go test -bench 'Benchmark' leveldb_test.go | tee output.txt
git stash
