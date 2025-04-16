package projectDB

import (
	"encoding/json"
	"fmt"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDB struct {
	entities                 keyValueDB.KeyValueDB
	indexManyCreatedByUserID keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		entities:                 keyValueDB.NewImplNamespaced(db, "project"),
		indexManyCreatedByUserID: keyValueDB.NewImplNamespaced(db, "project:indexByCreatedByUserID"),
	}
}

func projectKey(id projectID.ProjectID) string {
	return fmt.Sprintf("%s", id)
}

func (db *ImplKeyValueDB) GetByID(id projectID.ProjectID) (*project.Project, error) {
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

func (db *ImplKeyValueDB) Upsert(uow *uow.Uow, proj project.Project) error {
	if uow == nil {
		return fmt.Errorf("unit of work cannot be nil")
	}

	proj.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(proj)
	if err != nil {
		return fmt.Errorf("failed to marshal project: %w", err)
	}

	key := projectKey(proj.ID)
	if key == "" {
		return fmt.Errorf("invalid project ID")
	}

	// Check if project already exists to handle updating indexes
	existingProj, err := db.GetByID(proj.ID)
	if err != nil {
		return fmt.Errorf("failed to check existing project: %w", err)
	}

	// Store the project in the main entities store
	if err := db.entities.Put(uow, key, string(jsonData)); err != nil {
		return fmt.Errorf("failed to store project: %w", err)

	}

	// If project exists and CreatedByUserID changed, remove from old index
	if existingProj != nil && existingProj.CreatedByUserID != proj.CreatedByUserID {
		oldIndexKey := fmt.Sprintf("%s", existingProj.CreatedByUserID)
		oldProjectIDs, err := db.getProjectIDsByUserID(existingProj.CreatedByUserID)
		if err != nil {
			return fmt.Errorf("failed to get old index: %w", err)
		}

		// Remove project ID from old user's list
		newOldProjectIDs := make([]projectID.ProjectID, 0, len(oldProjectIDs))
		for _, id := range oldProjectIDs {
			if id != proj.ID {
				newOldProjectIDs = append(newOldProjectIDs, id)
			}
		}

		// Update old index
		oldIDsJSON, err := json.Marshal(newOldProjectIDs)
		if err != nil {
			return fmt.Errorf("failed to marshal old IDs: %w", err)
		}
		if err := db.indexManyCreatedByUserID.Put(uow, oldIndexKey, string(oldIDsJSON)); err != nil {
			return fmt.Errorf("failed to update old user index: %w", err)
		}
	}

	// Update the index for current CreatedByUserID
	indexKey := fmt.Sprintf("%s", proj.CreatedByUserID)
	projectIDs, err := db.getProjectIDsByUserID(proj.CreatedByUserID)
	if err != nil {
		return fmt.Errorf("failed to get current index: %w", err)
	}

	// Check if project ID already exists in the list
	found := false
	for _, id := range projectIDs {
		if id == proj.ID {
			found = true
			break
		}
	}

	// Add project ID if not already in the list
	if !found {
		projectIDs = append(projectIDs, proj.ID)
	}

	// Update index with new list of IDs
	idsJSON, err := json.Marshal(projectIDs)
	if err != nil {
		return fmt.Errorf("failed to marshal IDs: %w", err)
	}
	if err := db.indexManyCreatedByUserID.Put(uow, indexKey, string(idsJSON)); err != nil {
		return fmt.Errorf("failed to update user index: %w", err)
	}

	return nil
}

func (db *ImplKeyValueDB) ZapByID(uow *uow.Uow, id projectID.ProjectID) error {
	if uow == nil {
		return fmt.Errorf("unit of work cannot be nil")
	}

	key := projectKey(id)
	if key == "" {
		return fmt.Errorf("invalid project ID")
	}

	// Get the project first to find its CreatedByUserID for index cleanup
	proj, err := db.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to get project for index cleanup: %w", err)
	}

	if proj != nil {
		// Remove from the index
		indexKey := fmt.Sprintf("%s:%s", proj.CreatedByUserID, proj.ID)
		if err := db.indexManyCreatedByUserID.Zap(uow, indexKey); err != nil {
			return fmt.Errorf("failed to remove from user index: %w", err)
		}
	}

	// Remove from main entities store
	return db.entities.Zap(uow, key)
}

func (db *ImplKeyValueDB) GetByCreatedByUserID(createdByUserID userID.UserID) ([]*project.Project, error) {
	indexKey := fmt.Sprintf("%s", createdByUserID)
	value, err := db.indexManyCreatedByUserID.Get(indexKey)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return []*project.Project{}, nil
	}

	var projectIDs []projectID.ProjectID
	if err := json.Unmarshal([]byte(*value), &projectIDs); err != nil {
		return nil, err
	}

	var projects []*project.Project
	for _, id := range projectIDs {
		proj, err := db.GetByID(id)
		if err != nil {
			return nil, err
		}
		if proj != nil {
			projects = append(projects, proj)
		}
	}

	return projects, nil
}

func (db *ImplKeyValueDB) getProjectIDsByUserID(userID userID.UserID) ([]projectID.ProjectID, error) {
	indexKey := fmt.Sprintf("%s", userID)
	value, err := db.indexManyCreatedByUserID.Get(indexKey)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return []projectID.ProjectID{}, nil
	}

	var projectIDs []projectID.ProjectID
	if err := json.Unmarshal([]byte(*value), &projectIDs); err != nil {
		return nil, err
	}

	return projectIDs, nil
}

var _ ProjectDB = (*ImplKeyValueDB)(nil)
