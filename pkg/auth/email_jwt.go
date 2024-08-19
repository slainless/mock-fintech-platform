package auth

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
	"github.com/slainless/mock-fintech-platform/internal/util"
)

var headerValidator = util.NewHeaderValidator("HS256", "JWT")

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

// Validate implements platform.AuthService.
func (s *EmailJWTAuthService) Validate(ctx context.Context, credential any) (email string, err error) {
	v, ok := credential.(string)
	if !ok {
		return "", ErrInvalidCredential
	}

	token, err := jwt.VerifyWithHeaderValidator(jwt.HS256, s.secret, []byte(v), headerValidator)
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
