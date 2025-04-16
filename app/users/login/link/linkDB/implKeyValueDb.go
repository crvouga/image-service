package linkDB

import (
	"encoding/json"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/sessionID"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDB struct {
	links                   keyValueDB.KeyValueDB
	indexLinkIDsBySessionID keyValueDB.KeyValueDB
}

var _ LinkDB = ImplKeyValueDB{}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		links:                   keyValueDB.NewImplNamespaced(db, "link"),
		indexLinkIDsBySessionID: keyValueDB.NewImplNamespaced(db, "link:index:linkIDsBySessionID"),
	}
}

func (db ImplKeyValueDB) GetByLinkID(id linkID.LinkID) (*link.Link, error) {
	value, err := db.links.Get(string(id))
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

	if err := db.links.Put(uow, string(l.ID), string(jsonData)); err != nil {
		return err
	}

	// Update the session ID index
	linkIDs, err := db.getLinkIDsBySessionID(l.SessionID)
	if err != nil {
		return err
	}

	// Check if link ID already exists in the index
	linkExists := false
	for _, existingLinkID := range linkIDs {
		if existingLinkID == l.ID {
			linkExists = true
			break
		}
	}

	// If link ID doesn't exist in the index, add it
	if !linkExists {
		linkIDs = append(linkIDs, l.ID)
	}

	// Save the updated index
	indexData, err := json.Marshal(linkIDs)
	if err != nil {
		return err
	}

	return db.indexLinkIDsBySessionID.Put(uow, string(l.SessionID), string(indexData))
}

func (db ImplKeyValueDB) getLinkIDsBySessionID(sessionID sessionID.SessionID) ([]linkID.LinkID, error) {
	value, err := db.indexLinkIDsBySessionID.Get(string(sessionID))
	if err != nil {
		return nil, err
	}

	if value == nil {
		return []linkID.LinkID{}, nil
	}

	var linkIDs []linkID.LinkID
	if err := json.Unmarshal([]byte(*value), &linkIDs); err != nil {
		return nil, err
	}

	return linkIDs, nil
}

func (db ImplKeyValueDB) GetBySessionID(sessionID sessionID.SessionID) ([]*link.Link, error) {
	linkIDs, err := db.getLinkIDsBySessionID(sessionID)
	if err != nil {
		return nil, err
	}

	links := make([]*link.Link, 0, len(linkIDs))
	for _, id := range linkIDs {
		link, err := db.GetByLinkID(id)
		if err != nil {
			return nil, err
		}
		if link != nil {
			links = append(links, link)
		}
	}

	return links, nil
}

var _ LinkDB = (*ImplKeyValueDB)(nil)
