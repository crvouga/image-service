package userAccountDb

import (
	"encoding/json"
	"fmt"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDb struct {
	db         keyValueDb.KeyValueDb
	emailIndex keyValueDb.KeyValueDb
}

func NewImplKeyValueDb(db keyValueDb.KeyValueDb) *ImplKeyValueDb {
	return &ImplKeyValueDb{
		db:         keyValueDb.NewImplNamespaced(db, "userAccount"),
		emailIndex: keyValueDb.NewImplNamespaced(db, "userAccount:email"),
	}
}

func emailIndexKey(emailAddress emailAddress.EmailAddress) string {
	return fmt.Sprintf("%s", emailAddress)
}

func userAccountKey(id userID.UserID) string {
	return fmt.Sprintf("%s", id)
}

func (db ImplKeyValueDb) GetByUserID(id userID.UserID) (*userAccount.UserAccount, error) {
	value, err := db.db.Get(userAccountKey(id))
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	var account userAccount.UserAccount
	if err := json.Unmarshal([]byte(*value), &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (db ImplKeyValueDb) Upsert(uow *uow.Uow, account userAccount.UserAccount) error {
	account.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// Store the user account by ID
	if err := db.db.Put(uow, userAccountKey(account.ID), string(jsonData)); err != nil {
		return err
	}

	// Create an index entry for email address -> user ID
	return db.emailIndex.Put(uow, emailIndexKey(account.EmailAddress), string(account.ID))
}

func (db ImplKeyValueDb) GetByEmailAddress(emailAddress emailAddress.EmailAddress) (*userAccount.UserAccount, error) {
	// Get the user ID from the email index
	gotUserID, err := db.emailIndex.Get(emailIndexKey(emailAddress))
	if err != nil {
		return nil, err
	}

	if gotUserID == nil {
		return nil, nil
	}

	// Use the user ID to get the actual user account
	return db.GetByUserID(userID.UserID(*gotUserID))
}

var _ UserAccountDb = (*ImplKeyValueDb)(nil)
