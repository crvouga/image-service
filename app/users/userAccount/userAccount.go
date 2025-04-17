package userAccount

import (
	"imageresizerservice/app/users/userAccount/userRole"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"time"
)

type UserAccount struct {
	UserID         userID.UserID
	EmailAddress   emailAddress.EmailAddress
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Role           userRole.Role
	IsRoleAdmin    bool
	IsRoleStandard bool
}

func (u *UserAccount) EnsureComputed() UserAccount {
	u.IsRoleAdmin = u.Role == userRole.Admin
	u.IsRoleStandard = u.Role == userRole.Standard
	return *u
}
