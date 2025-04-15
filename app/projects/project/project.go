package project

import (
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/users/userID"
	"net/url"
	"time"
)

type Project struct {
	ID              projectID.ProjectID
	CreatedByUserID userID.UserID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            projectName.ProjectName
	AllowedDomains  []url.URL
}
