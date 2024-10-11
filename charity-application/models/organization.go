package models

import "gorm.io/gorm"

type Organization struct {
    gorm.Model
    Name string
    Email []string
    Address []string
    Description string
    TermsOfCooperation string
    FundraisingCampaigns []*FundraisingCampaign `gorm: "many2many:organization_fundraising_campaigns"`
    Projects []*Project `gorm: "many2many:organization_projects"`
}
