package scan

import (
	"time"

	"github.com/defendops/bedro-confuser/pkg/utils/types"
)

func CreateScanConfig(config types.CliConfig) types.Config {
	return types.Config{
		DefaultTimeout: time.Duration(config.Timeout) * time.Millisecond,
		CreatePackages: config.Takeover,
	}
}