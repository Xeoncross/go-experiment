package bboltrace

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"go.etcd.io/bbolt"
)

func TestUpdateRace(t *testing.T) {

	var err error

	err = cleanup()
	if err != nil {
		t.Log(err)
	}

	var db *bbolt.DB
	db, err = open()
	if err != nil {
		t.Fatal(err)
	}

	bucket := []byte("sample")
	key := []byte("item1")

	// Pre-create the bucket we will be using
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}

	wg := sync.WaitGroup{}

	// Start both about the same time
	startB := make(chan bool)

	wg.Add(1)
	// Process "A"
	go func() {
		var value []byte
		err := db.Update(func(tx *bbolt.Tx) error {

			// Step 1: "A" tries to read "key", does not exist
			b := tx.Bucket(bucket)
			// b.Get(key) // fails

			// Let "B" try to create this now before we get to it
			fmt.Println("A starting B")
			startB <- true
			fmt.Println("A sleeping...")
			time.Sleep(time.Millisecond * 500)

			fmt.Println("A .Put()")

			// Step 3: "A" creates "key"
			err := b.Put(key, []byte("A"))
			if err != nil {
				return fmt.Errorf("A: %w", err)
			}

			fmt.Printf("A set %q\n", key)
			value = []byte("A")
			return nil
		})
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("A got %q\n", value)
		wg.Done()
	}()

	// Process "B"
	var value []byte
	<-startB
	fmt.Println("B starting")
	err = db.Update(func(tx *bbolt.Tx) error {
		fmt.Println("B started")
		b := tx.Bucket(bucket)

		fmt.Println("B bucket loaded")

		// "B" will aways get the results of A's .Update() tx
		v := b.Get(key)
		if v != nil {
			return fmt.Errorf("B got unexpected existing value: %q", v)
		}

		fmt.Println("B .Put()")

		// Step 2: "B" creates "key"
		err := b.Put(key, []byte("B"))
		if err != nil {
			return fmt.Errorf("B: %w", err)
		}

		fmt.Printf("B set %q\n", key)
		value = []byte("B")

		return nil
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("B got %q\n", value)

	wg.Wait()
	db.Close()

}
