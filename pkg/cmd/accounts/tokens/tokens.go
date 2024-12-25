package tokens

import (
	"errors"
	"fmt"

	"github.com/defendops/bedro-confuser/pkg/registry/accounts"
	"github.com/defendops/bedro-confuser/pkg/utils/types"
	"github.com/spf13/cobra"
)

var (
	errParameters = errors.New("please specify token information")
)

func NewCmdRun() *cobra.Command {
	var config types.AddAccessTokenOptions;

	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "-> Manage Access Tokens",
		Long:  `BedroConfuser Tokens: Add access token to an account if authentication is prohibited`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if config.Token == "" || config.Email == ""{
				return errParameters
			}

			am := accounts.NewAccountManager()
			err := am.AddTokenToAccount(config.Email, config.Token)
			if err != nil {
				return err
			}

			fmt.Println("[i] Token Created and added to Account")

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&config.Token, "access-token", "t", "", "Account AccessToken")
	cmd.PersistentFlags().StringVarP(&config.Email, "email", "e", "", "Account Email")

	return cmd
}