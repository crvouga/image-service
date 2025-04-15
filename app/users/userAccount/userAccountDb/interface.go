package userAccountDb

import (
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/library/uow"
)

type UserAccountDb interface {
	GetById(id string) (*userAccount.UserAccount, error)
	GetByEmailAddress(emailAddress string) (*userAccount.UserAccount, error)
	Upsert(uow *uow.Uow, userAccount userAccount.UserAccount) error
}
