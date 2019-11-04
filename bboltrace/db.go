package bboltrace

import (
	"os"

	bolt "go.etcd.io/bbolt"
)

var bboltDBPath = "test.db"

// open database handle
func open() (*bolt.DB, error) {
	return bolt.Open(bboltDBPath, 0600, nil)
}

// Remove testing file
func cleanup() error {
	return os.Remove(bboltDBPath)
}
