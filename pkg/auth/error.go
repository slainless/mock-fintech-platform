package auth

import (
	"errors"
)

var ErrUnsupportedCredential = errors.New("unsupported credential")
var ErrEmptyCredential = errors.New("empty credential")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrUnsupportedHeader = errors.New("unsupported header")
