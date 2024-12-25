package accounts

import (
	createAccount "github.com/defendops/bedro-confuser/pkg/cmd/accounts/create"
	listAccounts "github.com/defendops/bedro-confuser/pkg/cmd/accounts/list"
	createToken "github.com/defendops/bedro-confuser/pkg/cmd/accounts/tokens"
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "-> Manage bedro-confuser accounts",
		Long:  `BedroConfuser Accounts: Manage your registered accounts`,
	}

	cmd.AddCommand(createAccount.NewCmdRun())
	cmd.AddCommand(listAccounts.NewCmdRun())
	cmd.AddCommand(createToken.NewCmdRun())

	return cmd
}