package token

import (
	"net/http"
	"strings"

	logger "github.com/bjr3ady/go-logger"
	"github.com/bjr3ady/simple-go-webapi/application/auth"
	"github.com/bjr3ady/simple-go-webapi/pkg/e"
)

//BearerMiddleware authenticate admin request.
func BearerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
			return
		}

		token := strings.Trim(strings.Split(authHeader, "Bearer")[1], " ")
		adminid := r.Header.Get("adminid")

		var adminAuth auth.Authler
		if token != "" || adminid == "" {
			adminAuth = &auth.AdminAuth{
				AdminID: adminid,
				Bearer:  token,
				Req:     r,
			}
			checked, err := adminAuth.Auth()
			if checked {
				next.ServeHTTP(w, r)
			} else {
				logger.Debug(err)
				http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
			}
		} else {
			http.Error(w, e.GetMsg(e.ERROR_AUTH_TOKEN), http.StatusUnauthorized)
		}
	})
}
