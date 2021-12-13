
[![terarkdb Workflow][terarkdb-badge]][terarkdb]
[![rocksdb Workflow][rocksdb-badge]][rocksdb]

Simple performance comparison between [TerarkDB](https://github.com/bytedance/terarkdb) and [RocksDB](https://github.com/facebook/rocksdb) 

# Backgrounds

Both of them with `jemalloc` open and with all the compression extensions.

Enable [`terark-zip`](https://github.com/bytedance/terark-zip) table according to [tutorial](https://github.com/bytedance/terarkdb#22-terarkziptable) says besides the commits before [f1b51f1](https://github.com/benchplus/terarkdb/commit/f1b51f172c2265dadb29c7c3acfdf7bdab8eb73d).

# Reports

[Continuous Bencmark Result (click here) ](https://benchplus.github.io/terarkdb/dev/bench/)

[terarkdb-badge]: https://github.com/benchplus/terarkdb/workflows/terarkdb/badge.svg
[rocksdb-badge]: https://github.com/benchplus/terarkdb/workflows/rocksdb/badge.svg
[terarkdb]: https://github.com/benchplus/terarkdb/actions?query=workflow%3A%22terarkdb%22
[rocksdb]: https://github.com/benchplus/terarkdb/actions?query=workflow%3A%22rocksdb%22
