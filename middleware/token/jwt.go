package token

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/bjr3ady/simple-go-webapi/models/orm"
	"github.com/bjr3ady/simple-go-webapi/pkg/e"
	"github.com/bjr3ady/simple-go-webapi/pkg/setting"
)

//GenerateToken generates new jwt token
func GenerateToken(admin *orm.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": admin.AdminID,
		//"exp": time.Now().Add(time.Hour * 2).Unix()
	})
	return token.SignedString([]byte(setting.JwtSecret))
}

//JWTMiddleware authenticate request with JWT
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
			return
		}
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
				return nil, fmt.Errorf("not authorized")
			}
			return []byte(setting.JwtSecret), nil
		})
		if !token.Valid {
			http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
