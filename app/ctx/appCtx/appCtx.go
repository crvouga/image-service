package appCtx

import (
	"database/sql"

	"imageresizerservice/app/users/login/link/linkDB"
	"imageresizerservice/app/users/userAccount/userAccountDB"
	"imageresizerservice/app/users/userSession/userSessionDB"
	"imageresizerservice/library/email/emailOutbox"
	"imageresizerservice/library/keyValueDB"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/uow"
	"log/slog"
)

type AppCtx struct {
	DB            *sql.DB
	Logger        *slog.Logger
	UowFactory    uow.UowFactory
	LinkDB        linkDB.LinkDB
	EmailOutbox   emailOutbox.EmailOutbox
	KeyValueDB    keyValueDB.KeyValueDB
	UserSessionDB userSessionDB.UserSessionDB
	UserAccountDB userAccountDB.UserAccountDB
}

func (appCtx *AppCtx) CleanUp() {
	appCtx.DB.Close()
}

func New() AppCtx {
	db := sqlite.New()

	keyValueDBFs := keyValueDB.NewImplFs("keyValueDB.json")

	return AppCtx{
		UowFactory:    *uow.NewFactory(db),
		DB:            db,
		Logger:        slog.Default(),
		KeyValueDB:    keyValueDB.NewImplNamespaced(keyValueDBFs, "app"),
		LinkDB:        linkDB.NewImplKeyValueDB(keyValueDBFs),
		EmailOutbox:   emailOutbox.NewImplKeyValueDB(keyValueDBFs),
		UserSessionDB: userSessionDB.NewImplKeyValueDB(keyValueDBFs),
		UserAccountDB: userAccountDB.NewImplKeyValueDB(keyValueDBFs),
	}

}

func NewTest() AppCtx {
	db := sqlite.New()

	keyValueDBHashMap := keyValueDB.ImplHashMap{}

	return AppCtx{
		UowFactory:    *uow.NewFactory(db),
		DB:            db,
		Logger:        slog.Default(),
		KeyValueDB:    &keyValueDBHashMap,
		LinkDB:        linkDB.NewImplKeyValueDB(&keyValueDBHashMap),
		EmailOutbox:   emailOutbox.NewImplKeyValueDB(&keyValueDBHashMap),
		UserSessionDB: userSessionDB.NewImplKeyValueDB(&keyValueDBHashMap),
		UserAccountDB: userAccountDB.NewImplKeyValueDB(&keyValueDBHashMap),
	}
}
