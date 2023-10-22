package utils

import "github.com/dgrijalva/jwt-go"
// структура для клаим
type AuthClaims struct {
    jwt.StandardClaims
    UserID int64 `json:"user_id"`
    Login  string `json:"login"`
}
