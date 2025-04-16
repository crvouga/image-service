package linkDB

import (
	"testing"

	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/login/link"
	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
)

type Fixture struct {
	UowFactory uow.UowFactory
	LinkDB     LinkDB
}

func newFixture() *Fixture {
	db := sqlite.New()

	return &Fixture{
		LinkDB:     NewImplKeyValueDB(keyValueDB.NewImplHashMap()),
		UowFactory: *uow.NewFactory(db),
	}
}

func Test_GetByLinkID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a link
	email, _ := emailAddress.New("test@example.com")
	testLink := link.New(email, sessionID.Gen())

	// Insert the link
	err := f.LinkDB.Upsert(uow, testLink)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Get the link
	retrieved, err := f.LinkDB.GetByLinkID(testLink.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve link, got nil")
	}

	if retrieved.ID != testLink.ID {
		t.Errorf("Expected ID to be %s, got %s", testLink.ID, retrieved.ID)
	}

	if retrieved.EmailAddress != testLink.EmailAddress {
		t.Errorf("Expected EmailAddress to be %s, got %s", testLink.EmailAddress, retrieved.EmailAddress)
	}

	uow.Commit()
}

func Test_GetByIDNonExistent(t *testing.T) {
	f := newFixture()

	// Try to get a link that doesn't exist
	nonexistentLinkID := linkID.Gen()

	retrieved, err := f.LinkDB.GetByLinkID(nonexistentLinkID)

	if err != nil {
		t.Errorf("Expected no error for nonexistent link, got %v", err)
	}

	if retrieved != nil {
		t.Errorf("Expected nil for nonexistent link, got %+v", retrieved)
	}
}

func Test_UpsertNewLink(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a link
	email, _ := emailAddress.New("test@example.com")
	testLink := link.New(email, sessionID.Gen())

	// Insert the link
	err := f.LinkDB.Upsert(uow, testLink)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Verify it exists
	retrieved, err := f.LinkDB.GetByLinkID(testLink.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve link, got nil")
	}

	if retrieved.ID != testLink.ID {
		t.Errorf("Expected ID to be %s, got %s", testLink.ID, retrieved.ID)
	}

	uow.Commit()
}

func Test_UpsertUpdateLink(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create initial link
	email, _ := emailAddress.New("test@example.com")
	testLink := link.New(email, sessionID.Gen())

	// Insert the link
	err := f.LinkDB.Upsert(uow, testLink)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Update the link by marking it as used
	updatedLink := link.MarkAsUsed(testLink)

	// Update the link
	err = f.LinkDB.Upsert(uow, updatedLink)
	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	// Verify it was updated
	retrieved, err := f.LinkDB.GetByLinkID(testLink.ID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected to retrieve link, got nil")
	}

	if retrieved.UsedAt.IsZero() {
		t.Error("Expected UsedAt to be set, but it's zero")
	}

	uow.Commit()
}

func Test_GetBySessionID(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a link with a known session ID
	email, _ := emailAddress.New("test@example.com")
	sessionID := sessionID.Gen()
	testLink := link.New(email, sessionID)

	// Insert the link
	err := f.LinkDB.Upsert(uow, testLink)
	if err != nil {
		t.Errorf("Expected no error on insert, got %v", err)
	}

	// Retrieve the link by session ID
	links, err := f.LinkDB.GetBySessionID(sessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	if links == nil || len(links) == 0 {
		t.Fatal("Expected to retrieve links, got none")
	}

	retrieved := links[0]
	if retrieved.ID != testLink.ID {
		t.Errorf("Expected ID to be %s, got %s", testLink.ID, retrieved.ID)
	}

	if retrieved.SessionID != sessionID {
		t.Errorf("Expected SessionID to be %s, got %s", sessionID, retrieved.SessionID)
	}

	uow.Commit()
}

func Test_GetBySessionIDMany(t *testing.T) {
	f := newFixture()
	uow, _ := f.UowFactory.Begin()

	// Create a shared session ID for multiple links
	sharedSessionID := sessionID.Gen()

	// Create multiple links with the same session ID
	email1, _ := emailAddress.New("test1@example.com")
	testLink1 := link.New(email1, sharedSessionID)

	email2, _ := emailAddress.New("test2@example.com")
	testLink2 := link.New(email2, sharedSessionID)

	email3, _ := emailAddress.New("test3@example.com")
	testLink3 := link.New(email3, sharedSessionID)

	// Insert all links
	if err := f.LinkDB.Upsert(uow, testLink1); err != nil {
		t.Errorf("Expected no error on insert for link1, got %v", err)
	}
	if err := f.LinkDB.Upsert(uow, testLink2); err != nil {
		t.Errorf("Expected no error on insert for link2, got %v", err)
	}
	if err := f.LinkDB.Upsert(uow, testLink3); err != nil {
		t.Errorf("Expected no error on insert for link3, got %v", err)
	}

	// Retrieve the links by session ID
	links, err := f.LinkDB.GetBySessionID(sharedSessionID)
	if err != nil {
		t.Errorf("Expected no error on retrieval, got %v", err)
	}

	// Verify we got all three links
	if links == nil {
		t.Fatal("Expected to retrieve links, got nil")
	}

	if len(links) != 3 {
		t.Fatalf("Expected to retrieve 3 links, got %d", len(links))
	}

	// Create a map of link IDs for easier verification
	linkMap := make(map[linkID.LinkID]bool)
	for _, link := range links {
		linkMap[link.ID] = true

		// Verify each link has the correct session ID
		if link.SessionID != sharedSessionID {
			t.Errorf("Expected SessionID to be %s, got %s", sharedSessionID, link.SessionID)
		}
	}

	// Verify all links were retrieved
	if !linkMap[testLink1.ID] {
		t.Errorf("Expected to find link1 with ID %s", testLink1.ID)
	}
	if !linkMap[testLink2.ID] {
		t.Errorf("Expected to find link2 with ID %s", testLink2.ID)
	}
	if !linkMap[testLink3.ID] {
		t.Errorf("Expected to find link3 with ID %s", testLink3.ID)
	}

	uow.Commit()
}
