package auth

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type EmailJWTAuthService struct {
	secret []byte
}

func NewEmailJWTAuthService(secret []byte) *EmailJWTAuthService {
	return &EmailJWTAuthService{
		secret: secret,
	}
}

// Credential implements platform.AuthService.
func (*EmailJWTAuthService) Credential(ctx *gin.Context) (any, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, ErrEmptyCredential
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		return token, nil
	} else {
		return nil, ErrUnsupportedCredential
	}
}

// ServiceID implements platform.AuthService.
func (*EmailJWTAuthService) ServiceID() string {
	return "supabase_jwt"
}

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

// Validate implements platform.AuthService.
func (s *EmailJWTAuthService) Validate(ctx context.Context, credential any) (email string, err error) {
	v, ok := credential.(string)
	if !ok {
		return "", ErrInvalidCredential
	}

	t, err := jwt.ParseWithClaims(v, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnsupportedHeader
		}
		return s.secret, nil
	})
	if err != nil {
		return "", nil
	}

	claim, ok := t.Claims.(*Claims)
	if !ok {
		return "", ErrInvalidCredential
	}

	return claim.Email, nil
}
