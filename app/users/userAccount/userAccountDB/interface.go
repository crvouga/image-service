package userAccountDB

import (
	"imageService/app/users/userAccount"
	"imageService/app/users/userAccount/userRole"
	"imageService/app/users/userID"
	"imageService/library/email/emailAddress"
	"imageService/library/uow"
)

type UserAccountDB interface {
	GetByUserID(id userID.UserID) (*userAccount.UserAccount, error)
	GetByEmailAddress(emailAddress emailAddress.EmailAddress) (*userAccount.UserAccount, error)
	GetByRole(role userRole.Role) ([]*userAccount.UserAccount, error)
	Upsert(uow *uow.Uow, userAccount userAccount.UserAccount) error
}
