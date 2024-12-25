package database

import (
	"os"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID				uint	`gorm:"primaryKey"`
	Email 			string
	Username 		string	`json:"username,omitempty"`
	Password		string	`gorm:"size:255" json:"password"`
	Registry		string
	CreatedPackages	[]CreatedPackage `gorm:"foreignKey:AccountID"`
	AccessTokens	[]AccessToken `gorm:"foreignKey:AccountID"`

	CreatedAt		time.Time
  	UpdatedAt		time.Time
}

type AccessToken struct {
	gorm.Model
	ID			uint	`gorm:"primaryKey"`
	AccessToken	string

	CreatedAt	time.Time
  	UpdatedAt	time.Time
	
	AccountID	uint
}

type CreatedPackage struct {
	gorm.Model
	ID			uint	`gorm:"primaryKey"`
	PackageName	string

	CreatedAt	time.Time
  	UpdatedAt	time.Time
	
	AccountID	uint
}

// AccessToken

func (a *AccessToken) BeforeSave(tx *gorm.DB) (err error) {
	if len(a.AccessToken) > 0 {
		encryptedToken, err := encrypt(a.AccessToken, os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			return err
		}
		tx.Statement.SetColumn("AccessToken", encryptedToken)
	}
	return nil
}

func (a *AccessToken) AfterFind(tx *gorm.DB) (err error) {
	if len(a.AccessToken) > 0 {
		decryptedToken, err := decrypt(a.AccessToken, os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			return err
		}
		a.AccessToken = decryptedToken
	}
	return nil
}

// Account

func (a *Account) BeforeSave(tx *gorm.DB) (err error) {
	if len(a.Password) > 0 {
		encryptedPassword, err := encrypt(a.Password, os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			return err
		}
		tx.Statement.SetColumn("Password", encryptedPassword)
	}
	return nil
}

func (a *Account) AfterFind(tx *gorm.DB) (err error) {
	if len(a.Password) > 0 {
		decryptedPassword, err := decrypt(a.Password, os.Getenv("ENCRYPTION_KEY"))
		if err != nil {
			return err
		}
		a.Password = decryptedPassword
	}
	return nil
}