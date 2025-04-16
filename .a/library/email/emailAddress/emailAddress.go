package emailAddress

import (
	"errors"
	"strings"
)

type EmailAddress string

func New(emailAddress string) (EmailAddress, error) {

	if emailAddress == "" {
		return "", errors.New("email cannot be empty")
	}

	if !strings.Contains(emailAddress, "@") {
		return "", errors.New("email must include @")
	}

	if !strings.Contains(emailAddress, ".") {
		return "", errors.New("email must include a dot")
	}

	return EmailAddress(emailAddress), nil
}

func NewElsePanic(emailAddress string) EmailAddress {
	email, err := New(emailAddress)
	if err != nil {
		panic(err)
	}
	return email
}
