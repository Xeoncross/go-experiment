## A Million Strings

How much memory do a million strings (10 bytes long each) take as both a
`[]string` slice and a lookup hash (`map[string]int`)?

```
Estimated           ~= 9.54MB (1000000 * 10 bytes / 1024 / 1024)
map[string]int      ~= 78MB (notice int32 cost)
map[string]struct{} ~= 61MB (empty struct{})
[]string            ~= 33MB
```

Furthermore, how much faster is a binary search over a map lookup?

With 1 million 20 byte strings

```
BenchmarkBinarySearch-8   	 3361562	       322 ns/op	      24 B/op	       2 allocs/op
BenchmarkMapSearch-8      	 4069456	       284 ns/op	      24 B/op	       2 allocs/op
```

And 1 million 40 byte strings

```
BenchmarkBinarySearch-8   	 2769397	       392 ns/op	      56 B/op	       2 allocs/op
BenchmarkMapSearch-8      	 3997479	       293 ns/op	      56 B/op	       2 allocs/op
```

And 1 million hashes

```
BenchmarkBinarySearch-8   	  922741	      1279 ns/op	     212 B/op	       5 allocs/op
BenchmarkMapSearch-8      	 2426865	       467 ns/op	     212 B/op	       5 allocs/op
```
