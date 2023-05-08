package middleware

import "github.com/golang-jwt/jwt/v4"

// type JwtCustomClaims struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// 	jwt.RegisteredClaims
// }

func CreateToken(id int, email string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["password"] = password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}
