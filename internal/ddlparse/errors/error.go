package errors

import "github.com/pkg/errors"

var (
	DDLParseFail = errors.New("DDL parsing failed")
)
