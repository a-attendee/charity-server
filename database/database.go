package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
    slogGorm "github.com/orandin/slog-gorm"
)

type DBInstance struct {
    Db *gorm.DB
}

var Database DBInstance

func InitDB() {
    db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database!\n", err.Error())
        os.Exit(2)
    }

    log.Println("Connected to the database: success!")

    db.Logger = slogGorm.New()

    Database = DBInstance{Db: db}
    
}
