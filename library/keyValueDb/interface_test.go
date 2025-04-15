package keyValueDb

import (
	"testing"

	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	KeyValueDb KeyValueDb
}

func newFixture() *Fixture {

	db := sqlite.New()

	return &Fixture{
		KeyValueDb: &ImplHashMap{},
		UowFactory: uow.UowFactory{Db: db},
	}

}

func Test_Interface(t *testing.T) {

	f := newFixture()

	uow, err := f.UowFactory.Begin()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	f.KeyValueDb.Put(uow, "key", "value")

	value, err := f.KeyValueDb.Get("key")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value == nil {
		t.Errorf("Expected value to be 'value', got nil")
	}

	uow.Commit()
}

func Test_UpdateValue(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Put initial value
	f.KeyValueDb.Put(uow, "key", "initial")

	// Update the value
	f.KeyValueDb.Put(uow, "key", "updated")

	// Verify the value was updated
	value, err := f.KeyValueDb.Get("key")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if value == nil {
		t.Errorf("Expected value to be 'updated', got nil")
	}

	uow.Commit()
}

func Test_GetNonExistentKey(t *testing.T) {
	f := newFixture()

	// Try to get a key that doesn't exist
	value, err := f.KeyValueDb.Get("nonexistent")

	if err != nil {
		t.Errorf("Expected no error for nonexistent key, got %v", err)
	}

	if value != nil {
		t.Errorf("Expected value to be nil for nonexistent key, got %v", value)
	}
}

func Test_ZapKey(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Put a value
	f.KeyValueDb.Put(uow, "key-to-zap", "value")
	uow.Commit()

	// Verify it exists
	value, err := f.KeyValueDb.Get("key-to-zap")
	if err != nil {
		t.Errorf("Expected key to exist, got error: %v", err)
	}

	if value == nil {
		t.Error("Expected key to exist, got nil")
	}

	// Zap the key
	uow, _ = f.UowFactory.Begin()
	err = f.KeyValueDb.Zap(uow, "key-to-zap")
	if err != nil {
		t.Errorf("Expected no error when zapping key, got %v", err)
	}
	uow.Commit()

	// Verify it no longer exists
	_, err = f.KeyValueDb.Get("key-to-zap")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}

func Test_ZapNonExistentKey(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Try to zap a key that doesn't exist
	err := f.KeyValueDb.Zap(uow, "nonexistent")

	// Should not error when zapping a non-existent key
	if err != nil {
		t.Errorf("Expected no error when zapping nonexistent key, got %v", err)
	}

	uow.Commit()
}
