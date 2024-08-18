package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

type EmailJWTAuthService struct {
	secret []byte
}

// Credential implements platform.AuthService.
func (*EmailJWTAuthService) Credential(ctx *gin.Context) (any, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, ErrEmptyCredential
	}

	return token, nil
}

// ServiceID implements platform.AuthService.
func (*EmailJWTAuthService) ServiceID() string {
	return "supabase_jwt"
}

// Validate implements platform.AuthService.
func (s *EmailJWTAuthService) Validate(ctx context.Context, credential any) (email string, err error) {
	v, ok := credential.(string)
	if !ok {
		return "", ErrInvalidCredential
	}

	token, err := jwt.Verify(jwt.HS256, s.secret, []byte(v))
	if err != nil {
		return "", err
	}

	var c struct {
		Email string `json:"email"`
	}
	err = token.Claims(&c)
	if err != nil {
		return "", err
	}

	return c.Email, nil
}
