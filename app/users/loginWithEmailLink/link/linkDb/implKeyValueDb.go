package linkDb

import (
	"encoding/json"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/app/users/loginWithEmailLink/link/linkID"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDb struct {
	Db keyValueDb.KeyValueDb
}

var _ LinkDb = ImplKeyValueDb{}

func (db ImplKeyValueDb) GetByLinkID(id linkID.LinkID) (*link.Link, error) {
	value, err := db.Db.Get(string(id))
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

	return db.Db.Put(uow, string(l.ID), string(jsonData))
}

var _ LinkDb = (*ImplKeyValueDb)(nil)
