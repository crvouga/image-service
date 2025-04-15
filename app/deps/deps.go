package deps

import (
	"database/sql"
	"imageresizerservice/app/users/loginWithEmailLink/link/linkDb"
	"imageresizerservice/library/email/emailOutbox"
	"imageresizerservice/library/email/sendEmail"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/uow"
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
