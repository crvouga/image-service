package email

import (
	"errors"
	"strings"
)

func Validate(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	if !strings.Contains(email, "@") {
		return errors.New("email must include @")
	}

	if !strings.Contains(email, ".") {
		return errors.New("email must include a dot")
	}

	return nil
}
