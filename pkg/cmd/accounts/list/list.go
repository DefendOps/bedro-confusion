package create

import (
	"fmt"
	"strings"

	"github.com/defendops/bedro-confuser/pkg/registry/accounts"
	"github.com/spf13/cobra"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "-> List Accounts",
		Long:  `BedroConfuser Accounts: List existing accounts for BedroConfuser`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			am := accounts.NewAccountManager()
			accounts, err := am.ListAccounts()
			if err != nil{
				return err
			}

			responseBuffer := "---------------------"
			for _, account := range accounts {
				responseBuffer += "\nEmail: "+account.Email+"\nUsername: "+account.Username+"\nPassword: "+account.Password+"\nRegistry: "+account.Registry+"\nCreatedPackages: "+strings.Join(account.CreatedPackages, ",")+"\nAccessTokens: "+strings.Join(account.AccessTokens, ",")
			}
			responseBuffer += "\n---------------------"

			fmt.Println(responseBuffer)

			return nil
		},
	}

	return cmd
}