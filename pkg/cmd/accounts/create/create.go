package create

import (
	"errors"
	"fmt"

	"github.com/defendops/bedro-confuser/pkg/registry/accounts"
	"github.com/defendops/bedro-confuser/pkg/utils"
	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/defendops/bedro-confuser/pkg/utils/types"
	"github.com/spf13/cobra"
)

var (
	errParameters = errors.New("please specify account information")
	errInvalidRegistryType = errors.New("invalid registry type")
)

func NewCmdRun() *cobra.Command {
	var config types.AccountCreationOptions;

	cmd := &cobra.Command{
		Use:   "create",
		Short: "-> Create an account",
		Long:  `BedroConfuser Accounts: Create an account for registry to be used in modules`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			allowedRegistries := []string{
				string(source.NPM),
				string(source.PyPI),
				string(source.GoMod),
			}

			if len(config.Email) < 6 || len(config.Password) < 6{
				return errParameters
			}

			if !utils.Contains(allowedRegistries, config.Registry){
				return errInvalidRegistryType
			}

			am := accounts.NewAccountManager()
			err := am.CreateAccount(config.Email, config.Username, config.Password, config.Registry)
			if err != nil {
				return err
			}

			fmt.Println("[i] Account Created")

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&config.Email, "email", "e", "", "Account Email")
	cmd.PersistentFlags().StringVarP(&config.Username, "username", "u", "", "Account Username")
	cmd.PersistentFlags().StringVarP(&config.Password, "password", "p", "", "Account Password")
	cmd.PersistentFlags().StringVarP(&config.Registry, "registry", "r", "", "Account Registry Type (npm, pypi, gomod, rubygems, etc)")

	return cmd
}