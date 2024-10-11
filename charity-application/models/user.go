package models

import (
	"database/sql"

)

type User struct {
    ID uint `gorm: "primaryKey"`
    
    FirstName string        `json: "first_name"`
    LastName string         `json: "last_name"`
    Email sql.NullString    `json: "email"`
    Money int               `json: "money"`

    FundraisingCampaigns []FundraisingCampaign `gorm: "many2many:user_fundraising_campaign; foreignKey:id; References:ID"`
    
    Projects []Project `gorm: "many2many:user_projects; foreignKey:id References:ID"`
}
