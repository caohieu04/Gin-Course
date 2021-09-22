package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/caohieu04/Gin-Course/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			convertTime := func(T interface{}) time.Time {
				var tm time.Time
				switch iat := T.(type) {
				case float64:
					tm = time.Unix(int64(iat), 0)
				}
				return tm
			}
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", convertTime(claims["iat"]))
			log.Println("Claims[ExpiresAt]: ", convertTime(claims["exp"]))
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
