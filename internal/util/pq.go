package util

import (
	"errors"

	"github.com/lib/pq"
)

func PQErrorCode(err error) string {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code.Name()
	}
	return ""
}

func PQError(err error) *pq.Error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr
	}
	return nil
}
