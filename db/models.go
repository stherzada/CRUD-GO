// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}