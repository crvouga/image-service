package userAccount

import (
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/email/emailAddress"
	"time"
)

type UserAccount struct {
	UserID       userID.UserID
	EmailAddress emailAddress.EmailAddress
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
