package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

const (
	// bearerWord 用于授权的承载关键字
	bearerWord string = "Bearer"
	// bearerFormat 授权令牌格式
	bearerFormat string = "Bearer %s"
)

var (
	ErrMissingJwtToken        = errors.New("JWT token is missing")
	ErrTokenInvalid           = errors.New("JWT token is invalid")
	ErrTokenExpired           = errors.New("JWT token is expired")
	ErrTokenParseFail         = errors.New("fail to parse JWT token")
	ErrUnSupportSigningMethod = errors.New("unsupported signing method")
)

type Options struct {
	SigningMethod jwt.SigningMethod
	SigningKey    []byte
	// token过期时长（s）
	ExpiresTime time.Duration
	Issuer      string
	Audience    []string
}

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}
type Option interface {
	Apply(*Options)
}

type optionFunc func(*Options)

func (f optionFunc) Apply(opts *Options) {
	f(opts)
}

func WithSigningMethod(signingMethod jwt.SigningMethod) Option {
	return optionFunc(func(opts *Options) {
		opts.SigningMethod = signingMethod
	})
}

func WithSigningKey(signingKey string) Option {
	return optionFunc(func(opts *Options) {
		opts.SigningKey = []byte(signingKey)
	})
}

func WithExpiresTime(expiresTime int64) Option {
	return optionFunc(func(opts *Options) {
		opts.ExpiresTime = time.Duration(expiresTime)
	})
}

func WithIssuer(issuer string) Option {
	return optionFunc(func(opts *Options) {
		opts.Issuer = issuer
	})
}

func WithAudience(audience []string) Option {
	return optionFunc(func(opts *Options) {
		opts.Audience = audience
	})
}

type Jwt struct {
	opts *Options
}

var defaultOptions = Options{
	SigningMethod: jwt.SigningMethodHS256,
}

func NewJwt(opts ...Option) *Jwt {

	for _, opt := range opts {
		opt.Apply(&defaultOptions)
	}
	return &Jwt{
		opts: &defaultOptions,
	}
}

func (j *Jwt) Generate(memberId, role string) (string, error) {
	claims := Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.opts.Issuer, // 发行人
			Subject:   memberId,      // 主题
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * j.opts.ExpiresTime)),
			Audience:  j.opts.Audience, // 接收人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.opts.SigningKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bearerFormat, signedString), nil
}

func (j *Jwt) Parse(jwtToken string) (*Claims, error) {
	auths := strings.SplitN(jwtToken, " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return nil, ErrMissingJwtToken
	}
	jwtToken = auths[1]
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (any, error) {
		return j.opts.SigningKey, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenUnverifiable) {
			return nil, ErrTokenInvalid
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) || errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenParseFail
	}
	if !token.Valid {
		return nil, ErrTokenInvalid
	}
	if token.Method != j.opts.SigningMethod {
		return nil, ErrUnSupportSigningMethod
	}
	return claims, nil
}
