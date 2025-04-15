package linkDb

import (
	"encoding/json"
	"imageresizerservice/app/users/loginWithEmailLink/link"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDb struct {
	Db keyValueDb.KeyValueDb
}

var _ LinkDb = ImplKeyValueDb{}

func (db ImplKeyValueDb) GetById(id string) (*link.Link, error) {
	time.Sleep(time.Second)

	value, err := db.Db.Get(id)
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	var l link.Link
	if err := json.Unmarshal([]byte(value), &l); err != nil {
		return nil, err
	}

	return &l, nil
}

func (db ImplKeyValueDb) Upsert(uow *uow.Uow, l link.Link) error {
	time.Sleep(time.Second)

	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return db.Db.Put(uow, l.Id, string(jsonData))
}

var _ LinkDb = (*ImplKeyValueDb)(nil)
