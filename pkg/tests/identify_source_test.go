package tests

import (
	"testing"

	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/stretchr/testify/assert"
)

func TestIdentifySource_PackageName(t *testing.T) {
	// Arrange
	srcIdentifier := source.SourceIdentifier{}
	sourceType := source.PackageNameSource

	testCases := []struct {
		name      string
		input     string
		expectErr bool
		expected  source.Source
	}{
		{
			name:      "Valid Not Existing Package Name",
			input:     "@testing/omarbdrn",
			expectErr: false,
			expected: source.Source{
				Type:     source.PackageNameSource,
				Registry: "npm",
				RawValue: "@testing/omarbdrn",
			},
		},
		{
			name:      "Valid Existing Package Name",
			input:     "@nextui-org/Input",
			expectErr: false,
			expected: source.Source{
				Type:     source.PackageNameSource,
				Registry: "npm",
				RawValue: "@nextui-org/Input",
			},
		},
		{
			name:      "Valid Package Name",
			input:     "github.com/omarbdrn/omarbdrn",
			expectErr: false,
			expected: source.Source{
				Type:     source.PackageNameSource,
				Registry: "gomod",
				RawValue: "github.com/omarbdrn/omarbdrn",
			},
		},
		{
			name:      "Invalid Package Name",
			input:     "invalid_package!",
			expectErr: true,
			expected: source.Source{
				Type:     source.PackageNameSource,
				Registry: "unknown",
				RawValue: "invalid_package!",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results, err := srcIdentifier.IdentifySource(tc.input, sourceType)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, result := range results{
				assert.Equal(t, source.PackageNameSource, result.Type)
				assert.Equal(t, tc.input, result.RawValue)
				assert.Contains(t, []source.Registry{source.NPM, source.PyPI, source.GoMod}, result.Registry)
			}
		})
	}
}

func TestIdentifySource_InvalidPackageName(t *testing.T) {
	srcIdentifier := source.SourceIdentifier{}
	input := "invalid_package!"
	sourceType := source.PackageNameSource

	_, err := srcIdentifier.IdentifySource(input, sourceType)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no eligible formats found")
}

func TestIdentifySource_URLSource(t *testing.T) {
	srcIdentifier := source.SourceIdentifier{}
	testCases := []struct {
		name      string
		input     string
		expectErr bool
		expected  source.Source
	}{
		{
			name:      "Valid HTTPS URL",
			input:     "https://example.com",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "unknown",
				RawValue: "https://example.com",
			},
		},
		{
			name:      "Valid HTTP URL",
			input:     "http://example.org",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "unknown",
				RawValue: "http://example.org",
			},
		},
		{
			name:      "Valid URL to PackageJSON",
			input:     "https://raw.githubusercontent.com/i5ting/stuq-koa/refs/heads/master/package.json",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "npm",
				RawValue: "https://raw.githubusercontent.com/i5ting/stuq-koa/refs/heads/master/package.json",
			},
		},
		{
			name:      "Valid URL to Random JSON File",
			input:     "https://raw.githubusercontent.com/i5ting/stuq-koa/refs/heads/master/book.json",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "unknown",
				RawValue: "https://raw.githubusercontent.com/i5ting/stuq-koa/refs/heads/master/book.json",
			},
		},
		{
			name:      "Valid URL to Go Mod File with No Modules",
			input:     "https://raw.githubusercontent.com/4dex/opentelemetry-operations-go1.18/refs/heads/main/go.mod",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "gomod",
				RawValue: "https://raw.githubusercontent.com/4dex/opentelemetry-operations-go1.18/refs/heads/main/go.mod",
			},
		},
		{
			name:      "Valid URL to Go Mod File With Modules",
			input:     "https://raw.githubusercontent.com/go101/golds/refs/heads/develop/go.mod",
			expectErr: false,
			expected: source.Source{
				Type:     source.URLSource,
				Registry: "gomod",
				RawValue: "https://raw.githubusercontent.com/go101/golds/refs/heads/develop/go.mod",
			},
		},
		{
			name:      "Invalid URL format",
			input:     "example.com",
			expectErr: true,
			expected:  source.Source{},
		},
		{
			name:      "Empty URL",
			input:     "",
			expectErr: true,
			expected:  source.Source{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results, err := srcIdentifier.IdentifySource(tc.input, source.URLSource)
			
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			for _, result := range results{
				assert.Equal(t, tc.expected.Type, result.Type)
				assert.Equal(t, tc.expected.Registry, result.Registry)
				assert.Equal(t, tc.expected.RawValue, result.RawValue)
			}
		})
	}
}
