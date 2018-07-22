package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/err0r500/go-realworld-clean/uc"
)

var tokenTimeToLive = time.Hour * 2

// tokenHandler handles JWT related request, implementing uc.AuthHandler interface
type tokenHandler struct {
	salt []byte
}

func NewTokenHandler(salt string) uc.AuthHandler {
	return tokenHandler{
		salt: []byte(salt),
	}
}

//GenToken (uc.Admin) : returns a signed token for an admin
func (tH tokenHandler) GenUserToken(userID string) (string, error) {
	if userID == "" {
		return "", errors.New("can't generate token for empty user")
	}

	return jwt.
		NewWithClaims(jwt.SigningMethodHS256, newUserClaims(userID, tokenTimeToLive)).
		SignedString(tH.salt)
}

func (tH tokenHandler) GetUserName(inToken string) (userID string, err error) {
	token, err := jwt.ParseWithClaims(
		inToken,
		&userClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tH.salt), nil
		},
	)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", errors.New("problem with jwt")
}

type userClaims struct {
	UserID string
	jwt.StandardClaims
}

// newUserClaims : constructor of userClaims
func newUserClaims(id string, ttl time.Duration) *userClaims {
	return &userClaims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Issuer:    "real-worl-demo-backend",
		},
	}
}
