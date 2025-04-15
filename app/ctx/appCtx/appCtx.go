package appCtx

import (
	"database/sql"
	"imageresizerservice/app/users/loginWithEmailLink/link/linkDb"
	"imageresizerservice/app/users/userAccount/userAccountDb"
	"imageresizerservice/app/users/userSession/userSessionDb"
	"imageresizerservice/library/email/emailOutbox"
	"imageresizerservice/library/email/sendEmail"
	"imageresizerservice/library/keyValueDb"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
	"log/slog"
)

type AppCtx struct {
	SendEmail     sendEmail.SendEmail
	LinkDb        linkDb.LinkDb
	UowFactory    uow.UowFactory
	EmailOutbox   emailOutbox.EmailOutbox
	KeyValueDb    keyValueDb.KeyValueDb
	UserSessionDb userSessionDb.UserSessionDb
	UserAccountDb userAccountDb.UserAccountDb
	Db            *sql.DB
	Logger        *slog.Logger
}

func (appCtx *AppCtx) CleanUp() {
	appCtx.Db.Close()
}

func New() AppCtx {
	db := sqlite.New()

	keyValueDbHashMap := keyValueDb.ImplHashMap{}

	return AppCtx{
		SendEmail:     &sendEmail.ImplFake{},
		LinkDb:        &linkDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UowFactory:    uow.UowFactory{Db: db},
		KeyValueDb:    &keyValueDbHashMap,
		EmailOutbox:   &emailOutbox.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UserSessionDb: &userSessionDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UserAccountDb: &userAccountDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		Db:            db,
		Logger:        slog.Default(),
	}

}

func NewTest() AppCtx {
	db := sqlite.New()

	keyValueDbHashMap := keyValueDb.ImplHashMap{}

	return AppCtx{
		SendEmail:     &sendEmail.ImplFake{},
		LinkDb:        &linkDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UowFactory:    uow.UowFactory{Db: db},
		KeyValueDb:    &keyValueDbHashMap,
		EmailOutbox:   &emailOutbox.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UserSessionDb: &userSessionDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		UserAccountDb: &userAccountDb.ImplKeyValueDb{Db: &keyValueDbHashMap},
		Db:            db,
		Logger:        slog.Default(),
	}
}
