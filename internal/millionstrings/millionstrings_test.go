package millionstrings

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"runtime"
	"sort"
	"testing"
)

func reportMemory(name string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%20s MB: %5.2f\n", name, float64(m.Alloc)/1024/1024)
}

var size = 1000000

func TestMillionStrings(t *testing.T) {

	fmt.Printf("Estimated size of %d strings: MB: %5.2f\n", size, float64(size)*10/1024/1024) // 10 bytes long

	var dict = make(map[string]int, size)
	for i := 0; i < size; i++ {
		dict[fmt.Sprintf("%10d", i)] = i
	}

	reportMemory("map[string]int")
	dict = make(map[string]int)
	runtime.GC()
	reportMemory("")

	var hash = make(map[string]struct{}, size)
	for i := 0; i < size; i++ {
		hash[fmt.Sprintf("%10d", i)] = struct{}{}
	}

	reportMemory("map[string]struct{}")

	hash = make(map[string]struct{})
	runtime.GC()

	reportMemory("")

	var list = make([]string, size)
	for i := 0; i < size; i++ {
		list[i] = fmt.Sprintf("%10d", i)
	}

	reportMemory("[]string")

}

func BenchmarkBinarySearch(b *testing.B) {

	var list = make([]string, size)
	for i := 0; i < size; i++ {
		// list[i] = fmt.Sprintf("%40d", i)
		list[i] = hashKey(i)
	}

	sort.Strings(list)

	b.ResetTimer()
	var found int
	for i := 0; i < b.N; i++ {
		// key := fmt.Sprintf("%40d", i)
		key := hashKey(i)

		i := sort.SearchStrings(list, key)
		if i < len(list) && list[i] == key {
			found++
		}
	}

	want := b.N
	if want > size {
		want = size
	}

	if found != want {
		b.Errorf("Search failure: got %d, want %d", found, want)
	}

}

func BenchmarkMapSearch(b *testing.B) {

	var hash = make(map[string]struct{}, size)
	for i := 0; i < size; i++ {
		// hash[fmt.Sprintf("%40d", i)] = struct{}{}
		hash[hashKey(i)] = struct{}{}
	}

	b.ResetTimer()
	var found int
	for i := 0; i < b.N; i++ {
		// key := fmt.Sprintf("%40d", i)
		key := hashKey(i)

		if _, ok := hash[key]; ok {
			found++
		}
	}

	want := b.N
	if want > size {
		want = size
	}

	if found != want {
		b.Errorf("Map search failure: got %d, want %d", found, want)
	}

}

func hashKey(i int) string {
	var b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	hasher := sha1.New()
	hasher.Write(b)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
