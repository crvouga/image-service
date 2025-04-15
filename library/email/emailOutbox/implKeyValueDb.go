package emailOutbox

import (
	"encoding/json"
	"errors"
	"imageresizerservice/library/email/email"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
)

type ImplKeyValueDb struct {
	Db keyValueDb.KeyValueDb
}

func (impl *ImplKeyValueDb) Add(uow *uow.Uow, email email.Email) error {
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return err
	}
	return impl.Db.Put(uow, "email", string(emailJSON))
}

func (impl *ImplKeyValueDb) GetUnsentEmails() ([]email.Email, error) {
	// This is a simplified implementation
	// In a real implementation, we would need to query all emails that are not marked as sent
	key := "unsent_emails"
	emailsJSON, err := impl.Db.Get(key)
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

func (impl *ImplKeyValueDb) MarkAsSent(uow *uow.Uow, email email.Email) error {
	// This is a simplified implementation
	// In a real implementation, we would need to update the specific email
	// For now, we'll just return an error indicating this needs to be implemented
	return errors.New("MarkAsSent not implemented")
}

// Ensure ImplKeyValueDb implements EmailOutbox interface
var _ EmailOutbox = (*ImplKeyValueDb)(nil)
