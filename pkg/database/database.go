package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))
    var err error
    db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
      })
    if err != nil {
        return err
    }

    return nil
}

func GetDB() *gorm.DB {
    return db
}