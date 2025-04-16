package keyValueDb

import (
	"encoding/json"
	"imageresizerservice/library/uow"
	"os"
	"path/filepath"
	"sync"
)

// ImplFs implements the KeyValueDb interface using a single JSON file
type ImplFs struct {
	filePath string
	data     map[string]string
	mutex    sync.RWMutex
}

// NewImplFs creates a new instance of ImplFs
func NewImplFs(dirPath string, fileName string) (*ImplFs, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(dirPath, fileName)

	db := &ImplFs{
		filePath: filePath,
		data:     make(map[string]string),
	}

	// Load existing data if file exists
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		if len(data) > 0 {
			if err := json.Unmarshal(data, &db.data); err != nil {
				return nil, err
			}
		}
	}

	return db, nil
}

// Get retrieves a value by key. Returns nil if key not found.
func (db *ImplFs) Get(key string) (*string, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.data[key]
	if !exists {
		return nil, nil
	}

	return &value, nil
}

// Put stores a key-value pair
func (db *ImplFs) Put(uow *uow.Uow, key string, value string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	// Initialize the map if it's nil
	if db.data == nil {
		db.data = make(map[string]string)
	}

	db.data[key] = value

	// Write the entire map to the JSON file
	data, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(db.filePath, data, 0644)
}

// Zap removes a key-value pair
func (db *ImplFs) Zap(uow *uow.Uow, key string) error {
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

	// Write the updated map to the JSON file
	data, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(db.filePath, data, 0644)
}

// Ensure ImplFs implements KeyValueDb interface
var _ KeyValueDb = (*ImplFs)(nil)
