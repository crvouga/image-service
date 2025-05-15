package userAccountDB

import (
	"encoding/json"
	"imageService/app/users/userAccount"
	"imageService/app/users/userAccount/userRole"
	"imageService/app/users/userID"
	"imageService/library/email/emailAddress"
	"imageService/library/keyValueDB"
	"imageService/library/uow"
	"time"
)

type ImplKeyValueDB struct {
	userAccounts              keyValueDB.KeyValueDB
	indexUserIDByEmailAddress keyValueDB.KeyValueDB
	indexUserIDsByRole        keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		userAccounts:              keyValueDB.NewImplNamespaced(db, "userAccount"),
		indexUserIDByEmailAddress: keyValueDB.NewImplNamespaced(db, "userAccount:index:userIDByEmailAddress"),
		indexUserIDsByRole:        keyValueDB.NewImplNamespaced(db, "userAccount:index:userIDsByRole"),
	}
}

func emailIndexKey(emailAddress emailAddress.EmailAddress) string {
	return string(emailAddress)
}

func userAccountKey(id userID.UserID) string {
	return string(id)
}

func roleIndexKey(role userRole.Role) string {
	return string(role)
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
	if err := db.indexUserIDByEmailAddress.Put(uow, emailIndexKey(account.EmailAddress), string(account.UserID)); err != nil {
		return err
	}

	// Create or update the role index
	roleKey := roleIndexKey(account.Role)
	roleIndex, err := db.indexUserIDsByRole.Get(roleKey)

	var userIDs []string
	if roleIndex != nil {
		if err := json.Unmarshal([]byte(*roleIndex), &userIDs); err != nil {
			return err
		}

		// Check if user ID already exists in the role index
		found := false
		for _, id := range userIDs {
			if id == string(account.UserID) {
				found = true
				break
			}
		}

		// Add user ID to role index if not already present
		if !found {
			userIDs = append(userIDs, string(account.UserID))
		}
	} else {
		// Create new role index with this user ID
		userIDs = []string{string(account.UserID)}
	}

	// Store updated role index
	roleIndexJSON, err := json.Marshal(userIDs)
	if err != nil {
		return err
	}

	return db.indexUserIDsByRole.Put(uow, roleKey, string(roleIndexJSON))
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

func (db ImplKeyValueDB) GetByRole(role userRole.Role) ([]*userAccount.UserAccount, error) {
	// Get the list of user IDs for this role
	roleIndex, err := db.indexUserIDsByRole.Get(roleIndexKey(role))
	if err != nil {
		return nil, err
	}

	if roleIndex == nil {
		return []*userAccount.UserAccount{}, nil
	}

	var userIDs []string
	if err := json.Unmarshal([]byte(*roleIndex), &userIDs); err != nil {
		return nil, err
	}

	// Fetch each user account by ID
	accounts := make([]*userAccount.UserAccount, 0, len(userIDs))
	for _, id := range userIDs {
		account, err := db.GetByUserID(userID.UserID(id))
		if err != nil {
			return nil, err
		}

		if account != nil {
			accounts = append(accounts, account)
		}
	}

	return accounts, nil
}

var _ UserAccountDB = (*ImplKeyValueDB)(nil)
