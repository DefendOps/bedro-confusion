package database

import "errors"

var (
	ErrRecordExists		= errors.New("record already exists")
	ErrInvalidArguments	= errors.New("invalid arguments to create model")
)