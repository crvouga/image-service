package keyValueDB

import (
	"testing"

	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	KeyValueDB KeyValueDB
}

func newFixtures() []*Fixture {
	db := sqlite.New()

	fixtures := make([]*Fixture, 0)

	keyValueDBs := []KeyValueDB{
		NewImplHashMap(),
		NewImplFs("test/keyValueDB.json"),
		NewImplNamespaced(NewImplFs("test/keyValueDB.json"), "test"),
	}

	for _, keyValueDB := range keyValueDBs {
		fixtures = append(fixtures, &Fixture{
			KeyValueDB: keyValueDB,
			UowFactory: *uow.NewFactory(db),
		})
	}

	return fixtures
}

func Test_NewFixtures(t *testing.T) {
	fixtures := newFixtures()

	if len(fixtures) == 0 {
		t.Errorf("Expected newFixtures to return non-empty fixtures, got empty slice")
	}

}

func Test_Interface(t *testing.T) {
	fixtures := newFixtures()

	for _, f := range fixtures {
		uow, err := f.UowFactory.Begin()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		f.KeyValueDB.Put(uow, "key", "value")

		value, err := f.KeyValueDB.Get("key")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if value == nil {
			t.Errorf("Expected value to be 'value', got nil")
		}

		uow.Commit()
	}
}

func Test_UpdateValue(t *testing.T) {
	fixtures := newFixtures()

	for _, f := range fixtures {
		uow, _ := f.UowFactory.Begin()

		// Put initial value
		f.KeyValueDB.Put(uow, "key", "initial")

		// Update the value
		f.KeyValueDB.Put(uow, "key", "updated")

		// Verify the value was updated
		value, err := f.KeyValueDB.Get("key")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if value == nil {
			t.Errorf("Expected value to be 'updated', got nil")
		}

		uow.Commit()
	}
}

func Test_GetNonExistentKey(t *testing.T) {
	fixtures := newFixtures()

	for _, f := range fixtures {
		// Try to get a key that doesn't exist
		value, err := f.KeyValueDB.Get("nonexistent")

		if err != nil {
			t.Errorf("Expected no error for nonexistent key, got %v", err)
		}

		if value != nil {
			t.Errorf("Expected value to be nil for nonexistent key, got %v", value)
		}
	}
}

func Test_ZapKey(t *testing.T) {
	fixtures := newFixtures()

	for _, f := range fixtures {
		uow, _ := f.UowFactory.Begin()

		// Put a value
		f.KeyValueDB.Put(uow, "key-to-zap", "value")
		uow.Commit()

		// Verify it exists
		value, err := f.KeyValueDB.Get("key-to-zap")
		if err != nil {
			t.Errorf("Expected key to exist, got error: %v", err)
		}

		if value == nil {
			t.Error("Expected key to exist, got nil")
		}

		// Zap the key
		uow, _ = f.UowFactory.Begin()
		err = f.KeyValueDB.Zap(uow, "key-to-zap")
		if err != nil {
			t.Errorf("Expected no error when zapping key, got %v", err)
		}
		uow.Commit()

		// Verify it no longer exists
		value, err = f.KeyValueDB.Get("key-to-zap")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if value != nil {
			t.Errorf("Expected value to be nil after zapping, got %v", *value)
		}
	}
}

func Test_ZapNonExistentKey(t *testing.T) {
	fixtures := newFixtures()

	for _, f := range fixtures {
		uow, _ := f.UowFactory.Begin()

		// Try to zap a key that doesn't exist
		err := f.KeyValueDB.Zap(uow, "nonexistent")

		// Should not error when zapping a non-existent key
		if err != nil {
			t.Errorf("Expected no error when zapping nonexistent key, got %v", err)
		}

		uow.Commit()
	}
}
