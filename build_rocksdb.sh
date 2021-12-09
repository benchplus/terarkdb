#!/bin/bash

BASE=$PWD

git submodule update --init --recursive

OUTPUT=$BASE/third-party/rocksdb/output
mkdir -p $OUTPUT

cd third-party/terarkdb/third-party/jemalloc

bash autogen.sh
CFLAGS=-fPIC CXXFLAGS=-fPIC LDFLAGS=-fPIC ./configure --enable-prof
make -j $(nproc)
sudo make install
sudo ldconfig

cd $OUTPUT

sudo apt-get install libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev liblz4-dev libzstd-dev -y

cmake ../ -DCMAKE_INSTALL_PREFIX=output -DCMAKE_BUILD_TYPE=Release -DWITH_JEMALLOC=ON -DWITH_SNAPPY=ON -DWITH_LZ4=ON -DWITH_ZLIB=ON -DWITH_ZSTD=ON -DROCKSDB_BUILD_SHARED=OFF
make -j $(nproc)
make install

cd $BASE
