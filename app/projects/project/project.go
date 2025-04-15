package project

import "time"

type ProjectID string

type Project struct {
	ID              ProjectID
	CreatedByUserID string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	AllowedDomains  []string
}
