package userAccountDB

import (
	"encoding/json"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDB struct {
	userAccounts              keyValueDB.KeyValueDB
	indexUserIDByEmailAddress keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		userAccounts:              keyValueDB.NewImplNamespaced(db, "userAccount"),
		indexUserIDByEmailAddress: keyValueDB.NewImplNamespaced(db, "userAccount:index:userIDByEmailAddress"),
	}
}

func emailIndexKey(emailAddress emailAddress.EmailAddress) string {
	return string(emailAddress)
}

func userAccountKey(id userID.UserID) string {
	return string(id)
}

func (db ImplKeyValueDB) GetByUserID(id userID.UserID) (*userAccount.UserAccount, error) {
	value, err := db.userAccounts.Get(userAccountKey(id))
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

func (db ImplKeyValueDB) Upsert(uow *uow.Uow, account userAccount.UserAccount) error {
	account.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// Store the user account by ID
	if err := db.userAccounts.Put(uow, userAccountKey(account.UserID), string(jsonData)); err != nil {
		return err
	}

	// Create an index entry for email address -> user ID
	return db.indexUserIDByEmailAddress.Put(uow, emailIndexKey(account.EmailAddress), string(account.UserID))
}

func (db ImplKeyValueDB) GetByEmailAddress(emailAddress emailAddress.EmailAddress) (*userAccount.UserAccount, error) {
	// Get the user ID from the email index
	gotUserID, err := db.indexUserIDByEmailAddress.Get(emailIndexKey(emailAddress))
	if err != nil {
		return nil, err
	}

	if gotUserID == nil {
		return nil, nil
	}

	// Use the user ID to get the actual user account
	return db.GetByUserID(userID.UserID(*gotUserID))
}

var _ UserAccountDB = (*ImplKeyValueDB)(nil)
