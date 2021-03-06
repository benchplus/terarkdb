#!/bin/bash

BASE=$PWD

git submodule update --init --recursive
sudo apt-get install libaio-dev

cd third-party/terarkdb
git remote add mine https://github.com/ez8-lab/terarkdb.git
git pull mine dev.1.4

OUTPUT=$BASE/third-party/terarkdb/output
mkdir -p $OUTPUT
cd $OUTPUT
cmake ../ -DCMAKE_INSTALL_PREFIX=output -DCMAKE_BUILD_TYPE=Release -DWITH_TESTS=OFF -DWITH_TERARK_ZIP=ON
make -j $(nproc)
make install

cp $OUTPUT/output/lib/libterarkdb.a $OUTPUT/lib/librocksdb.a

cd $BASE
