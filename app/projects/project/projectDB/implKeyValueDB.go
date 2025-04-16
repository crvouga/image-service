package projectDB

import (
	"encoding/json"
	"fmt"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDB struct {
	entities keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		entities: keyValueDB.NewImplNamespaced(db, "project"),
	}
}

func projectKey(id projectID.ProjectID) string {
	return fmt.Sprintf("%s", id)
}

func (db ImplKeyValueDB) GetByID(id projectID.ProjectID) (*project.Project, error) {
	value, err := db.entities.Get(projectKey(id))
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	var proj project.Project
	if err := json.Unmarshal([]byte(*value), &proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

func (db ImplKeyValueDB) Upsert(uow *uow.Uow, proj project.Project) error {
	proj.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(proj)
	if err != nil {
		return err
	}

	return db.entities.Put(uow, projectKey(proj.ID), string(jsonData))
}

var _ ProjectDB = (*ImplKeyValueDB)(nil)
