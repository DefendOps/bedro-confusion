package accounts

import (
	"errors"

	"github.com/defendops/bedro-confuser/pkg/database"
	"gorm.io/gorm"
)

func NewAccountManager() *AccountManager{
	return &AccountManager{}
}

func (am *AccountManager) Authenticate(account Account, RegistryType string){
	
}

func (am *AccountManager) ListAccounts() ([]Account, error){
	var Accounts []database.Account
	result := database.GetDB().Preload("AccessTokens").Preload("CreatedPackages").Find(&Accounts)
	if result.Error != nil{
		return nil, result.Error
	}

	var AccountsModel []Account

	for _, account := range Accounts {
		CreatedPackages := []string{}
		AccessTokens := []string{}

		for _, pkg := range account.CreatedPackages{
			CreatedPackages = append(CreatedPackages, pkg.PackageName)
		}

		for _, token := range account.AccessTokens{
			AccessTokens = append(AccessTokens, token.AccessToken)
		}

		AccountsModel = append(AccountsModel, Account{
			Email: account.Email,
			Username: account.Username,
			Password: account.Password,
			Registry: account.Registry,
			AccessTokens: AccessTokens,
			CreatedPackages: CreatedPackages,
		})
	}

	return AccountsModel, nil
}

func (am *AccountManager) GetAccount(RegistryType string) Account {
	var AccountObj database.Account
	result := database.GetDB().Preload("AccessTokens").Preload("CreatedPackages").Where("Registry = @registry", map[string]interface{}{"registry": RegistryType}).First(&AccountObj)	
	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Account{}
		}
	}

	var CreatedPackages []string
	var AccessTokens	[]string
	
	for _, pkg := range AccountObj.CreatedPackages {
		CreatedPackages = append(CreatedPackages, pkg.PackageName)
	}

	for _, token := range AccountObj.AccessTokens {
		AccessTokens = append(AccessTokens, token.AccessToken)
	}

	return Account{
		Email: AccountObj.Email,
		Username: AccountObj.Username,
		Password: AccountObj.Password,
		Registry: AccountObj.Registry,
		AccessTokens: AccessTokens,
		CreatedPackages: CreatedPackages,
	}
}

func (am *AccountManager) CreateAccount(email string, username string, password string, registry string) error {
	err := database.CreateAccount(email, username, password, registry)
	if err != nil {
		return err
	}

	return nil
}

func (am *AccountManager) AddTokenToAccount(email string, token string) error {
	err := database.AddAccessTokenToAccount(email, token)
	if err != nil {
		return err
	}

	return nil
}