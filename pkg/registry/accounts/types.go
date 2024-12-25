package accounts

type AccountManager struct{}

type Account struct {
	Email			string
	Username		string
	Password		string
	Registry		string
	AccessTokens	[]string
	CreatedPackages	[]string
}