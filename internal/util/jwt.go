package util

import (
	"bytes"
	"errors"

	"github.com/kataras/jwt"
	"github.com/valyala/fastjson"
)

var ErrUnsupportedHeader = errors.New("unsupported header")

func NewHeaderValidator(expectAlg string, expectType string) jwt.HeaderValidator {
	return func(alg string, headerDecoded []byte) (jwt.Alg, jwt.PublicKey, jwt.InjectFunc, error) {
		v, err := fastjson.ParseBytes(headerDecoded)
		if err != nil {
			return nil, nil, nil, err
		}

		if bytes.Equal(v.GetStringBytes("alg"), []byte(expectAlg)) && bytes.Equal(v.GetStringBytes("typ"), []byte(expectType)) {
			return nil, nil, nil, nil
		} else {
			return nil, nil, nil, ErrUnsupportedHeader
		}
	}
}
