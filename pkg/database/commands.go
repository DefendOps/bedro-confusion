package database

import (
	"errors"

	"gorm.io/gorm"
)

func CreateAccount(email string, username string, password string, registry string) error{
	if len(email) > 6 && len(password) > 6 && registry != "" {
		var AccountModel Account
		result := db.Where("Email = @email AND Registry = @registry", map[string]interface{}{"email": email, "registry": registry}).First(&AccountModel)
		if result.Error != nil{
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				AccountModel := &Account{
					Email: email,
					Username: username,
					Password: password,
					Registry: registry,
				}
								
				db.Create(AccountModel)
				return nil
			}
		}

		return ErrRecordExists
	}
	
	return ErrInvalidArguments
}

func AddAccessTokenToAccount(email string, token string) error{
	var AccountModel Account
	result := db.Preload("AccessTokens").Where("Email = ?", email).First(&AccountModel)
	if result.Error != nil{
		return result.Error
	}

	for _, actoken := range AccountModel.AccessTokens{
		if actoken.AccessToken == token{
			return errors.New("token already exists")
		}
	}

	err := db.Model(&AccountModel).Association("AccessTokens").Append(&AccessToken{AccessToken: token})
	if err != nil {
		return err
	}
	
	return nil
}

func MigrateModels() {
	db := GetDB()
	db.AutoMigrate(
		&Account{},
		&CreatedPackage{},
		&AccessToken{},
	)
}