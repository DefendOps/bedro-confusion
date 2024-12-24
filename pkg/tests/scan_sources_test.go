package tests

import (
	"testing"

	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/stretchr/testify/assert"
)

func TestScanSources_PackageNameSource(t *testing.T) {
	srcIdentifier := source.SourceIdentifier{}
	testCases := []struct {
		name      string
		input     string
		expectErr bool
		expected  source.Source
	}{
		{
			name:      "",
			input:     "",
			expectErr: true,
			expected: source.Source{},
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
				if tc.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tc.expected.Registry, result.Registry)
				}
			}
		})
	}
}
