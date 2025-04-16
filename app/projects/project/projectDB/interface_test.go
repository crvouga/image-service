package projectDB

import (
	"testing"
	"time"

	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
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
	proj := project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            "Test Project",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
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

	if retrieved.CreatedByUserID != proj.CreatedByUserID {
		t.Errorf("Expected UserID to be %s, got %s", proj.CreatedByUserID, retrieved.CreatedByUserID)
	}

	if retrieved.Name != proj.Name {
		t.Errorf("Expected Name to be %s, got %s", proj.Name, retrieved.Name)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a project that doesn't exist
	nonexistentProjectID := projectID.Gen()

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
	proj := project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            "Test Project",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
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
	proj := project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            "Test Project",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the project
	updatedProj := project.Project{
		ID:              proj.ID,
		CreatedByUserID: proj.CreatedByUserID,
		Name:            "Updated Project Name",
		CreatedAt:       proj.CreatedAt,
		UpdatedAt:       time.Now(),
	}

	// Update the project
	err = f.ProjectDB.Upsert(uow, updatedProj)
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

	if retrieved.Name != "Updated Project Name" {
		t.Errorf("Expected Name to be 'Updated Project Name', got %s", retrieved.Name)
	}

	uow.Commit()
}

func TestZapByID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial project
	proj := project.Project{
		ID:              projectID.Gen(),
		CreatedByUserID: userID.Gen(),
		Name:            "Test Project",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Insert the project
	err := f.ProjectDB.Upsert(uow, proj)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it was inserted
	retrieved, err := f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}
	if retrieved == nil {
		t.Fatal("Expected to retrieve project, got nil")
	}

	// Zap the project
	err = f.ProjectDB.ZapByID(uow, proj.ID)
	if err != nil {
		t.Errorf("Expected no error on zap, got %v", err)
	}

	// Verify it was zapped
	retrieved, err = f.ProjectDB.GetByID(proj.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval after zap, got %v", err)
	}
	if retrieved != nil {
		t.Errorf("Expected project to be zapped (nil), got %+v", retrieved)
	}

	uow.Commit()
}
