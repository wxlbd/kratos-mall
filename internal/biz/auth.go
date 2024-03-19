package biz

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"kratos-admin/api"
	"kratos-admin/internal/conf"
)

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the Jwt Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken = errors.Unauthorized(reason, "Jwt token is missing")
	ErrWrongContext    = errors.Unauthorized(reason, "Wrong context for middleware")
)

type AuthUseCase struct {
	key string
}

func NewAuthUseCase(data *conf.Data) *AuthUseCase {
	return &AuthUseCase{key: data.Jwt.Secret}
}

type MyClaims struct {
	MemberId int64 `json:"member_id"`
	jwtv5.RegisteredClaims
}

func (u *AuthUseCase) GenerateToken(memberId int64, username string) (string, error) {
	claims := MyClaims{
		MemberId: memberId,
		RegisteredClaims: jwtv5.RegisteredClaims{
			Issuer:    "kratos-admin",
			Subject:   username,
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Audience:  jwtv5.ClaimStrings{"kratos-admin", "kratos-admin-api"},
		},
	}
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(u.key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bearerFormat, signedString), nil
}

func (u *AuthUseCase) ParseToken(jwtToken string) (map[string]any, error) {
	claims := MyClaims{}
	token, err := jwtv5.ParseWithClaims(jwtToken, &claims, func(token *jwtv5.Token) (interface{}, error) {
		return []byte(u.key), nil
	})
	if err != nil {
		return nil, api.ErrorUnauthorized("invalid token").WithCause(err)
	}
	if token.Method != jwtv5.SigningMethodHS256 {
		return nil, api.ErrorUnauthorized("invalid token")
	}
	if !token.Valid {
		return nil, api.ErrorUnauthorized("invalid token")
	}

	return map[string]any{
		"member_id": claims.MemberId,
	}, nil
}

func (u *AuthUseCase) JwtMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, ErrMissingJwtToken
				}
				jwtToken := auths[1]
				tokenInfo, err := u.ParseToken(jwtToken)
				if err != nil {
					return nil, err
				}
				ctx = context.WithValue(ctx, "member_id", tokenInfo["member_id"])
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
}
