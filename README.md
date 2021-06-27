# bpool

minio 的字节池

```bash
$ make
==> Benchmark
goos: linux
goarch: amd64
pkg: github.com/maolonglong/bpool
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkBytePool-4         1642            719147 ns/op           40663 B/op       1155 allocs/op
BenchmarkSyncPool-4         1280            840009 ns/op           55862 B/op       1711 allocs/op
BenchmarkRaw-4               576           2198178 ns/op          572060 B/op       2001 allocs/op
PASS
ok      github.com/maolonglong/bpool    5.212s
```
