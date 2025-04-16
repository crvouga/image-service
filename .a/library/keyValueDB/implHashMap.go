package keyValueDB

import (
	"imageresizerservice/library/uow"
	"sync"
)

// ImplHashMap implements the KeyValueDB interface using a Go map
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

// Get retrieves a value by key. Returns nil if key not found.
func (db *ImplHashMap) Get(key string) (*string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.data[key]
	if !exists {
		return nil, nil
	}
	return &value, nil
}

// Put stores a key-value pair
func (db *ImplHashMap) Put(uow *uow.Uow, key string, value string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Initialize the map if it's nil
	if db.data == nil {
		db.data = make(map[string]string)
	}

	db.data[key] = value
	return nil
}

// Zap removes a key-value pair
func (db *ImplHashMap) Zap(uow *uow.Uow, key string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if db.data == nil {
		return nil
	}

	_, exists := db.data[key]

	if !exists {
		return nil
	}

	delete(db.data, key)
	return nil
}

// Ensure ImplHashMap implements KeyValueDB interface
var _ KeyValueDB = (*ImplHashMap)(nil)
