package userAccountDB

import (
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/uow"
)

type UserAccountDB interface {
	GetByUserID(id userID.UserID) (*userAccount.UserAccount, error)
	GetByEmailAddress(emailAddress emailAddress.EmailAddress) (*userAccount.UserAccount, error)
	Upsert(uow *uow.Uow, userAccount userAccount.UserAccount) error
}
