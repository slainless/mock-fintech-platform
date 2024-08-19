package auth

import (
	"errors"

	"github.com/slainless/mock-fintech-platform/internal/util"
)

var ErrUnsupportedCredential = errors.New("unsupported credential")
var ErrEmptyCredential = errors.New("empty credential")
var ErrInvalidCredential = errors.New("invalid credential")
var ErrUnsupportedHeader = util.ErrUnsupportedHeader
