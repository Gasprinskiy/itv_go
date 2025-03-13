package usecase

import (
	"itv_go/config"
	"itv_go/internal/entity/global"
	appjwt "itv_go/internal/entity/jwt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUsecase struct {
	*config.Config
}

func NewJwtUsecase(conf *config.Config) *JwtUsecase {
	return &JwtUsecase{conf}
}

func (u *JwtUsecase) GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(u.Config.JwtSecretLifeTime) * time.Hour)

	claims := appjwt.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(u.Config.JwtSecret)

	result, err := token.SignedString(secretKey)
	if err != nil {
		err = global.ErrInternalError
	}

	return result, err
}

func (u *JwtUsecase) ParseToken(tokenString string) (*appjwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &appjwt.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.Config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*appjwt.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
