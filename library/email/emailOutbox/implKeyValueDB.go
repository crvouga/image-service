package emailOutbox

import (
	"encoding/json"
	"errors"
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDB struct {
	entities keyValueDB.KeyValueDB
}

func NewImplKeyValueDB(db keyValueDB.KeyValueDB) *ImplKeyValueDB {
	return &ImplKeyValueDB{
		entities: keyValueDB.NewImplNamespaced(db, "emailOutbox"),
	}
}

func (impl *ImplKeyValueDB) Add(uow *uow.Uow, email email.Email) error {
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return err
	}
	return impl.entities.Put(uow, "email", string(emailJSON))
}

func (impl *ImplKeyValueDB) GetUnsentEmails() ([]email.Email, error) {
	// This is a simplified implementation
	// In a real implementation, we would need to query all emails that are not marked as sent
	key := "unsent_emails"
	emailsJSON, err := impl.entities.Get(key)
	if err != nil {
		return nil, err
	}

	var emails []email.Email
	if emailsJSON == nil {
		return emails, nil
	}

	err = json.Unmarshal([]byte(*emailsJSON), &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (impl *ImplKeyValueDB) MarkAsSent(uow *uow.Uow, email email.Email) error {
	// This is a simplified implementation
	// In a real implementation, we would need to update the specific email
	// For now, we'll just return an error indicating this needs to be implemented
	return errors.New("MarkAsSent not implemented")
}

// Ensure ImplKeyValueDB implements EmailOutbox interface
var _ EmailOutbox = (*ImplKeyValueDB)(nil)
