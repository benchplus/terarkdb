window.BENCHMARK_DATA = {
  "lastUpdate": 1638961002046,
  "repoUrl": "https://github.com/benchplus/terarkdb",
  "entries": {
    "terarkdb": [
      {
        "commit": {
          "author": {
            "email": "orca.zhang@yahoo.com",
            "name": "Orca",
            "username": "orca-zhang"
          },
          "committer": {
            "email": "orca.zhang@yahoo.com",
            "name": "Orca",
            "username": "orca-zhang"
          },
          "distinct": true,
          "id": "8f3cedfba15cfacfe2243132a2bc9cd44dc1efce",
          "message": ":bee: fix",
          "timestamp": "2021-12-08T06:06:37+08:00",
          "tree_id": "bc1f89a0fc542b665df148fc49fa10d63b844440",
          "url": "https://github.com/orcastor/orcas-engine-terarkdb/commit/8f3cedfba15cfacfe2243132a2bc9cd44dc1efce"
        },
        "date": 1638916490499,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSequenceWrite",
            "value": 25547,
            "unit": "ns/op",
            "extra": "200000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandWrite",
            "value": 28167,
            "unit": "ns/op",
            "extra": "100000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandRead",
            "value": 21048,
            "unit": "ns/op",
            "extra": "100000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandDel",
            "value": 6405,
            "unit": "ns/op",
            "extra": "300000 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "orca.zhang@yahoo.com",
            "name": "Orca",
            "username": "orca-zhang"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "1a76eff877e4dbf2da2e6a0792acfb829e32a2ee",
          "message": "🦉 add result link",
          "timestamp": "2021-12-08T12:31:20+08:00",
          "tree_id": "5c7128cf36911e0109ce9af9092b40e52fd6751d",
          "url": "https://github.com/orca-zhang/terarkdb/commit/1a76eff877e4dbf2da2e6a0792acfb829e32a2ee"
        },
        "date": 1638961001543,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkSequenceWrite",
            "value": 24832,
            "unit": "ns/op",
            "extra": "200000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandWrite",
            "value": 35999,
            "unit": "ns/op",
            "extra": "200000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandRead",
            "value": 16458,
            "unit": "ns/op",
            "extra": "200000 times\n2 procs"
          },
          {
            "name": "BenchmarkRandDel",
            "value": 5112,
            "unit": "ns/op",
            "extra": "300000 times\n2 procs"
          }
        ]
      }
    ]
  }
}