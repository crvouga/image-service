package projectDB

import (
	"encoding/json"
	"imageService/app/projects/project"
	"imageService/app/projects/project/projectID"
	"imageService/app/users/userID"
	"imageService/library/keyValueDB"
	"imageService/library/uow"
	"strings"
	"time"
)

type ImplKeyValueDB struct {
	projects                         keyValueDB.KeyValueDB
	indexProjectIDsByCreatedByUserID keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		projects:                         keyValueDB.NewImplNamespaced(db, "project"),
		indexProjectIDsByCreatedByUserID: keyValueDB.NewImplNamespaced(db, "project:index:projectIDsByCreatedByUserID"),
	}
}

func projectKey(id projectID.ProjectID) string {
	return string(id)
}

func userIndexKey(userID userID.UserID) string {
	return string(userID)
}

func (db ImplKeyValueDB) GetByID(projectID projectID.ProjectID) (*project.Project, error) {
	value, err := db.projects.Get(projectKey(projectID))
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

func (db ImplKeyValueDB) Upsert(uow *uow.Uow, project *project.Project) error {
	project.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(project)
	if err != nil {
		return err
	}

	// Store the project by ID
	if err := db.projects.Put(uow, projectKey(project.ID), string(jsonData)); err != nil {
		return err
	}

	// Update the index entry for user ID -> project IDs
	// Get existing project IDs for this user
	existingValue, err := db.indexProjectIDsByCreatedByUserID.Get(userIndexKey(project.CreatedByUserID))
	if err != nil {
		return err
	}

	var projectIDs []string
	if existingValue != nil {
		projectIDs = strings.Split(*existingValue, ",")
	}

	// Check if project ID already exists in the list
	projectIDStr := string(project.ID)
	found := false
	for _, id := range projectIDs {
		if id == projectIDStr {
			found = true
			break
		}
	}

	// If not found, add it to the list
	if !found {
		projectIDs = append(projectIDs, projectIDStr)
	}

	// Update the index
	return db.indexProjectIDsByCreatedByUserID.Put(uow, userIndexKey(project.CreatedByUserID), strings.Join(projectIDs, ","))
}

func (db ImplKeyValueDB) GetByCreatedByUserID(createdByUserID userID.UserID) ([]*project.Project, error) {
	// Get the list of project IDs for this user
	projectIDsValue, err := db.indexProjectIDsByCreatedByUserID.Get(userIndexKey(createdByUserID))
	if err != nil {
		return nil, err
	}

	if projectIDsValue == nil {
		return []*project.Project{}, nil
	}

	// Split the comma-separated list of project IDs
	projectIDStrings := strings.Split(*projectIDsValue, ",")

	var projects []*project.Project
	for _, idStr := range projectIDStrings {
		if idStr == "" {
			continue
		}

		projectIDVar, err := projectID.New(idStr)
		if err != nil {
			return nil, err
		}

		project, err := db.GetByID(projectIDVar)
		if err != nil {
			return nil, err
		}

		if project != nil {
			projects = append(projects, project)
		}
	}

	return projects, nil
}

func (db ImplKeyValueDB) ZapByID(uow *uow.Uow, projectID projectID.ProjectID) error {
	// Get the project first to remove from user index
	proj, err := db.GetByID(projectID)
	if err != nil {
		return err
	}

	if proj == nil {
		return nil // Already deleted or doesn't exist
	}

	// Delete from the main storage
	if err := db.projects.Zap(uow, projectKey(projectID)); err != nil {
		return err
	}

	// Remove the project ID from the user's project list
	existingValue, err := db.indexProjectIDsByCreatedByUserID.Get(userIndexKey(proj.CreatedByUserID))
	if err != nil {
		return err
	}

	if existingValue != nil {
		projectIDs := strings.Split(*existingValue, ",")
		projectIDStr := string(projectID)

		// Create a new list without the deleted project ID
		var updatedProjectIDs []string
		for _, id := range projectIDs {
			if id != projectIDStr {
				updatedProjectIDs = append(updatedProjectIDs, id)
			}
		}

		// Update the index with the new list
		return db.indexProjectIDsByCreatedByUserID.Put(uow, userIndexKey(proj.CreatedByUserID), strings.Join(updatedProjectIDs, ","))
	}

	return nil
}

var _ ProjectDB = (*ImplKeyValueDB)(nil)
