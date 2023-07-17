package storage

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("url exists")
)
