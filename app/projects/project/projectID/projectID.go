package projectID

import (
	"errors"
	"imageresizerservice/library/id"
)

type ProjectID string

func Gen() ProjectID {
	return ProjectID(id.Gen())
}

func New(id string) (ProjectID, error) {
	if id == "" {
		return "", errors.New("project ID is required")
	}
	return ProjectID(id), nil
}
