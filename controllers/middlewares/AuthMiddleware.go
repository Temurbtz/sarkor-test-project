package middlewares

import (
	"net/http"
	"test-project/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
// мидлвар где валидируется куки
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := c.Cookie("SESSTOKEN")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        claims := &utils.AuthClaims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("jwt_secret_key"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Set("claims", claims)
        c.Next()
    }
}
