package storage

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

// DB is the interface to a simple key/value store
type DB interface {
	// Get returns the value for the given key, ErrNotFound if the key doesn't exist,
	// or another error if the get failed
	Get(key string) ([]byte, error)
	// Set sets the value for the given key. Returns an error if the set failed.
	// If non-nil error is returned, the value was not updated
	Set(key string, val []byte) error
	List() (map[string][]byte, error)
}

type inMemoryDB struct {
	m   map[string][]byte
	lck sync.RWMutex
}

// NewInMemoryDB creates a new DB implementation that stores all data in memory.
// All operations are concurrency safe
func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

// Get is the interface implementation
func (d *inMemoryDB) Get(key string) ([]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

// Set is the interface implementation
func (d *inMemoryDB) Set(key string, val []byte) error {
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil
}

func (d *inMemoryDB) List() (map[string][]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	//returns 10 random elements in the map
	result := make(map[string][]byte)
	var limit int
	for key, coupon := range d.m {
		if limit < 10 {
			result[key] = coupon
			limit++
		}
		break
	}
	return result, nil
}
