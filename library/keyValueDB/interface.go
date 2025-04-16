package keyValueDB

import "imageresizerservice/library/uow"

type KeyValueDB interface {
	// Get retrieves a value by key. Returns nil if key not found.
	Get(key string) (*string, error)
	// Put stores a key-value pair. Returns nil if key not found.
	Put(uow *uow.Uow, key string, value string) error
	// Zap removes a key-value pair. Returns nil if key not found.
	Zap(uow *uow.Uow, key string) error
}
