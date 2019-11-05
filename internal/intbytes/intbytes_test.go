package intbytes

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func BenchmarkBinary(b *testing.B) {

	var values []uint16
	var i uint16
	for i = 0; i < 1000; i++ {
		values = append(values, i)
	}

	b.ResetTimer()

	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for _, value := range values {
			err := binary.Write(buf, binary.LittleEndian, uint16(value))
			if err != nil {
				b.Error(err)
			}
		}
	}

	// fmt.Printf("Length: %d\n", buf.Len())
	buf.Len()
}

func BenchmarkRaw(b *testing.B) {

	var values []uint16
	var i uint16
	for i = 0; i < 1000; i++ {
		values = append(values, i)
	}

	b.ResetTimer()

	var buf []byte
	v := make([]byte, 2)

	for i := 0; i < b.N; i++ {
		for _, value := range values {
			binary.BigEndian.PutUint16(v, uint16(value))
			buf = append(buf, v...)
		}
	}

	_ = len(buf)
}
