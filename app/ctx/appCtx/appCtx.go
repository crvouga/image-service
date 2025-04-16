package appCtx

import (
	"database/sql"
	"imageresizerservice/app/users/login/link/linkDb"
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

	keyValueDbFs := keyValueDb.NewImplFs("keyValueDb.json")

	return AppCtx{
		SendEmail:     &sendEmail.ImplFake{},
		UowFactory:    uow.UowFactory{Db: db},
		Db:            db,
		Logger:        slog.Default(),
		KeyValueDb:    keyValueDb.NewImplNamespaced(keyValueDbFs, "app"),
		LinkDb:        linkDb.NewImplKeyValueDb(keyValueDbFs),
		EmailOutbox:   emailOutbox.NewImplKeyValueDb(keyValueDbFs),
		UserSessionDb: userSessionDb.NewImplKeyValueDb(keyValueDbFs),
		UserAccountDb: userAccountDb.NewImplKeyValueDb(keyValueDbFs),
	}

}

func NewTest() AppCtx {
	db := sqlite.New()

	keyValueDbHashMap := keyValueDb.ImplHashMap{}

	return AppCtx{
		SendEmail:     &sendEmail.ImplFake{},
		UowFactory:    uow.UowFactory{Db: db},
		Db:            db,
		Logger:        slog.Default(),
		KeyValueDb:    &keyValueDbHashMap,
		LinkDb:        linkDb.NewImplKeyValueDb(&keyValueDbHashMap),
		EmailOutbox:   emailOutbox.NewImplKeyValueDb(&keyValueDbHashMap),
		UserSessionDb: userSessionDb.NewImplKeyValueDb(&keyValueDbHashMap),
		UserAccountDb: userAccountDb.NewImplKeyValueDb(&keyValueDbHashMap),
	}
}
