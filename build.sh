#!/bin/bash

BASE=$PWD

git submodule update --init --recursive

cd third-party/terarkdb

git remote add mine https://github.com/ez8-lab/terarkdb.git
git pull mine dev.1.4

OUTPUT=$BASE/third-party/terarkdb/output
mkdir -p $OUTPUT
cd $OUTPUT

sudo apt-get install libaio-dev -y
cmake ../ -DCMAKE_INSTALL_PREFIX=output -DCMAKE_BUILD_TYPE=Release -DWITH_TESTS=OFF -DWITH_TERARK_ZIP=OFF
make -j $(nproc)
make install

cp $OUTPUT/output/lib/libterarkdb.a $OUTPUT/lib/librocksdb.a

cd $BASE
