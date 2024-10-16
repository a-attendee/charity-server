package database

import (
	"os"

	"github.com/a-attendee/charity-server/models"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBInstance struct {
    Db *gorm.DB
}

var Database DBInstance

func InitDB() {
    db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})


    db.Logger = slogGorm.New()
    

    if err != nil {
        db.Logger.Error(nil, "Failed to connect to the database!\n", err.Error())
        os.Exit(2)
    }

    err = db.AutoMigrate(
                    &models.User{}, 
                    &models.Admin{},
                    &models.Project{},
     //               &models.Donater{},
     //               &models.Organizaion{},
     //               &models.HistoryDonations{},
    )

    if err != nil {
        db.Logger.Error(nil, "Failed to connect to the database!\n", err.Error())
        os.Exit(2)

    }
    Database = DBInstance{Db: db}

    db.Logger.Info(nil, "Connected to the database: success!")
}
