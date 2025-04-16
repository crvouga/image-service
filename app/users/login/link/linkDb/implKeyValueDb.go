package linkDb

import (
	"encoding/json"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDb struct {
	db keyValueDb.KeyValueDb
}

var _ LinkDb = ImplKeyValueDb{}

func NewImplKeyValueDb(db keyValueDb.KeyValueDb) *ImplKeyValueDb {
	return &ImplKeyValueDb{
		db: keyValueDb.NewImplNamespaced(db, "link"),
	}
}

func (db ImplKeyValueDb) GetByLinkID(id linkID.LinkID) (*link.Link, error) {
	value, err := db.db.Get(string(id))
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	var l link.Link
	if err := json.Unmarshal([]byte(*value), &l); err != nil {
		return nil, err
	}

	return &l, nil
}

func (db ImplKeyValueDb) Upsert(uow *uow.Uow, l link.Link) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return db.db.Put(uow, string(l.ID), string(jsonData))
}

var _ LinkDb = (*ImplKeyValueDb)(nil)
