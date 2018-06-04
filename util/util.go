package util

import uuid "github.com/satori/go.uuid"

// MakeUUID returns a random generated UUID
func MakeUUID() string {
	uuid := uuid.Must(uuid.NewV4())

	return uuid.String()
}
