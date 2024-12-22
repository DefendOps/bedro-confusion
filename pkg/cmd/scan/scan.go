package scan

import (
	scanPackage "github.com/defendops/bedro-confuser/pkg/cmd/scan/single_package"
	scanURL "github.com/defendops/bedro-confuser/pkg/cmd/scan/url"
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "scan",
		Short: "-> Scan a list of dependencies",
		Long:  `BedroConfuser Scan: Scan a list of dependencies`,
	}

	cmd.AddCommand(scanURL.NewCmdRun())
	cmd.AddCommand(scanPackage.NewCmdRun())

	return cmd
}