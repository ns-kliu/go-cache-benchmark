## Go Cache Benchmarks

Benchmarks of in-memory cache libraries for Go with expiration/TTL support.

## Run Test

    go get
    go test -bench=. -benchmem -benchtime=10000000x

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/Xeoncross/go-cache-benchmark
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkGoCache/set-16          	10000000	       891.3 ns/op	     310 B/op	       9 allocs/op
BenchmarkGoCache/get-16          	10000000	       312.0 ns/op	      53 B/op	       1 allocs/op
BenchmarkGoCacheParallel-16      	10000000	      1316 ns/op	     364 B/op	      11 allocs/op
BenchmarkFreecache/Set-16        	10000000	       600.1 ns/op	     203 B/op	       5 allocs/op
BenchmarkFreecache/Get-16        	10000000	       415.9 ns/op	     150 B/op	       3 allocs/op
BenchmarkFreecacheParallel-16    	10000000	       822.6 ns/op	     444 B/op	      11 allocs/op
BenchmarkFsWatcher/Set-16        	10000000	       510.8 ns/op	     132 B/op	       4 allocs/op
BenchmarkFsWatcher/Get-16        	10000000	       307.9 ns/op	      53 B/op	       1 allocs/op
BenchmarkFsWatcherParallel-16    	10000000	      1352 ns/op	     276 B/op	       8 allocs/op
PASS
```
