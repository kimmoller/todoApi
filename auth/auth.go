package auth

import (
	"context"
	"fmt"
	"time"
	"todoApi/database"

	"github.com/golang-jwt/jwt/v5"
)

type ValidToken struct {
	Token  jwt.Token
	Claims jwt.MapClaims
}

func generateJWT(username string, expTime time.Time) (string, error) {
	key := []byte("secretSigninKey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,
			"exp": expTime.Unix(),
		})
	return token.SignedString(key)
}

func Login(ctx context.Context, username string, password string) (string, error) {
	identity, err := database.Instance.GetIdentity(ctx, username)
	if err != nil {
		return "", fmt.Errorf("wrong username or password")
	}
	if check := CheckPasswordHash(password, identity.Password); !check {
		return "", fmt.Errorf("wrong username or password")
	}
	return generateJWT(identity.Username, time.Now().Add(5*time.Minute))
}

func ValidateToken(token string) (*ValidToken, error) {
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, &claims,
		func(t *jwt.Token) (interface{}, error) { return []byte("secretSigninKey"), nil },
		jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}
	return &ValidToken{*jwtToken, claims}, nil
}
