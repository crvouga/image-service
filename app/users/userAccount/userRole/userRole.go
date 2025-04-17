package userRole

import (
	"errors"
	"strings"
)

// Role represents a user role in the system
type Role string

const (
	// Admin has full access to all features
	Admin Role = "ADMIN"
	// Standard has standard access to features
	Standard Role = "STANDARD"
)

var (
	// ErrInvalidRole is returned when an invalid role is provided
	ErrInvalidRole = errors.New("invalid user role")
)

// AllRoles returns a slice of all valid roles
func AllRoles() []Role {
	return []Role{Admin, Standard}
}

// IsValid checks if the role is valid
func (r Role) IsValid() bool {
	for _, role := range AllRoles() {
		if r == role {
			return true
		}
	}
	return false
}

// String returns the string representation of the role
func (r Role) String() string {
	return string(r)
}

// ParseRole converts a string to a Role
func ParseRole(s string) (Role, error) {
	role := Role(strings.ToUpper(s))
	if !role.IsValid() {
		return "", ErrInvalidRole
	}
	return role, nil
}
