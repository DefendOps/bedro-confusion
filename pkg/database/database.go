package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
    // PSQL Connection
    // connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))
    var err error
    db, err = gorm.Open(sqlite.Open("bedro-confuser.db"), &gorm.Config{
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