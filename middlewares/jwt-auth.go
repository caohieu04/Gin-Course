package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
			f, _ := os.Open("token.log")
			f.WriteString(fmt.Sprint("Claims[Name]: ", claims["name"]))
			f.WriteString(fmt.Sprint("Claims[Admin]: ", claims["admin"]))
			f.WriteString(fmt.Sprint("Claims[Issuer]: ", claims["iss"]))
			f.WriteString(fmt.Sprint("Claims[IssuedAt]: ", convertTime(claims["iat"])))
			f.WriteString(fmt.Sprint("Claims[ExpiresAt]: ", convertTime(claims["exp"])))
			defer f.Close()
		} else {
			log.Println("jwt-auth", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
