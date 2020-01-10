package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	logger "github.com/bjr3ady/go-logger"
	"git.r3ady.com/golang/simple-go-webapi/controller"
)

//InitRouter routes all request to controllers.
func InitRouter() http.Handler {
	r := mux.NewRouter()
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recover:", r)
		}
	}()

	r.HandleFunc("/service-info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API-Service"))
	}).Methods("GET")

	r.HandleFunc("/api/testtoken", controller.TestToken).Methods("POST")
	
	r.HandleFunc("/api/v1/admin/login", controller.LoginAdmin).Methods("POST")

	HandleAdmin(r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	return c.Handler(r)
}
