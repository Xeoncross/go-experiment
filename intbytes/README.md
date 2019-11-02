# Integers to a []byte slice

Converting integers into bytes for storage.

Does using https://golang.org/pkg/encoding/binary/#Write help us?


## Result

No. Perhaps it would if we were writing larger slices of integers.

    go test -bench=. --benchmem

```
goos: darwin
goarch: amd64
pkg: github.com/Xeoncross/go-experiment/intbytes
BenchmarkBinary-8   	   27110	     42448 ns/op	    9104 B/op	    1999 allocs/op
BenchmarkRaw-8      	  437797	      8432 ns/op	   11272 B/op	       0 allocs/op
PASS
ok  	github.com/Xeoncross/go-experiment/intbytes	5.502s
```
