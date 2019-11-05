## A Million Strings

How much memory do a million strings (10 bytes long each) take as both a
`[]string` slice and a lookup hash (`map[string]int`)?

```
Estimated           ~= 9.54MB (1000000 * 10 bytes / 1024 / 1024)
map[string]int      ~= 78MB (notice int32 cost)
map[string]struct{} ~= 61MB (empty struct{})
[]string            ~= 33MB
```
