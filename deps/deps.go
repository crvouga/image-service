package deps

import (
	"database/sql"
	"imageresizerservice/email/emailOutbox"
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/keyValueDb"
	"imageresizerservice/uow"
	"imageresizerservice/users/loginWithEmailLink/link/linkDb"
)

type Deps struct {
	BaseUrl     string
	SendEmail   sendEmail.SendEmail
	LinkDb      linkDb.LinkDb
	UowFactory  uow.UowFactory
	EmailOutbox emailOutbox.EmailOutbox
	KeyValueDb  keyValueDb.KeyValueDb
}

func New(db *sql.DB, baseUrl string) Deps {

	keyValueDbHashMap := keyValueDb.ImplHashMap{}

	return Deps{
		BaseUrl:     baseUrl,
		SendEmail:   &sendEmail.ImplFake{},
		LinkDb:      &linkDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UowFactory:  uow.UowFactory{Db: db},
		KeyValueDb:  &keyValueDbHashMap,
		EmailOutbox: &emailOutbox.ImplKeyValueDb{Db: &keyValueDbHashMap},
	}

}
