package middleware

import (
	"fmt"
	"movie-app/internal/config"
	"movie-app/utils/auth"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserMiddlware struct {
	config config.Config
}

func NewUserMiddleware(cfg config.Config) (UserMiddlware, error) {
	return UserMiddlware{config: cfg}, nil
}

func (m *UserMiddlware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		jwToken := parts[1]
		token, err := jwt.ParseWithClaims(
			jwToken,
			&auth.AuthJWT{},
			func(token *jwt.Token) (any, error) {
				if jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg()) != token.Method {
					return nil, errors.Errorf("invalid signing method: %v", token.Header["alg"])
				}
				return []byte(m.config.JwtKey), nil
			},
		)
		fmt.Println(jwToken)
		if err == nil {
			claims := token.Claims.(*auth.AuthJWT)
			if claims.Role != "user" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				return
			}
			ctx.Set("AUTH_DATA", *claims)

			ctx.Next()
			return
		}
	} else {
		ctx.Next()
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}
