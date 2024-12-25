package root

import (
	"github.com/defendops/bedro-confuser/pkg/cmd/accounts"
	"github.com/defendops/bedro-confuser/pkg/cmd/scan"

	"github.com/spf13/cobra"
)

func NewCmdRoot() (*cobra.Command, error){
	rootCmd := &cobra.Command{
		Use:   "bedro-confuser <command> <subcommand> [flags]",
		Short: "BedroConfuser Dependency Confusion Helper",
		Long:  `BedroConfuser: A tool i wrote to have some fun.`,
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(scan.NewCmdConfig())
	rootCmd.AddCommand(accounts.NewCmdConfig())

	return rootCmd, nil
}