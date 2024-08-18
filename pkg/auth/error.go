package auth

import "errors"

var ErrEmptyCredential = errors.New("empty credential")
var ErrInvalidCredential = errors.New("invalid credential")
