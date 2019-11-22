package router

import (
	"github.com/gorilla/mux"

	"git.r3ady.com/golang/school-board/controller"
	"git.r3ady.com/golang/school-board/middleware/token"
)

//HandleAdmin handles admin-only authenticated routes
func HandleAdmin(baseRouter *mux.Router) {
	v1 := baseRouter.PathPrefix("/api/v1/admin").Subrouter()
	v1.Use(token.Admin)

	//Admin web apis
	v1.HandleFunc("/admin/all", controller.GetAllAdmins).Methods("GET")
	v1.HandleFunc("/admin", controller.CreateAdmin).Methods("POST")
	v1.HandleFunc("/admin/{id}", controller.GetOneAdmin).Methods("GET")
	v1.HandleFunc("/admin/{id}", controller.UpdateAdmin).Methods("PUT")
	v1.HandleFunc("/admin/pwd/{id}", controller.UpdateAdminPassword).Methods("PUT")
	v1.HandleFunc("/admin/{id}", controller.DeleteAdmin).Methods("DELETE")

	//Function web apis
	v1.HandleFunc("/function/all", controller.GetAllFuncs).Methods("GET")
	v1.HandleFunc("/function/default", controller.GetDefaultFunc).Methods("GET")
	v1.HandleFunc("/function", controller.CreateFunc).Methods("POST")
	v1.HandleFunc("/function/{id}", controller.GetOneFunc).Methods("GET")
	v1.HandleFunc("/function/{id}", controller.UpdateFunc).Methods("PUT")
	v1.HandleFunc("/function/{id}", controller.DeleteFunc).Methods("DELETE")

	//Role web apis
	v1.HandleFunc("/role/all", controller.GetAllRoles).Methods("GET")
	v1.HandleFunc("/role/default", controller.GetDefaultRole).Methods("GET")
	v1.HandleFunc("/role", controller.CreateRole).Methods("POST")
	v1.HandleFunc("/role/{id}", controller.GetOneRole).Methods("GET")
	v1.HandleFunc("/role/{id}", controller.UpdateRole).Methods("PUT")
	v1.HandleFunc("/role/{id}", controller.DeleteRole).Methods("DELETE")

	//Category web apis
	v1.HandleFunc("/category/all", controller.GetAllCategories).Methods("GET")
	v1.HandleFunc("/category/default", controller.GetDefaultCategory).Methods("GET")
	v1.HandleFunc("/category", controller.CreateCategory).Methods("POST")
	v1.HandleFunc("/category/{id}", controller.GetOneCategory).Methods("GET")
	v1.HandleFunc("/category/{id}", controller.UpdateCategory).Methods("PUT")
	v1.HandleFunc("/category/{id}", controller.DeleteCategory).Methods("DELETE")
}