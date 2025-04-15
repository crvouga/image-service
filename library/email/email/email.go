package email

import (
	"errors"
	"strings"
)

type Email struct {
	To      string
	From    string
	Subject string
	Body    string
}

func ValidateEmailAddress(emailAddress string) error {
	if emailAddress == "" {
		return errors.New("email cannot be empty")
	}

	if !strings.Contains(emailAddress, "@") {
		return errors.New("email must include @")
	}

	if !strings.Contains(emailAddress, ".") {
		return errors.New("email must include a dot")
	}

	return nil
}
