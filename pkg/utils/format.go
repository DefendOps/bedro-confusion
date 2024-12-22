package utils

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

type FormatValidator struct{}

type FormatType int

const (
	URLFormat FormatType = iota
	FileFormat
	NPMFormat
	PyPIFormat
	GoModFormat
)

var (
	ErrInvalidFormatType = errors.New("invalid format type")
	ErrInvalidFormat     = errors.New("format validation failed")
)

func (fv *FormatValidator) Validate(source string, _type FormatType) (bool, error) {
	switch _type {
	case URLFormat:
		_, err := url.ParseRequestURI(source)
		if err != nil {
			return false, fmt.Errorf("%w: invalid URL format", ErrInvalidFormat)
		}
		return true, nil

	case FileFormat:
		if matched, _ := regexp.MatchString(`^.*\.(json|txt|mod)$`, source); matched {
			return true, nil
		}
		return false, fmt.Errorf("%w: invalid File format", ErrInvalidFormat)

	case NPMFormat:
		if matched, _ := regexp.MatchString(`^(@[a-zA-Z0-9-]+\/[a-zA-Z0-9-_]+|[a-zA-Z0-9-_]+)$`, source); matched {
			return true, nil
		}
		return false, fmt.Errorf("%w: invalid NPM package name", ErrInvalidFormat)

	case PyPIFormat:
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9-_]+$`, source); matched {
			return true, nil
		}
		return false, fmt.Errorf("%w: invalid PyPI package name", ErrInvalidFormat)

	case GoModFormat:
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9.-]+\/[a-zA-Z0-9-\/]+$`, source); matched {
			return true, nil
		}
		return false, fmt.Errorf("%w: invalid Go module name", ErrInvalidFormat)

	default:
		return false, ErrInvalidFormatType
	}
}
