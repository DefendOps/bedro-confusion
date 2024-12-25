package scan

import (
	"context"

	"github.com/defendops/bedro-confuser/pkg/registry"
	utilsSource "github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/defendops/bedro-confuser/pkg/utils/types"
)

func ScanSources(sources []utilsSource.Source, scan_config types.Config, ctx *context.Context){
	for _, modules := range registry.RegistryModules{
		for _, module := range modules {
			for _, source := range sources{
				if module.GonnaBeExecuted(string(source.Registry)){
					module.Run(source, scan_config, ctx)
				}
			}
		}
	}
}