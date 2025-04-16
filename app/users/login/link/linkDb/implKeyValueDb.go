package linkDB

import (
	"encoding/json"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDB struct {
	db                 keyValueDB.KeyValueDB
	indexManySessionID keyValueDB.KeyValueDB
}

var _ LinkDB = ImplKeyValueDB{}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		db:                 keyValueDB.NewImplNamespaced(db, "link"),
		indexManySessionID: keyValueDB.NewImplNamespaced(db, "link:indexManySessionID"),
	}
}

func (db ImplKeyValueDB) GetByLinkID(id linkID.LinkID) (*link.Link, error) {
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

func (db ImplKeyValueDB) Upsert(uow *uow.Uow, l link.Link) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return err
	}

	if err := db.db.Put(uow, string(l.ID), string(jsonData)); err != nil {
		return err
	}

	// Update the session ID index
	links, err := db.GetBySessionID(l.SessionID)
	if err != nil {
		return err
	}

	// Check if link already exists in the index
	linkExists := false
	for i, existingLink := range links {
		if existingLink.ID == l.ID {
			// Update the existing link
			links[i] = &l
			linkExists = true
			break
		}
	}

	// If link doesn't exist in the index, add it
	if !linkExists {
		links = append(links, &l)
	}

	// Save the updated index
	indexData, err := json.Marshal(links)
	if err != nil {
		return err
	}

	return db.indexManySessionID.Put(uow, string(l.SessionID), string(indexData))
}

func (db ImplKeyValueDB) GetBySessionID(sessionID sessionID.SessionID) ([]*link.Link, error) {
	value, err := db.indexManySessionID.Get(string(sessionID))
	if err != nil {
		return nil, err
	}

	if value == nil {
		return []*link.Link{}, nil
	}

	var links []*link.Link
	if err := json.Unmarshal([]byte(*value), &links); err != nil {
		return nil, err
	}

	return links, nil
}

var _ LinkDB = (*ImplKeyValueDB)(nil)
