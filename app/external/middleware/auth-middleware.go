package middleware

import (
	"itv_go/internal/entity/global"
	"itv_go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtUsecase *usecase.JwtUsecase
}

func NewAuthMiddleware(jwtUsecase *usecase.JwtUsecase) *AuthMiddleware {
	return &AuthMiddleware{jwtUsecase}
}

func (m *AuthMiddleware) CheckAccesToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(global.ErrStatusCodes[global.ErrNotAllowedToUse], gin.H{"message": global.ErrNotAllowedToUse.Error()})
			c.Abort()
			return
		}

		_, err = m.jwtUsecase.ParseToken(token)
		if err != nil {
			c.JSON(global.ErrStatusCodes[global.ErrExpiredSesstion], gin.H{"message": global.ErrExpiredSesstion.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
