package models

import (
	"time"

)

type Project struct {
    ID uint `gorm: "primaryKey"`
    Name string
    Description string
    ExpectedFundraising int
    BeginFundraisingDate time.Time
    EndFundraisingDate time.Time
}
