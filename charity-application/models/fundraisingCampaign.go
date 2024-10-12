package models

import (
	"time"
)

type FundraisingCampaignStatus int8

const (
	Planned FundraisingCampaignStatus = iota
	Holding
	Completed
)

type FundraisingCampaign struct {
	ID uint `gorm: "primaryKey"`

	Name           string
	EventTimeStart time.Time
	Description    string
	ExpectedFundraising int
	ActualFundraising   int
	Status              FundraisingCampaignStatus

	Users []User `gorm: "many2many:user_fundraising_campaign; foreignKey:id: References:ID"`
}
