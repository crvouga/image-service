package projectName

import (
	"errors"
)

type ProjectName string

func New(name string) (ProjectName, error) {
	if name == "" {
		return "", errors.New("project name is required")
	}
	return ProjectName(name), nil
}
