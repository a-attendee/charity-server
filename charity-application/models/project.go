package models

import "time"

// This file contains models for GORM //
// This models is going to be used to create tables in database //

type Organizaion struct {
    ID uint `gorm: "primaryKey"`
    
    // ManyToMany relationship //
    ProjectID uint `gorm:"primaryKey"`
    Projects *[]Project `gorm: "many2many:project_organization;"`

    Name string
    Description string

    CreatedAt    time.Time
    UpdatedAt    time.Time
}




type Project struct {
    ID uint `gorm: "primaryKey"`

    // ManyToMany relationship //
    //AdminID uint `gorm:"primaryKey"`
    Admins []Admin `gorm: "many2many:admin_projects;"`
    //OrganizaionID uint `gorm:"primaryKey"`
    //Organizaions *[]Organizaion `gorm: "many2many:project_organization;"`

    //DonaterID uint `gorm:"primaryKey"`
    //Donaters *[]Donater `gorm: "many2many:history_donations;"`

    Name string
    Description string
    ExpectedMoneyRise int16
    ActualMoneyRise int16
    StartingDate time.Time
    EndingDate time.Time
    Status      string

    CreatedAt    time.Time
    UpdatedAt    time.Time 
}



// This struct is a join table, but it also contains  //
// information about how many money has been donated  //
type HistoryDonations struct {
    ID uint `gorm: "primaryKey"`

    // Has one relationship //
    DonaterID   uint `gorm:"primaryKey"`
    ProjectID   uint `gorm:"primaryKey"` 

    MoneyDonated int16

    CreatedAt    time.Time
    UpdatedAt    time.Time 
}

