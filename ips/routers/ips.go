package routers

import (
	"github.com/gorilla/mux"
	"rhodopsin-micros/ips/controllers"
)

// SetIpRoutes configures routes for ip entity
func SetIpRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ips", controllers.GetIps).Methods("GET")
	// todo: implement route for getting IPs by ipaddr value
	//router.HandleFunc("/ips/{id}", controllers.GetIpAddr()).Methods("GET")
	// todo: implement route for getting IPs by database ID value
	// ipRouter.HandleFunc("/ips/{id}", controllers.GetById).Methods("GET")

	return router
}
