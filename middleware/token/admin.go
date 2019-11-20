package token

import (
	"net/http"
	"strings"

	"git.r3ady.com/golang/school-board/application/auth"
	"git.r3ady.com/golang/school-board/pkg/e"
	logger "github.com/bjr3ady/go-logger"
)

//Admin authenticate admin request.
func Admin(next http.Handler) http.Handler {
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
				Req: r,
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
