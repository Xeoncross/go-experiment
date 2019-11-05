package millionstrings

import (
	"fmt"
	"runtime"
	"testing"
)

func reportMemory(name string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%20s MB: %5.2f\n", name, float64(m.Alloc)/1024/1024)
}

func TestMillionStrings(t *testing.T) {

	size := 1000000

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
