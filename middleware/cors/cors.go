package cors

import (
	"net/http"
)

//MiddleWare is the mux middleware to handle CORS request.
func MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) 
		w.Header().Set("Access-Control-Allow-Credentials", "true") 
		w.Header().Add("Access-Control-Allow-Method","POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE") 
		w.Header().Add("Access-Control-Allow-Headers","Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie") 
		w.Header().Set("Content-Type","application/json") 
		next.ServeHTTP(w, r)
	})
}