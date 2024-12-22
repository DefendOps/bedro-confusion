package source

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/defendops/bedro-confuser/pkg/utils"
	utilsRequester "github.com/defendops/bedro-confuser/pkg/utils/requester"
)

type SourceType int

const (
	UnknownSource SourceType = iota
	URLSource
	FileSource
	PackageNameSource
)

type Registry string

const (
	UnknownRegistry Registry = "unknown"
	NPM             Registry = "npm"
	PyPI            Registry = "pypi"
	GoMod           Registry = "gomod"
)

type Source struct {
	Type     SourceType // URL, File, or PackageName
	RawValue string     // Original input string
	Registry Registry   // Registry (NPM, PyPI, etc.)
	Metadata interface{} // Additional metadata (e.g., file content, parsed data)
}

var (
	ErrInvalidSource     = errors.New("invalid source provided")
	ErrUnknownSourceType = errors.New("unknown source type")
)

type SourceIdentifier struct {
	FormatValidator utils.FormatValidator
}

func (si *SourceIdentifier) IdentifySource(input string, sourceType SourceType) ([]Source, error) {
	var sources []Source

	switch sourceType {
	case URLSource:
		source, err := si.processURLSource(input)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)

	case FileSource:
		source, err := si.processFileSource(input)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)

	case PackageNameSource:
		sourcesList, err := si.processPackageNameSource(input, true)
		if err != nil {
			return nil, err
		}
		sources = append(sources, sourcesList...)

	default:
		return nil, ErrUnknownSourceType
	}

	return sources, nil
}

func (si *SourceIdentifier) processURLSource(url string) (Source, error) {
	requester := utilsRequester.HTTPRequester{}
	
	http_request := utilsRequester.HTTPRequest{
		BaseURL: url,
		Endpoint: "",
		Method: "GET",
		IsJson: false,
		Body: "",
		Headers: map[string]string{},
	}

	resp, err := requester.PerformRequest(http_request)
	if err != nil {
		return Source{}, fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Source{}, fmt.Errorf("error reading URL response: %v", err)
	}

	fileType := si.determineContentFormat(string(content))
	registry := si.determineRegistryByFileType(fileType)

	return Source{
		Type:     URLSource,
		RawValue: url,
		Registry: registry,
		Metadata: map[string]string{"FileType": fileType, "content": string(content)},
	}, nil
}

func (si *SourceIdentifier) processFileSource(filePath string) (Source, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Source{}, fmt.Errorf("error reading file: %v", err)
	}

	// Determine file format and registry
	fileType := si.determineContentFormat(string(content))
	registry := si.determineRegistryByFileType(fileType)

	return Source{
		Type:     FileSource,
		RawValue: filePath,
		Registry: registry,
		Metadata: map[string]string{"FileType": fileType},
	}, nil
}

func (si *SourceIdentifier) processPackageNameSource(packageName string, returnList bool) ([]Source, error) {
	var sources []Source

	if isNPM, _ := si.FormatValidator.Validate(packageName, utils.NPMFormat); isNPM {
		sources = append(sources, Source{
			Type:     PackageNameSource,
			RawValue: packageName,
			Registry: NPM,
		})
	}

	if isPyPI, _ := si.FormatValidator.Validate(packageName, utils.PyPIFormat); isPyPI {
		sources = append(sources, Source{
			Type:     PackageNameSource,
			RawValue: packageName,
			Registry: PyPI,
		})
	}

	if isGoMod, _ := si.FormatValidator.Validate(packageName, utils.GoModFormat); isGoMod {
		sources = append(sources, Source{
			Type:     PackageNameSource,
			RawValue: packageName,
			Registry: GoMod,
		})
	}

	if returnList {
		if len(sources) == 0 {
			return nil, fmt.Errorf("no eligible formats found for package name: %s", packageName)
		}
		return sources, nil
	}

	if len(sources) > 0 {
		return []Source{sources[0]}, nil
	}

	return nil, fmt.Errorf("unsupported package name format: %s", packageName)
}

func (si *SourceIdentifier) determineContentFormat(content string) string {
	// Placeholder: Add sophisticated pattern-based checking later
	if strings.Contains(content, "dependencies") {
		return "package.json"
	}
	if strings.Contains(content, "requires") {
		return "requirements.txt"
	}
	if strings.Contains(content, "module") {
		return "go.mod"
	}
	return "unknown"
}

func (si *SourceIdentifier) determineRegistryByFileType(fileType string) Registry {
	switch fileType {
	case "package.json":
		return NPM
	case "requirements.txt":
		return PyPI
	case "go.mod":
		return GoMod
	default:
		return UnknownRegistry
	}
}
