package routers

import (
	"rhodopsin-micros/ips/controllers"
	"github.com/gorilla/mux"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
router.HandleFunc("/ips", controllers.GetIps).Methods("GET")
router.HandleFunc("/ips/{id}", controllers.GetId).Methods("GET")
router.HandleFunc("/users/malicious", controllers.GetMalicious).Methods("GET")
return router
}