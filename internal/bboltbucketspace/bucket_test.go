// +build large

// go test --tags=large
package bboltbucketspace

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"testing"

	"go.etcd.io/bbolt"
)

var bucketDBPath = "bucket.db"
var keyDBPath = "key.db"

var topLevelKeys = 1000
var bottomLevelKeys = 100

// func toKey(a, b int) []byte {
//   var output string
//   for _, i := range ints {
//
//   return []byte(fmt.Sprintf("%d")
// }

func TestBucketStorage(t *testing.T) {

	os.Remove(bucketDBPath)
	os.Remove(bucketDBPath + ".copy")

	db, err := bbolt.Open(bucketDBPath, 0600, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {

		for i := 0; i < topLevelKeys; i++ {

			if i%100 == 0 {
				fmt.Printf("Bucket: Processed: %d\n", i)
			}

			// 8 byte long major key (represents a "word")
			majorKey := []byte(fmt.Sprintf("%8d", i))

			b, err := tx.CreateBucketIfNotExists(majorKey)
			if err != nil {
				return err
			}

			for j := 0; j < bottomLevelKeys; j++ {
				// err := b.Put([]byte(fmt.Sprintf("%d|%d", i, j)), value)

				err := b.Put(itob(j), itob(j))
				if err != nil {
					return err
				}
			}

		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	err = db.View(func(tx *bbolt.Tx) error {
		return tx.CopyFile(bucketDBPath+".copy", 0666)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	size, err := GetFileSize(bucketDBPath + ".copy")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%20s: %d KB\n", bucketDBPath+".copy", size/1024)
}

func TestKeyStorage(t *testing.T) {

	os.Remove(keyDBPath)
	os.Remove(keyDBPath + ".copy")

	db, err := bbolt.Open(keyDBPath, 0600, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {

		for i := 0; i < topLevelKeys; i++ {

			if i%100 == 0 {
				fmt.Printf("Key: Processed: %d\n", i)
			}

			b, err := tx.CreateBucketIfNotExists([]byte("x"))
			if err != nil {
				return err
			}

			// 8 byte long major key (represents a "word")
			majorKey := []byte(fmt.Sprintf("%8d", i))

			for j := 0; j < bottomLevelKeys; j++ {

				key := append(majorKey, append([]byte("|"), itob(j)...)...)

				err := b.Put(key, itob(j))
				if err != nil {
					return err
				}
			}

		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	err = db.View(func(tx *bbolt.Tx) error {
		return tx.CopyFile(keyDBPath+".copy", 0666)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	size, err := GetFileSize(keyDBPath + ".copy")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%20s: %d KB\n", keyDBPath+".copy", size/1024)
}

func itob(id int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(id))
	return b
}

func GetFileSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}
