#!/bin/bash

BASE=$PWD

git submodule update --init --recursive

cd third-party/rocksdb

OUTPUT=$BASE/third-party/rocksdb/output
mkdir -p $OUTPUT
cd $OUTPUT

sudo apt-get install libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev liblz4-dev libzstd-dev -y

cmake ../ -DCMAKE_INSTALL_PREFIX=output -DCMAKE_BUILD_TYPE=Release -DWITH_SNAPPY=ON -DWITH_LZ4=ON -DWITH_ZLIB=ON -DWITH_ZSTD=ON
make -j $(nproc)
make install

cd $BASE
