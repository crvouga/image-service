package projectDB

import (
	"net/url"
	"testing"
	"time"

	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	ProjectDB  ProjectDB
}

func newFixture() *Fixture {
	db := sqlite.New()

	return &Fixture{
		ProjectDB:  NewImplKeyValueDB(keyValueDB.NewImplHashMap()),
		UowFactory: *uow.NewFactory(db),
	}
}

func Test_GetByID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a project
	projectNameInst, _ := projectName.New("Test Project")
	proj := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            projectNameInst,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example.com"}},
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the project
	retrieved, err := f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve project, got nil")
	}

	if retrieved.ID != proj.ID {
		t.Errorf("Expected ID to be %s, got %s", proj.ID, retrieved.ID)
	}

	if retrieved.Name != proj.Name {
		t.Errorf("Expected Name to be %s, got %s", proj.Name, retrieved.Name)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a project that doesn't exist
	nonexistentProjectID, _ := projectID.New("nonexistent")

	retrieved, err := f.ProjectDB.GetByID(nonexistentProjectID)

	if err != nil {
		t.Errorf("Expected no error for nonexistent project, got %v", err)
	}

	if retrieved != nil {
		t.Errorf("Expected nil for nonexistent project, got %+v", retrieved)
	}
}

func Test_UpsertNewProject(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a project
	projectNameInst, _ := projectName.New("Test Project")
	proj := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            projectNameInst,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example.com"}},
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve project, got nil")
	}

	if retrieved.ID != proj.ID {
		t.Errorf("Expected ID to be %s, got %s", proj.ID, retrieved.ID)
	}

	uow.Commit()
}

func Test_UpsertUpdateProject(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial project
	projectNameInst, _ := projectName.New("Test Project")
	proj := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            projectNameInst,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example.com"}},
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the project
	updatedNameInst, _ := projectName.New("Updated Project")
	updatedProject := &project.Project{
		ID:              proj.ID,
		CreatedByUserID: proj.CreatedByUserID,
		Name:            updatedNameInst,
		CreatedAt:       proj.CreatedAt,
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example.com"}, {Host: "test.com"}},
	}

	// Update the project
	err = f.ProjectDB.Upsert(uow, updatedProject)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve project, got nil")
	}

	if retrieved.Name != updatedNameInst {
		t.Errorf("Expected Name to be %s, got %s", updatedNameInst, retrieved.Name)
	}

	if len(retrieved.AllowedDomains) != 2 {
		t.Errorf("Expected 2 allowed domains, got %d", len(retrieved.AllowedDomains))
	}

	uow.Commit()
}

func Test_GetByCreatedByUserID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	userIDInst := userID.Gen()

	// Create multiple projects for the same user
	projectName1, _ := projectName.New("Project 1")
	proj1 := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userIDInst,
		Name:            projectName1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example1.com"}},
	}

	projectName2, _ := projectName.New("Project 2")
	proj2 := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userIDInst,
		Name:            projectName2,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example2.com"}},
	}

	// Insert the projects
	err := f.ProjectDB.Upsert(uow, proj1)
	if err != nil {
		t.Errorf("Expected no error on insert proj1, got %v", err)
	}

	err = f.ProjectDB.Upsert(uow, proj2)
	if err != nil {
		t.Errorf("Expected no error on insert proj2, got %v", err)
	}

	uow.Commit()

	// Get projects by user ID
	projects, err := f.ProjectDB.GetByCreatedByUserID(userIDInst)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if len(projects) != 2 {
		t.Errorf("Expected 2 projects, got %d", len(projects))
	}
}

func Test_ZapByID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a project
	projectNameInst, _ := projectName.New("Test Project")
	proj := &project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            projectNameInst,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  []url.URL{{Host: "example.com"}},
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Commit the project
	uow.Commit()

	// Start a new UOW for the zap operation
	uow, _ = f.UowFactory.Begin()

	// Zap the project
	err = f.ProjectDB.ZapByID(uow, proj.ID)
	if err != nil {
		t.Errorf("Expected no error on zap, got %v", err)
	}

	// Commit the zap
	uow.Commit()

	// Verify it was removed
	retrieved, err := f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval after zap, got %v", err)
	}

	if retrieved != nil {
		t.Errorf("Expected project to be removed, but it was retrieved: %+v", retrieved)
	}
}
