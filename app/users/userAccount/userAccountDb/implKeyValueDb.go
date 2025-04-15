package userAccountDb

import (
	"encoding/json"
	"fmt"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
	"time"
)

type ImplKeyValueDb struct {
	Db keyValueDb.KeyValueDb
}

func emailIndexKey(emailAddress string) string {
	return fmt.Sprintf("userAccount:email:%s", emailAddress)
}

func userAccountKey(id string) string {
	return fmt.Sprintf("userAccount:%s", id)
}

func (db ImplKeyValueDb) GetById(id string) (*userAccount.UserAccount, error) {
	time.Sleep(time.Second)

	value, err := db.Db.Get(userAccountKey(id))
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
	time.Sleep(time.Second)

	account.UpdatedAt = time.Now()

	jsonData, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// Store the user account by ID
	if err := db.Db.Put(uow, userAccountKey(account.ID), string(jsonData)); err != nil {
		return err
	}

	// Create an index entry for email address -> user ID
	emailKey := emailIndexKey(account.EmailAddress)
	return db.Db.Put(uow, emailKey, account.ID)
}

func (db ImplKeyValueDb) GetByEmailAddress(emailAddress string) (*userAccount.UserAccount, error) {
	time.Sleep(time.Second)

	// Get the user ID from the email index
	emailKey := emailIndexKey(emailAddress)
	userId, err := db.Db.Get(emailKey)
	if err != nil {
		return nil, err
	}

	if userId == nil {
		return nil, nil
	}

	// Use the user ID to get the actual user account
	return db.GetById(*userId)
}

var _ UserAccountDb = (*ImplKeyValueDb)(nil)
