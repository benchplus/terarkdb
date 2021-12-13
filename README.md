
[![terarkdb Workflow][terarkdb-badge]][terarkdb]
[![rocksdb Workflow][rocksdb-badge]][rocksdb]

Simple performance comparison between [TerarkDB](https://github.com/bytedance/terarkdb) and [RocksDB](https://github.com/facebook/rocksdb) 

# Backgrounds

Both of them with `jemalloc` open and with all the compression extensions.

Commit #f1b51f1 uses terark-zip table as tutorial said and former cases without terark-zip.

# Reports

[Continuous Bencmark Result (click here) ](https://benchplus.github.io/terarkdb/dev/bench/)

[terarkdb-badge]: https://github.com/benchplus/terarkdb/workflows/terarkdb/badge.svg
[rocksdb-badge]: https://github.com/benchplus/terarkdb/workflows/rocksdb/badge.svg
[terarkdb]: https://github.com/benchplus/terarkdb/actions?query=workflow%3A%22terarkdb%22
[rocksdb]: https://github.com/benchplus/terarkdb/actions?query=workflow%3A%22rocksdb%22
