package keyValueDb

import (
	"errors"
	"imageresizerservice/uow"
	"sync"
)

// ImplHashMap implements the KeyValueDb interface using a Go map
type ImplHashMap struct {
	data  map[string]string
	mutex sync.RWMutex
}

// NewImplHashMap creates a new instance of ImplHashMap
func NewImplHashMap() *ImplHashMap {
	return &ImplHashMap{
		data: make(map[string]string),
	}
}

// Get retrieves a value by key
func (db *ImplHashMap) Get(key string) (string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.data[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

// Put stores a key-value pair
func (db *ImplHashMap) Put(uow *uow.Uow, key string, value string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.data[key] = value
	return nil
}

// Zap removes a key-value pair
func (db *ImplHashMap) Zap(uow *uow.Uow, key string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, exists := db.data[key]; !exists {
		return errors.New("key not found")
	}

	delete(db.data, key)
	return nil
}

// Ensure ImplHashMap implements KeyValueDb interface
var _ KeyValueDb = (*ImplHashMap)(nil)
