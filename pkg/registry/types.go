package registry

import (
	"context"

	"github.com/defendops/bedro-confuser/pkg/utils/source"
)

type Registry string

type RegistryModule interface {
	Run(src source.Source, ctx *context.Context) error	// Generic Run method for executing tasks
	GonnaBeExecuted(string) bool 						// Determine if this module is acceptable for that source
	Name() string             							// Name of the module
	Registry() Registry									// Registry type (e.g., NPM, PyPI)
	SourceAdapter(source.Source) (interface{}, error)	// SourceAdapter to convert the source into something readable by the module
	CreatePackage(args map[string]interface{}) error	// Create a package dynamically
}

type RegistryModuleID struct {
	name     string
	registry Registry // RepositoryType
}

type SourceAdapter struct{}