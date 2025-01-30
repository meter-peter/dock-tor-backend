package services

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateAMKA      = errors.New("duplicate AMKA")
	// Add other service errors here
)
