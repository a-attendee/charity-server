package models

import (
	"time"

	"github.com/google/uuid"
)

// This file contains models for GORM //
// This models is going to be used to create tables in database //

// This is the basic user representing fields that every user must have //
type User struct {
    ID  uuid.UUID `gorm: "primaryKey"` 
    FirstName string
    LastName string
    Email   string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
