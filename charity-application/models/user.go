package models

import (
	"time"
)

// This file contains models for GORM //
// This models is going to be used to create tables in database //

// This is the basic user representing fields that every user must have //
type User struct {
    ID uint `gorm: "primaryKey"` 

    FirstName string
    LastName string
    Email   string

    CreatedAt    time.Time
    UpdatedAt    time.Time
}

//This is the struct representing fields every admin must have //
type Admin struct {

    ID uint `gorm: "primaryKey"` 

    // ManyToMany relationship //
    Projects []Project `gorm: "many2many:admin_projects;"`

    UserID uint 
    User User
    // OrganizationID Organizaion

    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// This is the struct representing fields every donater must have //
type Donater struct {
    ID uint `gorm: "primaryKey"` 
    
    Money int16

    UserID uint 
    User User

    //ManyToMany relationship //
    Projects *[]Project `gorm: "many2many:history_donations;"`

    CreatedAt    time.Time
    UpdatedAt    time.Time
}
