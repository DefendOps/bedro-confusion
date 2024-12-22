package url

import (
	"errors"
	"fmt"

	"github.com/defendops/bedro-confuser/pkg/utils"
	"github.com/defendops/bedro-confuser/pkg/utils/scan"
	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/spf13/cobra"
)

var (
	errParameters = errors.New("please specify a URL source to scan")
)

func NewCmdRun() *cobra.Command {
	var config scan.CliConfig;

	cmd := &cobra.Command{
		Use:   "url",
		Short: "-> Scan a URL/Website for packages",
		Long:  `BedroConfuser Scan URL: Scan a URL/Website for packages`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return errParameters
			}

			url := args[0]
			scan_config := scan.CreateScanConfig(config)
			formatValidator := utils.FormatValidator{}
			_, err := formatValidator.Validate(url, utils.URLFormat)
			if err != nil {
				return err
			}

			identifier := source.SourceIdentifier{}
			sources, err := identifier.IdentifySource(url, source.URLSource)
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