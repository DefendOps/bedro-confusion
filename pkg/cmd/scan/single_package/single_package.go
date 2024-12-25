package single_package

import (
	"errors"
	"fmt"

	"github.com/defendops/bedro-confuser/pkg/utils/scan"
	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/defendops/bedro-confuser/pkg/utils/types"
	"github.com/spf13/cobra"
)

var (
	errParameters = errors.New("please specify a Package Name to scan")
)

func NewCmdRun() *cobra.Command {
	var config types.CliConfig;

	cmd := &cobra.Command{
		Use:   "package",
		Short: "-> Scan a Single Package",
		Long:  `BedroConfuser Scan URL: Scan a Single Package`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errParameters
			}

			package_name := args[0]
			scan_config := scan.CreateScanConfig(config)
			
			identifier := source.SourceIdentifier{}
			sources, err := identifier.IdentifySource(package_name, source.PackageNameSource)
			if err != nil {
				fmt.Printf("Error identifying source: %s\n", err)
			}
			
			ctx := cmd.Context()
			scan.ScanSources(sources, scan_config, &ctx)

			return nil
		},
	}

	cmd.PersistentFlags().Int16VarP(&config.Timeout, "input", "i", 360, "Scan Timeout (Seconds)")
	cmd.PersistentFlags().BoolVarP(&config.Takeover, "takeover", "t", false, "Try to takeover Packages (Default: false)")

	return cmd
}