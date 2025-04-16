package appCtx

import (
	"database/sql"

	"imageresizerservice/app/projects/project/projectDB"
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
	ProjectDB     projectDB.ProjectDB
}

func (ac *AppCtx) CleanUp() {
	ac.DB.Close()
}

func New() AppCtx {
	db := sqlite.New()

	keyValueDBFs := keyValueDB.NewImplFs("keyValueDB.json")

	return AppCtx{
		DB:            db,
		UowFactory:    *uow.NewFactory(db),
		Logger:        slog.Default(),
		KeyValueDB:    keyValueDB.NewImplNamespaced(keyValueDBFs, "app"),
		LinkDB:        linkDB.NewImplKeyValueDB(keyValueDBFs),
		EmailOutbox:   emailOutbox.NewImplKeyValueDB(keyValueDBFs),
		UserSessionDB: userSessionDB.NewImplKeyValueDB(keyValueDBFs),
		UserAccountDB: userAccountDB.NewImplKeyValueDB(keyValueDBFs),
		ProjectDB:     projectDB.NewImplKeyValueDB(keyValueDBFs),
	}
}

func NewTest() AppCtx {
	db := sqlite.New()

	keyValueDBHashMap := keyValueDB.ImplHashMap{}

	return AppCtx{
		DB:            db,
		UowFactory:    *uow.NewFactory(db),
		Logger:        slog.Default(),
		KeyValueDB:    &keyValueDBHashMap,
		LinkDB:        linkDB.NewImplKeyValueDB(&keyValueDBHashMap),
		EmailOutbox:   emailOutbox.NewImplKeyValueDB(&keyValueDBHashMap),
		UserSessionDB: userSessionDB.NewImplKeyValueDB(&keyValueDBHashMap),
		UserAccountDB: userAccountDB.NewImplKeyValueDB(&keyValueDBHashMap),
		ProjectDB:     projectDB.NewImplKeyValueDB(&keyValueDBHashMap),
	}
}
